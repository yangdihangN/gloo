package k8sadmisssion

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
	gwv1 "github.com/solo-io/gloo/projects/gateway/pkg/api/v1"
	v2 "github.com/solo-io/gloo/projects/gateway/pkg/api/v2"
	"github.com/solo-io/gloo/projects/gateway/pkg/validation"
	"github.com/solo-io/go-utils/contextutils"
	v1 "github.com/solo-io/solo-kit/pkg/api/v1/clients/kube/crd/solo.io/v1"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"github.com/solo-io/solo-kit/pkg/utils/kubeutils"
	"github.com/solo-io/solo-kit/pkg/utils/protoutils"
	"k8s.io/api/admission/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
)

const (
	ValidationPath    = "/validation"
	skipValidationKey = "gateway.solo.io/skip_validation"
)

var (
	runtimeScheme = runtime.NewScheme()
	codecs        = serializer.NewCodecFactory(runtimeScheme)
	deserializer  = codecs.UniversalDeserializer()
)

func skipValidationCheck(annotations map[string]string) bool {
	if annotations == nil {
		return false
	}
	return annotations[skipValidationKey] == "true"
}

func NewGatewayValidatingWebhook(ctx context.Context, validator validation.Validator, port int, serverCertPath, serverKeyPath string) (*http.Server, error) {
	keyPair, err := tls.LoadX509KeyPair(serverCertPath, serverKeyPath)
	if err != nil {
		return nil, errors.Wrapf(err, "loading x509 key pair")
	}

	handler := &gatewayValidationWebhook{
		ctx:       contextutils.WithLogger(ctx, "gateway-validation-webhook"),
		validator: validator,
	}

	mux := http.NewServeMux()
	mux.Handle(ValidationPath, handler)

	return &http.Server{
		Addr:      fmt.Sprintf(":%v", port),
		TLSConfig: &tls.Config{Certificates: []tls.Certificate{keyPair}},
		Handler:   mux,
	}, nil

}

type gatewayValidationWebhook struct {
	ctx       context.Context
	validator validation.Validator
}

func (wh *gatewayValidationWebhook) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	logger := contextutils.LoggerFrom(wh.ctx)

	logger.Infow("received validation request")

	// Verify the content type is accurate
	contentType := r.Header.Get("Content-Type")
	if contentType != "application/json" {
		logger.Errorf("contentType=%s, expecting application/json", contentType)
		http.Error(w, "empty body", http.StatusBadRequest)
		return
	}

	var body []byte
	if r.Body != nil {
		if data, err := ioutil.ReadAll(r.Body); err == nil {
			body = data
		}
		defer r.Body.Close()
	}
	if len(body) == 0 {
		logger.Errorf("empty body")
		http.Error(w, "empty body", http.StatusBadRequest)
		return
	}

	var (
		admissionResponse *v1beta1.AdmissionResponse
		review            v1beta1.AdmissionReview
	)
	if _, _, err := deserializer.Decode(body, nil, &review); err != nil {
		logger.Errorf("Can't decode body: %v", err)
		admissionResponse = &v1beta1.AdmissionResponse{
			Result: &metav1.Status{
				Message: err.Error(),
			},
		}
	} else {
		admissionResponse = wh.validate(wh.ctx, &review)
	}

	admissionReview := v1beta1.AdmissionReview{}
	if admissionResponse != nil {
		admissionReview.Response = admissionResponse
		if review.Request != nil {
			admissionReview.Response.UID = review.Request.UID
		}
	}

	resp, err := json.Marshal(admissionReview)
	if err != nil {
		logger.Errorf("Can't encode response: %v", err)
		http.Error(w, fmt.Sprintf("could not encode response: %v", err), http.StatusInternalServerError)
		return
	}
	logger.Infof("Ready to write response ...")
	if _, err := w.Write(resp); err != nil {
		logger.Errorf("Can't write response: %v", err)
		http.Error(w, fmt.Sprintf("could not write response: %v", err), http.StatusInternalServerError)
	}

	logger.Infof("responded with review: %#v", admissionResponse)
}
func (wh *gatewayValidationWebhook) validate(ctx context.Context, review *v1beta1.AdmissionReview) *v1beta1.AdmissionResponse {
	logger := contextutils.LoggerFrom(ctx)

	req := review.Request

	logger.Infow("AdmissionReview for Kind=%v, Namespace=%v Name=%v UID=%v patchOperation=%v UserInfo=%v",
		req.Kind, req.Namespace, req.Name, req.UID, req.Operation, req.UserInfo)

	gvk := schema.GroupVersionKind{
		Group:   req.Kind.Group,
		Version: req.Kind.Version,
		Kind:    req.Kind.Kind,
	}

	var valdationErr error

	isDelete := req.Operation == v1beta1.Delete

	ref := core.ResourceRef{
		Namespace: req.Namespace,
		Name:      req.Name,
	}

	switch gvk {
	case v2.GatewayGVK:
		if isDelete {
			// we don't validate gateway deletion
			break
		}
		valdationErr = wh.validateGateway(ctx, req.Object.Raw)
	case gwv1.VirtualServiceGVK:
		if isDelete {
			valdationErr = wh.validator.ValidateDeleteVirtualService(ctx, ref)
		} else {
			valdationErr = wh.validateVirtualService(ctx, req.Object.Raw)
		}
	case gwv1.RouteTableGVK:
		if isDelete {
			valdationErr = wh.validator.ValidateDeleteRouteTable(ctx, ref)
		} else {
			valdationErr = wh.validateRouteTable(ctx, req.Object.Raw)
		}
	}

	if valdationErr != nil {
		logger.Errorf("Validation failed: %v", valdationErr)
		return &v1beta1.AdmissionResponse{
			Result: &metav1.Status{
				Reason:  metav1.StatusReasonInvalid,
				Message: valdationErr.Error(),
			},
		}
	}

	logger.Debug("Succeeded")

	return &v1beta1.AdmissionResponse{
		Allowed: true,
	}
}

func (wh *gatewayValidationWebhook) validateGateway(ctx context.Context, rawJson []byte) error {
	var gw v2.Gateway
	if err := unmarshalResource(rawJson, &gw); err != nil {
		return errors.Wrapf(err, "could not unmarshal raw object")
	}
	if skipValidationCheck(gw.Metadata.Annotations) {
		return nil
	}
	if err := wh.validator.ValidateGateway(ctx, &gw); err != nil {
		return errors.Wrapf(err, "Validating %T failed", gw)
	}
	return nil
}

func (wh *gatewayValidationWebhook) validateVirtualService(ctx context.Context, rawJson []byte) error {
	var vs gwv1.VirtualService
	if err := unmarshalResource(rawJson, &vs); err != nil {
		return errors.Wrapf(err, "could not unmarshal raw object")
	}
	if skipValidationCheck(vs.Metadata.Annotations) {
		return nil
	}
	if err := wh.validator.ValidateVirtualService(ctx, &vs); err != nil {
		return errors.Wrapf(err, "Validating %T failed", vs)
	}
	return nil
}

func (wh *gatewayValidationWebhook) validateRouteTable(ctx context.Context, rawJson []byte) error {
	var rt gwv1.RouteTable
	if err := unmarshalResource(rawJson, &rt); err != nil {
		return errors.Wrapf(err, "could not unmarshal raw object")
	}
	if skipValidationCheck(rt.Metadata.Annotations) {
		return nil
	}
	if err := wh.validator.ValidateRouteTable(ctx, &rt); err != nil {
		return errors.Wrapf(err, "Validating %T failed", rt)
	}
	return nil
}

func unmarshalResource(kubeJson []byte, resource resources.Resource) error {
	var resourceCrd v1.Resource
	if err := json.Unmarshal(kubeJson, &resourceCrd); err != nil {
		return errors.Wrapf(err, "unmarshalling from raw json")
	}
	resource.SetMetadata(kubeutils.FromKubeMeta(resourceCrd.ObjectMeta))
	if withStatus, ok := resource.(resources.InputResource); ok {
		resources.UpdateStatus(withStatus, func(status *core.Status) {
			*status = resourceCrd.Status
		})
	}

	if resourceCrd.Spec != nil {
		if err := protoutils.UnmarshalMap(*resourceCrd.Spec, resource); err != nil {
			return errors.Wrapf(err, "parsing resource from crd spec %v in namespace %v into %T", resourceCrd.Name, resourceCrd.Namespace, resource)
		}
	}

	return nil
}

package validation

import (
	"context"
	"sync"

	"github.com/pkg/errors"
	v1 "github.com/solo-io/gloo/projects/gateway/pkg/api/v1"
	v2 "github.com/solo-io/gloo/projects/gateway/pkg/api/v2"
	"github.com/solo-io/gloo/projects/gateway/pkg/translator"
	"github.com/solo-io/gloo/projects/gateway/pkg/utils"
	"github.com/solo-io/gloo/projects/gloo/pkg/api/grpc/validation"
	validationutils "github.com/solo-io/gloo/projects/gloo/pkg/utils/validation"
	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/reporter"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"go.uber.org/zap"
)

type Validator interface {
	v2.ApiSyncer
	Ready() bool
	ValidateGateway(ctx context.Context, gw *v2.Gateway) error
	ValidateVirtualService(ctx context.Context, vs *v1.VirtualService) error
	ValidateRouteTable(ctx context.Context, rt *v1.RouteTable) error
	ValidateDeleteRouteTable(ctx context.Context, rt core.ResourceRef) error
	ValidateDeleteVirtualService(ctx context.Context, vs core.ResourceRef) error
}

type validator struct {
	l                sync.RWMutex
	latestSnapshot   *v2.ApiSnapshot
	translator       translator.Translator
	validationClient validation.ProxyValidationServiceClient
	writeNamespace   string
}

func NewValidator(translator translator.Translator, validationClient validation.ProxyValidationServiceClient, writeNamespace string) *validator {
	return &validator{translator: translator, validationClient: validationClient, writeNamespace: writeNamespace}
}

func (v *validator) Ready() bool {
	return v.latestSnapshot != nil
}

func (v *validator) Sync(_ context.Context, snap *v2.ApiSnapshot) error {
	snapCopy := snap.Clone()
	v.l.Lock()
	v.latestSnapshot = &snapCopy
	v.l.Unlock()
	return nil
}

func (v *validator) ValidateVirtualService(ctx context.Context, vs *v1.VirtualService) error {
	if !v.Ready() {
		return errors.Errorf("Gateway validation is yet not available. Waiting for first snapshot")
	}
	v.l.RLock()
	snap := v.latestSnapshot.Clone()
	v.l.RUnlock()

	vsRef := vs.GetMetadata().Ref()

	var isUpdate bool
	for i, existingVs := range snap.VirtualServices {
		if vsRef == existingVs.GetMetadata().Ref() {
			// replace the existing virtual service in the snapshot
			snap.VirtualServices[i] = vs
			isUpdate = true
			break
		}
	}
	if !isUpdate {
		snap.VirtualServices = append(snap.VirtualServices, vs)
		snap.VirtualServices.Sort()
	}

	gatewaysByProxy := utils.GatewaysByProxyName(snap.Gateways)

	for proxyName, gatewayList := range gatewaysByProxy {
		if !gatewayListContainsVirtualService(gatewayList, snap.VirtualServices, vs) {
			// we only care about validating this proxy if it contains this virtual service
			continue
		}

		proxy, resourceErrs := v.translator.Translate(ctx, proxyName, v.writeNamespace, &snap, gatewayList)
		if err := resourceErrs.Validate(); err != nil {
			return errors.Wrapf(err, "could not render proxy from %T %v", vs, vsRef)
		}

		if v.validationClient == nil {
			contextutils.LoggerFrom(ctx).Warnf("skipping proxy validation check as the " +
				"Proxy validation client has not been initialized. check to ensure that the gateway and gloo processes " +
				"are configured to communicate.")
			return nil
		}

		// validate the proxy with gloo
		proxyReport, err := v.validationClient.ValidateProxy(ctx, proxy)
		if err != nil {
			contextutils.LoggerFrom(ctx).Errorw("failed to validate Proxy with Gloo validation server.", zap.Error(err))
			return errors.Wrapf(err, "failed to validate Proxy with Gloo validation server")
		}

		if proxyErr := validationutils.GetProxyError(proxyReport); proxyErr != nil {
			return errors.Wrapf(err, "rendered proxy had errors")
		}
	}

	contextutils.LoggerFrom(ctx).Debugw("Accepted %T %v", vs, vsRef)

	return nil
}

func (v *validator) ValidateRouteTable(ctx context.Context, rt *v1.RouteTable) error {
	panic("implement me")
}

func (v *validator) ValidateDeleteRouteTable(ctx context.Context, rt core.ResourceRef) error {
	panic("implement me")
}

func (v *validator) ValidateGateway(ctx context.Context, gw *v2.Gateway) error {
	panic("implement me")
}

func (v *validator) ValidateDeleteVirtualService(ctx context.Context, vs core.ResourceRef) error {
	panic("implement me")
}

func gatewayListContainsVirtualService(gwList v2.GatewayList, vsList v1.VirtualServiceList, vs *v1.VirtualService) bool {
	for _, gw := range gwList {
		// we don't care about validating the gateways here, pass an anonymous ResourceErrs
		vssForGateway := translator.GetVirtualServicesForGateway(gw, vsList, reporter.ResourceErrors{})

		if _, err := vssForGateway.Find(vs.Metadata.Ref().Strings()); err == nil {
			return true
		}
	}

	return false
}

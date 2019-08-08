package check

import (
	"fmt"

	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/solo-io/gloo/projects/gloo/cli/pkg/helpers"
	"github.com/solo-io/gloo/projects/gloo/pkg/plugins/aws/ec2"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"

	"github.com/solo-io/gloo/projects/gloo/cli/pkg/cmd/options"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/common"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
)

func checkUpstreams(opts *options.Options, args []string) (*Response, error) {
	upstreams, err := common.GetUpstreams(common.GetName(args, opts), opts)
	if err != nil {
		return nil, err
	}
	secrets, err := getSecrets(opts)
	if err != nil {
		return nil, err
	}
	resp := NewResponse("upstreams")
	for _, upstream := range upstreams {
		upstreamResp, err := checkUpstream(upstream, secrets)
		if err != nil {
			return nil, err
		}
		resp.AddResponse(upstreamResp)
	}
	return resp, nil
}

func checkUpstream(upstream *v1.Upstream, secrets v1.SecretList) (*Response, error) {
	response := NewResponseFromMetadata(upstream.Metadata)
	response.AddResponse(statusCheck(upstream.Status))
	typeSpecificResponse, err := upstreamTypeSpecificChecks(upstream, secrets)
	if err != nil {
		return nil, err
	}
	response.AddResponse(typeSpecificResponse)
	return response, nil
}

func statusCheck(status core.Status) *Response {
	resp := NewResponse("status")
	if status.State == core.Status_Rejected {
		resp.LevelError = status.Reason
	}
	return resp
}

func upstreamTypeSpecificChecks(upstream *v1.Upstream, secrets v1.SecretList) (*Response, error) {
	switch upstream.UpstreamSpec.UpstreamType.(type) {
	case *v1.UpstreamSpec_AwsEc2:
		resp := NewResponse("EC2")
		instances, err := ec2.InstancesForUpstream(upstream, secrets)
		if err != nil {
			return nil, err
		}
		resp.Details = fmt.Sprintf("instances:\n%v", ec2.SummarizeInstances(instances))
		return resp, nil
	default:
		return NewResponse("no type-specific info"), nil
	}
}

func getSecrets(opts *options.Options) (v1.SecretList, error) {
	var secrets v1.SecretList
	var err error
	var nsList []string
	if opts.Check.AllNamespaces {
		kubeClient, err := helpers.KubeClient()
		if err != nil {
			return nil, err
		}
		nss, err := kubeClient.CoreV1().Namespaces().List(v12.ListOptions{})
		if err != nil {
			return nil, err
		}
		for _, ns := range nss.Items {
			nsList = append(nsList, ns.Name)
		}
	} else if len(opts.Check.NamespaceList) > 0 {
		nsList = opts.Check.NamespaceList
	}
	secretClient := helpers.MustSecretClient()
	listOpts := clients.ListOpts{
		Ctx: opts.Top.Ctx,
	}
	if len(nsList) > 0 {
		for _, ns := range nsList {
			nsSecrets, err := secretClient.List(ns, listOpts)
			if err != nil {
				return nil, err
			}
			secrets = append(secrets, nsSecrets...)
		}
	} else {
		secrets, err = secretClient.List(opts.Metadata.Namespace, listOpts)
		if err != nil {
			return nil, err
		}
	}
	return secrets, nil
}

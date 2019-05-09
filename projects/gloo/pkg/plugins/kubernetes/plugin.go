package kubernetes

import (
	"fmt"
	"net/url"

	envoyapi "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	envoycore "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	types "github.com/gogo/protobuf/types"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gloo/pkg/plugins"
	"github.com/solo-io/gloo/projects/gloo/pkg/xds"
	"k8s.io/client-go/kubernetes"
)

type plugin struct {
	kube kubernetes.Interface

	kubeShareFactory KubePluginSharedFactory

	UpstreamConverter UpstreamConverter
}

func (p *plugin) Resolve(u *v1.Upstream) (*url.URL, error) {
	kubeSpec, ok := u.UpstreamSpec.UpstreamType.(*v1.UpstreamSpec_Kube)
	if !ok {
		return nil, nil
	}

	return url.Parse(fmt.Sprintf("tcp://%v.%v.svc.cluster.local:%v", kubeSpec.Kube.ServiceName, kubeSpec.Kube.ServiceNamespace, kubeSpec.Kube.ServicePort))
}

func NewPlugin(kube kubernetes.Interface) plugins.Plugin {
	return &plugin{
		kube:              kube,
		UpstreamConverter: DefaultUpstreamConverter(),
	}
}

func (p *plugin) Init(params plugins.InitParams) error {
	return nil
}

func (p *plugin) ProcessUpstream(params plugins.Params, in *v1.Upstream, out *envoyapi.Cluster) error {
	// not ours
	up, ok := in.UpstreamSpec.UpstreamType.(*v1.UpstreamSpec_Kube)
	if !ok {
		return nil
	}

	out.Metadata = &envoycore.Metadata{
		FilterMetadata: map[string]*types.Struct{
			"envoy.filters.http.tap": &types.Struct{
				Fields: map[string]*types.Value{
					"svc":   tostring(up.Kube.GetServiceName()),
					"svcns": tostring(up.Kube.GetServiceNamespace()),
				},
			},
		},
	}

	// configure the cluster to use EDS:ADS and call it a day
	xds.SetEdsOnCluster(out)
	return nil
}

func tostring(s string) *types.Value {
	return &types.Value{
		Kind: &types.Value_StringValue{
			StringValue: s,
		},
	}
}

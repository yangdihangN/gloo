package kube

import (
	"fmt"
	"strings"

	gloov1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gloo/pkg/api/v1/plugins"
	grpcplugin "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/plugins/grpc"
	kubeplugin "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/plugins/kubernetes"
	snapkube "github.com/solo-io/solo-kit/pkg/api/v1/resources/common/kubernetes"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"github.com/solo-io/solo-kit/pkg/utils/kubeutils"
	kubev1 "k8s.io/api/core/v1"
)

const GlooH2Annotation = "gloo.solo.io/h2_service"

func ShouldStore(u *gloov1.Upstream) bool {
	return !strings.HasPrefix(u.Metadata.Name, "svc:")
}

func SvcRefToUpstreamRef(name, namespace string, port int32) core.ResourceRef {
	return core.ResourceRef{
		Name:      fmt.Sprintf("svc:%s-%d", name, port),
		Namespace: namespace,
	}
}

func Combine(svcs snapkube.ServiceList, ups []*gloov1.Upstream) []*gloov1.Upstream {
	ups2 := make([]*gloov1.Upstream, 0, len(svcs))
	for _, svc := range svcs {
		for _, p := range svc.Spec.Ports {
			kubeSvc := kubev1.Service(svc.Service)
			ups2 = append(ups2, SvcToUpstream(&kubeSvc, p))
		}
	}

	return append(ups, ups2...)
}

func SvcToUpstream(svc *kubev1.Service, port kubev1.ServicePort) *gloov1.Upstream {
	meta := svc.ObjectMeta
	coremeta := kubeutils.FromKubeMeta(meta)
	coremeta.ResourceVersion = ""

	// invalid name so it fails to be written if it is ever attempted.
	svcRef := SvcRefToUpstreamRef(meta.Name, meta.Namespace, port.Port)
	coremeta.Name = svcRef.Name
	coremeta.Namespace = svcRef.Namespace
	return &gloov1.Upstream{
		Metadata: coremeta,
		UpstreamSpec: &v1.UpstreamSpec{
			UpstreamType: &v1.UpstreamSpec_Kube{
				Kube: &kubeplugin.UpstreamSpec{
					ServiceSpec:      getServiceSpec(svc, port),
					ServiceName:      meta.Name,
					ServiceNamespace: meta.Namespace,
					ServicePort:      uint32(port.Port),
				},
			},
		},
		DiscoveryMetadata: &v1.DiscoveryMetadata{},
	}
}

func getServiceSpec(svc *kubev1.Service, port kubev1.ServicePort) *plugins.ServiceSpec {
	grpcSpec := &plugins.ServiceSpec{
		PluginType: &plugins.ServiceSpec_Grpc{
			Grpc: &grpcplugin.ServiceSpec{},
		},
	}

	if svc.Annotations != nil {
		if svc.Annotations[GlooH2Annotation] == "true" {
			return grpcSpec
		} else if svc.Annotations[GlooH2Annotation] == "false" {
			return nil
		}
	}

	if strings.HasPrefix(port.Name, "grpc") || strings.HasPrefix(port.Name, "h2") {
		return grpcSpec
	}

	return nil
}

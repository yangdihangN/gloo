package utils

// this package exists outside of the glooec2 and v1 api packages to avoid import issues

import (
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gloo/pkg/api/v1/plugins/aws/glooec2"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
)

type InvertedEc2Upstream struct {
	Spec     *glooec2.UpstreamSpec
	Upstream *v1.Upstream
}

type InvertedEc2UpstreamRefMap map[core.ResourceRef]*InvertedEc2Upstream

func BuildInvertedUpstreamRefMap(upstreams v1.UpstreamList) InvertedEc2UpstreamRefMap {
	upstreamSpecs := make(InvertedEc2UpstreamRefMap)
	typeOk := true
	for _, upstream := range upstreams {
		inverted := InvertEc2Upstream(upstream, &typeOk)
		// only care about ec2 upstreams
		if !typeOk {
			continue
		}
		ref := upstream.Metadata.Ref()
		upstreamSpecs[ref] = inverted
	}
	return upstreamSpecs
}

// InvertEc2Upstream is a helper for working with EC2 Upstreams.
// if you know that you are working EC2 Upstreams you can pass nil and ignore the type cast response
func InvertEc2Upstream(upstream *v1.Upstream, typeOk *bool) *InvertedEc2Upstream {
	var ec2Spec *glooec2.UpstreamSpec
	spec, ok := upstream.UpstreamSpec.UpstreamType.(*v1.UpstreamSpec_AwsEc2)
	// only care about ec2 upstreams
	if ok {
		ec2Spec = spec.AwsEc2
	}
	if typeOk != nil {
		*typeOk = ok
	}
	return &InvertedEc2Upstream{
		Spec:     ec2Spec,
		Upstream: upstream,
	}
}

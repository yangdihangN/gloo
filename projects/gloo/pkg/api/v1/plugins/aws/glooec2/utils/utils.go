package utils

// this package exists outside of the glooec2 and v1 api packages to avoid import issues

import (
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gloo/pkg/api/v1/plugins/aws/glooec2"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
)

type InvertedEc2Upstream struct {
	AwsEc2Spec *glooec2.UpstreamSpec
	Base       *v1.Upstream
}

type InvertedEc2UpstreamRefMap map[core.ResourceRef]*InvertedEc2Upstream

// BuildInvertedUpstreamRefMap is a helper for working with EC2 Upstreams.
// it ignores any upstreams that are not EC2 Upstreams
func BuildInvertedUpstreamRefMap(upstreams v1.UpstreamList) InvertedEc2UpstreamRefMap {
	upstreamSpecs := make(InvertedEc2UpstreamRefMap)
	for _, upstream := range upstreams {
		inverted, ok := invertEc2Upstream(upstream)
		// only care about ec2 upstreams
		if !ok {
			continue
		}
		ref := upstream.Metadata.Ref()
		upstreamSpecs[ref] = inverted
	}
	return upstreamSpecs
}

func invertEc2Upstream(upstream *v1.Upstream) (*InvertedEc2Upstream, bool) {
	spec, ok := upstream.UpstreamSpec.UpstreamType.(*v1.UpstreamSpec_AwsEc2)
	if !ok {
		return nil, false
	}
	return &InvertedEc2Upstream{
		AwsEc2Spec: spec.AwsEc2,
		Base:       upstream,
	}, true
}

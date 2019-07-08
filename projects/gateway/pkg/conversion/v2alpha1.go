package conversion

import (
	"reflect"

	"github.com/pkg/errors"
	v1 "github.com/solo-io/gloo/projects/gateway/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gateway/pkg/api/v2alpha1"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
)

var (
	ExpectedV1SrcError = func(actualType string) error {
		return errors.Errorf("Expected *v1.Gateway as src, got %v", actualType)
	}

	ExpectedV2alpha1DstError = func(actualType string) error {
		return errors.Errorf("Expected *v2alpha1.Gateway as src, got %v", actualType)
	}
)

type v2alpha1Converter struct{}

func NewV2alpha1Converter() Converter {
	return &v2alpha1Converter{}
}

func (c *v2alpha1Converter) Convert(src, dst resources.Resource) error {
	srcGateway, ok := src.(*v1.Gateway)
	if !ok {
		return ExpectedV1SrcError(reflect.TypeOf(src).String())
	}

	dstGateway, ok := src.(*v2alpha1.Gateway)
	if !ok {
		return ExpectedV2alpha1DstError(reflect.TypeOf(dst).String())
	}

	dstGateway.Ssl = srcGateway.Ssl
	dstGateway.BindAddress = srcGateway.BindAddress
	dstGateway.BindPort = srcGateway.BindPort
	dstGateway.UseProxyProto = srcGateway.UseProxyProto
	dstGateway.GatewayType = &v2alpha1.Gateway_HttpGateway{
		HttpGateway: &v2alpha1.HttpGateway{
			VirtualServices: srcGateway.VirtualServices,
			Plugins:         srcGateway.Plugins,
		},
	}

	return nil
}

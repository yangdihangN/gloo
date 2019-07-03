package conversion

import (
	"context"

	"github.com/solo-io/gloo/projects/gatewayinit/pkg/conversion/convertv2alpha1"
	"github.com/solo-io/gloo/projects/gatewayinit/pkg/setup"
)

type Converter interface {
	Convert()
}

/*
List of converters that make up the "ladder" of resource conversion.
*/
type ConverterList []Converter

func NewConverterList(ctx context.Context, clientSet setup.ClientSet, namespace string) ConverterList {
	return []Converter{
		convertv2alpha1.NewConverter(ctx, clientSet.V1Gateway, clientSet.V2alpha1Gateway, convertv2alpha1.NewGatewayConverter(), namespace),
	}
}

func (l ConverterList) ConvertAll() {
	for _, c := range l {
		c.Convert()
	}
}

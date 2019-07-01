package v1

import (
	v1 "github.com/solo-io/gloo/projects/gateway/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gateway/pkg/api/v2alpha1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	kscheme "k8s.io/client-go/kubernetes/scheme"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

var (
	api        scheme.Builder
	kubescheme *runtime.Scheme
)

func init() {
	kubescheme := kscheme.Scheme
	api.AddToScheme(kubescheme)
	api.Register(&v2alpha1.Gateway{})
}

type Gateway struct {
	v1.Gateway
}

type Hub struct {
}

func (*Hub) GetObjectKind() schema.ObjectKind {
	panic("implement me")
}

func (*Hub) DeepCopyObject() runtime.Object {
	panic("implement me")
}

func (*Hub) Hub() {
	panic("implement me")
}

func (g *Gateway) ConvertTo(dst conversion.Hub) error {
	panic("implement me")
}

func (g *Gateway) ConvertFrom(src conversion.Hub) error {
	panic("implement me")
}

func (*Gateway) ConvertToTyped(dst *v2alpha1.Gateway) error {

}

func (*Gateway) ConvertFromTyped(src *v2alpha1.Gateway) error {

}

/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	gatewaysoloiov1 "github.com/solo-io/gloo/projects/gateway/pkg/api/v1/kube/apis/gateway.solo.io/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeGateways implements GatewayInterface
type FakeGateways struct {
	Fake *FakeGatewayV1
	ns   string
}

var gatewaysResource = schema.GroupVersionResource{Group: "gateway.solo.io", Version: "v1", Resource: "gateways"}

var gatewaysKind = schema.GroupVersionKind{Group: "gateway.solo.io", Version: "v1", Kind: "Gateway"}

// Get takes name of the gateway, and returns the corresponding gateway object, and an error if there is any.
func (c *FakeGateways) Get(name string, options v1.GetOptions) (result *gatewaysoloiov1.Gateway, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(gatewaysResource, c.ns, name), &gatewaysoloiov1.Gateway{})

	if obj == nil {
		return nil, err
	}
	return obj.(*gatewaysoloiov1.Gateway), err
}

// List takes label and field selectors, and returns the list of Gateways that match those selectors.
func (c *FakeGateways) List(opts v1.ListOptions) (result *gatewaysoloiov1.GatewayList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(gatewaysResource, gatewaysKind, c.ns, opts), &gatewaysoloiov1.GatewayList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &gatewaysoloiov1.GatewayList{ListMeta: obj.(*gatewaysoloiov1.GatewayList).ListMeta}
	for _, item := range obj.(*gatewaysoloiov1.GatewayList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested gateways.
func (c *FakeGateways) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(gatewaysResource, c.ns, opts))

}

// Create takes the representation of a gateway and creates it.  Returns the server's representation of the gateway, and an error, if there is any.
func (c *FakeGateways) Create(gateway *gatewaysoloiov1.Gateway) (result *gatewaysoloiov1.Gateway, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(gatewaysResource, c.ns, gateway), &gatewaysoloiov1.Gateway{})

	if obj == nil {
		return nil, err
	}
	return obj.(*gatewaysoloiov1.Gateway), err
}

// Update takes the representation of a gateway and updates it. Returns the server's representation of the gateway, and an error, if there is any.
func (c *FakeGateways) Update(gateway *gatewaysoloiov1.Gateway) (result *gatewaysoloiov1.Gateway, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(gatewaysResource, c.ns, gateway), &gatewaysoloiov1.Gateway{})

	if obj == nil {
		return nil, err
	}
	return obj.(*gatewaysoloiov1.Gateway), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeGateways) UpdateStatus(gateway *gatewaysoloiov1.Gateway) (*gatewaysoloiov1.Gateway, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(gatewaysResource, "status", c.ns, gateway), &gatewaysoloiov1.Gateway{})

	if obj == nil {
		return nil, err
	}
	return obj.(*gatewaysoloiov1.Gateway), err
}

// Delete takes name of the gateway and deletes it. Returns an error if one occurs.
func (c *FakeGateways) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(gatewaysResource, c.ns, name), &gatewaysoloiov1.Gateway{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGateways) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(gatewaysResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &gatewaysoloiov1.GatewayList{})
	return err
}

// Patch applies the patch and returns the patched gateway.
func (c *FakeGateways) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *gatewaysoloiov1.Gateway, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(gatewaysResource, c.ns, name, pt, data, subresources...), &gatewaysoloiov1.Gateway{})

	if obj == nil {
		return nil, err
	}
	return obj.(*gatewaysoloiov1.Gateway), err
}

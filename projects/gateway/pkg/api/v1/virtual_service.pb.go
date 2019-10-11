// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gateway/api/v1/virtual_service.proto

package v1

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

//
//
// The **VirtualService** is the root Routing object for the Gloo Gateway.
// A virtual service describes the set of routes to match for a set of domains.
//
// It defines:
// - a set of domains
// - the root set of routes for those domains
// - an optional SSL configuration for server TLS Termination
// - VirtualHostPlugins that will apply configuration to all routes that live on the VirtualService.
//
// Domains must be unique across all virtual services within a gateway (i.e. no overlap between sets).
//
// VirtualServices can delegate routing behavior to the RouteTable resource by using the `delegateAction` on routes.
//
// An example configuration using two VirtualServices (one with TLS termination and one without) which share
// a RouteTable looks as follows:
//
// ```yaml
// # HTTP VirtualService:
// apiVersion: gateway.solo.io/v1
// kind: VirtualService
// metadata:
//   name: 'http'
//   namespace: 'usernamespace'
// spec:
//   virtualHost:
//     domains:
//     - '*.mydomain.com'
//     - 'mydomain.com'
//     routes:
//     - matcher:
//         prefix: '/'
//       # delegate all traffic to the `shared-routes` RouteTable
//       delegateAction:
//         name: 'shared-routes'
//         namespace: 'usernamespace'
//
// ```
//
// ```yaml
// # HTTPS VirtualService:
// apiVersion: gateway.solo.io/v1
// kind: VirtualService
// metadata:
//   name: 'https'
//   namespace: 'usernamespace'
// spec:
//   virtualHost:
//     domains:
//     - '*.mydomain.com'
//     - 'mydomain.com'
//     routes:
//     - matcher:
//         prefix: '/'
//       # delegate all traffic to the `shared-routes` RouteTable
//       delegateAction:
//         name: 'shared-routes'
//         namespace: 'usernamespace'
//   sslConfig:
//     secretRef:
//       name: gateway-tls
//       namespace: gloo-system
//
// ```
//
// ```yaml
// # the RouteTable shared by both VirtualServices:
// apiVersion: gateway.solo.io/v1
// kind: RouteTable
// metadata:
//   name: 'shared-routes'
//   namespace: 'usernamespace'
// spec:
//   routes:
//     - matcher:
//         prefix: '/some-route'
//       routeAction:
//         single:
//           upstream:
//             name: 'some-upstream'
//      ...
// ```
//
// **Delegated Routes** are routes that use the `delegateAction` routing action. Delegated Routes obey the following
// constraints:
//
// - delegate routes must use `prefix` path matchers
// - delegated routes cannot specify header, query, or methods portion of the normal route matcher.
// - `routePlugin` configuration will be inherited from parent routes, but can be overridden by the child
//
type VirtualService struct {
	// The VirtualHost contains the
	// The list of HTTP routes define routing actions to be taken
	// for incoming HTTP requests whose host header matches
	// this virtual host. If the request matches more than one route in the list, the first route matched will be selected.
	// If the list of routes is empty, the virtual host will be ignored by Gloo.
	VirtualHost *VirtualHost `protobuf:"bytes,1,opt,name=virtual_host,json=virtualHost,proto3" json:"virtual_host,omitempty"`
	// If provided, the Gateway will serve TLS/SSL traffic for this set of routes
	SslConfig *v1.SslConfig `protobuf:"bytes,2,opt,name=ssl_config,json=sslConfig,proto3" json:"ssl_config,omitempty"`
	// Display only, optional descriptive name.
	// Unlike metadata.name, DisplayName can be any string
	// and can be changed after creating the resource.
	DisplayName string `protobuf:"bytes,3,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Status indicates the validation status of this resource.
	// Status is read-only by clients, and set by gloo during validation
	Status core.Status `protobuf:"bytes,6,opt,name=status,proto3" json:"status" testdiff:"ignore"`
	// Metadata contains the object metadata for this resource
	Metadata             core.Metadata `protobuf:"bytes,7,opt,name=metadata,proto3" json:"metadata"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *VirtualService) Reset()         { *m = VirtualService{} }
func (m *VirtualService) String() string { return proto.CompactTextString(m) }
func (*VirtualService) ProtoMessage()    {}
func (*VirtualService) Descriptor() ([]byte, []int) {
	return fileDescriptor_93fa9472926a2049, []int{0}
}
func (m *VirtualService) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualService.Unmarshal(m, b)
}
func (m *VirtualService) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualService.Marshal(b, m, deterministic)
}
func (m *VirtualService) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualService.Merge(m, src)
}
func (m *VirtualService) XXX_Size() int {
	return xxx_messageInfo_VirtualService.Size(m)
}
func (m *VirtualService) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualService.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualService proto.InternalMessageInfo

func (m *VirtualService) GetVirtualHost() *VirtualHost {
	if m != nil {
		return m.VirtualHost
	}
	return nil
}

func (m *VirtualService) GetSslConfig() *v1.SslConfig {
	if m != nil {
		return m.SslConfig
	}
	return nil
}

func (m *VirtualService) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *VirtualService) GetStatus() core.Status {
	if m != nil {
		return m.Status
	}
	return core.Status{}
}

func (m *VirtualService) GetMetadata() core.Metadata {
	if m != nil {
		return m.Metadata
	}
	return core.Metadata{}
}

//
//Virtual Hosts serve an ordered list of routes for a set of domains.
//
//An HTTP request is first matched to a virtual host based on its host header, then to a route within the virtual host.
//
//If a request is not matched to any virtual host or a route therein, the target proxy will reply with a 404.
//
//Unlike the [Gloo Virtual Host]({{< ref "/api/github.com/solo-io/gloo/projects/gloo/api/v1/proxy.proto.sk.md" >}}/#virtualhost),
//Gateway* Virtual Hosts can **delegate** their routes to `RouteTables`.
//
type VirtualHost struct {
	// deprecated. this field is ignored
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The list of domains (i.e.: matching the `Host` header of a request) that belong to this virtual host.
	// Note that the wildcard will not match the empty string. e.g. “*-bar.foo.com” will match “baz-bar.foo.com”
	// but not “-bar.foo.com”. Additionally, a special entry “*” is allowed which will match any host/authority header.
	// Only a single virtual host on a gateway can match on “*”. A domain must be unique across all
	// virtual hosts on a gateway or the config will be invalidated by Gloo
	// Domains on virtual hosts obey the same rules as [Envoy Virtual Hosts](https://github.com/envoyproxy/envoy/blob/master/api/envoy/api/v2/route/route.proto)
	Domains []string `protobuf:"bytes,2,rep,name=domains,proto3" json:"domains,omitempty"`
	// The list of HTTP routes define routing actions to be taken for incoming HTTP requests whose host header matches
	// this virtual host. If the request matches more than one route in the list, the first route matched will be selected.
	// If the list of routes is empty, the virtual host will be ignored by Gloo.
	Routes []*Route `protobuf:"bytes,3,rep,name=routes,proto3" json:"routes,omitempty"`
	// Virtual host plugins contain additional configuration to be applied to all traffic served by the Virtual Host.
	// Some configuration here can be overridden by Route Plugins.
	VirtualHostPlugins *v1.VirtualHostPlugins `protobuf:"bytes,4,opt,name=virtual_host_plugins,json=virtualHostPlugins,proto3" json:"virtual_host_plugins,omitempty"`
	// Defines a CORS policy for the virtual host
	// If a CORS policy is also defined on the route matched by the request, the policies are merged.
	// DEPRECATED set cors policy through the Virtual Host Plugin
	CorsPolicy           *v1.CorsPolicy `protobuf:"bytes,5,opt,name=cors_policy,json=corsPolicy,proto3" json:"cors_policy,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *VirtualHost) Reset()         { *m = VirtualHost{} }
func (m *VirtualHost) String() string { return proto.CompactTextString(m) }
func (*VirtualHost) ProtoMessage()    {}
func (*VirtualHost) Descriptor() ([]byte, []int) {
	return fileDescriptor_93fa9472926a2049, []int{1}
}
func (m *VirtualHost) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualHost.Unmarshal(m, b)
}
func (m *VirtualHost) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualHost.Marshal(b, m, deterministic)
}
func (m *VirtualHost) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualHost.Merge(m, src)
}
func (m *VirtualHost) XXX_Size() int {
	return xxx_messageInfo_VirtualHost.Size(m)
}
func (m *VirtualHost) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualHost.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualHost proto.InternalMessageInfo

func (m *VirtualHost) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *VirtualHost) GetDomains() []string {
	if m != nil {
		return m.Domains
	}
	return nil
}

func (m *VirtualHost) GetRoutes() []*Route {
	if m != nil {
		return m.Routes
	}
	return nil
}

func (m *VirtualHost) GetVirtualHostPlugins() *v1.VirtualHostPlugins {
	if m != nil {
		return m.VirtualHostPlugins
	}
	return nil
}

func (m *VirtualHost) GetCorsPolicy() *v1.CorsPolicy {
	if m != nil {
		return m.CorsPolicy
	}
	return nil
}

//
//
// A route specifies how to match a request and what action to take when the request is matched.
//
//
// When a request matches on a route, the route can perform one of the following actions:
// - *Route* the request to a destination
// - Reply with a *Direct Response*
// - Send a *Redirect* response to the client
// - *Delegate* the action for the request to a top-level [`RouteTable`]({{< ref "/api/github.com/solo-io/gloo/projects/gateway/api/v1/route_table.proto.sk.md" >}}) resource
// DelegateActions can be used to delegate the behavior for a set out routes with a given *prefix* to
// a top-level `RouteTable` resource.
//
type Route struct {
	// The matcher contains parameters for matching requests (i.e.: based on HTTP path, headers, etc.)
	// For delegated routes, the matcher must contain only a `prefix` path matcher and no other config
	Matcher *v1.Matcher `protobuf:"bytes,1,opt,name=matcher,proto3" json:"matcher,omitempty"`
	// The Route Action Defines what action the proxy should take when a request matches the route.
	//
	// Types that are valid to be assigned to Action:
	//	*Route_RouteAction
	//	*Route_RedirectAction
	//	*Route_DirectResponseAction
	//	*Route_DelegateAction
	Action isRoute_Action `protobuf_oneof:"action"`
	// Route Plugins extend the behavior of routes.
	// Route plugins include configuration such as retries, rate limiting, and request/response transformation.
	// RoutePlugin behavior will be inherited by delegated routes which do not specify their own `routePlugins`
	RoutePlugins         *v1.RoutePlugins `protobuf:"bytes,6,opt,name=route_plugins,json=routePlugins,proto3" json:"route_plugins,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Route) Reset()         { *m = Route{} }
func (m *Route) String() string { return proto.CompactTextString(m) }
func (*Route) ProtoMessage()    {}
func (*Route) Descriptor() ([]byte, []int) {
	return fileDescriptor_93fa9472926a2049, []int{2}
}
func (m *Route) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Route.Unmarshal(m, b)
}
func (m *Route) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Route.Marshal(b, m, deterministic)
}
func (m *Route) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Route.Merge(m, src)
}
func (m *Route) XXX_Size() int {
	return xxx_messageInfo_Route.Size(m)
}
func (m *Route) XXX_DiscardUnknown() {
	xxx_messageInfo_Route.DiscardUnknown(m)
}

var xxx_messageInfo_Route proto.InternalMessageInfo

type isRoute_Action interface {
	isRoute_Action()
	Equal(interface{}) bool
}

type Route_RouteAction struct {
	RouteAction *v1.RouteAction `protobuf:"bytes,2,opt,name=route_action,json=routeAction,proto3,oneof" json:"route_action,omitempty"`
}
type Route_RedirectAction struct {
	RedirectAction *v1.RedirectAction `protobuf:"bytes,3,opt,name=redirect_action,json=redirectAction,proto3,oneof" json:"redirect_action,omitempty"`
}
type Route_DirectResponseAction struct {
	DirectResponseAction *v1.DirectResponseAction `protobuf:"bytes,4,opt,name=direct_response_action,json=directResponseAction,proto3,oneof" json:"direct_response_action,omitempty"`
}
type Route_DelegateAction struct {
	DelegateAction *core.ResourceRef `protobuf:"bytes,5,opt,name=delegate_action,json=delegateAction,proto3,oneof" json:"delegate_action,omitempty"`
}

func (*Route_RouteAction) isRoute_Action()          {}
func (*Route_RedirectAction) isRoute_Action()       {}
func (*Route_DirectResponseAction) isRoute_Action() {}
func (*Route_DelegateAction) isRoute_Action()       {}

func (m *Route) GetAction() isRoute_Action {
	if m != nil {
		return m.Action
	}
	return nil
}

func (m *Route) GetMatcher() *v1.Matcher {
	if m != nil {
		return m.Matcher
	}
	return nil
}

func (m *Route) GetRouteAction() *v1.RouteAction {
	if x, ok := m.GetAction().(*Route_RouteAction); ok {
		return x.RouteAction
	}
	return nil
}

func (m *Route) GetRedirectAction() *v1.RedirectAction {
	if x, ok := m.GetAction().(*Route_RedirectAction); ok {
		return x.RedirectAction
	}
	return nil
}

func (m *Route) GetDirectResponseAction() *v1.DirectResponseAction {
	if x, ok := m.GetAction().(*Route_DirectResponseAction); ok {
		return x.DirectResponseAction
	}
	return nil
}

func (m *Route) GetDelegateAction() *core.ResourceRef {
	if x, ok := m.GetAction().(*Route_DelegateAction); ok {
		return x.DelegateAction
	}
	return nil
}

func (m *Route) GetRoutePlugins() *v1.RoutePlugins {
	if m != nil {
		return m.RoutePlugins
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Route) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Route_RouteAction)(nil),
		(*Route_RedirectAction)(nil),
		(*Route_DirectResponseAction)(nil),
		(*Route_DelegateAction)(nil),
	}
}

func init() {
	proto.RegisterType((*VirtualService)(nil), "gateway.solo.io.VirtualService")
	proto.RegisterType((*VirtualHost)(nil), "gateway.solo.io.VirtualHost")
	proto.RegisterType((*Route)(nil), "gateway.solo.io.Route")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gateway/api/v1/virtual_service.proto", fileDescriptor_93fa9472926a2049)
}

var fileDescriptor_93fa9472926a2049 = []byte{
	// 677 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xc1, 0x4e, 0x1b, 0x3d,
	0x10, 0x26, 0xc9, 0x12, 0x88, 0xc3, 0x0f, 0x7f, 0xad, 0x94, 0x2e, 0xa8, 0x82, 0x28, 0x97, 0x72,
	0x68, 0x77, 0x45, 0x91, 0x10, 0x70, 0x28, 0x6a, 0xa0, 0x82, 0x0b, 0x15, 0x32, 0x52, 0x0f, 0x5c,
	0x22, 0xe3, 0x75, 0x16, 0x97, 0xcd, 0xce, 0xca, 0x76, 0xd2, 0xe6, 0x56, 0xf1, 0x34, 0x3d, 0xf5,
	0x39, 0xfa, 0x02, 0xbd, 0x72, 0xe8, 0x1b, 0xd0, 0x63, 0x4f, 0xd5, 0x7a, 0xed, 0xc0, 0x52, 0xa4,
	0xc2, 0x69, 0x3d, 0x9e, 0xf9, 0x3e, 0x8f, 0xbf, 0x6f, 0xd6, 0xe8, 0x5d, 0x2c, 0xf4, 0xf9, 0xf0,
	0x2c, 0x60, 0x30, 0x08, 0x15, 0x24, 0xf0, 0x4a, 0x40, 0x18, 0x27, 0x00, 0x61, 0x26, 0xe1, 0x23,
	0x67, 0x5a, 0x85, 0x31, 0xd5, 0xfc, 0x13, 0x1d, 0x87, 0x34, 0x13, 0xe1, 0x68, 0x3d, 0x1c, 0x09,
	0xa9, 0x87, 0x34, 0xe9, 0x29, 0x2e, 0x47, 0x82, 0xf1, 0x20, 0x93, 0xa0, 0x01, 0x2f, 0xd8, 0xaa,
	0x20, 0xe7, 0x08, 0x04, 0x2c, 0xb7, 0x62, 0x88, 0xc1, 0xe4, 0xc2, 0x7c, 0x55, 0x94, 0x2d, 0xaf,
	0xdf, 0x73, 0x9a, 0xf9, 0x5e, 0x08, 0xed, 0x0e, 0x18, 0x70, 0x4d, 0x23, 0xaa, 0xa9, 0x85, 0x84,
	0x0f, 0x80, 0x28, 0x4d, 0xf5, 0x50, 0x59, 0xc0, 0xcb, 0x07, 0x00, 0x24, 0xef, 0x3f, 0xa2, 0x23,
	0x17, 0x5b, 0xc8, 0xe6, 0xbf, 0x25, 0xcb, 0x23, 0x07, 0x56, 0x89, 0xc5, 0x6d, 0x3d, 0x0a, 0x97,
	0x49, 0xf8, 0x3c, 0xb6, 0xc8, 0x9d, 0xc7, 0x21, 0x93, 0x61, 0x2c, 0x52, 0x2b, 0x47, 0xe7, 0x47,
	0x15, 0xcd, 0x7f, 0x28, 0x3c, 0x3b, 0x29, 0x2c, 0xc3, 0xbb, 0x68, 0xce, 0xb9, 0x78, 0x0e, 0x4a,
	0xfb, 0x95, 0x76, 0x65, 0xad, 0xf9, 0xfa, 0x79, 0x70, 0xc7, 0xc3, 0xc0, 0xc2, 0x0e, 0x41, 0x69,
	0xd2, 0x1c, 0xdd, 0x04, 0x78, 0x13, 0x21, 0xa5, 0x92, 0x1e, 0x83, 0xb4, 0x2f, 0x62, 0xbf, 0x6a,
	0xe0, 0xcf, 0x82, 0xbc, 0x87, 0x09, 0xf6, 0x44, 0x25, 0x7b, 0x26, 0x4d, 0x1a, 0xca, 0x2d, 0xf1,
	0x0b, 0x34, 0x17, 0x09, 0x95, 0x25, 0x74, 0xdc, 0x4b, 0xe9, 0x80, 0xfb, 0xb5, 0x76, 0x65, 0xad,
	0xd1, 0xf5, 0xbe, 0x5c, 0x7b, 0x15, 0xd2, 0xb4, 0x99, 0xf7, 0x74, 0xc0, 0xf1, 0x01, 0xaa, 0x17,
	0x9e, 0xfa, 0x75, 0x43, 0xde, 0x0a, 0x18, 0x48, 0x7e, 0x43, 0x6e, 0x72, 0xdd, 0xa5, 0xef, 0x57,
	0xab, 0x53, 0xbf, 0xae, 0x56, 0x9f, 0x68, 0xae, 0x74, 0x24, 0xfa, 0xfd, 0x9d, 0x8e, 0x88, 0x53,
	0x90, 0xbc, 0x43, 0x2c, 0x1c, 0x6f, 0xa1, 0x59, 0x37, 0x4f, 0xfe, 0x8c, 0xa1, 0x5a, 0x2c, 0x53,
	0x1d, 0xd9, 0x6c, 0xd7, 0xcb, 0xc9, 0xc8, 0xa4, 0x7a, 0x67, 0xe5, 0xf2, 0xda, 0xf3, 0x50, 0x75,
	0xa4, 0x2e, 0xaf, 0x3d, 0x8c, 0xff, 0xbf, 0x33, 0xf6, 0xaa, 0xf3, 0xbb, 0x82, 0x9a, 0xb7, 0x04,
	0xc2, 0x18, 0x79, 0xe6, 0x4e, 0xb9, 0x98, 0x0d, 0x62, 0xd6, 0xd8, 0x47, 0x33, 0x11, 0x0c, 0xa8,
	0x48, 0x95, 0x5f, 0x6d, 0xd7, 0xd6, 0x1a, 0xc4, 0x85, 0x38, 0x40, 0x75, 0x09, 0x43, 0xcd, 0x95,
	0x5f, 0x6b, 0xd7, 0x4c, 0x57, 0x77, 0xc5, 0x27, 0x79, 0x9a, 0xd8, 0x2a, 0x4c, 0x50, 0xeb, 0xb6,
	0x65, 0x3d, 0xeb, 0xb1, 0xef, 0x99, 0x3b, 0xb5, 0xcb, 0xda, 0xdf, 0x6a, 0xeb, 0xb8, 0xa8, 0x23,
	0x78, 0xf4, 0xd7, 0x1e, 0xde, 0x46, 0x4d, 0x06, 0x52, 0xf5, 0x32, 0x48, 0x04, 0x1b, 0xfb, 0xd3,
	0x86, 0xca, 0x2f, 0x53, 0xed, 0x81, 0x54, 0xc7, 0x26, 0x4f, 0x10, 0x9b, 0xac, 0x3b, 0xdf, 0x6a,
	0x68, 0xda, 0x34, 0x88, 0x43, 0x34, 0x33, 0xa0, 0x9a, 0x9d, 0x73, 0x69, 0xc7, 0xe8, 0x69, 0x99,
	0xe0, 0xa8, 0x48, 0x12, 0x57, 0x85, 0xdf, 0xa0, 0x39, 0x73, 0xa7, 0x1e, 0x65, 0x5a, 0x40, 0x6a,
	0xa7, 0x67, 0xa9, 0x8c, 0x32, 0xdc, 0x6f, 0x4d, 0xc1, 0xe1, 0x14, 0x69, 0xca, 0x9b, 0x10, 0x1f,
	0xa0, 0x05, 0xc9, 0x23, 0x21, 0x39, 0xd3, 0x8e, 0xa2, 0xe6, 0xe6, 0xb7, 0x44, 0x61, 0x8b, 0x26,
	0x2c, 0xf3, 0xb2, 0xb4, 0x83, 0x4f, 0xd1, 0xa2, 0xa5, 0x91, 0x5c, 0x65, 0x90, 0xaa, 0x49, 0x4b,
	0x85, 0xa8, 0x9d, 0x32, 0xdf, 0xbe, 0xa9, 0x25, 0xb6, 0x74, 0xc2, 0xda, 0x8a, 0xee, 0xd9, 0xc7,
	0xfb, 0x68, 0x21, 0xe2, 0x09, 0xcf, 0x3d, 0x75, 0xa4, 0xd3, 0xf6, 0x9e, 0xa5, 0xe9, 0x23, 0x5c,
	0xc1, 0x50, 0x32, 0x4e, 0x78, 0x3f, 0xef, 0xd0, 0x61, 0x2c, 0xcb, 0x2e, 0xfa, 0xaf, 0x90, 0xca,
	0xb9, 0x5d, 0xfc, 0x0c, 0xcb, 0xf7, 0x68, 0xe5, 0x7c, 0x2e, 0xb4, 0xb5, 0x51, 0x77, 0x16, 0xd5,
	0x8b, 0xd3, 0xbb, 0xdb, 0x5f, 0x7f, 0xae, 0x54, 0x4e, 0x37, 0x1e, 0xfc, 0xd8, 0x67, 0x17, 0xb1,
	0x7d, 0x4e, 0xce, 0xea, 0xe6, 0x1d, 0xd9, 0xf8, 0x13, 0x00, 0x00, 0xff, 0xff, 0xdb, 0x9a, 0xad,
	0xcf, 0x2a, 0x06, 0x00, 0x00,
}

func (this *VirtualService) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*VirtualService)
	if !ok {
		that2, ok := that.(VirtualService)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.VirtualHost.Equal(that1.VirtualHost) {
		return false
	}
	if !this.SslConfig.Equal(that1.SslConfig) {
		return false
	}
	if this.DisplayName != that1.DisplayName {
		return false
	}
	if !this.Status.Equal(&that1.Status) {
		return false
	}
	if !this.Metadata.Equal(&that1.Metadata) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *VirtualHost) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*VirtualHost)
	if !ok {
		that2, ok := that.(VirtualHost)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if len(this.Domains) != len(that1.Domains) {
		return false
	}
	for i := range this.Domains {
		if this.Domains[i] != that1.Domains[i] {
			return false
		}
	}
	if len(this.Routes) != len(that1.Routes) {
		return false
	}
	for i := range this.Routes {
		if !this.Routes[i].Equal(that1.Routes[i]) {
			return false
		}
	}
	if !this.VirtualHostPlugins.Equal(that1.VirtualHostPlugins) {
		return false
	}
	if !this.CorsPolicy.Equal(that1.CorsPolicy) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Route) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Route)
	if !ok {
		that2, ok := that.(Route)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Matcher.Equal(that1.Matcher) {
		return false
	}
	if that1.Action == nil {
		if this.Action != nil {
			return false
		}
	} else if this.Action == nil {
		return false
	} else if !this.Action.Equal(that1.Action) {
		return false
	}
	if !this.RoutePlugins.Equal(that1.RoutePlugins) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Route_RouteAction) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Route_RouteAction)
	if !ok {
		that2, ok := that.(Route_RouteAction)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.RouteAction.Equal(that1.RouteAction) {
		return false
	}
	return true
}
func (this *Route_RedirectAction) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Route_RedirectAction)
	if !ok {
		that2, ok := that.(Route_RedirectAction)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.RedirectAction.Equal(that1.RedirectAction) {
		return false
	}
	return true
}
func (this *Route_DirectResponseAction) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Route_DirectResponseAction)
	if !ok {
		that2, ok := that.(Route_DirectResponseAction)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.DirectResponseAction.Equal(that1.DirectResponseAction) {
		return false
	}
	return true
}
func (this *Route_DelegateAction) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Route_DelegateAction)
	if !ok {
		that2, ok := that.(Route_DelegateAction)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.DelegateAction.Equal(that1.DelegateAction) {
		return false
	}
	return true
}

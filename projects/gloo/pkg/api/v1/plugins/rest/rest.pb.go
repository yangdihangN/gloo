// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/plugins/rest/rest.proto

package rest

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	transformation "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/plugins/transformation"
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

type ServiceSpec struct {
	Transformations      map[string]*transformation.TransformationTemplate `protobuf:"bytes,1,rep,name=transformations,proto3" json:"transformations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	SwaggerInfo          *ServiceSpec_SwaggerInfo                          `protobuf:"bytes,2,opt,name=swagger_info,json=swaggerInfo,proto3" json:"swagger_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                          `json:"-"`
	XXX_unrecognized     []byte                                            `json:"-"`
	XXX_sizecache        int32                                             `json:"-"`
}

func (m *ServiceSpec) Reset()         { *m = ServiceSpec{} }
func (m *ServiceSpec) String() string { return proto.CompactTextString(m) }
func (*ServiceSpec) ProtoMessage()    {}
func (*ServiceSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_10f084fc89ebe515, []int{0}
}
func (m *ServiceSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceSpec.Unmarshal(m, b)
}
func (m *ServiceSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceSpec.Marshal(b, m, deterministic)
}
func (m *ServiceSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceSpec.Merge(m, src)
}
func (m *ServiceSpec) XXX_Size() int {
	return xxx_messageInfo_ServiceSpec.Size(m)
}
func (m *ServiceSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceSpec.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceSpec proto.InternalMessageInfo

func (m *ServiceSpec) GetTransformations() map[string]*transformation.TransformationTemplate {
	if m != nil {
		return m.Transformations
	}
	return nil
}

func (m *ServiceSpec) GetSwaggerInfo() *ServiceSpec_SwaggerInfo {
	if m != nil {
		return m.SwaggerInfo
	}
	return nil
}

type ServiceSpec_SwaggerInfo struct {
	// Types that are valid to be assigned to SwaggerSpec:
	//	*ServiceSpec_SwaggerInfo_Url
	//	*ServiceSpec_SwaggerInfo_Inline
	SwaggerSpec          isServiceSpec_SwaggerInfo_SwaggerSpec `protobuf_oneof:"swagger_spec"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *ServiceSpec_SwaggerInfo) Reset()         { *m = ServiceSpec_SwaggerInfo{} }
func (m *ServiceSpec_SwaggerInfo) String() string { return proto.CompactTextString(m) }
func (*ServiceSpec_SwaggerInfo) ProtoMessage()    {}
func (*ServiceSpec_SwaggerInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_10f084fc89ebe515, []int{0, 1}
}
func (m *ServiceSpec_SwaggerInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceSpec_SwaggerInfo.Unmarshal(m, b)
}
func (m *ServiceSpec_SwaggerInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceSpec_SwaggerInfo.Marshal(b, m, deterministic)
}
func (m *ServiceSpec_SwaggerInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceSpec_SwaggerInfo.Merge(m, src)
}
func (m *ServiceSpec_SwaggerInfo) XXX_Size() int {
	return xxx_messageInfo_ServiceSpec_SwaggerInfo.Size(m)
}
func (m *ServiceSpec_SwaggerInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceSpec_SwaggerInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceSpec_SwaggerInfo proto.InternalMessageInfo

type isServiceSpec_SwaggerInfo_SwaggerSpec interface {
	isServiceSpec_SwaggerInfo_SwaggerSpec()
	Equal(interface{}) bool
}

type ServiceSpec_SwaggerInfo_Url struct {
	Url string `protobuf:"bytes,1,opt,name=url,proto3,oneof" json:"url,omitempty"`
}
type ServiceSpec_SwaggerInfo_Inline struct {
	Inline string `protobuf:"bytes,2,opt,name=inline,proto3,oneof" json:"inline,omitempty"`
}

func (*ServiceSpec_SwaggerInfo_Url) isServiceSpec_SwaggerInfo_SwaggerSpec()    {}
func (*ServiceSpec_SwaggerInfo_Inline) isServiceSpec_SwaggerInfo_SwaggerSpec() {}

func (m *ServiceSpec_SwaggerInfo) GetSwaggerSpec() isServiceSpec_SwaggerInfo_SwaggerSpec {
	if m != nil {
		return m.SwaggerSpec
	}
	return nil
}

func (m *ServiceSpec_SwaggerInfo) GetUrl() string {
	if x, ok := m.GetSwaggerSpec().(*ServiceSpec_SwaggerInfo_Url); ok {
		return x.Url
	}
	return ""
}

func (m *ServiceSpec_SwaggerInfo) GetInline() string {
	if x, ok := m.GetSwaggerSpec().(*ServiceSpec_SwaggerInfo_Inline); ok {
		return x.Inline
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ServiceSpec_SwaggerInfo) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ServiceSpec_SwaggerInfo_Url)(nil),
		(*ServiceSpec_SwaggerInfo_Inline)(nil),
	}
}

// This is only for upstream with REST service spec
type DestinationSpec struct {
	FunctionName           string                                 `protobuf:"bytes,1,opt,name=function_name,json=functionName,proto3" json:"function_name,omitempty"`
	Parameters             *transformation.Parameters             `protobuf:"bytes,2,opt,name=parameters,proto3" json:"parameters,omitempty"`
	ResponseTransformation *transformation.TransformationTemplate `protobuf:"bytes,3,opt,name=response_transformation,json=responseTransformation,proto3" json:"response_transformation,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}                               `json:"-"`
	XXX_unrecognized       []byte                                 `json:"-"`
	XXX_sizecache          int32                                  `json:"-"`
}

func (m *DestinationSpec) Reset()         { *m = DestinationSpec{} }
func (m *DestinationSpec) String() string { return proto.CompactTextString(m) }
func (*DestinationSpec) ProtoMessage()    {}
func (*DestinationSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_10f084fc89ebe515, []int{1}
}
func (m *DestinationSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DestinationSpec.Unmarshal(m, b)
}
func (m *DestinationSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DestinationSpec.Marshal(b, m, deterministic)
}
func (m *DestinationSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DestinationSpec.Merge(m, src)
}
func (m *DestinationSpec) XXX_Size() int {
	return xxx_messageInfo_DestinationSpec.Size(m)
}
func (m *DestinationSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_DestinationSpec.DiscardUnknown(m)
}

var xxx_messageInfo_DestinationSpec proto.InternalMessageInfo

func (m *DestinationSpec) GetFunctionName() string {
	if m != nil {
		return m.FunctionName
	}
	return ""
}

func (m *DestinationSpec) GetParameters() *transformation.Parameters {
	if m != nil {
		return m.Parameters
	}
	return nil
}

func (m *DestinationSpec) GetResponseTransformation() *transformation.TransformationTemplate {
	if m != nil {
		return m.ResponseTransformation
	}
	return nil
}

func init() {
	proto.RegisterType((*ServiceSpec)(nil), "rest.plugins.gloo.solo.io.ServiceSpec")
	proto.RegisterMapType((map[string]*transformation.TransformationTemplate)(nil), "rest.plugins.gloo.solo.io.ServiceSpec.TransformationsEntry")
	proto.RegisterType((*ServiceSpec_SwaggerInfo)(nil), "rest.plugins.gloo.solo.io.ServiceSpec.SwaggerInfo")
	proto.RegisterType((*DestinationSpec)(nil), "rest.plugins.gloo.solo.io.DestinationSpec")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/v1/plugins/rest/rest.proto", fileDescriptor_10f084fc89ebe515)
}

var fileDescriptor_10f084fc89ebe515 = []byte{
	// 440 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0xd1, 0x8a, 0xd4, 0x30,
	0x14, 0x86, 0xed, 0x16, 0x17, 0x4c, 0x57, 0x57, 0xc2, 0xa2, 0x75, 0x2e, 0x64, 0x58, 0x6f, 0xe6,
	0xc6, 0x54, 0xc7, 0x1b, 0x51, 0xbc, 0x59, 0x57, 0x51, 0x04, 0x95, 0xce, 0x08, 0xe2, 0xcd, 0x90,
	0x2d, 0xa7, 0xd9, 0xb8, 0x69, 0x4e, 0x48, 0xd2, 0xca, 0xbc, 0x84, 0xcf, 0xe1, 0x73, 0xf9, 0x08,
	0x3e, 0x81, 0xb4, 0xe9, 0x3a, 0x33, 0x65, 0x84, 0x65, 0xd8, 0x9b, 0x92, 0x3f, 0xe9, 0xff, 0xfd,
	0x3d, 0x27, 0xa7, 0xe4, 0x54, 0x48, 0x7f, 0x5e, 0x9f, 0xb1, 0x02, 0xab, 0xcc, 0xa1, 0xc2, 0xc7,
	0x12, 0x33, 0xa1, 0x10, 0x33, 0x63, 0xf1, 0x3b, 0x14, 0xde, 0x05, 0xc5, 0x8d, 0xcc, 0x9a, 0xa7,
	0x99, 0x51, 0xb5, 0x90, 0xda, 0x65, 0x16, 0x9c, 0xef, 0x1e, 0xcc, 0x58, 0xf4, 0x48, 0x1f, 0x84,
	0x75, 0x38, 0x65, 0xad, 0x83, 0xb5, 0x30, 0x26, 0x71, 0x74, 0x24, 0x50, 0x60, 0xf7, 0x56, 0xd6,
	0xae, 0x82, 0x61, 0xf4, 0x75, 0xa7, 0x58, 0x6f, 0xb9, 0x76, 0x25, 0xda, 0x8a, 0x7b, 0x89, 0x7a,
	0x20, 0x7b, 0xf2, 0xfc, 0x3a, 0xc8, 0x86, 0x5b, 0x5e, 0x81, 0x07, 0xeb, 0x02, 0xf5, 0xf8, 0x67,
	0x4c, 0x92, 0x19, 0xd8, 0x46, 0x16, 0x30, 0x33, 0x50, 0x50, 0x20, 0x87, 0x9b, 0x16, 0x97, 0x46,
	0xe3, 0x78, 0x92, 0x4c, 0x5f, 0xb2, 0xff, 0xb6, 0x82, 0xad, 0x01, 0xd8, 0x7c, 0xd3, 0xfd, 0x46,
	0x7b, 0xbb, 0xcc, 0x87, 0x4c, 0xfa, 0x85, 0x1c, 0xb8, 0x1f, 0x5c, 0x08, 0xb0, 0x0b, 0xa9, 0x4b,
	0x4c, 0xf7, 0xc6, 0xd1, 0x24, 0x99, 0x4e, 0xaf, 0x98, 0x31, 0x0b, 0xd6, 0xf7, 0xba, 0xc4, 0x3c,
	0x71, 0x2b, 0x31, 0xf2, 0xe4, 0x68, 0x5b, 0x3e, 0xbd, 0x4b, 0xe2, 0x0b, 0x58, 0xa6, 0xd1, 0x38,
	0x9a, 0xdc, 0xca, 0xdb, 0x25, 0x7d, 0x4b, 0x6e, 0x36, 0x5c, 0xd5, 0xd0, 0x27, 0x3f, 0x61, 0xa0,
	0x1b, 0x5c, 0x32, 0x6e, 0x24, 0x6b, 0xa6, 0xac, 0x94, 0xca, 0x83, 0x65, 0xe7, 0xde, 0x9b, 0x41,
	0x41, 0x73, 0xa8, 0x8c, 0xe2, 0x1e, 0xf2, 0x60, 0x7f, 0xb1, 0xf7, 0x3c, 0x1a, 0x7d, 0x20, 0xc9,
	0xda, 0x17, 0x51, 0x4a, 0xe2, 0xda, 0xaa, 0x10, 0xf6, 0xee, 0x46, 0xde, 0x0a, 0x9a, 0x92, 0x7d,
	0xa9, 0x95, 0xd4, 0x21, 0xaf, 0xdd, 0xee, 0xf5, 0xc9, 0x9d, 0x55, 0x27, 0x9c, 0x81, 0xe2, 0xf8,
	0x4f, 0x44, 0x0e, 0x4f, 0xc1, 0x79, 0xa9, 0xbb, 0xbc, 0xee, 0x52, 0x1e, 0x91, 0xdb, 0x65, 0xad,
	0x8b, 0x56, 0x2f, 0x34, 0xaf, 0xa0, 0x2f, 0xe4, 0xe0, 0x72, 0xf3, 0x23, 0xaf, 0x80, 0x7e, 0x22,
	0x64, 0x75, 0xbb, 0x7d, 0x59, 0x19, 0x1b, 0x8e, 0xd2, 0xb6, 0xd6, 0x7e, 0xfe, 0x67, 0xcb, 0xd7,
	0x10, 0x54, 0x92, 0xfb, 0x16, 0x9c, 0x41, 0xed, 0x60, 0xb1, 0x89, 0x49, 0xe3, 0x1d, 0x9b, 0x76,
	0xef, 0x12, 0xb8, 0x79, 0x7e, 0xf2, 0xfa, 0xd7, 0xef, 0x87, 0xd1, 0xb7, 0x57, 0x57, 0x9b, 0x70,
	0x73, 0x21, 0xb6, 0xfd, 0xb6, 0x67, 0xfb, 0xdd, 0x44, 0x3f, 0xfb, 0x1b, 0x00, 0x00, 0xff, 0xff,
	0xab, 0x66, 0xe3, 0x5e, 0xfa, 0x03, 0x00, 0x00,
}

func (this *ServiceSpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ServiceSpec)
	if !ok {
		that2, ok := that.(ServiceSpec)
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
	if len(this.Transformations) != len(that1.Transformations) {
		return false
	}
	for i := range this.Transformations {
		if !this.Transformations[i].Equal(that1.Transformations[i]) {
			return false
		}
	}
	if !this.SwaggerInfo.Equal(that1.SwaggerInfo) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *ServiceSpec_SwaggerInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ServiceSpec_SwaggerInfo)
	if !ok {
		that2, ok := that.(ServiceSpec_SwaggerInfo)
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
	if that1.SwaggerSpec == nil {
		if this.SwaggerSpec != nil {
			return false
		}
	} else if this.SwaggerSpec == nil {
		return false
	} else if !this.SwaggerSpec.Equal(that1.SwaggerSpec) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *ServiceSpec_SwaggerInfo_Url) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ServiceSpec_SwaggerInfo_Url)
	if !ok {
		that2, ok := that.(ServiceSpec_SwaggerInfo_Url)
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
	if this.Url != that1.Url {
		return false
	}
	return true
}
func (this *ServiceSpec_SwaggerInfo_Inline) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ServiceSpec_SwaggerInfo_Inline)
	if !ok {
		that2, ok := that.(ServiceSpec_SwaggerInfo_Inline)
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
	if this.Inline != that1.Inline {
		return false
	}
	return true
}
func (this *DestinationSpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*DestinationSpec)
	if !ok {
		that2, ok := that.(DestinationSpec)
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
	if this.FunctionName != that1.FunctionName {
		return false
	}
	if !this.Parameters.Equal(that1.Parameters) {
		return false
	}
	if !this.ResponseTransformation.Equal(that1.ResponseTransformation) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}

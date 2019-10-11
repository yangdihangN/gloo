// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/artifact.proto

package v1

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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
//@solo-kit:resource.short_name=art
//@solo-kit:resource.plural_name=artifacts
//
//Gloo Artifacts are used by Gloo to store small bits of binary or file data.
//
//Certain plugins such as the gRPC plugin read and write artifacts to one of Gloo's configured
//storage layer.
//
//Artifacts can be backed by files on disk, Kubernetes ConfigMaps, and Consul Key/Value pairs.
//
//Supported artifact backends can be selected in Gloo's boostrap options.
type Artifact struct {
	// Raw data data being stored
	Data map[string]string `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Metadata contains the object metadata for this resource
	Metadata             core.Metadata `protobuf:"bytes,7,opt,name=metadata,proto3" json:"metadata"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Artifact) Reset()         { *m = Artifact{} }
func (m *Artifact) String() string { return proto.CompactTextString(m) }
func (*Artifact) ProtoMessage()    {}
func (*Artifact) Descriptor() ([]byte, []int) {
	return fileDescriptor_c52f0e475c0b5a35, []int{0}
}
func (m *Artifact) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Artifact.Unmarshal(m, b)
}
func (m *Artifact) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Artifact.Marshal(b, m, deterministic)
}
func (m *Artifact) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Artifact.Merge(m, src)
}
func (m *Artifact) XXX_Size() int {
	return xxx_messageInfo_Artifact.Size(m)
}
func (m *Artifact) XXX_DiscardUnknown() {
	xxx_messageInfo_Artifact.DiscardUnknown(m)
}

var xxx_messageInfo_Artifact proto.InternalMessageInfo

func (m *Artifact) GetData() map[string]string {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Artifact) GetMetadata() core.Metadata {
	if m != nil {
		return m.Metadata
	}
	return core.Metadata{}
}

func init() {
	proto.RegisterType((*Artifact)(nil), "gloo.solo.io.Artifact")
	proto.RegisterMapType((map[string]string)(nil), "gloo.solo.io.Artifact.DataEntry")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/v1/artifact.proto", fileDescriptor_c52f0e475c0b5a35)
}

var fileDescriptor_c52f0e475c0b5a35 = []byte{
	// 293 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x4e, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x2f, 0xce, 0xcf, 0xc9, 0xd7, 0xcd, 0xcc, 0xd7, 0x4f,
	0xcf, 0xc9, 0xcf, 0xd7, 0x2f, 0x28, 0xca, 0xcf, 0x4a, 0x4d, 0x2e, 0x29, 0x86, 0xf0, 0x12, 0x0b,
	0x32, 0xf5, 0xcb, 0x0c, 0xf5, 0x13, 0x8b, 0x4a, 0x32, 0xd3, 0x12, 0x93, 0x4b, 0xf4, 0x0a, 0x8a,
	0xf2, 0x4b, 0xf2, 0x85, 0x78, 0x40, 0x72, 0x7a, 0x20, 0x6d, 0x7a, 0x99, 0xf9, 0x52, 0x22, 0xe9,
	0xf9, 0xe9, 0xf9, 0x60, 0x09, 0x7d, 0x10, 0x0b, 0xa2, 0x46, 0xca, 0x10, 0x8b, 0x05, 0x60, 0x3a,
	0x3b, 0xb3, 0x04, 0x66, 0x6c, 0x6e, 0x6a, 0x49, 0x62, 0x4a, 0x62, 0x49, 0x22, 0x09, 0x5a, 0x60,
	0x7c, 0x88, 0x16, 0xa5, 0xcb, 0x8c, 0x5c, 0x1c, 0x8e, 0x50, 0xc7, 0x09, 0x99, 0x70, 0xb1, 0x80,
	0x4c, 0x93, 0x60, 0x54, 0x60, 0xd6, 0xe0, 0x36, 0x52, 0xd0, 0x43, 0x76, 0xa5, 0x1e, 0x4c, 0x95,
	0x9e, 0x4b, 0x62, 0x49, 0xa2, 0x6b, 0x5e, 0x49, 0x51, 0x65, 0x10, 0x58, 0xb5, 0x90, 0x05, 0x17,
	0x07, 0xcc, 0x1d, 0x12, 0xec, 0x0a, 0x8c, 0x1a, 0xdc, 0x46, 0x62, 0x7a, 0xc9, 0xf9, 0x45, 0xa9,
	0x70, 0x9d, 0xbe, 0x50, 0x59, 0x27, 0x96, 0x13, 0xf7, 0xe4, 0x19, 0x82, 0xe0, 0xaa, 0xa5, 0xcc,
	0xb9, 0x38, 0xe1, 0x86, 0x09, 0x09, 0x70, 0x31, 0x67, 0xa7, 0x56, 0x4a, 0x30, 0x2a, 0x30, 0x6a,
	0x70, 0x06, 0x81, 0x98, 0x42, 0x22, 0x5c, 0xac, 0x65, 0x89, 0x39, 0xa5, 0xa9, 0x12, 0x4c, 0x60,
	0x31, 0x08, 0xc7, 0x8a, 0xc9, 0x82, 0xd1, 0x4a, 0xae, 0xe9, 0x23, 0x0b, 0x2b, 0x17, 0x73, 0x62,
	0x51, 0x49, 0xd3, 0x47, 0x16, 0x6e, 0x21, 0x4e, 0x58, 0xf0, 0x16, 0x37, 0x7d, 0x64, 0x61, 0xd2,
	0x60, 0x74, 0x32, 0x5b, 0xf1, 0x48, 0x8e, 0x31, 0xca, 0x80, 0xb8, 0x28, 0x2a, 0xc8, 0x4e, 0x87,
	0x06, 0x4e, 0x12, 0x1b, 0x38, 0x50, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x48, 0x20, 0x2f,
	0x9b, 0xdd, 0x01, 0x00, 0x00,
}

func (this *Artifact) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Artifact)
	if !ok {
		that2, ok := that.(Artifact)
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
	if len(this.Data) != len(that1.Data) {
		return false
	}
	for i := range this.Data {
		if this.Data[i] != that1.Data[i] {
			return false
		}
	}
	if !this.Metadata.Equal(&that1.Metadata) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/plugins/als/als.proto

package als

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	_ "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
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

// Contains various settings for Envoy's access logging service.
// See here for more information: https://www.envoyproxy.io/docs/envoy/latest/api-v2/config/filter/accesslog/v2/accesslog.proto#envoy-api-msg-config-filter-accesslog-v2-accesslog
type AccessLoggingService struct {
	AccessLog            []*AccessLog `protobuf:"bytes,1,rep,name=access_log,json=accessLog,proto3" json:"access_log,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *AccessLoggingService) Reset()         { *m = AccessLoggingService{} }
func (m *AccessLoggingService) String() string { return proto.CompactTextString(m) }
func (*AccessLoggingService) ProtoMessage()    {}
func (*AccessLoggingService) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd8d2602efe636cc, []int{0}
}
func (m *AccessLoggingService) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccessLoggingService.Unmarshal(m, b)
}
func (m *AccessLoggingService) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccessLoggingService.Marshal(b, m, deterministic)
}
func (m *AccessLoggingService) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessLoggingService.Merge(m, src)
}
func (m *AccessLoggingService) XXX_Size() int {
	return xxx_messageInfo_AccessLoggingService.Size(m)
}
func (m *AccessLoggingService) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessLoggingService.DiscardUnknown(m)
}

var xxx_messageInfo_AccessLoggingService proto.InternalMessageInfo

func (m *AccessLoggingService) GetAccessLog() []*AccessLog {
	if m != nil {
		return m.AccessLog
	}
	return nil
}

type AccessLog struct {
	// type of Access Logging service to implement
	//
	// Types that are valid to be assigned to OutputDestination:
	//	*AccessLog_FileSink
	//	*AccessLog_GrpcService
	OutputDestination    isAccessLog_OutputDestination `protobuf_oneof:"OutputDestination"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *AccessLog) Reset()         { *m = AccessLog{} }
func (m *AccessLog) String() string { return proto.CompactTextString(m) }
func (*AccessLog) ProtoMessage()    {}
func (*AccessLog) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd8d2602efe636cc, []int{1}
}
func (m *AccessLog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccessLog.Unmarshal(m, b)
}
func (m *AccessLog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccessLog.Marshal(b, m, deterministic)
}
func (m *AccessLog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessLog.Merge(m, src)
}
func (m *AccessLog) XXX_Size() int {
	return xxx_messageInfo_AccessLog.Size(m)
}
func (m *AccessLog) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessLog.DiscardUnknown(m)
}

var xxx_messageInfo_AccessLog proto.InternalMessageInfo

type isAccessLog_OutputDestination interface {
	isAccessLog_OutputDestination()
	Equal(interface{}) bool
}

type AccessLog_FileSink struct {
	FileSink *FileSink `protobuf:"bytes,2,opt,name=file_sink,json=fileSink,proto3,oneof" json:"file_sink,omitempty"`
}
type AccessLog_GrpcService struct {
	GrpcService *GrpcService `protobuf:"bytes,3,opt,name=grpc_service,json=grpcService,proto3,oneof" json:"grpc_service,omitempty"`
}

func (*AccessLog_FileSink) isAccessLog_OutputDestination()    {}
func (*AccessLog_GrpcService) isAccessLog_OutputDestination() {}

func (m *AccessLog) GetOutputDestination() isAccessLog_OutputDestination {
	if m != nil {
		return m.OutputDestination
	}
	return nil
}

func (m *AccessLog) GetFileSink() *FileSink {
	if x, ok := m.GetOutputDestination().(*AccessLog_FileSink); ok {
		return x.FileSink
	}
	return nil
}

func (m *AccessLog) GetGrpcService() *GrpcService {
	if x, ok := m.GetOutputDestination().(*AccessLog_GrpcService); ok {
		return x.GrpcService
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*AccessLog) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*AccessLog_FileSink)(nil),
		(*AccessLog_GrpcService)(nil),
	}
}

type FileSink struct {
	// the file path to which the file access logging service will sink
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	// the format which the logs should be outputted by
	//
	// Types that are valid to be assigned to OutputFormat:
	//	*FileSink_StringFormat
	//	*FileSink_JsonFormat
	OutputFormat         isFileSink_OutputFormat `protobuf_oneof:"output_format"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *FileSink) Reset()         { *m = FileSink{} }
func (m *FileSink) String() string { return proto.CompactTextString(m) }
func (*FileSink) ProtoMessage()    {}
func (*FileSink) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd8d2602efe636cc, []int{2}
}
func (m *FileSink) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileSink.Unmarshal(m, b)
}
func (m *FileSink) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileSink.Marshal(b, m, deterministic)
}
func (m *FileSink) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileSink.Merge(m, src)
}
func (m *FileSink) XXX_Size() int {
	return xxx_messageInfo_FileSink.Size(m)
}
func (m *FileSink) XXX_DiscardUnknown() {
	xxx_messageInfo_FileSink.DiscardUnknown(m)
}

var xxx_messageInfo_FileSink proto.InternalMessageInfo

type isFileSink_OutputFormat interface {
	isFileSink_OutputFormat()
	Equal(interface{}) bool
}

type FileSink_StringFormat struct {
	StringFormat string `protobuf:"bytes,2,opt,name=string_format,json=stringFormat,proto3,oneof" json:"string_format,omitempty"`
}
type FileSink_JsonFormat struct {
	JsonFormat *types.Struct `protobuf:"bytes,3,opt,name=json_format,json=jsonFormat,proto3,oneof" json:"json_format,omitempty"`
}

func (*FileSink_StringFormat) isFileSink_OutputFormat() {}
func (*FileSink_JsonFormat) isFileSink_OutputFormat()   {}

func (m *FileSink) GetOutputFormat() isFileSink_OutputFormat {
	if m != nil {
		return m.OutputFormat
	}
	return nil
}

func (m *FileSink) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *FileSink) GetStringFormat() string {
	if x, ok := m.GetOutputFormat().(*FileSink_StringFormat); ok {
		return x.StringFormat
	}
	return ""
}

func (m *FileSink) GetJsonFormat() *types.Struct {
	if x, ok := m.GetOutputFormat().(*FileSink_JsonFormat); ok {
		return x.JsonFormat
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*FileSink) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*FileSink_StringFormat)(nil),
		(*FileSink_JsonFormat)(nil),
	}
}

type GrpcService struct {
	// name of log stream
	LogName string `protobuf:"bytes,1,opt,name=log_name,json=logName,proto3" json:"log_name,omitempty"`
	// The static cluster defined in bootstrap config to route to
	//
	// Types that are valid to be assigned to ServiceRef:
	//	*GrpcService_StaticClusterName
	ServiceRef                      isGrpcService_ServiceRef `protobuf_oneof:"service_ref"`
	AdditionalRequestHeadersToLog   []string                 `protobuf:"bytes,4,rep,name=additional_request_headers_to_log,json=additionalRequestHeadersToLog,proto3" json:"additional_request_headers_to_log,omitempty"`
	AdditionalResponseHeadersToLog  []string                 `protobuf:"bytes,5,rep,name=additional_response_headers_to_log,json=additionalResponseHeadersToLog,proto3" json:"additional_response_headers_to_log,omitempty"`
	AdditionalResponseTrailersToLog []string                 `protobuf:"bytes,6,rep,name=additional_response_trailers_to_log,json=additionalResponseTrailersToLog,proto3" json:"additional_response_trailers_to_log,omitempty"`
	XXX_NoUnkeyedLiteral            struct{}                 `json:"-"`
	XXX_unrecognized                []byte                   `json:"-"`
	XXX_sizecache                   int32                    `json:"-"`
}

func (m *GrpcService) Reset()         { *m = GrpcService{} }
func (m *GrpcService) String() string { return proto.CompactTextString(m) }
func (*GrpcService) ProtoMessage()    {}
func (*GrpcService) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd8d2602efe636cc, []int{3}
}
func (m *GrpcService) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GrpcService.Unmarshal(m, b)
}
func (m *GrpcService) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GrpcService.Marshal(b, m, deterministic)
}
func (m *GrpcService) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GrpcService.Merge(m, src)
}
func (m *GrpcService) XXX_Size() int {
	return xxx_messageInfo_GrpcService.Size(m)
}
func (m *GrpcService) XXX_DiscardUnknown() {
	xxx_messageInfo_GrpcService.DiscardUnknown(m)
}

var xxx_messageInfo_GrpcService proto.InternalMessageInfo

type isGrpcService_ServiceRef interface {
	isGrpcService_ServiceRef()
	Equal(interface{}) bool
}

type GrpcService_StaticClusterName struct {
	StaticClusterName string `protobuf:"bytes,2,opt,name=static_cluster_name,json=staticClusterName,proto3,oneof" json:"static_cluster_name,omitempty"`
}

func (*GrpcService_StaticClusterName) isGrpcService_ServiceRef() {}

func (m *GrpcService) GetServiceRef() isGrpcService_ServiceRef {
	if m != nil {
		return m.ServiceRef
	}
	return nil
}

func (m *GrpcService) GetLogName() string {
	if m != nil {
		return m.LogName
	}
	return ""
}

func (m *GrpcService) GetStaticClusterName() string {
	if x, ok := m.GetServiceRef().(*GrpcService_StaticClusterName); ok {
		return x.StaticClusterName
	}
	return ""
}

func (m *GrpcService) GetAdditionalRequestHeadersToLog() []string {
	if m != nil {
		return m.AdditionalRequestHeadersToLog
	}
	return nil
}

func (m *GrpcService) GetAdditionalResponseHeadersToLog() []string {
	if m != nil {
		return m.AdditionalResponseHeadersToLog
	}
	return nil
}

func (m *GrpcService) GetAdditionalResponseTrailersToLog() []string {
	if m != nil {
		return m.AdditionalResponseTrailersToLog
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*GrpcService) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*GrpcService_StaticClusterName)(nil),
	}
}

func init() {
	proto.RegisterType((*AccessLoggingService)(nil), "als.plugins.gloo.solo.io.AccessLoggingService")
	proto.RegisterType((*AccessLog)(nil), "als.plugins.gloo.solo.io.AccessLog")
	proto.RegisterType((*FileSink)(nil), "als.plugins.gloo.solo.io.FileSink")
	proto.RegisterType((*GrpcService)(nil), "als.plugins.gloo.solo.io.GrpcService")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/v1/plugins/als/als.proto", fileDescriptor_dd8d2602efe636cc)
}

var fileDescriptor_dd8d2602efe636cc = []byte{
	// 536 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xd1, 0x6e, 0xd3, 0x3c,
	0x14, 0x80, 0x9b, 0x6d, 0xff, 0xfe, 0xe5, 0x64, 0x15, 0x5a, 0x36, 0x89, 0x32, 0xc1, 0x28, 0x99,
	0x26, 0xf5, 0x02, 0x12, 0x18, 0x77, 0x88, 0x9b, 0x05, 0x34, 0xa2, 0x69, 0x02, 0x29, 0xdd, 0xd5,
	0x6e, 0x22, 0x37, 0x75, 0x5c, 0xaf, 0x6e, 0x4e, 0xb0, 0x9d, 0x3d, 0x08, 0x4f, 0xc1, 0x1d, 0x4f,
	0xc2, 0x4b, 0xf0, 0x24, 0x28, 0x76, 0xb2, 0x75, 0xd0, 0x4a, 0x5c, 0x54, 0x3d, 0xf6, 0xf9, 0xfc,
	0xf9, 0x1c, 0x9d, 0x18, 0x62, 0xc6, 0xf5, 0xac, 0x9e, 0x84, 0x39, 0x2e, 0x22, 0x85, 0x02, 0x5f,
	0x71, 0x8c, 0x98, 0x40, 0x8c, 0x2a, 0x89, 0x37, 0x34, 0xd7, 0xca, 0xae, 0x48, 0xc5, 0xa3, 0xdb,
	0x37, 0x51, 0x25, 0x6a, 0xc6, 0x4b, 0x15, 0x11, 0x61, 0x7e, 0x61, 0x25, 0x51, 0xa3, 0x3f, 0x30,
	0xa1, 0x4d, 0x85, 0x0d, 0x1e, 0x36, 0xa6, 0x90, 0xe3, 0xe1, 0x01, 0x43, 0x86, 0x06, 0x8a, 0x9a,
	0xc8, 0xf2, 0x87, 0x2f, 0x57, 0xdc, 0x69, 0xfe, 0xe7, 0x5c, 0x77, 0x37, 0x49, 0x5a, 0xb4, 0xf4,
	0x53, 0x86, 0xc8, 0x04, 0x8d, 0xcc, 0x6a, 0x52, 0x17, 0x91, 0xd2, 0xb2, 0xce, 0xb5, 0xcd, 0x06,
	0xd7, 0x70, 0x70, 0x96, 0xe7, 0x54, 0xa9, 0x4b, 0x64, 0x8c, 0x97, 0x6c, 0x4c, 0xe5, 0x2d, 0xcf,
	0xa9, 0x1f, 0x03, 0x10, 0xb3, 0x9f, 0x09, 0x64, 0x03, 0x67, 0xb8, 0x39, 0xf2, 0x4e, 0x8f, 0xc3,
	0x75, 0x85, 0x86, 0x77, 0x8e, 0xd4, 0x25, 0x5d, 0x18, 0xfc, 0x70, 0xc0, 0xbd, 0x4b, 0xf8, 0x67,
	0xe0, 0x16, 0x5c, 0xd0, 0x4c, 0xf1, 0x72, 0x3e, 0xd8, 0x18, 0x3a, 0x23, 0xef, 0x34, 0x58, 0x2f,
	0x3c, 0xe7, 0x82, 0x8e, 0x79, 0x39, 0x4f, 0x7a, 0xe9, 0x4e, 0xd1, 0xc6, 0xfe, 0x05, 0xec, 0x32,
	0x59, 0xe5, 0x99, 0xb2, 0x45, 0x0e, 0x36, 0x8d, 0xe5, 0x64, 0xbd, 0xe5, 0x93, 0xac, 0xf2, 0xb6,
	0xa3, 0xa4, 0x97, 0x7a, 0xec, 0x7e, 0x19, 0xef, 0xc3, 0xde, 0x97, 0x5a, 0x57, 0xb5, 0xfe, 0x48,
	0x95, 0xe6, 0x25, 0xd1, 0x1c, 0xcb, 0xe0, 0x9b, 0x03, 0x3b, 0xdd, 0xcd, 0xbe, 0x0f, 0x5b, 0x15,
	0xd1, 0xb3, 0x81, 0x33, 0x74, 0x46, 0x6e, 0x6a, 0x62, 0xff, 0x04, 0xfa, 0x4a, 0x4b, 0x5e, 0xb2,
	0xac, 0x40, 0xb9, 0x20, 0xda, 0x34, 0xe2, 0x26, 0xbd, 0x74, 0xd7, 0x6e, 0x9f, 0x9b, 0x5d, 0xff,
	0x1d, 0x78, 0x37, 0x0a, 0xcb, 0x0e, 0xb2, 0x75, 0x3e, 0x0e, 0xed, 0x24, 0xc2, 0x6e, 0x12, 0xe1,
	0xd8, 0x4c, 0x22, 0xe9, 0xa5, 0xd0, 0xd0, 0xf6, 0x6c, 0xfc, 0x08, 0xfa, 0x68, 0x0a, 0x6b, 0x4f,
	0x07, 0x3f, 0x37, 0xc0, 0x5b, 0x6a, 0xc4, 0x7f, 0x02, 0x3b, 0x02, 0x59, 0x56, 0x92, 0x05, 0x6d,
	0x6b, 0xfb, 0x5f, 0x20, 0xfb, 0x4c, 0x16, 0xd4, 0x7f, 0x0d, 0xfb, 0x4a, 0x13, 0xcd, 0xf3, 0x2c,
	0x17, 0xb5, 0xd2, 0x54, 0x5a, 0xaa, 0x2b, 0x72, 0xcf, 0x26, 0x3f, 0xd8, 0x9c, 0x39, 0x91, 0xc0,
	0x0b, 0x32, 0x9d, 0xf2, 0xa6, 0x7b, 0x22, 0x32, 0x49, 0xbf, 0xd6, 0x54, 0xe9, 0x6c, 0x46, 0xc9,
	0x94, 0x4a, 0x95, 0x69, 0x34, 0xe3, 0xdf, 0x1a, 0x6e, 0x8e, 0xdc, 0xf4, 0xd9, 0x3d, 0x98, 0x5a,
	0x2e, 0xb1, 0xd8, 0x15, 0x36, 0xf3, 0xbd, 0x80, 0xe0, 0x81, 0x49, 0x55, 0x58, 0x2a, 0xfa, 0xa7,
	0xea, 0x3f, 0xa3, 0x3a, 0x5a, 0x56, 0x59, 0xf0, 0x81, 0xeb, 0x12, 0x8e, 0x57, 0xb9, 0xb4, 0x24,
	0x5c, 0x2c, 0xc9, 0xb6, 0x8d, 0xec, 0xf9, 0xdf, 0xb2, 0xab, 0x16, 0x34, 0xb6, 0xb8, 0x0f, 0x5e,
	0xfb, 0xc5, 0x64, 0x92, 0x16, 0x71, 0xfc, 0xfd, 0xd7, 0x91, 0x73, 0xfd, 0xfe, 0xdf, 0x1e, 0x6e,
	0x35, 0x67, 0x2b, 0x1e, 0xef, 0x64, 0xdb, 0xcc, 0xf0, 0xed, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xb1, 0x43, 0xa9, 0x88, 0xff, 0x03, 0x00, 0x00,
}

func (this *AccessLoggingService) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*AccessLoggingService)
	if !ok {
		that2, ok := that.(AccessLoggingService)
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
	if len(this.AccessLog) != len(that1.AccessLog) {
		return false
	}
	for i := range this.AccessLog {
		if !this.AccessLog[i].Equal(that1.AccessLog[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *AccessLog) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*AccessLog)
	if !ok {
		that2, ok := that.(AccessLog)
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
	if that1.OutputDestination == nil {
		if this.OutputDestination != nil {
			return false
		}
	} else if this.OutputDestination == nil {
		return false
	} else if !this.OutputDestination.Equal(that1.OutputDestination) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *AccessLog_FileSink) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*AccessLog_FileSink)
	if !ok {
		that2, ok := that.(AccessLog_FileSink)
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
	if !this.FileSink.Equal(that1.FileSink) {
		return false
	}
	return true
}
func (this *AccessLog_GrpcService) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*AccessLog_GrpcService)
	if !ok {
		that2, ok := that.(AccessLog_GrpcService)
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
	if !this.GrpcService.Equal(that1.GrpcService) {
		return false
	}
	return true
}
func (this *FileSink) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*FileSink)
	if !ok {
		that2, ok := that.(FileSink)
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
	if this.Path != that1.Path {
		return false
	}
	if that1.OutputFormat == nil {
		if this.OutputFormat != nil {
			return false
		}
	} else if this.OutputFormat == nil {
		return false
	} else if !this.OutputFormat.Equal(that1.OutputFormat) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *FileSink_StringFormat) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*FileSink_StringFormat)
	if !ok {
		that2, ok := that.(FileSink_StringFormat)
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
	if this.StringFormat != that1.StringFormat {
		return false
	}
	return true
}
func (this *FileSink_JsonFormat) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*FileSink_JsonFormat)
	if !ok {
		that2, ok := that.(FileSink_JsonFormat)
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
	if !this.JsonFormat.Equal(that1.JsonFormat) {
		return false
	}
	return true
}
func (this *GrpcService) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GrpcService)
	if !ok {
		that2, ok := that.(GrpcService)
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
	if this.LogName != that1.LogName {
		return false
	}
	if that1.ServiceRef == nil {
		if this.ServiceRef != nil {
			return false
		}
	} else if this.ServiceRef == nil {
		return false
	} else if !this.ServiceRef.Equal(that1.ServiceRef) {
		return false
	}
	if len(this.AdditionalRequestHeadersToLog) != len(that1.AdditionalRequestHeadersToLog) {
		return false
	}
	for i := range this.AdditionalRequestHeadersToLog {
		if this.AdditionalRequestHeadersToLog[i] != that1.AdditionalRequestHeadersToLog[i] {
			return false
		}
	}
	if len(this.AdditionalResponseHeadersToLog) != len(that1.AdditionalResponseHeadersToLog) {
		return false
	}
	for i := range this.AdditionalResponseHeadersToLog {
		if this.AdditionalResponseHeadersToLog[i] != that1.AdditionalResponseHeadersToLog[i] {
			return false
		}
	}
	if len(this.AdditionalResponseTrailersToLog) != len(that1.AdditionalResponseTrailersToLog) {
		return false
	}
	for i := range this.AdditionalResponseTrailersToLog {
		if this.AdditionalResponseTrailersToLog[i] != that1.AdditionalResponseTrailersToLog[i] {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *GrpcService_StaticClusterName) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GrpcService_StaticClusterName)
	if !ok {
		that2, ok := that.(GrpcService_StaticClusterName)
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
	if this.StaticClusterName != that1.StaticClusterName {
		return false
	}
	return true
}

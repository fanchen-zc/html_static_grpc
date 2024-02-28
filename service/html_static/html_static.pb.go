// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: html_static.proto

package html_static

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type HTMLRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *HTMLRequest) Reset() {
	*x = HTMLRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_html_static_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HTMLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HTMLRequest) ProtoMessage() {}

func (x *HTMLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_html_static_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HTMLRequest.ProtoReflect.Descriptor instead.
func (*HTMLRequest) Descriptor() ([]byte, []int) {
	return file_html_static_proto_rawDescGZIP(), []int{0}
}

func (x *HTMLRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

// 定义服务端响应的数据格式
type HTMLResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content string `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *HTMLResponse) Reset() {
	*x = HTMLResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_html_static_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HTMLResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HTMLResponse) ProtoMessage() {}

func (x *HTMLResponse) ProtoReflect() protoreflect.Message {
	mi := &file_html_static_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HTMLResponse.ProtoReflect.Descriptor instead.
func (*HTMLResponse) Descriptor() ([]byte, []int) {
	return file_html_static_proto_rawDescGZIP(), []int{1}
}

func (x *HTMLResponse) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

var File_html_static_proto protoreflect.FileDescriptor

var file_html_static_proto_rawDesc = []byte{
	0x0a, 0x11, 0x68, 0x74, 0x6d, 0x6c, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x68, 0x74, 0x6d, 0x6c, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63,
	0x22, 0x1f, 0x0a, 0x0b, 0x48, 0x54, 0x4d, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72,
	0x6c, 0x22, 0x28, 0x0a, 0x0c, 0x48, 0x54, 0x4d, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x32, 0x54, 0x0a, 0x0b, 0x48,
	0x54, 0x4d, 0x4c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x45, 0x0a, 0x0e, 0x47, 0x65,
	0x74, 0x48, 0x54, 0x4d, 0x4c, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x2e, 0x68,
	0x74, 0x6d, 0x6c, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x2e, 0x48, 0x54, 0x4d, 0x4c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x68, 0x74, 0x6d, 0x6c, 0x5f, 0x73, 0x74,
	0x61, 0x74, 0x69, 0x63, 0x2e, 0x48, 0x54, 0x4d, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_html_static_proto_rawDescOnce sync.Once
	file_html_static_proto_rawDescData = file_html_static_proto_rawDesc
)

func file_html_static_proto_rawDescGZIP() []byte {
	file_html_static_proto_rawDescOnce.Do(func() {
		file_html_static_proto_rawDescData = protoimpl.X.CompressGZIP(file_html_static_proto_rawDescData)
	})
	return file_html_static_proto_rawDescData
}

var file_html_static_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_html_static_proto_goTypes = []interface{}{
	(*HTMLRequest)(nil),  // 0: html_static.HTMLRequest
	(*HTMLResponse)(nil), // 1: html_static.HTMLResponse
}
var file_html_static_proto_depIdxs = []int32{
	0, // 0: html_static.HTMLService.GetHTMLContent:input_type -> html_static.HTMLRequest
	1, // 1: html_static.HTMLService.GetHTMLContent:output_type -> html_static.HTMLResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_html_static_proto_init() }
func file_html_static_proto_init() {
	if File_html_static_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_html_static_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HTMLRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_html_static_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HTMLResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_html_static_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_html_static_proto_goTypes,
		DependencyIndexes: file_html_static_proto_depIdxs,
		MessageInfos:      file_html_static_proto_msgTypes,
	}.Build()
	File_html_static_proto = out.File
	file_html_static_proto_rawDesc = nil
	file_html_static_proto_goTypes = nil
	file_html_static_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// HTMLServiceClient is the client API for HTMLService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HTMLServiceClient interface {
	// 定义服务内的 GetHTMLContent 远程调用
	GetHTMLContent(ctx context.Context, in *HTMLRequest, opts ...grpc.CallOption) (*HTMLResponse, error)
}

type hTMLServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHTMLServiceClient(cc grpc.ClientConnInterface) HTMLServiceClient {
	return &hTMLServiceClient{cc}
}

func (c *hTMLServiceClient) GetHTMLContent(ctx context.Context, in *HTMLRequest, opts ...grpc.CallOption) (*HTMLResponse, error) {
	out := new(HTMLResponse)
	err := c.cc.Invoke(ctx, "/html_static.HTMLService/GetHTMLContent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HTMLServiceServer is the server API for HTMLService service.
type HTMLServiceServer interface {
	// 定义服务内的 GetHTMLContent 远程调用
	GetHTMLContent(context.Context, *HTMLRequest) (*HTMLResponse, error)
}

// UnimplementedHTMLServiceServer can be embedded to have forward compatible implementations.
type UnimplementedHTMLServiceServer struct {
}

func (*UnimplementedHTMLServiceServer) GetHTMLContent(context.Context, *HTMLRequest) (*HTMLResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHTMLContent not implemented")
}

func RegisterHTMLServiceServer(s *grpc.Server, srv HTMLServiceServer) {
	s.RegisterService(&_HTMLService_serviceDesc, srv)
}

func _HTMLService_GetHTMLContent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HTMLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HTMLServiceServer).GetHTMLContent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/html_static.HTMLService/GetHTMLContent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HTMLServiceServer).GetHTMLContent(ctx, req.(*HTMLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HTMLService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "html_static.HTMLService",
	HandlerType: (*HTMLServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetHTMLContent",
			Handler:    _HTMLService_GetHTMLContent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "html_static.proto",
}

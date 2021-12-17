// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.14.0
// source: policy.proto

package protobuf

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

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int32 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_policy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_policy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_policy_proto_rawDescGZIP(), []int{0}
}

func (x *Response) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type Policy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Policy []byte `protobuf:"bytes,1,opt,name=policy,proto3" json:"policy,omitempty"`
}

func (x *Policy) Reset() {
	*x = Policy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_policy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Policy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Policy) ProtoMessage() {}

func (x *Policy) ProtoReflect() protoreflect.Message {
	mi := &file_policy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Policy.ProtoReflect.Descriptor instead.
func (*Policy) Descriptor() ([]byte, []int) {
	return file_policy_proto_rawDescGZIP(), []int{1}
}

func (x *Policy) GetPolicy() []byte {
	if x != nil {
		return x.Policy
	}
	return nil
}

var File_policy_proto protoreflect.FileDescriptor

var file_policy_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x22, 0x22, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x20, 0x0a, 0x06, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x32, 0x3f, 0x0a, 0x0d,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2e, 0x0a,
	0x0a, 0x68, 0x6f, 0x73, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x0e, 0x2e, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x1a, 0x10, 0x2e, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x29, 0x5a,
	0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x75, 0x62, 0x65,
	0x61, 0x72, 0x6d, 0x6f, 0x72, 0x2f, 0x4b, 0x75, 0x62, 0x65, 0x41, 0x72, 0x6d, 0x6f, 0x72, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_policy_proto_rawDescOnce sync.Once
	file_policy_proto_rawDescData = file_policy_proto_rawDesc
)

func file_policy_proto_rawDescGZIP() []byte {
	file_policy_proto_rawDescOnce.Do(func() {
		file_policy_proto_rawDescData = protoimpl.X.CompressGZIP(file_policy_proto_rawDescData)
	})
	return file_policy_proto_rawDescData
}

var file_policy_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_policy_proto_goTypes = []interface{}{
	(*Response)(nil), // 0: policy.response
	(*Policy)(nil),   // 1: policy.policy
}
var file_policy_proto_depIdxs = []int32{
	1, // 0: policy.PolicyService.hostPolicy:input_type -> policy.policy
	0, // 1: policy.PolicyService.hostPolicy:output_type -> policy.response
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_policy_proto_init() }
func file_policy_proto_init() {
	if File_policy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_policy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_policy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Policy); i {
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
			RawDescriptor: file_policy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_policy_proto_goTypes,
		DependencyIndexes: file_policy_proto_depIdxs,
		MessageInfos:      file_policy_proto_msgTypes,
	}.Build()
	File_policy_proto = out.File
	file_policy_proto_rawDesc = nil
	file_policy_proto_goTypes = nil
	file_policy_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PolicyServiceClient is the client API for PolicyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PolicyServiceClient interface {
	HostPolicy(ctx context.Context, in *Policy, opts ...grpc.CallOption) (*Response, error)
}

type policyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPolicyServiceClient(cc grpc.ClientConnInterface) PolicyServiceClient {
	return &policyServiceClient{cc}
}

func (c *policyServiceClient) HostPolicy(ctx context.Context, in *Policy, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/policy.PolicyService/hostPolicy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PolicyServiceServer is the server API for PolicyService service.
type PolicyServiceServer interface {
	HostPolicy(context.Context, *Policy) (*Response, error)
}

// UnimplementedPolicyServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPolicyServiceServer struct {
}

func (*UnimplementedPolicyServiceServer) HostPolicy(context.Context, *Policy) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HostPolicy not implemented")
}

func RegisterPolicyServiceServer(s *grpc.Server, srv PolicyServiceServer) {
	s.RegisterService(&_PolicyService_serviceDesc, srv)
}

func _PolicyService_HostPolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Policy)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolicyServiceServer).HostPolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/policy.PolicyService/HostPolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolicyServiceServer).HostPolicy(ctx, req.(*Policy))
	}
	return interceptor(ctx, in, info, handler)
}

var _PolicyService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "policy.PolicyService",
	HandlerType: (*PolicyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "hostPolicy",
			Handler:    _PolicyService_HostPolicy_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "policy.proto",
}
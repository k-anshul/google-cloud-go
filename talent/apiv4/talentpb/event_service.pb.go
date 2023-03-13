// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: google/cloud/talent/v4/event_service.proto

package talentpb

import (
	context "context"
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The report event request.
type CreateClientEventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Resource name of the tenant under which the event is created.
	//
	// The format is "projects/{project_id}/tenants/{tenant_id}", for example,
	// "projects/foo/tenants/bar".
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Required. Events issued when end user interacts with customer's application
	// that uses Cloud Talent Solution.
	ClientEvent *ClientEvent `protobuf:"bytes,2,opt,name=client_event,json=clientEvent,proto3" json:"client_event,omitempty"`
}

func (x *CreateClientEventRequest) Reset() {
	*x = CreateClientEventRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_talent_v4_event_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateClientEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateClientEventRequest) ProtoMessage() {}

func (x *CreateClientEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_talent_v4_event_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateClientEventRequest.ProtoReflect.Descriptor instead.
func (*CreateClientEventRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_talent_v4_event_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateClientEventRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *CreateClientEventRequest) GetClientEvent() *ClientEvent {
	if x != nil {
		return x.ClientEvent
	}
	return nil
}

var File_google_cloud_talent_v4_event_service_proto protoreflect.FileDescriptor

var file_google_cloud_talent_v4_event_service_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x74,
	0x61, 0x6c, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x34, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x61, 0x6c, 0x65, 0x6e,
	0x74, 0x2e, 0x76, 0x34, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65,
	0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x74, 0x61, 0x6c, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x34, 0x2f,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa3, 0x01, 0x0a, 0x18,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3a, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x22, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x1c,
	0x0a, 0x1a, 0x6a, 0x6f, 0x62, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x06, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x12, 0x4b, 0x0a, 0x0c, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x61, 0x6c, 0x65, 0x6e, 0x74,
	0x2e, 0x76, 0x34, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x42,
	0x03, 0xe0, 0x41, 0x02, 0x52, 0x0b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x32, 0xc5, 0x02, 0x0a, 0x0c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0xc6, 0x01, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x30, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x61, 0x6c, 0x65, 0x6e, 0x74, 0x2e, 0x76,
	0x34, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x61, 0x6c, 0x65, 0x6e, 0x74,
	0x2e, 0x76, 0x34, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x22,
	0x5a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x3e, 0x22, 0x2e, 0x2f, 0x76, 0x34, 0x2f, 0x7b, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f,
	0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x3a, 0x0c, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0xda, 0x41, 0x13, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x2c, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x1a, 0x6c, 0xca, 0x41, 0x13,
	0x6a, 0x6f, 0x62, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0xd2, 0x41, 0x53, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77,
	0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x2c, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x61, 0x75, 0x74, 0x68, 0x2f, 0x6a, 0x6f, 0x62, 0x73, 0x42, 0x6b, 0x0a, 0x1a, 0x63, 0x6f, 0x6d,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x61,
	0x6c, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x34, 0x42, 0x11, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x32, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67,
	0x6f, 0x2f, 0x74, 0x61, 0x6c, 0x65, 0x6e, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x34, 0x2f, 0x74,
	0x61, 0x6c, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x3b, 0x74, 0x61, 0x6c, 0x65, 0x6e, 0x74, 0x70, 0x62,
	0xa2, 0x02, 0x03, 0x43, 0x54, 0x53, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_talent_v4_event_service_proto_rawDescOnce sync.Once
	file_google_cloud_talent_v4_event_service_proto_rawDescData = file_google_cloud_talent_v4_event_service_proto_rawDesc
)

func file_google_cloud_talent_v4_event_service_proto_rawDescGZIP() []byte {
	file_google_cloud_talent_v4_event_service_proto_rawDescOnce.Do(func() {
		file_google_cloud_talent_v4_event_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_talent_v4_event_service_proto_rawDescData)
	})
	return file_google_cloud_talent_v4_event_service_proto_rawDescData
}

var file_google_cloud_talent_v4_event_service_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_cloud_talent_v4_event_service_proto_goTypes = []interface{}{
	(*CreateClientEventRequest)(nil), // 0: google.cloud.talent.v4.CreateClientEventRequest
	(*ClientEvent)(nil),              // 1: google.cloud.talent.v4.ClientEvent
}
var file_google_cloud_talent_v4_event_service_proto_depIdxs = []int32{
	1, // 0: google.cloud.talent.v4.CreateClientEventRequest.client_event:type_name -> google.cloud.talent.v4.ClientEvent
	0, // 1: google.cloud.talent.v4.EventService.CreateClientEvent:input_type -> google.cloud.talent.v4.CreateClientEventRequest
	1, // 2: google.cloud.talent.v4.EventService.CreateClientEvent:output_type -> google.cloud.talent.v4.ClientEvent
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_google_cloud_talent_v4_event_service_proto_init() }
func file_google_cloud_talent_v4_event_service_proto_init() {
	if File_google_cloud_talent_v4_event_service_proto != nil {
		return
	}
	file_google_cloud_talent_v4_event_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_talent_v4_event_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateClientEventRequest); i {
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
			RawDescriptor: file_google_cloud_talent_v4_event_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_cloud_talent_v4_event_service_proto_goTypes,
		DependencyIndexes: file_google_cloud_talent_v4_event_service_proto_depIdxs,
		MessageInfos:      file_google_cloud_talent_v4_event_service_proto_msgTypes,
	}.Build()
	File_google_cloud_talent_v4_event_service_proto = out.File
	file_google_cloud_talent_v4_event_service_proto_rawDesc = nil
	file_google_cloud_talent_v4_event_service_proto_goTypes = nil
	file_google_cloud_talent_v4_event_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// EventServiceClient is the client API for EventService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EventServiceClient interface {
	// Report events issued when end user interacts with customer's application
	// that uses Cloud Talent Solution. You may inspect the created events in
	// [self service
	// tools](https://console.cloud.google.com/talent-solution/overview).
	// [Learn
	// more](https://cloud.google.com/talent-solution/docs/management-tools)
	// about self service tools.
	CreateClientEvent(ctx context.Context, in *CreateClientEventRequest, opts ...grpc.CallOption) (*ClientEvent, error)
}

type eventServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEventServiceClient(cc grpc.ClientConnInterface) EventServiceClient {
	return &eventServiceClient{cc}
}

func (c *eventServiceClient) CreateClientEvent(ctx context.Context, in *CreateClientEventRequest, opts ...grpc.CallOption) (*ClientEvent, error) {
	out := new(ClientEvent)
	err := c.cc.Invoke(ctx, "/google.cloud.talent.v4.EventService/CreateClientEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventServiceServer is the server API for EventService service.
type EventServiceServer interface {
	// Report events issued when end user interacts with customer's application
	// that uses Cloud Talent Solution. You may inspect the created events in
	// [self service
	// tools](https://console.cloud.google.com/talent-solution/overview).
	// [Learn
	// more](https://cloud.google.com/talent-solution/docs/management-tools)
	// about self service tools.
	CreateClientEvent(context.Context, *CreateClientEventRequest) (*ClientEvent, error)
}

// UnimplementedEventServiceServer can be embedded to have forward compatible implementations.
type UnimplementedEventServiceServer struct {
}

func (*UnimplementedEventServiceServer) CreateClientEvent(context.Context, *CreateClientEventRequest) (*ClientEvent, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateClientEvent not implemented")
}

func RegisterEventServiceServer(s *grpc.Server, srv EventServiceServer) {
	s.RegisterService(&_EventService_serviceDesc, srv)
}

func _EventService_CreateClientEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateClientEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).CreateClientEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.talent.v4.EventService/CreateClientEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).CreateClientEvent(ctx, req.(*CreateClientEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _EventService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.talent.v4.EventService",
	HandlerType: (*EventServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateClientEvent",
			Handler:    _EventService_CreateClientEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/talent/v4/event_service.proto",
}

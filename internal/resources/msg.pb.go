// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: msg.proto

package resources

import (
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

type Msg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid *UUID      `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Ts   *Timestamp `protobuf:"bytes,2,opt,name=ts,proto3" json:"ts,omitempty"`
	// Types that are assignable to Payload:
	//	*Msg_Error
	//	*Msg_AuthUserRequest
	//	*Msg_AuthUserResponse
	Payload isMsg_Payload `protobuf_oneof:"payload"`
}

func (x *Msg) Reset() {
	*x = Msg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Msg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Msg) ProtoMessage() {}

func (x *Msg) ProtoReflect() protoreflect.Message {
	mi := &file_msg_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Msg.ProtoReflect.Descriptor instead.
func (*Msg) Descriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{0}
}

func (x *Msg) GetUuid() *UUID {
	if x != nil {
		return x.Uuid
	}
	return nil
}

func (x *Msg) GetTs() *Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

func (m *Msg) GetPayload() isMsg_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *Msg) GetError() *Error {
	if x, ok := x.GetPayload().(*Msg_Error); ok {
		return x.Error
	}
	return nil
}

func (x *Msg) GetAuthUserRequest() *AuthUserRequest {
	if x, ok := x.GetPayload().(*Msg_AuthUserRequest); ok {
		return x.AuthUserRequest
	}
	return nil
}

func (x *Msg) GetAuthUserResponse() *AuthUserResponse {
	if x, ok := x.GetPayload().(*Msg_AuthUserResponse); ok {
		return x.AuthUserResponse
	}
	return nil
}

type isMsg_Payload interface {
	isMsg_Payload()
}

type Msg_Error struct {
	Error *Error `protobuf:"bytes,3,opt,name=error,proto3,oneof"`
}

type Msg_AuthUserRequest struct {
	AuthUserRequest *AuthUserRequest `protobuf:"bytes,4,opt,name=auth_user_request,json=authUserRequest,proto3,oneof"`
}

type Msg_AuthUserResponse struct {
	AuthUserResponse *AuthUserResponse `protobuf:"bytes,5,opt,name=auth_user_response,json=authUserResponse,proto3,oneof"`
}

func (*Msg_Error) isMsg_Payload() {}

func (*Msg_AuthUserRequest) isMsg_Payload() {}

func (*Msg_AuthUserResponse) isMsg_Payload() {}

var File_msg_proto protoreflect.FileDescriptor

var file_msg_proto_rawDesc = []byte{
	0x0a, 0x09, 0x6d, 0x73, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x1a, 0x0d, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x0f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x9c, 0x02, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x12, 0x23, 0x0a, 0x04, 0x75,
	0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x55, 0x55, 0x49, 0x44, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64,
	0x12, 0x24, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x02, 0x74, 0x73, 0x12, 0x28, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x12, 0x48, 0x0a, 0x11, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x0f, 0x61, 0x75, 0x74, 0x68, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4b, 0x0a, 0x12, 0x61, 0x75,
	0x74, 0x68, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x10, 0x61, 0x75, 0x74, 0x68, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x42, 0x0c, 0x5a, 0x0a, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_msg_proto_rawDescOnce sync.Once
	file_msg_proto_rawDescData = file_msg_proto_rawDesc
)

func file_msg_proto_rawDescGZIP() []byte {
	file_msg_proto_rawDescOnce.Do(func() {
		file_msg_proto_rawDescData = protoimpl.X.CompressGZIP(file_msg_proto_rawDescData)
	})
	return file_msg_proto_rawDescData
}

var file_msg_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_msg_proto_goTypes = []interface{}{
	(*Msg)(nil),              // 0: resources.Msg
	(*UUID)(nil),             // 1: resources.UUID
	(*Timestamp)(nil),        // 2: resources.Timestamp
	(*Error)(nil),            // 3: resources.Error
	(*AuthUserRequest)(nil),  // 4: resources.AuthUserRequest
	(*AuthUserResponse)(nil), // 5: resources.AuthUserResponse
}
var file_msg_proto_depIdxs = []int32{
	1, // 0: resources.Msg.uuid:type_name -> resources.UUID
	2, // 1: resources.Msg.ts:type_name -> resources.Timestamp
	3, // 2: resources.Msg.error:type_name -> resources.Error
	4, // 3: resources.Msg.auth_user_request:type_name -> resources.AuthUserRequest
	5, // 4: resources.Msg.auth_user_response:type_name -> resources.AuthUserResponse
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_msg_proto_init() }
func file_msg_proto_init() {
	if File_msg_proto != nil {
		return
	}
	file_generic_proto_init()
	file_error_proto_init()
	file_auth_user_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_msg_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Msg); i {
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
	file_msg_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Msg_Error)(nil),
		(*Msg_AuthUserRequest)(nil),
		(*Msg_AuthUserResponse)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_msg_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_msg_proto_goTypes,
		DependencyIndexes: file_msg_proto_depIdxs,
		MessageInfos:      file_msg_proto_msgTypes,
	}.Build()
	File_msg_proto = out.File
	file_msg_proto_rawDesc = nil
	file_msg_proto_goTypes = nil
	file_msg_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: user/v1/user.proto

package userv1

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

// UserMessage represents a User.
// This is the shape of the data artifact that will move across the API boundary.
type UserMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FirstName string `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName  string `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Email     string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *UserMessage) Reset() {
	*x = UserMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_v1_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserMessage) ProtoMessage() {}

func (x *UserMessage) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserMessage.ProtoReflect.Descriptor instead.
func (*UserMessage) Descriptor() ([]byte, []int) {
	return file_user_v1_user_proto_rawDescGZIP(), []int{0}
}

func (x *UserMessage) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserMessage) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *UserMessage) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *UserMessage) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

// FetchUserByIDRequest is for fetching a user by unique id.
type FetchUserByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *FetchUserByIDRequest) Reset() {
	*x = FetchUserByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_v1_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchUserByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchUserByIDRequest) ProtoMessage() {}

func (x *FetchUserByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchUserByIDRequest.ProtoReflect.Descriptor instead.
func (*FetchUserByIDRequest) Descriptor() ([]byte, []int) {
	return file_user_v1_user_proto_rawDescGZIP(), []int{1}
}

func (x *FetchUserByIDRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

// FetchUserByIDResponse contains User data.
type FetchUserByIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *UserMessage `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *FetchUserByIDResponse) Reset() {
	*x = FetchUserByIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_v1_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchUserByIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchUserByIDResponse) ProtoMessage() {}

func (x *FetchUserByIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchUserByIDResponse.ProtoReflect.Descriptor instead.
func (*FetchUserByIDResponse) Descriptor() ([]byte, []int) {
	return file_user_v1_user_proto_rawDescGZIP(), []int{2}
}

func (x *FetchUserByIDResponse) GetUser() *UserMessage {
	if x != nil {
		return x.User
	}
	return nil
}

var File_user_v1_user_proto protoreflect.FileDescriptor

var file_user_v1_user_proto_rawDesc = []byte{
	0x0a, 0x12, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x22, 0x6f, 0x0a,
	0x0b, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a,
	0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c,
	0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x26,
	0x0a, 0x14, 0x46, 0x65, 0x74, 0x63, 0x68, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x44, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x41, 0x0a, 0x15, 0x46, 0x65, 0x74, 0x63, 0x68, 0x55,
	0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x28, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x32, 0x5f, 0x0a, 0x0b, 0x55, 0x73, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x50, 0x0a, 0x0d, 0x46, 0x65, 0x74, 0x63,
	0x68, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x44, 0x12, 0x1d, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49,
	0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x44,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x88, 0x01, 0x0a, 0x0b, 0x63,
	0x6f, 0x6d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x09, 0x55, 0x73, 0x65, 0x72,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x69, 0x6b, 0x65, 0x64, 0x6f, 0x6e, 0x6e, 0x69, 0x63, 0x69, 0x2f,
	0x6d, 0x6f, 0x6e, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x2f, 0x76, 0x31, 0x3b, 0x75, 0x73, 0x65, 0x72, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x55, 0x58, 0x58,
	0xaa, 0x02, 0x07, 0x55, 0x73, 0x65, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x07, 0x55, 0x73, 0x65,
	0x72, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x13, 0x55, 0x73, 0x65, 0x72, 0x5c, 0x56, 0x31, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x08, 0x55, 0x73, 0x65,
	0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_v1_user_proto_rawDescOnce sync.Once
	file_user_v1_user_proto_rawDescData = file_user_v1_user_proto_rawDesc
)

func file_user_v1_user_proto_rawDescGZIP() []byte {
	file_user_v1_user_proto_rawDescOnce.Do(func() {
		file_user_v1_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_v1_user_proto_rawDescData)
	})
	return file_user_v1_user_proto_rawDescData
}

var file_user_v1_user_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_user_v1_user_proto_goTypes = []interface{}{
	(*UserMessage)(nil),           // 0: user.v1.UserMessage
	(*FetchUserByIDRequest)(nil),  // 1: user.v1.FetchUserByIDRequest
	(*FetchUserByIDResponse)(nil), // 2: user.v1.FetchUserByIDResponse
}
var file_user_v1_user_proto_depIdxs = []int32{
	0, // 0: user.v1.FetchUserByIDResponse.user:type_name -> user.v1.UserMessage
	1, // 1: user.v1.UserService.FetchUserByID:input_type -> user.v1.FetchUserByIDRequest
	2, // 2: user.v1.UserService.FetchUserByID:output_type -> user.v1.FetchUserByIDResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_user_v1_user_proto_init() }
func file_user_v1_user_proto_init() {
	if File_user_v1_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_v1_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserMessage); i {
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
		file_user_v1_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchUserByIDRequest); i {
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
		file_user_v1_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchUserByIDResponse); i {
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
			RawDescriptor: file_user_v1_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_v1_user_proto_goTypes,
		DependencyIndexes: file_user_v1_user_proto_depIdxs,
		MessageInfos:      file_user_v1_user_proto_msgTypes,
	}.Build()
	File_user_v1_user_proto = out.File
	file_user_v1_user_proto_rawDesc = nil
	file_user_v1_user_proto_goTypes = nil
	file_user_v1_user_proto_depIdxs = nil
}

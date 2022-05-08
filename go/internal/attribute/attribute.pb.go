// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: attribute.proto

package attribute

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

// Attribute represents an attribute in its proto message form.
// This is the shape of the data artifact that will move across the API boundary.
type AttributeRaw struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *AttributeRaw) Reset() {
	*x = AttributeRaw{}
	if protoimpl.UnsafeEnabled {
		mi := &file_attribute_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttributeRaw) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttributeRaw) ProtoMessage() {}

func (x *AttributeRaw) ProtoReflect() protoreflect.Message {
	mi := &file_attribute_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttributeRaw.ProtoReflect.Descriptor instead.
func (*AttributeRaw) Descriptor() ([]byte, []int) {
	return file_attribute_proto_rawDescGZIP(), []int{0}
}

func (x *AttributeRaw) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AttributeRaw) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *AttributeRaw) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// FetchAttributeRequest contains one or more fields used for locating an attribute
type FetchAttributeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   *int64  `protobuf:"varint,1,opt,name=id,proto3,oneof" json:"id,omitempty"`
	Type *string `protobuf:"bytes,2,opt,name=type,proto3,oneof" json:"type,omitempty"`
}

func (x *FetchAttributeRequest) Reset() {
	*x = FetchAttributeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_attribute_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchAttributeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchAttributeRequest) ProtoMessage() {}

func (x *FetchAttributeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_attribute_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchAttributeRequest.ProtoReflect.Descriptor instead.
func (*FetchAttributeRequest) Descriptor() ([]byte, []int) {
	return file_attribute_proto_rawDescGZIP(), []int{1}
}

func (x *FetchAttributeRequest) GetId() int64 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *FetchAttributeRequest) GetType() string {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return ""
}

// FetchAttributeResponse is returned
type FetchAttributeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Attribute *AttributeRaw `protobuf:"bytes,1,opt,name=attribute,proto3" json:"attribute,omitempty"`
}

func (x *FetchAttributeResponse) Reset() {
	*x = FetchAttributeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_attribute_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchAttributeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchAttributeResponse) ProtoMessage() {}

func (x *FetchAttributeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_attribute_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchAttributeResponse.ProtoReflect.Descriptor instead.
func (*FetchAttributeResponse) Descriptor() ([]byte, []int) {
	return file_attribute_proto_rawDescGZIP(), []int{2}
}

func (x *FetchAttributeResponse) GetAttribute() *AttributeRaw {
	if x != nil {
		return x.Attribute
	}
	return nil
}

var File_attribute_proto protoreflect.FileDescriptor

var file_attribute_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x22, 0x46, 0x0a, 0x0c,
	0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x52, 0x61, 0x77, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x55, 0x0a, 0x15, 0x46, 0x65, 0x74, 0x63, 0x68, 0x41, 0x74, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x13, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x02, 0x69, 0x64, 0x88,
	0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x01, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x42, 0x05, 0x0a, 0x03, 0x5f,
	0x69, 0x64, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x22, 0x4f, 0x0a, 0x16, 0x46,
	0x65, 0x74, 0x63, 0x68, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x09, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x61, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x52, 0x61,
	0x77, 0x52, 0x09, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x32, 0x6b, 0x0a, 0x10,
	0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x57, 0x0a, 0x0e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x12, 0x20, 0x2e, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x2e, 0x46,
	0x65, 0x74, 0x63, 0x68, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x69, 0x6b, 0x65, 0x64, 0x6f, 0x6e, 0x6e,
	0x69, 0x63, 0x69, 0x2f, 0x6d, 0x6f, 0x6e, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_attribute_proto_rawDescOnce sync.Once
	file_attribute_proto_rawDescData = file_attribute_proto_rawDesc
)

func file_attribute_proto_rawDescGZIP() []byte {
	file_attribute_proto_rawDescOnce.Do(func() {
		file_attribute_proto_rawDescData = protoimpl.X.CompressGZIP(file_attribute_proto_rawDescData)
	})
	return file_attribute_proto_rawDescData
}

var file_attribute_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_attribute_proto_goTypes = []interface{}{
	(*AttributeRaw)(nil),           // 0: attribute.AttributeRaw
	(*FetchAttributeRequest)(nil),  // 1: attribute.FetchAttributeRequest
	(*FetchAttributeResponse)(nil), // 2: attribute.FetchAttributeResponse
}
var file_attribute_proto_depIdxs = []int32{
	0, // 0: attribute.FetchAttributeResponse.attribute:type_name -> attribute.AttributeRaw
	1, // 1: attribute.AttributeService.FetchAttribute:input_type -> attribute.FetchAttributeRequest
	2, // 2: attribute.AttributeService.FetchAttribute:output_type -> attribute.FetchAttributeResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_attribute_proto_init() }
func file_attribute_proto_init() {
	if File_attribute_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_attribute_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttributeRaw); i {
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
		file_attribute_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchAttributeRequest); i {
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
		file_attribute_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchAttributeResponse); i {
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
	file_attribute_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_attribute_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_attribute_proto_goTypes,
		DependencyIndexes: file_attribute_proto_depIdxs,
		MessageInfos:      file_attribute_proto_msgTypes,
	}.Build()
	File_attribute_proto = out.File
	file_attribute_proto_rawDesc = nil
	file_attribute_proto_goTypes = nil
	file_attribute_proto_depIdxs = nil
}

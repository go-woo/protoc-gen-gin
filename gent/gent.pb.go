// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: gent/gent.proto

package gent

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type FieldRules struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rules string `protobuf:"bytes,1,opt,name=rules,proto3" json:"rules,omitempty"`
}

func (x *FieldRules) Reset() {
	*x = FieldRules{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gent_gent_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FieldRules) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FieldRules) ProtoMessage() {}

func (x *FieldRules) ProtoReflect() protoreflect.Message {
	mi := &file_gent_gent_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FieldRules.ProtoReflect.Descriptor instead.
func (*FieldRules) Descriptor() ([]byte, []int) {
	return file_gent_gent_proto_rawDescGZIP(), []int{0}
}

func (x *FieldRules) GetRules() string {
	if x != nil {
		return x.Rules
	}
	return ""
}

var file_gent_gent_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*FieldRules)(nil),
		Field:         1171,
		Name:          "gent.field",
		Tag:           "bytes,1171,opt,name=field",
		Filename:      "gent/gent.proto",
	},
}

// Extension fields to descriptorpb.FieldOptions.
var (
	// Rules specify the validations to be performed on this field. By default,
	// no validation is performed against a field.
	//
	// optional gent.FieldRules field = 1171;
	E_Field = &file_gent_gent_proto_extTypes[0]
)

var File_gent_gent_proto protoreflect.FileDescriptor

var file_gent_gent_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x67, 0x65, 0x6e, 0x74, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x22, 0x0a, 0x0a, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x3a, 0x49, 0x0a,
	0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x93, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x67,
	0x65, 0x6e, 0x74, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x52, 0x05,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x88, 0x01, 0x01, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2f, 0x67, 0x6f, 0x2d, 0x77, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x65, 0x6e, 0x74, 0x2f, 0x67, 0x65, 0x6e, 0x74, 0x3b, 0x67, 0x65,
	0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gent_gent_proto_rawDescOnce sync.Once
	file_gent_gent_proto_rawDescData = file_gent_gent_proto_rawDesc
)

func file_gent_gent_proto_rawDescGZIP() []byte {
	file_gent_gent_proto_rawDescOnce.Do(func() {
		file_gent_gent_proto_rawDescData = protoimpl.X.CompressGZIP(file_gent_gent_proto_rawDescData)
	})
	return file_gent_gent_proto_rawDescData
}

var file_gent_gent_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_gent_gent_proto_goTypes = []interface{}{
	(*FieldRules)(nil),                // 0: gent.FieldRules
	(*descriptorpb.FieldOptions)(nil), // 1: google.protobuf.FieldOptions
}
var file_gent_gent_proto_depIdxs = []int32{
	1, // 0: gent.field:extendee -> google.protobuf.FieldOptions
	0, // 1: gent.field:type_name -> gent.FieldRules
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	1, // [1:2] is the sub-list for extension type_name
	0, // [0:1] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_gent_gent_proto_init() }
func file_gent_gent_proto_init() {
	if File_gent_gent_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gent_gent_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FieldRules); i {
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
			RawDescriptor: file_gent_gent_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_gent_gent_proto_goTypes,
		DependencyIndexes: file_gent_gent_proto_depIdxs,
		MessageInfos:      file_gent_gent_proto_msgTypes,
		ExtensionInfos:    file_gent_gent_proto_extTypes,
	}.Build()
	File_gent_gent_proto = out.File
	file_gent_gent_proto_rawDesc = nil
	file_gent_gent_proto_goTypes = nil
	file_gent_gent_proto_depIdxs = nil
}
// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: oidc/v1/oidc.proto

package oidcv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_oidc_v1_oidc_proto protoreflect.FileDescriptor

var file_oidc_v1_oidc_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6f, 0x69, 0x64, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x69, 0x64, 0x63, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6f, 0x69, 0x64, 0x63, 0x2e, 0x76, 0x31, 0x32, 0x17, 0x0a,
	0x15, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x42, 0x85, 0x01, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x2e, 0x6f,
	0x69, 0x64, 0x63, 0x2e, 0x76, 0x31, 0x42, 0x09, 0x4f, 0x69, 0x64, 0x63, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x70, 0x31, 0x61, 0x73, 0x73, 0x2f, 0x69, 0x64, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64,
	0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x6f, 0x69, 0x64, 0x63, 0x2f, 0x76, 0x31, 0x3b, 0x6f, 0x69, 0x64,
	0x63, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x4f, 0x58, 0x58, 0xaa, 0x02, 0x07, 0x4f, 0x69, 0x64, 0x63,
	0x2e, 0x56, 0x31, 0xca, 0x02, 0x07, 0x4f, 0x69, 0x64, 0x63, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x13,
	0x4f, 0x69, 0x64, 0x63, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x08, 0x4f, 0x69, 0x64, 0x63, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_oidc_v1_oidc_proto_goTypes = []interface{}{}
var file_oidc_v1_oidc_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_oidc_v1_oidc_proto_init() }
func file_oidc_v1_oidc_proto_init() {
	if File_oidc_v1_oidc_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_oidc_v1_oidc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_oidc_v1_oidc_proto_goTypes,
		DependencyIndexes: file_oidc_v1_oidc_proto_depIdxs,
	}.Build()
	File_oidc_v1_oidc_proto = out.File
	file_oidc_v1_oidc_proto_rawDesc = nil
	file_oidc_v1_oidc_proto_goTypes = nil
	file_oidc_v1_oidc_proto_depIdxs = nil
}

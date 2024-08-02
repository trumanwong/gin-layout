// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.24.3
// source: api/app/v1/app_error.proto

package v1

import (
	_ "github.com/trumanwong/gin-transport/transport/errors"
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

type AppServiceErrorReason int32

const (
	AppServiceErrorReason_UNKNOWN_ERROR AppServiceErrorReason = 0
	AppServiceErrorReason_LOGIN_FAILED  AppServiceErrorReason = 1
	AppServiceErrorReason_FORBIDDEN     AppServiceErrorReason = 2
	AppServiceErrorReason_NotFound      AppServiceErrorReason = 3
	AppServiceErrorReason_NotAcceptable AppServiceErrorReason = 4
	AppServiceErrorReason_Conflict      AppServiceErrorReason = 5
	AppServiceErrorReason_Internal      AppServiceErrorReason = 6
)

// Enum value maps for AppServiceErrorReason.
var (
	AppServiceErrorReason_name = map[int32]string{
		0: "UNKNOWN_ERROR",
		1: "LOGIN_FAILED",
		2: "FORBIDDEN",
		3: "NotFound",
		4: "NotAcceptable",
		5: "Conflict",
		6: "Internal",
	}
	AppServiceErrorReason_value = map[string]int32{
		"UNKNOWN_ERROR": 0,
		"LOGIN_FAILED":  1,
		"FORBIDDEN":     2,
		"NotFound":      3,
		"NotAcceptable": 4,
		"Conflict":      5,
		"Internal":      6,
	}
)

func (x AppServiceErrorReason) Enum() *AppServiceErrorReason {
	p := new(AppServiceErrorReason)
	*p = x
	return p
}

func (x AppServiceErrorReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AppServiceErrorReason) Descriptor() protoreflect.EnumDescriptor {
	return file_api_app_v1_app_error_proto_enumTypes[0].Descriptor()
}

func (AppServiceErrorReason) Type() protoreflect.EnumType {
	return &file_api_app_v1_app_error_proto_enumTypes[0]
}

func (x AppServiceErrorReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AppServiceErrorReason.Descriptor instead.
func (AppServiceErrorReason) EnumDescriptor() ([]byte, []int) {
	return file_api_app_v1_app_error_proto_rawDescGZIP(), []int{0}
}

var File_api_app_v1_app_error_proto protoreflect.FileDescriptor

var file_api_app_v1_app_error_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x70,
	0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x61, 0x70,
	0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x76, 0x31, 0x1a, 0x13, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73,
	0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xb2, 0x01,
	0x0a, 0x15, 0x41, 0x70, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x11, 0x0a, 0x0d, 0x55, 0x4e, 0x4b, 0x4e, 0x4f,
	0x57, 0x4e, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x0c, 0x4c, 0x4f,
	0x47, 0x49, 0x4e, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x01, 0x1a, 0x04, 0xa8, 0x45,
	0x91, 0x03, 0x12, 0x13, 0x0a, 0x09, 0x46, 0x4f, 0x52, 0x42, 0x49, 0x44, 0x44, 0x45, 0x4e, 0x10,
	0x02, 0x1a, 0x04, 0xa8, 0x45, 0x93, 0x03, 0x12, 0x12, 0x0a, 0x08, 0x4e, 0x6f, 0x74, 0x46, 0x6f,
	0x75, 0x6e, 0x64, 0x10, 0x03, 0x1a, 0x04, 0xa8, 0x45, 0x94, 0x03, 0x12, 0x17, 0x0a, 0x0d, 0x4e,
	0x6f, 0x74, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x10, 0x04, 0x1a, 0x04,
	0xa8, 0x45, 0x96, 0x03, 0x12, 0x12, 0x0a, 0x08, 0x43, 0x6f, 0x6e, 0x66, 0x6c, 0x69, 0x63, 0x74,
	0x10, 0x05, 0x1a, 0x04, 0xa8, 0x45, 0x99, 0x03, 0x12, 0x12, 0x0a, 0x08, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x10, 0x06, 0x1a, 0x04, 0xa8, 0x45, 0xf4, 0x03, 0x1a, 0x04, 0xa0, 0x45,
	0x90, 0x03, 0x42, 0x28, 0x0a, 0x0a, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x76, 0x31,
	0x50, 0x01, 0x5a, 0x18, 0x67, 0x69, 0x6e, 0x2d, 0x6c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_app_v1_app_error_proto_rawDescOnce sync.Once
	file_api_app_v1_app_error_proto_rawDescData = file_api_app_v1_app_error_proto_rawDesc
)

func file_api_app_v1_app_error_proto_rawDescGZIP() []byte {
	file_api_app_v1_app_error_proto_rawDescOnce.Do(func() {
		file_api_app_v1_app_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_app_v1_app_error_proto_rawDescData)
	})
	return file_api_app_v1_app_error_proto_rawDescData
}

var file_api_app_v1_app_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_app_v1_app_error_proto_goTypes = []any{
	(AppServiceErrorReason)(0), // 0: api.app.v1.AppServiceErrorReason
}
var file_api_app_v1_app_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_app_v1_app_error_proto_init() }
func file_api_app_v1_app_error_proto_init() {
	if File_api_app_v1_app_error_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_app_v1_app_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_app_v1_app_error_proto_goTypes,
		DependencyIndexes: file_api_app_v1_app_error_proto_depIdxs,
		EnumInfos:         file_api_app_v1_app_error_proto_enumTypes,
	}.Build()
	File_api_app_v1_app_error_proto = out.File
	file_api_app_v1_app_error_proto_rawDesc = nil
	file_api_app_v1_app_error_proto_goTypes = nil
	file_api_app_v1_app_error_proto_depIdxs = nil
}

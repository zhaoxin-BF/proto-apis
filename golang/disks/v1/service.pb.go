// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: disks/v1/service.proto

package disks

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/fieldmaskpb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_disks_v1_service_proto protoreflect.FileDescriptor

var file_disks_v1_service_proto_rawDesc = []byte{
	0x0a, 0x16, 0x64, 0x69, 0x73, 0x6b, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x64, 0x69, 0x73, 0x6b, 0x73, 0x2e,
	0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61,
	0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x64, 0x69, 0x73, 0x6b, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x64, 0x69, 0x73, 0x6b, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xfe, 0x08,
	0x0a, 0x0c, 0x44, 0x69, 0x73, 0x6b, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xa4,
	0x01, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x69, 0x73, 0x6b, 0x73, 0x12, 0x1c,
	0x2e, 0x64, 0x69, 0x73, 0x6b, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x44, 0x69, 0x73, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x64,
	0x69, 0x73, 0x6b, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x69,
	0x73, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x58, 0x92, 0x41, 0x1b,
	0x12, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x64, 0x69, 0x73, 0x6b, 0x73, 0x2a, 0x0b,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x69, 0x73, 0x6b, 0x73, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x34, 0x3a, 0x04, 0x64, 0x69, 0x73, 0x6b, 0x22, 0x2c, 0x2f, 0x64, 0x69, 0x73, 0x6b, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x64, 0x69, 0x73, 0x6b, 0x2f, 0x7b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x3d,
	0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x2a, 0x7d, 0x12, 0x94, 0x01, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x69,
	0x73, 0x6b, 0x73, 0x12, 0x1a, 0x2e, 0x64, 0x69, 0x73, 0x6b, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x44, 0x69, 0x73, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1b, 0x2e, 0x64, 0x69, 0x73, 0x6b, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44,
	0x69, 0x73, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4e, 0x92, 0x41,
	0x17, 0x12, 0x0a, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x64, 0x69, 0x73, 0x6b, 0x73, 0x2a, 0x09, 0x4c,
	0x69, 0x73, 0x74, 0x44, 0x69, 0x73, 0x6b, 0x73, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2e, 0x12, 0x2c,
	0x2f, 0x64, 0x69, 0x73, 0x6b, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x69, 0x73, 0x6b, 0x2f, 0x7b,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x3d, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x73, 0x2f, 0x2a,
	0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x7d, 0x12, 0x90, 0x01, 0x0a,
	0x07, 0x47, 0x65, 0x74, 0x44, 0x69, 0x73, 0x6b, 0x12, 0x18, 0x2e, 0x64, 0x69, 0x73, 0x6b, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x69, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x19, 0x2e, 0x64, 0x69, 0x73, 0x6b, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x44, 0x69, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x50, 0x92,
	0x41, 0x13, 0x12, 0x08, 0x67, 0x65, 0x74, 0x20, 0x64, 0x69, 0x73, 0x6b, 0x2a, 0x07, 0x47, 0x65,
	0x74, 0x44, 0x69, 0x73, 0x6b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x34, 0x12, 0x32, 0x2f, 0x64, 0x69,
	0x73, 0x6b, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x69, 0x73, 0x6b, 0x2f, 0x7b, 0x6e, 0x61, 0x6d,
	0x65, 0x3d, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x64, 0x69, 0x73, 0x6b, 0x73, 0x2f, 0x2a, 0x7d, 0x12,
	0xaa, 0x01, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x69, 0x73, 0x6b, 0x12, 0x1b,
	0x2e, 0x64, 0x69, 0x73, 0x6b, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x44, 0x69, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x64, 0x69,
	0x73, 0x6b, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x69, 0x73,
	0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x61, 0x92, 0x41, 0x19, 0x12, 0x0b,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x64, 0x69, 0x73, 0x6b, 0x2a, 0x0a, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x44, 0x69, 0x73, 0x6b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x3f, 0x3a, 0x04, 0x64,
	0x69, 0x73, 0x6b, 0x32, 0x37, 0x2f, 0x64, 0x69, 0x73, 0x6b, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x64,
	0x69, 0x73, 0x6b, 0x2f, 0x7b, 0x64, 0x69, 0x73, 0x6b, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x74,
	0x65, 0x6e, 0x61, 0x6e, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x73, 0x2f, 0x2a, 0x2f, 0x64, 0x69, 0x73, 0x6b, 0x73, 0x2f, 0x2a, 0x7d, 0x12, 0xa1, 0x01, 0x0a,
	0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x69, 0x73, 0x6b, 0x12, 0x1c, 0x2e, 0x64, 0x69,
	0x73, 0x6b, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x69, 0x73,
	0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x64, 0x69, 0x73, 0x6b,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x69, 0x73, 0x6b, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x56, 0x92, 0x41, 0x19, 0x12, 0x0b, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x20, 0x64, 0x69, 0x73, 0x6b, 0x2a, 0x0a, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x44, 0x69, 0x73, 0x6b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x34, 0x2a, 0x32, 0x2f, 0x64,
	0x69, 0x73, 0x6b, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x69, 0x73, 0x6b, 0x2f, 0x7b, 0x6e, 0x61,
	0x6d, 0x65, 0x3d, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x64, 0x69, 0x73, 0x6b, 0x73, 0x2f, 0x2a, 0x7d,
	0x12, 0xa4, 0x01, 0x0a, 0x0a, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x44, 0x69, 0x73, 0x6b, 0x12,
	0x1b, 0x2e, 0x64, 0x69, 0x73, 0x6b, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x74, 0x74, 0x61, 0x63,
	0x68, 0x44, 0x69, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x64,
	0x69, 0x73, 0x6b, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x44, 0x69,
	0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5b, 0x92, 0x41, 0x19, 0x12,
	0x0b, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x20, 0x64, 0x69, 0x73, 0x6b, 0x2a, 0x0a, 0x41, 0x74,
	0x74, 0x61, 0x63, 0x68, 0x44, 0x69, 0x73, 0x6b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x39, 0x3a, 0x01,
	0x2a, 0x22, 0x34, 0x2f, 0x64, 0x69, 0x73, 0x6b, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x74, 0x74,
	0x61, 0x63, 0x68, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x73, 0x2f, 0x2a, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x64,
	0x69, 0x73, 0x6b, 0x73, 0x2f, 0x2a, 0x7d, 0x12, 0xa4, 0x01, 0x0a, 0x0a, 0x44, 0x65, 0x74, 0x61,
	0x63, 0x68, 0x44, 0x69, 0x73, 0x6b, 0x12, 0x1b, 0x2e, 0x64, 0x69, 0x73, 0x6b, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x65, 0x74, 0x61, 0x63, 0x68, 0x44, 0x69, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x64, 0x69, 0x73, 0x6b, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44,
	0x65, 0x74, 0x61, 0x63, 0x68, 0x44, 0x69, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x5b, 0x92, 0x41, 0x19, 0x12, 0x0b, 0x64, 0x65, 0x74, 0x61, 0x63, 0x68, 0x20, 0x64,
	0x69, 0x73, 0x6b, 0x2a, 0x0a, 0x44, 0x65, 0x74, 0x61, 0x63, 0x68, 0x44, 0x69, 0x73, 0x6b, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x39, 0x3a, 0x01, 0x2a, 0x22, 0x34, 0x2f, 0x64, 0x69, 0x73, 0x6b, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x74, 0x61, 0x63, 0x68, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65,
	0x3d, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x64, 0x69, 0x73, 0x6b, 0x73, 0x2f, 0x2a, 0x7d, 0x42, 0x38,
	0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a, 0x68, 0x61,
	0x6f, 0x78, 0x69, 0x6e, 0x2d, 0x42, 0x46, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x64, 0x69, 0x73, 0x6b, 0x73, 0x2f,
	0x76, 0x31, 0x3b, 0x64, 0x69, 0x73, 0x6b, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_disks_v1_service_proto_goTypes = []interface{}{
	(*CreateDisksRequest)(nil),  // 0: disks.v1.CreateDisksRequest
	(*ListDisksRequest)(nil),    // 1: disks.v1.ListDisksRequest
	(*GetDiskRequest)(nil),      // 2: disks.v1.GetDiskRequest
	(*UpdateDiskRequest)(nil),   // 3: disks.v1.UpdateDiskRequest
	(*DeleteDisksRequest)(nil),  // 4: disks.v1.DeleteDisksRequest
	(*AttachDiskRequest)(nil),   // 5: disks.v1.AttachDiskRequest
	(*DetachDiskRequest)(nil),   // 6: disks.v1.DetachDiskRequest
	(*CreateDisksResponse)(nil), // 7: disks.v1.CreateDisksResponse
	(*ListDisksResponse)(nil),   // 8: disks.v1.ListDisksResponse
	(*GetDiskResponse)(nil),     // 9: disks.v1.GetDiskResponse
	(*UpdateDiskResponse)(nil),  // 10: disks.v1.UpdateDiskResponse
	(*DeleteDisksResponse)(nil), // 11: disks.v1.DeleteDisksResponse
	(*AttachDiskResponse)(nil),  // 12: disks.v1.AttachDiskResponse
	(*DetachDiskResponse)(nil),  // 13: disks.v1.DetachDiskResponse
}
var file_disks_v1_service_proto_depIdxs = []int32{
	0,  // 0: disks.v1.DisksService.CreateDisks:input_type -> disks.v1.CreateDisksRequest
	1,  // 1: disks.v1.DisksService.ListDisks:input_type -> disks.v1.ListDisksRequest
	2,  // 2: disks.v1.DisksService.GetDisk:input_type -> disks.v1.GetDiskRequest
	3,  // 3: disks.v1.DisksService.UpdateDisk:input_type -> disks.v1.UpdateDiskRequest
	4,  // 4: disks.v1.DisksService.DeleteDisk:input_type -> disks.v1.DeleteDisksRequest
	5,  // 5: disks.v1.DisksService.AttachDisk:input_type -> disks.v1.AttachDiskRequest
	6,  // 6: disks.v1.DisksService.DetachDisk:input_type -> disks.v1.DetachDiskRequest
	7,  // 7: disks.v1.DisksService.CreateDisks:output_type -> disks.v1.CreateDisksResponse
	8,  // 8: disks.v1.DisksService.ListDisks:output_type -> disks.v1.ListDisksResponse
	9,  // 9: disks.v1.DisksService.GetDisk:output_type -> disks.v1.GetDiskResponse
	10, // 10: disks.v1.DisksService.UpdateDisk:output_type -> disks.v1.UpdateDiskResponse
	11, // 11: disks.v1.DisksService.DeleteDisk:output_type -> disks.v1.DeleteDisksResponse
	12, // 12: disks.v1.DisksService.AttachDisk:output_type -> disks.v1.AttachDiskResponse
	13, // 13: disks.v1.DisksService.DetachDisk:output_type -> disks.v1.DetachDiskResponse
	7,  // [7:14] is the sub-list for method output_type
	0,  // [0:7] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_disks_v1_service_proto_init() }
func file_disks_v1_service_proto_init() {
	if File_disks_v1_service_proto != nil {
		return
	}
	file_disks_v1_disks_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_disks_v1_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_disks_v1_service_proto_goTypes,
		DependencyIndexes: file_disks_v1_service_proto_depIdxs,
	}.Build()
	File_disks_v1_service_proto = out.File
	file_disks_v1_service_proto_rawDesc = nil
	file_disks_v1_service_proto_goTypes = nil
	file_disks_v1_service_proto_depIdxs = nil
}

// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: v1/variable.proto

package proto

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

type VariableAncestorsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// DCID of a stat var (group)
	Node string `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
}

func (x *VariableAncestorsRequest) Reset() {
	*x = VariableAncestorsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_variable_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VariableAncestorsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VariableAncestorsRequest) ProtoMessage() {}

func (x *VariableAncestorsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_variable_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VariableAncestorsRequest.ProtoReflect.Descriptor instead.
func (*VariableAncestorsRequest) Descriptor() ([]byte, []int) {
	return file_v1_variable_proto_rawDescGZIP(), []int{0}
}

func (x *VariableAncestorsRequest) GetNode() string {
	if x != nil {
		return x.Node
	}
	return ""
}

type VariableAncestorsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A list of ancestor stat var (group) DCIDs from the queried stat var to the
	// root of stat var hierarchy.
	//
	// This is one of many combinations in the hierarchy. When there are multiple
	// ancestor stat var (groups) to select at a level, the first one in
	// alphabetical sorted order is selected.
	Ancestors []string `protobuf:"bytes,1,rep,name=ancestors,proto3" json:"ancestors,omitempty"`
}

func (x *VariableAncestorsResponse) Reset() {
	*x = VariableAncestorsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_variable_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VariableAncestorsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VariableAncestorsResponse) ProtoMessage() {}

func (x *VariableAncestorsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_variable_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VariableAncestorsResponse.ProtoReflect.Descriptor instead.
func (*VariableAncestorsResponse) Descriptor() ([]byte, []int) {
	return file_v1_variable_proto_rawDescGZIP(), []int{1}
}

func (x *VariableAncestorsResponse) GetAncestors() []string {
	if x != nil {
		return x.Ancestors
	}
	return nil
}

var File_v1_variable_proto protoreflect.FileDescriptor

var file_v1_variable_proto_rawDesc = []byte{
	0x0a, 0x11, 0x76, 0x31, 0x2f, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73,
	0x2e, 0x76, 0x31, 0x22, 0x2e, 0x0a, 0x18, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x41,
	0x6e, 0x63, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x6f, 0x64, 0x65, 0x22, 0x39, 0x0a, 0x19, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x41,
	0x6e, 0x63, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x09, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x73, 0x42, 0x09,
	0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_v1_variable_proto_rawDescOnce sync.Once
	file_v1_variable_proto_rawDescData = file_v1_variable_proto_rawDesc
)

func file_v1_variable_proto_rawDescGZIP() []byte {
	file_v1_variable_proto_rawDescOnce.Do(func() {
		file_v1_variable_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_variable_proto_rawDescData)
	})
	return file_v1_variable_proto_rawDescData
}

var file_v1_variable_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_v1_variable_proto_goTypes = []interface{}{
	(*VariableAncestorsRequest)(nil),  // 0: datacommons.v1.VariableAncestorsRequest
	(*VariableAncestorsResponse)(nil), // 1: datacommons.v1.VariableAncestorsResponse
}
var file_v1_variable_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_v1_variable_proto_init() }
func file_v1_variable_proto_init() {
	if File_v1_variable_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_variable_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VariableAncestorsRequest); i {
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
		file_v1_variable_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VariableAncestorsResponse); i {
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
			RawDescriptor: file_v1_variable_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_variable_proto_goTypes,
		DependencyIndexes: file_v1_variable_proto_depIdxs,
		MessageInfos:      file_v1_variable_proto_msgTypes,
	}.Build()
	File_v1_variable_proto = out.File
	file_v1_variable_proto_rawDesc = nil
	file_v1_variable_proto_goTypes = nil
	file_v1_variable_proto_depIdxs = nil
}

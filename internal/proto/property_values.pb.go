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
// source: v1/property_values.proto

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

// A page of entities. The page number starts from 0, and is in the cache key.
// Page size is set by ::datacommons::prophet::kPageSize.
type PagedEntities struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A list of EntityInfo messages for PagedPropVal{In|Out} cache result.
	Entities       []*EntityInfo `protobuf:"bytes,1,rep,name=entities,proto3" json:"entities,omitempty"`
	TotalPageCount float64       `protobuf:"fixed64,2,opt,name=total_page_count,json=totalPageCount,proto3" json:"total_page_count,omitempty"`
}

func (x *PagedEntities) Reset() {
	*x = PagedEntities{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_property_values_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PagedEntities) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PagedEntities) ProtoMessage() {}

func (x *PagedEntities) ProtoReflect() protoreflect.Message {
	mi := &file_v1_property_values_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PagedEntities.ProtoReflect.Descriptor instead.
func (*PagedEntities) Descriptor() ([]byte, []int) {
	return file_v1_property_values_proto_rawDescGZIP(), []int{0}
}

func (x *PagedEntities) GetEntities() []*EntityInfo {
	if x != nil {
		return x.Entities
	}
	return nil
}

func (x *PagedEntities) GetTotalPageCount() float64 {
	if x != nil {
		return x.TotalPageCount
	}
	return 0
}

type InPropertyValuesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Property string `protobuf:"bytes,1,opt,name=property,proto3" json:"property,omitempty"`
	Entity   string `protobuf:"bytes,2,opt,name=entity,proto3" json:"entity,omitempty"`
	// [Optional]
	// The limit of the number of values to return. The maximium limit is 1000.
	// If not specified, the default limit is 1000.
	Limit int32 `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	// [Optional]
	// The pagination token for getting the next set of entries. This is empty
	// for the first request and needs to be set in the subsequent request.
	// This is the value returned from a prior call to InPropertyValuesResponse
	NextToken string `protobuf:"bytes,4,opt,name=next_token,json=nextToken,proto3" json:"next_token,omitempty"`
}

func (x *InPropertyValuesRequest) Reset() {
	*x = InPropertyValuesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_property_values_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InPropertyValuesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InPropertyValuesRequest) ProtoMessage() {}

func (x *InPropertyValuesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_property_values_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InPropertyValuesRequest.ProtoReflect.Descriptor instead.
func (*InPropertyValuesRequest) Descriptor() ([]byte, []int) {
	return file_v1_property_values_proto_rawDescGZIP(), []int{1}
}

func (x *InPropertyValuesRequest) GetProperty() string {
	if x != nil {
		return x.Property
	}
	return ""
}

func (x *InPropertyValuesRequest) GetEntity() string {
	if x != nil {
		return x.Entity
	}
	return ""
}

func (x *InPropertyValuesRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *InPropertyValuesRequest) GetNextToken() string {
	if x != nil {
		return x.NextToken
	}
	return ""
}

type InPropertyValuesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*EntityInfo `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	// The pagination token for getting the next set of entries.
	NextToken string `protobuf:"bytes,2,opt,name=next_token,json=nextToken,proto3" json:"next_token,omitempty"`
}

func (x *InPropertyValuesResponse) Reset() {
	*x = InPropertyValuesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_property_values_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InPropertyValuesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InPropertyValuesResponse) ProtoMessage() {}

func (x *InPropertyValuesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_property_values_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InPropertyValuesResponse.ProtoReflect.Descriptor instead.
func (*InPropertyValuesResponse) Descriptor() ([]byte, []int) {
	return file_v1_property_values_proto_rawDescGZIP(), []int{2}
}

func (x *InPropertyValuesResponse) GetData() []*EntityInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *InPropertyValuesResponse) GetNextToken() string {
	if x != nil {
		return x.NextToken
	}
	return ""
}

var File_v1_property_values_proto protoreflect.FileDescriptor

var file_v1_property_values_proto_rawDesc = []byte{
	0x0a, 0x18, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x5f, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x64, 0x61, 0x74, 0x61,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x0c, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6e, 0x0a, 0x0d, 0x50, 0x61, 0x67, 0x65,
	0x64, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x33, 0x0a, 0x08, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x28,
	0x0a, 0x10, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0e, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50,
	0x61, 0x67, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x82, 0x01, 0x0a, 0x17, 0x49, 0x6e, 0x50,
	0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79,
	0x12, 0x16, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x6e, 0x65, 0x78, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x66, 0x0a,
	0x18, 0x49, 0x6e, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x65, 0x78, 0x74,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_property_values_proto_rawDescOnce sync.Once
	file_v1_property_values_proto_rawDescData = file_v1_property_values_proto_rawDesc
)

func file_v1_property_values_proto_rawDescGZIP() []byte {
	file_v1_property_values_proto_rawDescOnce.Do(func() {
		file_v1_property_values_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_property_values_proto_rawDescData)
	})
	return file_v1_property_values_proto_rawDescData
}

var file_v1_property_values_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_v1_property_values_proto_goTypes = []interface{}{
	(*PagedEntities)(nil),            // 0: datacommons.v1.PagedEntities
	(*InPropertyValuesRequest)(nil),  // 1: datacommons.v1.InPropertyValuesRequest
	(*InPropertyValuesResponse)(nil), // 2: datacommons.v1.InPropertyValuesResponse
	(*EntityInfo)(nil),               // 3: datacommons.EntityInfo
}
var file_v1_property_values_proto_depIdxs = []int32{
	3, // 0: datacommons.v1.PagedEntities.entities:type_name -> datacommons.EntityInfo
	3, // 1: datacommons.v1.InPropertyValuesResponse.data:type_name -> datacommons.EntityInfo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_v1_property_values_proto_init() }
func file_v1_property_values_proto_init() {
	if File_v1_property_values_proto != nil {
		return
	}
	file_entity_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_v1_property_values_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PagedEntities); i {
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
		file_v1_property_values_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InPropertyValuesRequest); i {
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
		file_v1_property_values_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InPropertyValuesResponse); i {
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
			RawDescriptor: file_v1_property_values_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_property_values_proto_goTypes,
		DependencyIndexes: file_v1_property_values_proto_depIdxs,
		MessageInfos:      file_v1_property_values_proto_msgTypes,
	}.Build()
	File_v1_property_values_proto = out.File
	file_v1_property_values_proto_rawDesc = nil
	file_v1_property_values_proto_goTypes = nil
	file_v1_property_values_proto_depIdxs = nil
}

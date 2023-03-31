// Copyright 2023 Google LLC
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
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: v1/observation_existence.proto

package v1

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

type ExistenceByEntity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Key is entity DCID, and value specifies whether observation exists for this
	// entity.
	Entity map[string]bool `protobuf:"bytes,1,rep,name=entity,proto3" json:"entity,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *ExistenceByEntity) Reset() {
	*x = ExistenceByEntity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_observation_existence_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExistenceByEntity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExistenceByEntity) ProtoMessage() {}

func (x *ExistenceByEntity) ProtoReflect() protoreflect.Message {
	mi := &file_v1_observation_existence_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExistenceByEntity.ProtoReflect.Descriptor instead.
func (*ExistenceByEntity) Descriptor() ([]byte, []int) {
	return file_v1_observation_existence_proto_rawDescGZIP(), []int{0}
}

func (x *ExistenceByEntity) GetEntity() map[string]bool {
	if x != nil {
		return x.Entity
	}
	return nil
}

type BulkObservationExistenceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Variables to query for
	Variables []string `protobuf:"bytes,1,rep,name=variables,proto3" json:"variables,omitempty"`
	// Entities to query for
	Entities []string `protobuf:"bytes,2,rep,name=entities,proto3" json:"entities,omitempty"`
}

func (x *BulkObservationExistenceRequest) Reset() {
	*x = BulkObservationExistenceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_observation_existence_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BulkObservationExistenceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BulkObservationExistenceRequest) ProtoMessage() {}

func (x *BulkObservationExistenceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_observation_existence_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BulkObservationExistenceRequest.ProtoReflect.Descriptor instead.
func (*BulkObservationExistenceRequest) Descriptor() ([]byte, []int) {
	return file_v1_observation_existence_proto_rawDescGZIP(), []int{1}
}

func (x *BulkObservationExistenceRequest) GetVariables() []string {
	if x != nil {
		return x.Variables
	}
	return nil
}

func (x *BulkObservationExistenceRequest) GetEntities() []string {
	if x != nil {
		return x.Entities
	}
	return nil
}

type BulkObservationExistenceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Key is variable DCID
	Variable map[string]*ExistenceByEntity `protobuf:"bytes,2,rep,name=variable,proto3" json:"variable,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *BulkObservationExistenceResponse) Reset() {
	*x = BulkObservationExistenceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_observation_existence_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BulkObservationExistenceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BulkObservationExistenceResponse) ProtoMessage() {}

func (x *BulkObservationExistenceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_observation_existence_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BulkObservationExistenceResponse.ProtoReflect.Descriptor instead.
func (*BulkObservationExistenceResponse) Descriptor() ([]byte, []int) {
	return file_v1_observation_existence_proto_rawDescGZIP(), []int{2}
}

func (x *BulkObservationExistenceResponse) GetVariable() map[string]*ExistenceByEntity {
	if x != nil {
		return x.Variable
	}
	return nil
}

var File_v1_observation_existence_proto protoreflect.FileDescriptor

var file_v1_observation_existence_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x76, 0x31, 0x2f, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x65, 0x78, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31,
	0x22, 0x95, 0x01, 0x0a, 0x11, 0x45, 0x78, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x42, 0x79,
	0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x45, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63,
	0x65, 0x42, 0x79, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x1a, 0x39, 0x0a,
	0x0b, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x5b, 0x0a, 0x1f, 0x42, 0x75, 0x6c, 0x6b,
	0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x78, 0x69, 0x73, 0x74,
	0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x76,
	0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09,
	0x76, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x69, 0x65, 0x73, 0x22, 0xde, 0x01, 0x0a, 0x20, 0x42, 0x75, 0x6c, 0x6b, 0x4f, 0x62,
	0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x78, 0x69, 0x73, 0x74, 0x65, 0x6e,
	0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5a, 0x0a, 0x08, 0x76, 0x61,
	0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3e, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x75,
	0x6c, 0x6b, 0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x78, 0x69,
	0x73, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x56,
	0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x76, 0x61,
	0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x1a, 0x5e, 0x0a, 0x0d, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62,
	0x6c, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x37, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x69, 0x73, 0x74, 0x65,
	0x6e, 0x63, 0x65, 0x42, 0x79, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73,
	0x6f, 0x72, 0x67, 0x2f, 0x6d, 0x69, 0x78, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_v1_observation_existence_proto_rawDescOnce sync.Once
	file_v1_observation_existence_proto_rawDescData = file_v1_observation_existence_proto_rawDesc
)

func file_v1_observation_existence_proto_rawDescGZIP() []byte {
	file_v1_observation_existence_proto_rawDescOnce.Do(func() {
		file_v1_observation_existence_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_observation_existence_proto_rawDescData)
	})
	return file_v1_observation_existence_proto_rawDescData
}

var file_v1_observation_existence_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_v1_observation_existence_proto_goTypes = []interface{}{
	(*ExistenceByEntity)(nil),                // 0: datacommons.v1.ExistenceByEntity
	(*BulkObservationExistenceRequest)(nil),  // 1: datacommons.v1.BulkObservationExistenceRequest
	(*BulkObservationExistenceResponse)(nil), // 2: datacommons.v1.BulkObservationExistenceResponse
	nil,                                      // 3: datacommons.v1.ExistenceByEntity.EntityEntry
	nil,                                      // 4: datacommons.v1.BulkObservationExistenceResponse.VariableEntry
}
var file_v1_observation_existence_proto_depIdxs = []int32{
	3, // 0: datacommons.v1.ExistenceByEntity.entity:type_name -> datacommons.v1.ExistenceByEntity.EntityEntry
	4, // 1: datacommons.v1.BulkObservationExistenceResponse.variable:type_name -> datacommons.v1.BulkObservationExistenceResponse.VariableEntry
	0, // 2: datacommons.v1.BulkObservationExistenceResponse.VariableEntry.value:type_name -> datacommons.v1.ExistenceByEntity
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_v1_observation_existence_proto_init() }
func file_v1_observation_existence_proto_init() {
	if File_v1_observation_existence_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_observation_existence_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExistenceByEntity); i {
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
		file_v1_observation_existence_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BulkObservationExistenceRequest); i {
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
		file_v1_observation_existence_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BulkObservationExistenceResponse); i {
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
			RawDescriptor: file_v1_observation_existence_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_observation_existence_proto_goTypes,
		DependencyIndexes: file_v1_observation_existence_proto_depIdxs,
		MessageInfos:      file_v1_observation_existence_proto_msgTypes,
	}.Build()
	File_v1_observation_existence_proto = out.File
	file_v1_observation_existence_proto_rawDesc = nil
	file_v1_observation_existence_proto_goTypes = nil
	file_v1_observation_existence_proto_depIdxs = nil
}
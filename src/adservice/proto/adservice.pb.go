// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.20.0
// source: proto/adservice.proto

package hipstershop

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

type AdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of important key words from the current page describing the context.
	ContextKeys []string `protobuf:"bytes,1,rep,name=context_keys,json=contextKeys,proto3" json:"context_keys,omitempty"`
}

func (x *AdRequest) Reset() {
	*x = AdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_adservice_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdRequest) ProtoMessage() {}

func (x *AdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_adservice_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdRequest.ProtoReflect.Descriptor instead.
func (*AdRequest) Descriptor() ([]byte, []int) {
	return file_proto_adservice_proto_rawDescGZIP(), []int{0}
}

func (x *AdRequest) GetContextKeys() []string {
	if x != nil {
		return x.ContextKeys
	}
	return nil
}

type AdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ads []*Ad `protobuf:"bytes,1,rep,name=ads,proto3" json:"ads,omitempty"`
}

func (x *AdResponse) Reset() {
	*x = AdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_adservice_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdResponse) ProtoMessage() {}

func (x *AdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_adservice_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdResponse.ProtoReflect.Descriptor instead.
func (*AdResponse) Descriptor() ([]byte, []int) {
	return file_proto_adservice_proto_rawDescGZIP(), []int{1}
}

func (x *AdResponse) GetAds() []*Ad {
	if x != nil {
		return x.Ads
	}
	return nil
}

type Ad struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// url to redirect to when an ad is clicked.
	RedirectUrl string `protobuf:"bytes,1,opt,name=redirect_url,json=redirectUrl,proto3" json:"redirect_url,omitempty"`
	// short advertisement text to display.
	Text string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *Ad) Reset() {
	*x = Ad{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_adservice_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ad) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ad) ProtoMessage() {}

func (x *Ad) ProtoReflect() protoreflect.Message {
	mi := &file_proto_adservice_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ad.ProtoReflect.Descriptor instead.
func (*Ad) Descriptor() ([]byte, []int) {
	return file_proto_adservice_proto_rawDescGZIP(), []int{2}
}

func (x *Ad) GetRedirectUrl() string {
	if x != nil {
		return x.RedirectUrl
	}
	return ""
}

func (x *Ad) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

var File_proto_adservice_proto protoreflect.FileDescriptor

var file_proto_adservice_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x64, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x68, 0x69, 0x70, 0x73, 0x74, 0x65, 0x72,
	0x73, 0x68, 0x6f, 0x70, 0x22, 0x2e, 0x0a, 0x09, 0x41, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x6b, 0x65, 0x79,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74,
	0x4b, 0x65, 0x79, 0x73, 0x22, 0x2f, 0x0a, 0x0a, 0x41, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x21, 0x0a, 0x03, 0x61, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x68, 0x69, 0x70, 0x73, 0x74, 0x65, 0x72, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x41, 0x64,
	0x52, 0x03, 0x61, 0x64, 0x73, 0x22, 0x3b, 0x0a, 0x02, 0x41, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x72,
	0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x55, 0x72, 0x6c, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65,
	0x78, 0x74, 0x32, 0x48, 0x0a, 0x09, 0x41, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x3b, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x64, 0x73, 0x12, 0x16, 0x2e, 0x68, 0x69, 0x70, 0x73,
	0x74, 0x65, 0x72, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x41, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x17, 0x2e, 0x68, 0x69, 0x70, 0x73, 0x74, 0x65, 0x72, 0x73, 0x68, 0x6f, 0x70, 0x2e,
	0x41, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x15, 0x5a, 0x13,
	0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x68, 0x69, 0x70, 0x73, 0x74, 0x65, 0x72, 0x73,
	0x68, 0x6f, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_adservice_proto_rawDescOnce sync.Once
	file_proto_adservice_proto_rawDescData = file_proto_adservice_proto_rawDesc
)

func file_proto_adservice_proto_rawDescGZIP() []byte {
	file_proto_adservice_proto_rawDescOnce.Do(func() {
		file_proto_adservice_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_adservice_proto_rawDescData)
	})
	return file_proto_adservice_proto_rawDescData
}

var file_proto_adservice_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_adservice_proto_goTypes = []interface{}{
	(*AdRequest)(nil),  // 0: hipstershop.AdRequest
	(*AdResponse)(nil), // 1: hipstershop.AdResponse
	(*Ad)(nil),         // 2: hipstershop.Ad
}
var file_proto_adservice_proto_depIdxs = []int32{
	2, // 0: hipstershop.AdResponse.ads:type_name -> hipstershop.Ad
	0, // 1: hipstershop.AdService.GetAds:input_type -> hipstershop.AdRequest
	1, // 2: hipstershop.AdService.GetAds:output_type -> hipstershop.AdResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_adservice_proto_init() }
func file_proto_adservice_proto_init() {
	if File_proto_adservice_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_adservice_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdRequest); i {
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
		file_proto_adservice_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdResponse); i {
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
		file_proto_adservice_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ad); i {
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
			RawDescriptor: file_proto_adservice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_adservice_proto_goTypes,
		DependencyIndexes: file_proto_adservice_proto_depIdxs,
		MessageInfos:      file_proto_adservice_proto_msgTypes,
	}.Build()
	File_proto_adservice_proto = out.File
	file_proto_adservice_proto_rawDesc = nil
	file_proto_adservice_proto_goTypes = nil
	file_proto_adservice_proto_depIdxs = nil
}

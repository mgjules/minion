// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: words/words.proto

package words

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type AddWordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Word string `protobuf:"bytes,1,opt,name=word,proto3" json:"word,omitempty"`
}

func (x *AddWordRequest) Reset() {
	*x = AddWordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_words_words_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddWordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddWordRequest) ProtoMessage() {}

func (x *AddWordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_words_words_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddWordRequest.ProtoReflect.Descriptor instead.
func (*AddWordRequest) Descriptor() ([]byte, []int) {
	return file_words_words_proto_rawDescGZIP(), []int{0}
}

func (x *AddWordRequest) GetWord() string {
	if x != nil {
		return x.Word
	}
	return ""
}

type AddWordResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Word string `protobuf:"bytes,2,opt,name=word,proto3" json:"word,omitempty"`
}

func (x *AddWordResponse) Reset() {
	*x = AddWordResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_words_words_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddWordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddWordResponse) ProtoMessage() {}

func (x *AddWordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_words_words_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddWordResponse.ProtoReflect.Descriptor instead.
func (*AddWordResponse) Descriptor() ([]byte, []int) {
	return file_words_words_proto_rawDescGZIP(), []int{1}
}

func (x *AddWordResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AddWordResponse) GetWord() string {
	if x != nil {
		return x.Word
	}
	return ""
}

type RandomWordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RandomWordRequest) Reset() {
	*x = RandomWordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_words_words_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RandomWordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RandomWordRequest) ProtoMessage() {}

func (x *RandomWordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_words_words_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RandomWordRequest.ProtoReflect.Descriptor instead.
func (*RandomWordRequest) Descriptor() ([]byte, []int) {
	return file_words_words_proto_rawDescGZIP(), []int{2}
}

type RandomWordResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Word string `protobuf:"bytes,2,opt,name=word,proto3" json:"word,omitempty"`
}

func (x *RandomWordResponse) Reset() {
	*x = RandomWordResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_words_words_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RandomWordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RandomWordResponse) ProtoMessage() {}

func (x *RandomWordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_words_words_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RandomWordResponse.ProtoReflect.Descriptor instead.
func (*RandomWordResponse) Descriptor() ([]byte, []int) {
	return file_words_words_proto_rawDescGZIP(), []int{3}
}

func (x *RandomWordResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *RandomWordResponse) GetWord() string {
	if x != nil {
		return x.Word
	}
	return ""
}

type SearchWordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query string `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
}

func (x *SearchWordRequest) Reset() {
	*x = SearchWordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_words_words_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchWordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchWordRequest) ProtoMessage() {}

func (x *SearchWordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_words_words_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchWordRequest.ProtoReflect.Descriptor instead.
func (*SearchWordRequest) Descriptor() ([]byte, []int) {
	return file_words_words_proto_rawDescGZIP(), []int{4}
}

func (x *SearchWordRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

type SearchWordResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Words []string `protobuf:"bytes,1,rep,name=words,proto3" json:"words,omitempty"`
}

func (x *SearchWordResponse) Reset() {
	*x = SearchWordResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_words_words_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchWordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchWordResponse) ProtoMessage() {}

func (x *SearchWordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_words_words_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchWordResponse.ProtoReflect.Descriptor instead.
func (*SearchWordResponse) Descriptor() ([]byte, []int) {
	return file_words_words_proto_rawDescGZIP(), []int{5}
}

func (x *SearchWordResponse) GetWords() []string {
	if x != nil {
		return x.Words
	}
	return nil
}

type HealthCheckRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HealthCheckRequest) Reset() {
	*x = HealthCheckRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_words_words_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthCheckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthCheckRequest) ProtoMessage() {}

func (x *HealthCheckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_words_words_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthCheckRequest.ProtoReflect.Descriptor instead.
func (*HealthCheckRequest) Descriptor() ([]byte, []int) {
	return file_words_words_proto_rawDescGZIP(), []int{6}
}

type HealthCheckResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HealthCheckResponse) Reset() {
	*x = HealthCheckResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_words_words_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthCheckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthCheckResponse) ProtoMessage() {}

func (x *HealthCheckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_words_words_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthCheckResponse.ProtoReflect.Descriptor instead.
func (*HealthCheckResponse) Descriptor() ([]byte, []int) {
	return file_words_words_proto_rawDescGZIP(), []int{7}
}

var File_words_words_proto protoreflect.FileDescriptor

var file_words_words_proto_rawDesc = []byte{
	0x0a, 0x11, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x2f, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x24, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x57,
	0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x35,
	0x0a, 0x0f, 0x41, 0x64, 0x64, 0x57, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x13, 0x0a, 0x11, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x57,
	0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x38, 0x0a, 0x12, 0x52, 0x61,
	0x6e, 0x64, 0x6f, 0x6d, 0x57, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x77, 0x6f, 0x72, 0x64, 0x22, 0x29, 0x0a, 0x11, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x57, 0x6f,
	0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x22,
	0x2a, 0x0a, 0x12, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x57, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x22, 0x14, 0x0a, 0x12, 0x48,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x15, 0x0a, 0x13, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xc9, 0x03, 0x0a, 0x0c, 0x57, 0x6f, 0x72,
	0x64, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6c, 0x0a, 0x07, 0x41, 0x64, 0x64,
	0x57, 0x6f, 0x72, 0x64, 0x12, 0x15, 0x2e, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x2e, 0x41, 0x64, 0x64,
	0x57, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x77, 0x6f,
	0x72, 0x64, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x57, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x32, 0x92, 0x41, 0x1b, 0x12, 0x08, 0x41, 0x64, 0x64, 0x20, 0x77, 0x6f,
	0x72, 0x64, 0x1a, 0x0f, 0x41, 0x64, 0x64, 0x20, 0x61, 0x20, 0x6e, 0x65, 0x77, 0x20, 0x77, 0x6f,
	0x72, 0x64, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x22, 0x09, 0x2f, 0x76, 0x31, 0x2f, 0x77,
	0x6f, 0x72, 0x64, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x8f, 0x01, 0x0a, 0x0a, 0x52, 0x61, 0x6e, 0x64,
	0x6f, 0x6d, 0x57, 0x6f, 0x72, 0x64, 0x12, 0x18, 0x2e, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x2e, 0x52,
	0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x57, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x19, 0x2e, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x2e, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x57,
	0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4c, 0x92, 0x41, 0x31,
	0x12, 0x0b, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x20, 0x77, 0x6f, 0x72, 0x64, 0x1a, 0x22, 0x52,
	0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x20, 0x61, 0x20, 0x72, 0x61, 0x6e, 0x64, 0x6f, 0x6d,
	0x6c, 0x79, 0x20, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x20, 0x77, 0x6f, 0x72, 0x64,
	0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x12, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x6f, 0x72,
	0x64, 0x73, 0x2f, 0x72, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x12, 0xb8, 0x01, 0x0a, 0x0a, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x57, 0x6f, 0x72, 0x64, 0x12, 0x18, 0x2e, 0x77, 0x6f, 0x72, 0x64, 0x73,
	0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x57, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x19, 0x2e, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x57, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x75, 0x92,
	0x41, 0x5a, 0x12, 0x0b, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x20, 0x77, 0x6f, 0x72, 0x64, 0x1a,
	0x4b, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x61, 0x6c, 0x6c, 0x20,
	0x77, 0x6f, 0x72, 0x64, 0x73, 0x20, 0x74, 0x68, 0x61, 0x74, 0x20, 0x68, 0x61, 0x76, 0x65, 0x20,
	0x74, 0x68, 0x65, 0x20, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x20, 0x6f, 0x66, 0x20, 0x74,
	0x68, 0x65, 0x20, 0x71, 0x75, 0x65, 0x72, 0x79, 0x20, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x20, 0x61, 0x73, 0x20, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x2e, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x12, 0x12, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x2f, 0x73, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x42, 0x86, 0x01, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x67, 0x6a, 0x75, 0x6c, 0x65, 0x73, 0x2f, 0x6d, 0x69, 0x6e, 0x69,
	0x6f, 0x6e, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x92, 0x41, 0x50, 0x12, 0x4e, 0x0a,
	0x05, 0x57, 0x6f, 0x72, 0x64, 0x73, 0x22, 0x3e, 0x0a, 0x17, 0x4d, 0x69, 0x63, 0x68, 0x61, 0xc3,
	0xab, 0x6c, 0x20, 0x47, 0x69, 0x6f, 0x76, 0x61, 0x6e, 0x6e, 0x69, 0x20, 0x4a, 0x75, 0x6c, 0x65,
	0x73, 0x12, 0x13, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x6d, 0x67, 0x6a, 0x75, 0x6c,
	0x65, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x1a, 0x0e, 0x68, 0x69, 0x40, 0x6d, 0x67, 0x6a, 0x75, 0x6c,
	0x65, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x32, 0x05, 0x30, 0x2e, 0x31, 0x2e, 0x30, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_words_words_proto_rawDescOnce sync.Once
	file_words_words_proto_rawDescData = file_words_words_proto_rawDesc
)

func file_words_words_proto_rawDescGZIP() []byte {
	file_words_words_proto_rawDescOnce.Do(func() {
		file_words_words_proto_rawDescData = protoimpl.X.CompressGZIP(file_words_words_proto_rawDescData)
	})
	return file_words_words_proto_rawDescData
}

var file_words_words_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_words_words_proto_goTypes = []interface{}{
	(*AddWordRequest)(nil),      // 0: words.AddWordRequest
	(*AddWordResponse)(nil),     // 1: words.AddWordResponse
	(*RandomWordRequest)(nil),   // 2: words.RandomWordRequest
	(*RandomWordResponse)(nil),  // 3: words.RandomWordResponse
	(*SearchWordRequest)(nil),   // 4: words.SearchWordRequest
	(*SearchWordResponse)(nil),  // 5: words.SearchWordResponse
	(*HealthCheckRequest)(nil),  // 6: words.HealthCheckRequest
	(*HealthCheckResponse)(nil), // 7: words.HealthCheckResponse
}
var file_words_words_proto_depIdxs = []int32{
	0, // 0: words.WordsService.AddWord:input_type -> words.AddWordRequest
	2, // 1: words.WordsService.RandomWord:input_type -> words.RandomWordRequest
	4, // 2: words.WordsService.SearchWord:input_type -> words.SearchWordRequest
	1, // 3: words.WordsService.AddWord:output_type -> words.AddWordResponse
	3, // 4: words.WordsService.RandomWord:output_type -> words.RandomWordResponse
	5, // 5: words.WordsService.SearchWord:output_type -> words.SearchWordResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_words_words_proto_init() }
func file_words_words_proto_init() {
	if File_words_words_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_words_words_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddWordRequest); i {
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
		file_words_words_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddWordResponse); i {
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
		file_words_words_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RandomWordRequest); i {
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
		file_words_words_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RandomWordResponse); i {
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
		file_words_words_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchWordRequest); i {
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
		file_words_words_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchWordResponse); i {
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
		file_words_words_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthCheckRequest); i {
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
		file_words_words_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthCheckResponse); i {
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
			RawDescriptor: file_words_words_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_words_words_proto_goTypes,
		DependencyIndexes: file_words_words_proto_depIdxs,
		MessageInfos:      file_words_words_proto_msgTypes,
	}.Build()
	File_words_words_proto = out.File
	file_words_words_proto_rawDesc = nil
	file_words_words_proto_goTypes = nil
	file_words_words_proto_depIdxs = nil
}
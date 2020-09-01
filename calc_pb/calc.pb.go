// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: calc_pb/calc.proto

package calcpb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type SumRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number1 int32 `protobuf:"varint,1,opt,name=number1,proto3" json:"number1,omitempty"`
	Number2 int32 `protobuf:"varint,2,opt,name=number2,proto3" json:"number2,omitempty"`
}

func (x *SumRequest) Reset() {
	*x = SumRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calc_pb_calc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SumRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SumRequest) ProtoMessage() {}

func (x *SumRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calc_pb_calc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SumRequest.ProtoReflect.Descriptor instead.
func (*SumRequest) Descriptor() ([]byte, []int) {
	return file_calc_pb_calc_proto_rawDescGZIP(), []int{0}
}

func (x *SumRequest) GetNumber1() int32 {
	if x != nil {
		return x.Number1
	}
	return 0
}

func (x *SumRequest) GetNumber2() int32 {
	if x != nil {
		return x.Number2
	}
	return 0
}

type SumResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sum int64 `protobuf:"varint,1,opt,name=sum,proto3" json:"sum,omitempty"`
}

func (x *SumResponse) Reset() {
	*x = SumResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calc_pb_calc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SumResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SumResponse) ProtoMessage() {}

func (x *SumResponse) ProtoReflect() protoreflect.Message {
	mi := &file_calc_pb_calc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SumResponse.ProtoReflect.Descriptor instead.
func (*SumResponse) Descriptor() ([]byte, []int) {
	return file_calc_pb_calc_proto_rawDescGZIP(), []int{1}
}

func (x *SumResponse) GetSum() int64 {
	if x != nil {
		return x.Sum
	}
	return 0
}

type DiffRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number1 float32 `protobuf:"fixed32,1,opt,name=number1,proto3" json:"number1,omitempty"`
	Number2 float32 `protobuf:"fixed32,2,opt,name=number2,proto3" json:"number2,omitempty"`
}

func (x *DiffRequest) Reset() {
	*x = DiffRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calc_pb_calc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiffRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiffRequest) ProtoMessage() {}

func (x *DiffRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calc_pb_calc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiffRequest.ProtoReflect.Descriptor instead.
func (*DiffRequest) Descriptor() ([]byte, []int) {
	return file_calc_pb_calc_proto_rawDescGZIP(), []int{2}
}

func (x *DiffRequest) GetNumber1() float32 {
	if x != nil {
		return x.Number1
	}
	return 0
}

func (x *DiffRequest) GetNumber2() float32 {
	if x != nil {
		return x.Number2
	}
	return 0
}

type DiffResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Difference float32 `protobuf:"fixed32,1,opt,name=difference,proto3" json:"difference,omitempty"`
}

func (x *DiffResponse) Reset() {
	*x = DiffResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calc_pb_calc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiffResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiffResponse) ProtoMessage() {}

func (x *DiffResponse) ProtoReflect() protoreflect.Message {
	mi := &file_calc_pb_calc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiffResponse.ProtoReflect.Descriptor instead.
func (*DiffResponse) Descriptor() ([]byte, []int) {
	return file_calc_pb_calc_proto_rawDescGZIP(), []int{3}
}

func (x *DiffResponse) GetDifference() float32 {
	if x != nil {
		return x.Difference
	}
	return 0
}

type MultiplierRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number1 float32 `protobuf:"fixed32,1,opt,name=number1,proto3" json:"number1,omitempty"`
	Number2 float32 `protobuf:"fixed32,2,opt,name=number2,proto3" json:"number2,omitempty"`
	Number3 float32 `protobuf:"fixed32,3,opt,name=number3,proto3" json:"number3,omitempty"`
}

func (x *MultiplierRequest) Reset() {
	*x = MultiplierRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calc_pb_calc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultiplierRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultiplierRequest) ProtoMessage() {}

func (x *MultiplierRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calc_pb_calc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultiplierRequest.ProtoReflect.Descriptor instead.
func (*MultiplierRequest) Descriptor() ([]byte, []int) {
	return file_calc_pb_calc_proto_rawDescGZIP(), []int{4}
}

func (x *MultiplierRequest) GetNumber1() float32 {
	if x != nil {
		return x.Number1
	}
	return 0
}

func (x *MultiplierRequest) GetNumber2() float32 {
	if x != nil {
		return x.Number2
	}
	return 0
}

func (x *MultiplierRequest) GetNumber3() float32 {
	if x != nil {
		return x.Number3
	}
	return 0
}

type MultiplierResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resp float32 `protobuf:"fixed32,1,opt,name=resp,proto3" json:"resp,omitempty"`
}

func (x *MultiplierResponse) Reset() {
	*x = MultiplierResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calc_pb_calc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultiplierResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultiplierResponse) ProtoMessage() {}

func (x *MultiplierResponse) ProtoReflect() protoreflect.Message {
	mi := &file_calc_pb_calc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultiplierResponse.ProtoReflect.Descriptor instead.
func (*MultiplierResponse) Descriptor() ([]byte, []int) {
	return file_calc_pb_calc_proto_rawDescGZIP(), []int{5}
}

func (x *MultiplierResponse) GetResp() float32 {
	if x != nil {
		return x.Resp
	}
	return 0
}

type GigaByteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GbValue string `protobuf:"bytes,1,opt,name=gbValue,proto3" json:"gbValue,omitempty"`
}

func (x *GigaByteRequest) Reset() {
	*x = GigaByteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calc_pb_calc_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GigaByteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GigaByteRequest) ProtoMessage() {}

func (x *GigaByteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calc_pb_calc_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GigaByteRequest.ProtoReflect.Descriptor instead.
func (*GigaByteRequest) Descriptor() ([]byte, []int) {
	return file_calc_pb_calc_proto_rawDescGZIP(), []int{6}
}

func (x *GigaByteRequest) GetGbValue() string {
	if x != nil {
		return x.GbValue
	}
	return ""
}

type GigaByteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resp string `protobuf:"bytes,1,opt,name=resp,proto3" json:"resp,omitempty"`
}

func (x *GigaByteResponse) Reset() {
	*x = GigaByteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calc_pb_calc_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GigaByteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GigaByteResponse) ProtoMessage() {}

func (x *GigaByteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_calc_pb_calc_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GigaByteResponse.ProtoReflect.Descriptor instead.
func (*GigaByteResponse) Descriptor() ([]byte, []int) {
	return file_calc_pb_calc_proto_rawDescGZIP(), []int{7}
}

func (x *GigaByteResponse) GetResp() string {
	if x != nil {
		return x.Resp
	}
	return ""
}

var File_calc_pb_calc_proto protoreflect.FileDescriptor

var file_calc_pb_calc_proto_rawDesc = []byte{
	0x0a, 0x12, 0x63, 0x61, 0x6c, 0x63, 0x5f, 0x70, 0x62, 0x2f, 0x63, 0x61, 0x6c, 0x63, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63, 0x61, 0x6c, 0x63, 0x22, 0x40, 0x0a, 0x0a, 0x73, 0x75,
	0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x31, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x32, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x07, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x32, 0x22, 0x1f, 0x0a, 0x0b,
	0x73, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x73,
	0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x73, 0x75, 0x6d, 0x22, 0x41, 0x0a,
	0x0b, 0x64, 0x69, 0x66, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x31, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x32,
	0x22, 0x2e, 0x0a, 0x0c, 0x64, 0x69, 0x66, 0x66, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x69, 0x66, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x64, 0x69, 0x66, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x22, 0x61, 0x0a, 0x11, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x31,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x31, 0x12,
	0x18, 0x0a, 0x07, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x07, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x32, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x33, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x33, 0x22, 0x28, 0x0a, 0x12, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x69, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x65, 0x73,
	0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x72, 0x65, 0x73, 0x70, 0x22, 0x2b, 0x0a,
	0x0f, 0x67, 0x69, 0x67, 0x61, 0x42, 0x79, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x67, 0x62, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x67, 0x62, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x26, 0x0a, 0x10, 0x67, 0x69,
	0x67, 0x61, 0x42, 0x79, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x72, 0x65, 0x73, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x65,
	0x73, 0x70, 0x32, 0xfd, 0x01, 0x0a, 0x11, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2e, 0x0a, 0x05, 0x64, 0x6f, 0x53, 0x75,
	0x6d, 0x12, 0x10, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x2e, 0x73, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x2e, 0x73, 0x75, 0x6d, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x08, 0x63, 0x61, 0x6c, 0x63,
	0x44, 0x69, 0x66, 0x66, 0x12, 0x11, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x2e, 0x64, 0x69, 0x66, 0x66,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x2e, 0x64,
	0x69, 0x66, 0x66, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3f, 0x0a,
	0x08, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x79, 0x12, 0x17, 0x2e, 0x63, 0x61, 0x6c, 0x63,
	0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x18, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70,
	0x6c, 0x69, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x42,
	0x0a, 0x0d, 0x67, 0x65, 0x74, 0x42, 0x79, 0x74, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12,
	0x15, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x2e, 0x67, 0x69, 0x67, 0x61, 0x42, 0x79, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x2e, 0x67, 0x69,
	0x67, 0x61, 0x42, 0x79, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x30, 0x01, 0x42, 0x10, 0x5a, 0x0e, 0x63, 0x61, 0x6c, 0x63, 0x5f, 0x70, 0x62, 0x3b, 0x63, 0x61,
	0x6c, 0x63, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_calc_pb_calc_proto_rawDescOnce sync.Once
	file_calc_pb_calc_proto_rawDescData = file_calc_pb_calc_proto_rawDesc
)

func file_calc_pb_calc_proto_rawDescGZIP() []byte {
	file_calc_pb_calc_proto_rawDescOnce.Do(func() {
		file_calc_pb_calc_proto_rawDescData = protoimpl.X.CompressGZIP(file_calc_pb_calc_proto_rawDescData)
	})
	return file_calc_pb_calc_proto_rawDescData
}

var file_calc_pb_calc_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_calc_pb_calc_proto_goTypes = []interface{}{
	(*SumRequest)(nil),         // 0: calc.sumRequest
	(*SumResponse)(nil),        // 1: calc.sumResponse
	(*DiffRequest)(nil),        // 2: calc.diffRequest
	(*DiffResponse)(nil),       // 3: calc.diffResponse
	(*MultiplierRequest)(nil),  // 4: calc.multiplierRequest
	(*MultiplierResponse)(nil), // 5: calc.multiplierResponse
	(*GigaByteRequest)(nil),    // 6: calc.gigaByteRequest
	(*GigaByteResponse)(nil),   // 7: calc.gigaByteResponse
}
var file_calc_pb_calc_proto_depIdxs = []int32{
	0, // 0: calc.calculatorService.doSum:input_type -> calc.sumRequest
	2, // 1: calc.calculatorService.calcDiff:input_type -> calc.diffRequest
	4, // 2: calc.calculatorService.multiply:input_type -> calc.multiplierRequest
	6, // 3: calc.calculatorService.getByteValues:input_type -> calc.gigaByteRequest
	1, // 4: calc.calculatorService.doSum:output_type -> calc.sumResponse
	3, // 5: calc.calculatorService.calcDiff:output_type -> calc.diffResponse
	5, // 6: calc.calculatorService.multiply:output_type -> calc.multiplierResponse
	7, // 7: calc.calculatorService.getByteValues:output_type -> calc.gigaByteResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_calc_pb_calc_proto_init() }
func file_calc_pb_calc_proto_init() {
	if File_calc_pb_calc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_calc_pb_calc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SumRequest); i {
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
		file_calc_pb_calc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SumResponse); i {
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
		file_calc_pb_calc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DiffRequest); i {
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
		file_calc_pb_calc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DiffResponse); i {
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
		file_calc_pb_calc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MultiplierRequest); i {
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
		file_calc_pb_calc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MultiplierResponse); i {
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
		file_calc_pb_calc_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GigaByteRequest); i {
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
		file_calc_pb_calc_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GigaByteResponse); i {
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
			RawDescriptor: file_calc_pb_calc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_calc_pb_calc_proto_goTypes,
		DependencyIndexes: file_calc_pb_calc_proto_depIdxs,
		MessageInfos:      file_calc_pb_calc_proto_msgTypes,
	}.Build()
	File_calc_pb_calc_proto = out.File
	file_calc_pb_calc_proto_rawDesc = nil
	file_calc_pb_calc_proto_goTypes = nil
	file_calc_pb_calc_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CalculatorServiceClient is the client API for CalculatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalculatorServiceClient interface {
	// Unary
	DoSum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error)
	CalcDiff(ctx context.Context, in *DiffRequest, opts ...grpc.CallOption) (*DiffResponse, error)
	Multiply(ctx context.Context, in *MultiplierRequest, opts ...grpc.CallOption) (*MultiplierResponse, error)
	// Server Streaming
	GetByteValues(ctx context.Context, in *GigaByteRequest, opts ...grpc.CallOption) (CalculatorService_GetByteValuesClient, error)
}

type calculatorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCalculatorServiceClient(cc grpc.ClientConnInterface) CalculatorServiceClient {
	return &calculatorServiceClient{cc}
}

func (c *calculatorServiceClient) DoSum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error) {
	out := new(SumResponse)
	err := c.cc.Invoke(ctx, "/calc.calculatorService/doSum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorServiceClient) CalcDiff(ctx context.Context, in *DiffRequest, opts ...grpc.CallOption) (*DiffResponse, error) {
	out := new(DiffResponse)
	err := c.cc.Invoke(ctx, "/calc.calculatorService/calcDiff", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorServiceClient) Multiply(ctx context.Context, in *MultiplierRequest, opts ...grpc.CallOption) (*MultiplierResponse, error) {
	out := new(MultiplierResponse)
	err := c.cc.Invoke(ctx, "/calc.calculatorService/multiply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorServiceClient) GetByteValues(ctx context.Context, in *GigaByteRequest, opts ...grpc.CallOption) (CalculatorService_GetByteValuesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalculatorService_serviceDesc.Streams[0], "/calc.calculatorService/getByteValues", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServiceGetByteValuesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CalculatorService_GetByteValuesClient interface {
	Recv() (*GigaByteResponse, error)
	grpc.ClientStream
}

type calculatorServiceGetByteValuesClient struct {
	grpc.ClientStream
}

func (x *calculatorServiceGetByteValuesClient) Recv() (*GigaByteResponse, error) {
	m := new(GigaByteResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CalculatorServiceServer is the server API for CalculatorService service.
type CalculatorServiceServer interface {
	// Unary
	DoSum(context.Context, *SumRequest) (*SumResponse, error)
	CalcDiff(context.Context, *DiffRequest) (*DiffResponse, error)
	Multiply(context.Context, *MultiplierRequest) (*MultiplierResponse, error)
	// Server Streaming
	GetByteValues(*GigaByteRequest, CalculatorService_GetByteValuesServer) error
}

// UnimplementedCalculatorServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCalculatorServiceServer struct {
}

func (*UnimplementedCalculatorServiceServer) DoSum(context.Context, *SumRequest) (*SumResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoSum not implemented")
}
func (*UnimplementedCalculatorServiceServer) CalcDiff(context.Context, *DiffRequest) (*DiffResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CalcDiff not implemented")
}
func (*UnimplementedCalculatorServiceServer) Multiply(context.Context, *MultiplierRequest) (*MultiplierResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Multiply not implemented")
}
func (*UnimplementedCalculatorServiceServer) GetByteValues(*GigaByteRequest, CalculatorService_GetByteValuesServer) error {
	return status.Errorf(codes.Unimplemented, "method GetByteValues not implemented")
}

func RegisterCalculatorServiceServer(s *grpc.Server, srv CalculatorServiceServer) {
	s.RegisterService(&_CalculatorService_serviceDesc, srv)
}

func _CalculatorService_DoSum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).DoSum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calc.calculatorService/DoSum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).DoSum(ctx, req.(*SumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculatorService_CalcDiff_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiffRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).CalcDiff(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calc.calculatorService/CalcDiff",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).CalcDiff(ctx, req.(*DiffRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculatorService_Multiply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MultiplierRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).Multiply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calc.calculatorService/Multiply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).Multiply(ctx, req.(*MultiplierRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculatorService_GetByteValues_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GigaByteRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CalculatorServiceServer).GetByteValues(m, &calculatorServiceGetByteValuesServer{stream})
}

type CalculatorService_GetByteValuesServer interface {
	Send(*GigaByteResponse) error
	grpc.ServerStream
}

type calculatorServiceGetByteValuesServer struct {
	grpc.ServerStream
}

func (x *calculatorServiceGetByteValuesServer) Send(m *GigaByteResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _CalculatorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "calc.calculatorService",
	HandlerType: (*CalculatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "doSum",
			Handler:    _CalculatorService_DoSum_Handler,
		},
		{
			MethodName: "calcDiff",
			Handler:    _CalculatorService_CalcDiff_Handler,
		},
		{
			MethodName: "multiply",
			Handler:    _CalculatorService_Multiply_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "getByteValues",
			Handler:       _CalculatorService_GetByteValues_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "calc_pb/calc.proto",
}

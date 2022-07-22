// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: pb/face_api.proto

package pb

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

type FaceVerifyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Img1 []byte `protobuf:"bytes,1,opt,name=img1,proto3" json:"img1,omitempty"`
	Img2 []byte `protobuf:"bytes,2,opt,name=img2,proto3" json:"img2,omitempty"`
}

func (x *FaceVerifyRequest) Reset() {
	*x = FaceVerifyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_face_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FaceVerifyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FaceVerifyRequest) ProtoMessage() {}

func (x *FaceVerifyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_face_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FaceVerifyRequest.ProtoReflect.Descriptor instead.
func (*FaceVerifyRequest) Descriptor() ([]byte, []int) {
	return file_pb_face_api_proto_rawDescGZIP(), []int{0}
}

func (x *FaceVerifyRequest) GetImg1() []byte {
	if x != nil {
		return x.Img1
	}
	return nil
}

func (x *FaceVerifyRequest) GetImg2() []byte {
	if x != nil {
		return x.Img2
	}
	return nil
}

type FaceVerifyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32   `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Score   float32 `protobuf:"fixed32,2,opt,name=score,proto3" json:"score,omitempty"`
	Message string  `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *FaceVerifyResponse) Reset() {
	*x = FaceVerifyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_face_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FaceVerifyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FaceVerifyResponse) ProtoMessage() {}

func (x *FaceVerifyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_face_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FaceVerifyResponse.ProtoReflect.Descriptor instead.
func (*FaceVerifyResponse) Descriptor() ([]byte, []int) {
	return file_pb_face_api_proto_rawDescGZIP(), []int{1}
}

func (x *FaceVerifyResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *FaceVerifyResponse) GetScore() float32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *FaceVerifyResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ObjectFeature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type    string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Version int32  `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
	Flags   int32  `protobuf:"varint,3,opt,name=flags,proto3" json:"flags,omitempty"`
	Blob    []byte `protobuf:"bytes,4,opt,name=blob,proto3" json:"blob,omitempty"`
}

func (x *ObjectFeature) Reset() {
	*x = ObjectFeature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_face_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObjectFeature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectFeature) ProtoMessage() {}

func (x *ObjectFeature) ProtoReflect() protoreflect.Message {
	mi := &file_pb_face_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectFeature.ProtoReflect.Descriptor instead.
func (*ObjectFeature) Descriptor() ([]byte, []int) {
	return file_pb_face_api_proto_rawDescGZIP(), []int{2}
}

func (x *ObjectFeature) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ObjectFeature) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *ObjectFeature) GetFlags() int32 {
	if x != nil {
		return x.Flags
	}
	return 0
}

func (x *ObjectFeature) GetBlob() []byte {
	if x != nil {
		return x.Blob
	}
	return nil
}

type FaceExtractRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Img []byte `protobuf:"bytes,1,opt,name=img,proto3" json:"img,omitempty"`
}

func (x *FaceExtractRequest) Reset() {
	*x = FaceExtractRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_face_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FaceExtractRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FaceExtractRequest) ProtoMessage() {}

func (x *FaceExtractRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_face_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FaceExtractRequest.ProtoReflect.Descriptor instead.
func (*FaceExtractRequest) Descriptor() ([]byte, []int) {
	return file_pb_face_api_proto_rawDescGZIP(), []int{3}
}

func (x *FaceExtractRequest) GetImg() []byte {
	if x != nil {
		return x.Img
	}
	return nil
}

type FaceExtractResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Blob    []byte `protobuf:"bytes,2,opt,name=blob,proto3" json:"blob,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *FaceExtractResponse) Reset() {
	*x = FaceExtractResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_face_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FaceExtractResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FaceExtractResponse) ProtoMessage() {}

func (x *FaceExtractResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_face_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FaceExtractResponse.ProtoReflect.Descriptor instead.
func (*FaceExtractResponse) Descriptor() ([]byte, []int) {
	return file_pb_face_api_proto_rawDescGZIP(), []int{4}
}

func (x *FaceExtractResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *FaceExtractResponse) GetBlob() []byte {
	if x != nil {
		return x.Blob
	}
	return nil
}

func (x *FaceExtractResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type FeatVerifyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Feat1 []byte `protobuf:"bytes,1,opt,name=feat1,proto3" json:"feat1,omitempty"`
	Feat2 []byte `protobuf:"bytes,2,opt,name=feat2,proto3" json:"feat2,omitempty"`
}

func (x *FeatVerifyRequest) Reset() {
	*x = FeatVerifyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_face_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeatVerifyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeatVerifyRequest) ProtoMessage() {}

func (x *FeatVerifyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_face_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeatVerifyRequest.ProtoReflect.Descriptor instead.
func (*FeatVerifyRequest) Descriptor() ([]byte, []int) {
	return file_pb_face_api_proto_rawDescGZIP(), []int{5}
}

func (x *FeatVerifyRequest) GetFeat1() []byte {
	if x != nil {
		return x.Feat1
	}
	return nil
}

func (x *FeatVerifyRequest) GetFeat2() []byte {
	if x != nil {
		return x.Feat2
	}
	return nil
}

type FeatVerifyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32   `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Score   float32 `protobuf:"fixed32,2,opt,name=score,proto3" json:"score,omitempty"`
	Message string  `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *FeatVerifyResponse) Reset() {
	*x = FeatVerifyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_face_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeatVerifyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeatVerifyResponse) ProtoMessage() {}

func (x *FeatVerifyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_face_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeatVerifyResponse.ProtoReflect.Descriptor instead.
func (*FeatVerifyResponse) Descriptor() ([]byte, []int) {
	return file_pb_face_api_proto_rawDescGZIP(), []int{6}
}

func (x *FeatVerifyResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *FeatVerifyResponse) GetScore() float32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *FeatVerifyResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_pb_face_api_proto protoreflect.FileDescriptor

var file_pb_face_api_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x62, 0x2f, 0x66, 0x61, 0x63, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x3b, 0x0a, 0x11, 0x46, 0x61, 0x63, 0x65, 0x56,
	0x65, 0x72, 0x69, 0x66, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x69, 0x6d, 0x67, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x69, 0x6d, 0x67, 0x31,
	0x12, 0x12, 0x0a, 0x04, 0x69, 0x6d, 0x67, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04,
	0x69, 0x6d, 0x67, 0x32, 0x22, 0x58, 0x0a, 0x12, 0x46, 0x61, 0x63, 0x65, 0x56, 0x65, 0x72, 0x69,
	0x66, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x73,
	0x63, 0x6f, 0x72, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x67,
	0x0a, 0x0d, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a,
	0x05, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x66, 0x6c,
	0x61, 0x67, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6c, 0x6f, 0x62, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x04, 0x62, 0x6c, 0x6f, 0x62, 0x22, 0x26, 0x0a, 0x12, 0x46, 0x61, 0x63, 0x65, 0x45,
	0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x69, 0x6d, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x69, 0x6d, 0x67, 0x22,
	0x57, 0x0a, 0x13, 0x46, 0x61, 0x63, 0x65, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6c,
	0x6f, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x62, 0x6c, 0x6f, 0x62, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x3f, 0x0a, 0x11, 0x46, 0x65, 0x61, 0x74,
	0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x66, 0x65, 0x61, 0x74, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x66, 0x65,
	0x61, 0x74, 0x31, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x65, 0x61, 0x74, 0x32, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x05, 0x66, 0x65, 0x61, 0x74, 0x32, 0x22, 0x58, 0x0a, 0x12, 0x46, 0x65, 0x61,
	0x74, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x32, 0xc9, 0x01, 0x0a, 0x07, 0x46, 0x61, 0x63, 0x65, 0x41, 0x70, 0x69, 0x12,
	0x3d, 0x0a, 0x0a, 0x46, 0x61, 0x63, 0x65, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x12, 0x15, 0x2e,
	0x70, 0x62, 0x2e, 0x46, 0x61, 0x63, 0x65, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x61, 0x63, 0x65, 0x56, 0x65,
	0x72, 0x69, 0x66, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x40,
	0x0a, 0x0b, 0x46, 0x61, 0x63, 0x65, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x12, 0x16, 0x2e,
	0x70, 0x62, 0x2e, 0x46, 0x61, 0x63, 0x65, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x61, 0x63, 0x65, 0x45,
	0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x3d, 0x0a, 0x0a, 0x46, 0x65, 0x61, 0x74, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x12, 0x15,
	0x2e, 0x70, 0x62, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x56,
	0x65, 0x72, 0x69, 0x66, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x05, 0x5a, 0x03, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_face_api_proto_rawDescOnce sync.Once
	file_pb_face_api_proto_rawDescData = file_pb_face_api_proto_rawDesc
)

func file_pb_face_api_proto_rawDescGZIP() []byte {
	file_pb_face_api_proto_rawDescOnce.Do(func() {
		file_pb_face_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_face_api_proto_rawDescData)
	})
	return file_pb_face_api_proto_rawDescData
}

var file_pb_face_api_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_pb_face_api_proto_goTypes = []interface{}{
	(*FaceVerifyRequest)(nil),   // 0: pb.FaceVerifyRequest
	(*FaceVerifyResponse)(nil),  // 1: pb.FaceVerifyResponse
	(*ObjectFeature)(nil),       // 2: pb.ObjectFeature
	(*FaceExtractRequest)(nil),  // 3: pb.FaceExtractRequest
	(*FaceExtractResponse)(nil), // 4: pb.FaceExtractResponse
	(*FeatVerifyRequest)(nil),   // 5: pb.FeatVerifyRequest
	(*FeatVerifyResponse)(nil),  // 6: pb.FeatVerifyResponse
}
var file_pb_face_api_proto_depIdxs = []int32{
	0, // 0: pb.FaceApi.FaceVerify:input_type -> pb.FaceVerifyRequest
	3, // 1: pb.FaceApi.FaceExtract:input_type -> pb.FaceExtractRequest
	5, // 2: pb.FaceApi.FeatVerify:input_type -> pb.FeatVerifyRequest
	1, // 3: pb.FaceApi.FaceVerify:output_type -> pb.FaceVerifyResponse
	4, // 4: pb.FaceApi.FaceExtract:output_type -> pb.FaceExtractResponse
	6, // 5: pb.FaceApi.FeatVerify:output_type -> pb.FeatVerifyResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pb_face_api_proto_init() }
func file_pb_face_api_proto_init() {
	if File_pb_face_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_face_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FaceVerifyRequest); i {
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
		file_pb_face_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FaceVerifyResponse); i {
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
		file_pb_face_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObjectFeature); i {
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
		file_pb_face_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FaceExtractRequest); i {
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
		file_pb_face_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FaceExtractResponse); i {
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
		file_pb_face_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeatVerifyRequest); i {
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
		file_pb_face_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeatVerifyResponse); i {
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
			RawDescriptor: file_pb_face_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_face_api_proto_goTypes,
		DependencyIndexes: file_pb_face_api_proto_depIdxs,
		MessageInfos:      file_pb_face_api_proto_msgTypes,
	}.Build()
	File_pb_face_api_proto = out.File
	file_pb_face_api_proto_rawDesc = nil
	file_pb_face_api_proto_goTypes = nil
	file_pb_face_api_proto_depIdxs = nil
}

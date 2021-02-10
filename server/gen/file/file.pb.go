// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: file.proto

package file

import (
	proto "github.com/golang/protobuf/proto"
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

type FileMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FileMetadata) Reset() {
	*x = FileMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileMetadata) ProtoMessage() {}

func (x *FileMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_file_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileMetadata.ProtoReflect.Descriptor instead.
func (*FileMetadata) Descriptor() ([]byte, []int) {
	return file_file_proto_rawDescGZIP(), []int{0}
}

type AwsS3FileDestination struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Region     string `protobuf:"bytes,2,opt,name=region,proto3" json:"region,omitempty"`
	BucketName string `protobuf:"bytes,3,opt,name=bucketName,proto3" json:"bucketName,omitempty"`
	Prefix     string `protobuf:"bytes,4,opt,name=prefix,proto3" json:"prefix,omitempty"`
}

func (x *AwsS3FileDestination) Reset() {
	*x = AwsS3FileDestination{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AwsS3FileDestination) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AwsS3FileDestination) ProtoMessage() {}

func (x *AwsS3FileDestination) ProtoReflect() protoreflect.Message {
	mi := &file_file_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AwsS3FileDestination.ProtoReflect.Descriptor instead.
func (*AwsS3FileDestination) Descriptor() ([]byte, []int) {
	return file_file_proto_rawDescGZIP(), []int{1}
}

func (x *AwsS3FileDestination) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AwsS3FileDestination) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *AwsS3FileDestination) GetBucketName() string {
	if x != nil {
		return x.BucketName
	}
	return ""
}

func (x *AwsS3FileDestination) GetPrefix() string {
	if x != nil {
		return x.Prefix
	}
	return ""
}

type FileDestination struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Destination:
	//	*FileDestination_Aws
	Destination isFileDestination_Destination `protobuf_oneof:"destination"`
}

func (x *FileDestination) Reset() {
	*x = FileDestination{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileDestination) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileDestination) ProtoMessage() {}

func (x *FileDestination) ProtoReflect() protoreflect.Message {
	mi := &file_file_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileDestination.ProtoReflect.Descriptor instead.
func (*FileDestination) Descriptor() ([]byte, []int) {
	return file_file_proto_rawDescGZIP(), []int{2}
}

func (m *FileDestination) GetDestination() isFileDestination_Destination {
	if m != nil {
		return m.Destination
	}
	return nil
}

func (x *FileDestination) GetAws() *AwsS3FileDestination {
	if x, ok := x.GetDestination().(*FileDestination_Aws); ok {
		return x.Aws
	}
	return nil
}

type isFileDestination_Destination interface {
	isFileDestination_Destination()
}

type FileDestination_Aws struct {
	Aws *AwsS3FileDestination `protobuf:"bytes,1,opt,name=aws,proto3,oneof"`
}

func (*FileDestination_Aws) isFileDestination_Destination() {}

type UploadFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ChunkOffset uint32 `protobuf:"varint,2,opt,name=chunkOffset,proto3" json:"chunkOffset,omitempty"`
	// Types that are assignable to Data:
	//	*UploadFileRequest_Chunk
	//	*UploadFileRequest_Metadata
	Data isUploadFileRequest_Data `protobuf_oneof:"data"`
}

func (x *UploadFileRequest) Reset() {
	*x = UploadFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadFileRequest) ProtoMessage() {}

func (x *UploadFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_file_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadFileRequest.ProtoReflect.Descriptor instead.
func (*UploadFileRequest) Descriptor() ([]byte, []int) {
	return file_file_proto_rawDescGZIP(), []int{3}
}

func (x *UploadFileRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UploadFileRequest) GetChunkOffset() uint32 {
	if x != nil {
		return x.ChunkOffset
	}
	return 0
}

func (m *UploadFileRequest) GetData() isUploadFileRequest_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *UploadFileRequest) GetChunk() []byte {
	if x, ok := x.GetData().(*UploadFileRequest_Chunk); ok {
		return x.Chunk
	}
	return nil
}

func (x *UploadFileRequest) GetMetadata() *FileMetadata {
	if x, ok := x.GetData().(*UploadFileRequest_Metadata); ok {
		return x.Metadata
	}
	return nil
}

type isUploadFileRequest_Data interface {
	isUploadFileRequest_Data()
}

type UploadFileRequest_Chunk struct {
	Chunk []byte `protobuf:"bytes,3,opt,name=chunk,proto3,oneof"`
}

type UploadFileRequest_Metadata struct {
	Metadata *FileMetadata `protobuf:"bytes,4,opt,name=metadata,proto3,oneof"`
}

func (*UploadFileRequest_Chunk) isUploadFileRequest_Data() {}

func (*UploadFileRequest_Metadata) isUploadFileRequest_Data() {}

type UploadFileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	DestinationName string `protobuf:"bytes,2,opt,name=destinationName,proto3" json:"destinationName,omitempty"`
}

func (x *UploadFileResponse) Reset() {
	*x = UploadFileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadFileResponse) ProtoMessage() {}

func (x *UploadFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_file_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadFileResponse.ProtoReflect.Descriptor instead.
func (*UploadFileResponse) Descriptor() ([]byte, []int) {
	return file_file_proto_rawDescGZIP(), []int{4}
}

func (x *UploadFileResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UploadFileResponse) GetDestinationName() string {
	if x != nil {
		return x.DestinationName
	}
	return ""
}

type BeginUploadFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BeginUploadFileRequest) Reset() {
	*x = BeginUploadFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BeginUploadFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BeginUploadFileRequest) ProtoMessage() {}

func (x *BeginUploadFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_file_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BeginUploadFileRequest.ProtoReflect.Descriptor instead.
func (*BeginUploadFileRequest) Descriptor() ([]byte, []int) {
	return file_file_proto_rawDescGZIP(), []int{5}
}

type BeginUploadFileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string             `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Destinations []*FileDestination `protobuf:"bytes,2,rep,name=destinations,proto3" json:"destinations,omitempty"`
}

func (x *BeginUploadFileResponse) Reset() {
	*x = BeginUploadFileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BeginUploadFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BeginUploadFileResponse) ProtoMessage() {}

func (x *BeginUploadFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_file_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BeginUploadFileResponse.ProtoReflect.Descriptor instead.
func (*BeginUploadFileResponse) Descriptor() ([]byte, []int) {
	return file_file_proto_rawDescGZIP(), []int{6}
}

func (x *BeginUploadFileResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *BeginUploadFileResponse) GetDestinations() []*FileDestination {
	if x != nil {
		return x.Destinations
	}
	return nil
}

var File_file_proto protoreflect.FileDescriptor

var file_file_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x64, 0x6f,
	0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x22, 0x0e, 0x0a, 0x0c,
	0x46, 0x69, 0x6c, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0x7a, 0x0a, 0x14,
	0x41, 0x77, 0x73, 0x53, 0x33, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x12, 0x1e, 0x0a, 0x0a, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x22, 0x5a, 0x0a, 0x0f, 0x46, 0x69, 0x6c, 0x65,
	0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x38, 0x0a, 0x03, 0x61,
	0x77, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x64, 0x6f, 0x63, 0x6c, 0x6f,
	0x63, 0x6b, 0x65, 0x72, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x41, 0x77, 0x73, 0x53, 0x33, 0x46,
	0x69, 0x6c, 0x65, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00,
	0x52, 0x03, 0x61, 0x77, 0x73, 0x42, 0x0d, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0xa1, 0x01, 0x0a, 0x11, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46,
	0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x68,
	0x75, 0x6e, 0x6b, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0b, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x16, 0x0a, 0x05,
	0x63, 0x68, 0x75, 0x6e, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x05, 0x63,
	0x68, 0x75, 0x6e, 0x6b, 0x12, 0x3a, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x64, 0x6f, 0x63, 0x6c, 0x6f, 0x63, 0x6b,
	0x65, 0x72, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x42, 0x06, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x4e, 0x0a, 0x12, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x28,
	0x0a, 0x0f, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x18, 0x0a, 0x16, 0x42, 0x65, 0x67, 0x69,
	0x6e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x6e, 0x0a, 0x17, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x43, 0x0a,
	0x0c, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x64, 0x6f, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x2e,
	0x66, 0x69, 0x6c, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x32, 0xca, 0x01, 0x0a, 0x0b, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x62, 0x0a, 0x0f, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x26, 0x2e, 0x64, 0x6f, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x65,
	0x72, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e,
	0x64, 0x6f, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x42,
	0x65, 0x67, 0x69, 0x6e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x57, 0x0a, 0x0a, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x46, 0x69, 0x6c, 0x65, 0x12, 0x21, 0x2e, 0x64, 0x6f, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x72,
	0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x64, 0x6f, 0x63, 0x6c, 0x6f, 0x63,
	0x6b, 0x65, 0x72, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46,
	0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x42,
	0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74,
	0x65, 0x70, 0x68, 0x65, 0x6e, 0x7a, 0x73, 0x79, 0x2f, 0x64, 0x6f, 0x63, 0x2d, 0x6c, 0x6f, 0x63,
	0x6b, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x66,
	0x69, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_file_proto_rawDescOnce sync.Once
	file_file_proto_rawDescData = file_file_proto_rawDesc
)

func file_file_proto_rawDescGZIP() []byte {
	file_file_proto_rawDescOnce.Do(func() {
		file_file_proto_rawDescData = protoimpl.X.CompressGZIP(file_file_proto_rawDescData)
	})
	return file_file_proto_rawDescData
}

var file_file_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_file_proto_goTypes = []interface{}{
	(*FileMetadata)(nil),            // 0: doclocker.file.FileMetadata
	(*AwsS3FileDestination)(nil),    // 1: doclocker.file.AwsS3FileDestination
	(*FileDestination)(nil),         // 2: doclocker.file.FileDestination
	(*UploadFileRequest)(nil),       // 3: doclocker.file.UploadFileRequest
	(*UploadFileResponse)(nil),      // 4: doclocker.file.UploadFileResponse
	(*BeginUploadFileRequest)(nil),  // 5: doclocker.file.BeginUploadFileRequest
	(*BeginUploadFileResponse)(nil), // 6: doclocker.file.BeginUploadFileResponse
}
var file_file_proto_depIdxs = []int32{
	1, // 0: doclocker.file.FileDestination.aws:type_name -> doclocker.file.AwsS3FileDestination
	0, // 1: doclocker.file.UploadFileRequest.metadata:type_name -> doclocker.file.FileMetadata
	2, // 2: doclocker.file.BeginUploadFileResponse.destinations:type_name -> doclocker.file.FileDestination
	5, // 3: doclocker.file.FileService.BeginUploadFile:input_type -> doclocker.file.BeginUploadFileRequest
	3, // 4: doclocker.file.FileService.UploadFile:input_type -> doclocker.file.UploadFileRequest
	6, // 5: doclocker.file.FileService.BeginUploadFile:output_type -> doclocker.file.BeginUploadFileResponse
	4, // 6: doclocker.file.FileService.UploadFile:output_type -> doclocker.file.UploadFileResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_file_proto_init() }
func file_file_proto_init() {
	if File_file_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_file_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileMetadata); i {
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
		file_file_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AwsS3FileDestination); i {
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
		file_file_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileDestination); i {
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
		file_file_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadFileRequest); i {
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
		file_file_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadFileResponse); i {
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
		file_file_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BeginUploadFileRequest); i {
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
		file_file_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BeginUploadFileResponse); i {
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
	file_file_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*FileDestination_Aws)(nil),
	}
	file_file_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*UploadFileRequest_Chunk)(nil),
		(*UploadFileRequest_Metadata)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_file_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_file_proto_goTypes,
		DependencyIndexes: file_file_proto_depIdxs,
		MessageInfos:      file_file_proto_msgTypes,
	}.Build()
	File_file_proto = out.File
	file_file_proto_rawDesc = nil
	file_file_proto_goTypes = nil
	file_file_proto_depIdxs = nil
}
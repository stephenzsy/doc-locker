// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: configurations.proto

package configurations

import (
	proto "github.com/golang/protobuf/proto"
	file "github.com/stephenzsy/doc-locker/server/gen/file"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type SiteConfigurationsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SiteConfigurationsJson string `protobuf:"bytes,1,opt,name=siteConfigurationsJson,proto3" json:"siteConfigurationsJson,omitempty"`
}

func (x *SiteConfigurationsResponse) Reset() {
	*x = SiteConfigurationsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_configurations_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SiteConfigurationsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SiteConfigurationsResponse) ProtoMessage() {}

func (x *SiteConfigurationsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_configurations_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SiteConfigurationsResponse.ProtoReflect.Descriptor instead.
func (*SiteConfigurationsResponse) Descriptor() ([]byte, []int) {
	return file_configurations_proto_rawDescGZIP(), []int{0}
}

func (x *SiteConfigurationsResponse) GetSiteConfigurationsJson() string {
	if x != nil {
		return x.SiteConfigurationsJson
	}
	return ""
}

type UserProfileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Provider   string `protobuf:"bytes,1,opt,name=provider,proto3" json:"provider,omitempty"`
	ConfigPath string `protobuf:"bytes,2,opt,name=configPath,proto3" json:"configPath,omitempty"`
}

func (x *UserProfileRequest) Reset() {
	*x = UserProfileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_configurations_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserProfileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserProfileRequest) ProtoMessage() {}

func (x *UserProfileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_configurations_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserProfileRequest.ProtoReflect.Descriptor instead.
func (*UserProfileRequest) Descriptor() ([]byte, []int) {
	return file_configurations_proto_rawDescGZIP(), []int{1}
}

func (x *UserProfileRequest) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

func (x *UserProfileRequest) GetConfigPath() string {
	if x != nil {
		return x.ConfigPath
	}
	return ""
}

type UserProfileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LastUpdated      *timestamppb.Timestamp  `protobuf:"bytes,1,opt,name=lastUpdated,proto3" json:"lastUpdated,omitempty"`
	FileDestinations []*file.FileDestination `protobuf:"bytes,2,rep,name=fileDestinations,proto3" json:"fileDestinations,omitempty"`
}

func (x *UserProfileResponse) Reset() {
	*x = UserProfileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_configurations_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserProfileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserProfileResponse) ProtoMessage() {}

func (x *UserProfileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_configurations_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserProfileResponse.ProtoReflect.Descriptor instead.
func (*UserProfileResponse) Descriptor() ([]byte, []int) {
	return file_configurations_proto_rawDescGZIP(), []int{2}
}

func (x *UserProfileResponse) GetLastUpdated() *timestamppb.Timestamp {
	if x != nil {
		return x.LastUpdated
	}
	return nil
}

func (x *UserProfileResponse) GetFileDestinations() []*file.FileDestination {
	if x != nil {
		return x.FileDestinations
	}
	return nil
}

var File_configurations_proto protoreflect.FileDescriptor

var file_configurations_proto_rawDesc = []byte{
	0x0a, 0x14, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x64, 0x6f, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x65,
	0x72, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a,
	0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x54, 0x0a, 0x1a, 0x53, 0x69,
	0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x16, 0x73, 0x69, 0x74, 0x65,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x4a, 0x73,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x73, 0x69, 0x74, 0x65, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x4a, 0x73, 0x6f, 0x6e,
	0x22, 0x50, 0x0a, 0x12, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x50, 0x61, 0x74, 0x68,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x50, 0x61,
	0x74, 0x68, 0x22, 0xa0, 0x01, 0x0a, 0x13, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x0b, 0x6c, 0x61,
	0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x6c, 0x61, 0x73,
	0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x4b, 0x0a, 0x10, 0x66, 0x69, 0x6c, 0x65,
	0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x64, 0x6f, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x66,
	0x69, 0x6c, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x10, 0x66, 0x69, 0x6c, 0x65, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x32, 0xe7, 0x01, 0x0a, 0x15, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x62, 0x0a, 0x12, 0x53, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x34, 0x2e,
	0x64, 0x6f, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x53, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x6a, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x12, 0x2c, 0x2e, 0x64, 0x6f, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2d, 0x2e, 0x64, 0x6f, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x3c, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74,
	0x65, 0x70, 0x68, 0x65, 0x6e, 0x7a, 0x73, 0x79, 0x2f, 0x64, 0x6f, 0x63, 0x2d, 0x6c, 0x6f, 0x63,
	0x6b, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_configurations_proto_rawDescOnce sync.Once
	file_configurations_proto_rawDescData = file_configurations_proto_rawDesc
)

func file_configurations_proto_rawDescGZIP() []byte {
	file_configurations_proto_rawDescOnce.Do(func() {
		file_configurations_proto_rawDescData = protoimpl.X.CompressGZIP(file_configurations_proto_rawDescData)
	})
	return file_configurations_proto_rawDescData
}

var file_configurations_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_configurations_proto_goTypes = []interface{}{
	(*SiteConfigurationsResponse)(nil), // 0: doclocker.configurations.SiteConfigurationsResponse
	(*UserProfileRequest)(nil),         // 1: doclocker.configurations.UserProfileRequest
	(*UserProfileResponse)(nil),        // 2: doclocker.configurations.UserProfileResponse
	(*timestamppb.Timestamp)(nil),      // 3: google.protobuf.Timestamp
	(*file.FileDestination)(nil),       // 4: doclocker.file.FileDestination
	(*emptypb.Empty)(nil),              // 5: google.protobuf.Empty
}
var file_configurations_proto_depIdxs = []int32{
	3, // 0: doclocker.configurations.UserProfileResponse.lastUpdated:type_name -> google.protobuf.Timestamp
	4, // 1: doclocker.configurations.UserProfileResponse.fileDestinations:type_name -> doclocker.file.FileDestination
	5, // 2: doclocker.configurations.ConfigurationsService.SiteConfigurations:input_type -> google.protobuf.Empty
	1, // 3: doclocker.configurations.ConfigurationsService.UserProfile:input_type -> doclocker.configurations.UserProfileRequest
	0, // 4: doclocker.configurations.ConfigurationsService.SiteConfigurations:output_type -> doclocker.configurations.SiteConfigurationsResponse
	2, // 5: doclocker.configurations.ConfigurationsService.UserProfile:output_type -> doclocker.configurations.UserProfileResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_configurations_proto_init() }
func file_configurations_proto_init() {
	if File_configurations_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_configurations_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SiteConfigurationsResponse); i {
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
		file_configurations_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserProfileRequest); i {
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
		file_configurations_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserProfileResponse); i {
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
			RawDescriptor: file_configurations_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_configurations_proto_goTypes,
		DependencyIndexes: file_configurations_proto_depIdxs,
		MessageInfos:      file_configurations_proto_msgTypes,
	}.Build()
	File_configurations_proto = out.File
	file_configurations_proto_rawDesc = nil
	file_configurations_proto_goTypes = nil
	file_configurations_proto_depIdxs = nil
}
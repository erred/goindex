// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: goindex.proto

package goindex

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

type VersionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Project string `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	Semver  bool   `protobuf:"varint,2,opt,name=semver,proto3" json:"semver,omitempty"`
}

func (x *VersionsRequest) Reset() {
	*x = VersionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goindex_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VersionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VersionsRequest) ProtoMessage() {}

func (x *VersionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_goindex_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VersionsRequest.ProtoReflect.Descriptor instead.
func (*VersionsRequest) Descriptor() ([]byte, []int) {
	return file_goindex_proto_rawDescGZIP(), []int{0}
}

func (x *VersionsRequest) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

func (x *VersionsRequest) GetSemver() bool {
	if x != nil {
		return x.Semver
	}
	return false
}

type IndexRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path      string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Version   string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Timestamp string `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *IndexRecord) Reset() {
	*x = IndexRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goindex_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IndexRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IndexRecord) ProtoMessage() {}

func (x *IndexRecord) ProtoReflect() protoreflect.Message {
	mi := &file_goindex_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IndexRecord.ProtoReflect.Descriptor instead.
func (*IndexRecord) Descriptor() ([]byte, []int) {
	return file_goindex_proto_rawDescGZIP(), []int{1}
}

func (x *IndexRecord) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *IndexRecord) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *IndexRecord) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

type ProjectVersions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Project  string         `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	Versions []*IndexRecord `protobuf:"bytes,2,rep,name=versions,proto3" json:"versions,omitempty"`
}

func (x *ProjectVersions) Reset() {
	*x = ProjectVersions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goindex_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectVersions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectVersions) ProtoMessage() {}

func (x *ProjectVersions) ProtoReflect() protoreflect.Message {
	mi := &file_goindex_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectVersions.ProtoReflect.Descriptor instead.
func (*ProjectVersions) Descriptor() ([]byte, []int) {
	return file_goindex_proto_rawDescGZIP(), []int{2}
}

func (x *ProjectVersions) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

func (x *ProjectVersions) GetVersions() []*IndexRecord {
	if x != nil {
		return x.Versions
	}
	return nil
}

var File_goindex_proto protoreflect.FileDescriptor

var file_goindex_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x67, 0x6f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x06, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x22, 0x43, 0x0a, 0x0f, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6d, 0x76, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x65, 0x6d, 0x76, 0x65, 0x72, 0x22, 0x59, 0x0a, 0x0b,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12,
	0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x5c, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x12, 0x2f, 0x0a, 0x08, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x08, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x32, 0x49, 0x0a, 0x07, 0x47, 0x6f, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x12, 0x3e, 0x0a, 0x08, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x17, 0x2e, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x00,
	0x42, 0x1b, 0x5a, 0x19, 0x67, 0x6f, 0x2e, 0x73, 0x65, 0x61, 0x6e, 0x6b, 0x68, 0x6c, 0x69, 0x61,
	0x6f, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_goindex_proto_rawDescOnce sync.Once
	file_goindex_proto_rawDescData = file_goindex_proto_rawDesc
)

func file_goindex_proto_rawDescGZIP() []byte {
	file_goindex_proto_rawDescOnce.Do(func() {
		file_goindex_proto_rawDescData = protoimpl.X.CompressGZIP(file_goindex_proto_rawDescData)
	})
	return file_goindex_proto_rawDescData
}

var file_goindex_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_goindex_proto_goTypes = []interface{}{
	(*VersionsRequest)(nil), // 0: stream.VersionsRequest
	(*IndexRecord)(nil),     // 1: stream.IndexRecord
	(*ProjectVersions)(nil), // 2: stream.ProjectVersions
}
var file_goindex_proto_depIdxs = []int32{
	1, // 0: stream.ProjectVersions.versions:type_name -> stream.IndexRecord
	0, // 1: stream.Goindex.Versions:input_type -> stream.VersionsRequest
	2, // 2: stream.Goindex.Versions:output_type -> stream.ProjectVersions
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_goindex_proto_init() }
func file_goindex_proto_init() {
	if File_goindex_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_goindex_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VersionsRequest); i {
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
		file_goindex_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IndexRecord); i {
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
		file_goindex_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectVersions); i {
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
			RawDescriptor: file_goindex_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_goindex_proto_goTypes,
		DependencyIndexes: file_goindex_proto_depIdxs,
		MessageInfos:      file_goindex_proto_msgTypes,
	}.Build()
	File_goindex_proto = out.File
	file_goindex_proto_rawDesc = nil
	file_goindex_proto_goTypes = nil
	file_goindex_proto_depIdxs = nil
}

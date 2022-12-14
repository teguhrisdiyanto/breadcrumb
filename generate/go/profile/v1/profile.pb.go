// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: proto/profile/v1/profile.proto

package profilev1

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

// Profile profile message represent entity job in talent.
type Profile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	IdUser          uint64   `protobuf:"varint,2,opt,name=id_user,json=idUser,proto3" json:"id_user,omitempty"`
	FullName        string   `protobuf:"bytes,3,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	Description     string   `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Headline        string   `protobuf:"bytes,5,opt,name=headline,proto3" json:"headline,omitempty"`
	CurrentDomicile string   `protobuf:"bytes,6,opt,name=current_domicile,json=currentDomicile,proto3" json:"current_domicile,omitempty"`
	ListOfSkills    []string `protobuf:"bytes,7,rep,name=list_of_skills,json=listOfSkills,proto3" json:"list_of_skills,omitempty"`
	CreatedAt       int64    `protobuf:"fixed64,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt       int64    `protobuf:"fixed64,9,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt       int64    `protobuf:"fixed64,10,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
}

func (x *Profile) Reset() {
	*x = Profile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_profile_v1_profile_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Profile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Profile) ProtoMessage() {}

func (x *Profile) ProtoReflect() protoreflect.Message {
	mi := &file_proto_profile_v1_profile_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Profile.ProtoReflect.Descriptor instead.
func (*Profile) Descriptor() ([]byte, []int) {
	return file_proto_profile_v1_profile_proto_rawDescGZIP(), []int{0}
}

func (x *Profile) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Profile) GetIdUser() uint64 {
	if x != nil {
		return x.IdUser
	}
	return 0
}

func (x *Profile) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *Profile) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Profile) GetHeadline() string {
	if x != nil {
		return x.Headline
	}
	return ""
}

func (x *Profile) GetCurrentDomicile() string {
	if x != nil {
		return x.CurrentDomicile
	}
	return ""
}

func (x *Profile) GetListOfSkills() []string {
	if x != nil {
		return x.ListOfSkills
	}
	return nil
}

func (x *Profile) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Profile) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *Profile) GetDeletedAt() int64 {
	if x != nil {
		return x.DeletedAt
	}
	return 0
}

var File_proto_profile_v1_profile_proto protoreflect.FileDescriptor

var file_proto_profile_v1_profile_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2f,
	0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e,
	0x76, 0x31, 0x22, 0xbb, 0x02, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17,
	0x0a, 0x07, 0x69, 0x64, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x06, 0x69, 0x64, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x75, 0x6c, 0x6c, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x65, 0x61, 0x64, 0x6c, 0x69,
	0x6e, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x65, 0x61, 0x64, 0x6c, 0x69,
	0x6e, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x64, 0x6f,
	0x6d, 0x69, 0x63, 0x69, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x44, 0x6f, 0x6d, 0x69, 0x63, 0x69, 0x6c, 0x65, 0x12, 0x24, 0x0a,
	0x0e, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x6f, 0x66, 0x5f, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x18,
	0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x6c, 0x69, 0x73, 0x74, 0x4f, 0x66, 0x53, 0x6b, 0x69,
	0x6c, 0x6c, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x10, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x10, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x10, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x42, 0x86, 0x01, 0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x2e, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x42,
	0x0c, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x20, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x76,
	0x31, 0xa2, 0x02, 0x03, 0x54, 0x50, 0x58, 0xaa, 0x02, 0x16, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x2e, 0x67, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x56, 0x31,
	0xca, 0x02, 0x16, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x5c, 0x67, 0x6f, 0x5c, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5c, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_proto_profile_v1_profile_proto_rawDescOnce sync.Once
	file_proto_profile_v1_profile_proto_rawDescData = file_proto_profile_v1_profile_proto_rawDesc
)

func file_proto_profile_v1_profile_proto_rawDescGZIP() []byte {
	file_proto_profile_v1_profile_proto_rawDescOnce.Do(func() {
		file_proto_profile_v1_profile_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_profile_v1_profile_proto_rawDescData)
	})
	return file_proto_profile_v1_profile_proto_rawDescData
}

var file_proto_profile_v1_profile_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_profile_v1_profile_proto_goTypes = []interface{}{
	(*Profile)(nil), // 0: proto.profile.v1.Profile
}
var file_proto_profile_v1_profile_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_profile_v1_profile_proto_init() }
func file_proto_profile_v1_profile_proto_init() {
	if File_proto_profile_v1_profile_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_profile_v1_profile_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Profile); i {
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
			RawDescriptor: file_proto_profile_v1_profile_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_profile_v1_profile_proto_goTypes,
		DependencyIndexes: file_proto_profile_v1_profile_proto_depIdxs,
		MessageInfos:      file_proto_profile_v1_profile_proto_msgTypes,
	}.Build()
	File_proto_profile_v1_profile_proto = out.File
	file_proto_profile_v1_profile_proto_rawDesc = nil
	file_proto_profile_v1_profile_proto_goTypes = nil
	file_proto_profile_v1_profile_proto_depIdxs = nil
}

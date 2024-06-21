// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: multi/v1/game_type.proto

package multiv1

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

type Game struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GameId        int64  `protobuf:"varint,1,opt,name=game_id,json=gameId,proto3" json:"game_id,omitempty"`
	Name          string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Password      string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	HostIpAddress string `protobuf:"bytes,4,opt,name=host_ip_address,json=hostIpAddress,proto3" json:"host_ip_address,omitempty"`
	MapId         int64  `protobuf:"varint,5,opt,name=map_id,json=mapId,proto3" json:"map_id,omitempty"`
	CreatedBy     int64  `protobuf:"varint,6,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	HostUserId    int64  `protobuf:"varint,7,opt,name=host_user_id,json=hostUserId,proto3" json:"host_user_id,omitempty"`
}

func (x *Game) Reset() {
	*x = Game{}
	if protoimpl.UnsafeEnabled {
		mi := &file_multi_v1_game_type_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Game) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Game) ProtoMessage() {}

func (x *Game) ProtoReflect() protoreflect.Message {
	mi := &file_multi_v1_game_type_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Game.ProtoReflect.Descriptor instead.
func (*Game) Descriptor() ([]byte, []int) {
	return file_multi_v1_game_type_proto_rawDescGZIP(), []int{0}
}

func (x *Game) GetGameId() int64 {
	if x != nil {
		return x.GameId
	}
	return 0
}

func (x *Game) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Game) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *Game) GetHostIpAddress() string {
	if x != nil {
		return x.HostIpAddress
	}
	return ""
}

func (x *Game) GetMapId() int64 {
	if x != nil {
		return x.MapId
	}
	return 0
}

func (x *Game) GetCreatedBy() int64 {
	if x != nil {
		return x.CreatedBy
	}
	return 0
}

func (x *Game) GetHostUserId() int64 {
	if x != nil {
		return x.HostUserId
	}
	return 0
}

type Player struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId        int64  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Username      string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	CharacterName string `protobuf:"bytes,3,opt,name=character_name,json=characterName,proto3" json:"character_name,omitempty"`
	ClassType     int64  `protobuf:"varint,4,opt,name=class_type,json=classType,proto3" json:"class_type,omitempty"`
	IpAddress     string `protobuf:"bytes,5,opt,name=ip_address,json=ipAddress,proto3" json:"ip_address,omitempty"`
}

func (x *Player) Reset() {
	*x = Player{}
	if protoimpl.UnsafeEnabled {
		mi := &file_multi_v1_game_type_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Player) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Player) ProtoMessage() {}

func (x *Player) ProtoReflect() protoreflect.Message {
	mi := &file_multi_v1_game_type_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Player.ProtoReflect.Descriptor instead.
func (*Player) Descriptor() ([]byte, []int) {
	return file_multi_v1_game_type_proto_rawDescGZIP(), []int{1}
}

func (x *Player) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Player) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Player) GetCharacterName() string {
	if x != nil {
		return x.CharacterName
	}
	return ""
}

func (x *Player) GetClassType() int64 {
	if x != nil {
		return x.ClassType
	}
	return 0
}

func (x *Player) GetIpAddress() string {
	if x != nil {
		return x.IpAddress
	}
	return ""
}

var File_multi_v1_game_type_proto protoreflect.FileDescriptor

var file_multi_v1_game_type_proto_rawDesc = []byte{
	0x0a, 0x18, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6d, 0x75, 0x6c, 0x74,
	0x69, 0x2e, 0x76, 0x31, 0x22, 0xcf, 0x01, 0x0a, 0x04, 0x47, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a,
	0x07, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x67, 0x61, 0x6d, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x69,
	0x70, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x68, 0x6f, 0x73, 0x74, 0x49, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x15,
	0x0a, 0x06, 0x6d, 0x61, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x6d, 0x61, 0x70, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x62, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x42, 0x79, 0x12, 0x20, 0x0a, 0x0c, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x68, 0x6f, 0x73, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0xa2, 0x01, 0x0a, 0x06, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63,
	0x74, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x69, 0x70, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x69, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x92, 0x01, 0x0a, 0x0c,
	0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x2e, 0x76, 0x31, 0x42, 0x0d, 0x47, 0x61,
	0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x32, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x6d, 0x73, 0x70, 0x65,
	0x6c, 0x6c, 0x2f, 0x67, 0x6c, 0x61, 0x64, 0x69, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x67, 0x65, 0x6e,
	0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x2f, 0x76, 0x31, 0x3b, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x76,
	0x31, 0xa2, 0x02, 0x03, 0x4d, 0x58, 0x58, 0xaa, 0x02, 0x08, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x2e,
	0x56, 0x31, 0xca, 0x02, 0x08, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x14,
	0x4d, 0x75, 0x6c, 0x74, 0x69, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x09, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x3a, 0x3a, 0x56, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_multi_v1_game_type_proto_rawDescOnce sync.Once
	file_multi_v1_game_type_proto_rawDescData = file_multi_v1_game_type_proto_rawDesc
)

func file_multi_v1_game_type_proto_rawDescGZIP() []byte {
	file_multi_v1_game_type_proto_rawDescOnce.Do(func() {
		file_multi_v1_game_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_multi_v1_game_type_proto_rawDescData)
	})
	return file_multi_v1_game_type_proto_rawDescData
}

var file_multi_v1_game_type_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_multi_v1_game_type_proto_goTypes = []any{
	(*Game)(nil),   // 0: multi.v1.Game
	(*Player)(nil), // 1: multi.v1.Player
}
var file_multi_v1_game_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_multi_v1_game_type_proto_init() }
func file_multi_v1_game_type_proto_init() {
	if File_multi_v1_game_type_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_multi_v1_game_type_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Game); i {
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
		file_multi_v1_game_type_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Player); i {
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
			RawDescriptor: file_multi_v1_game_type_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_multi_v1_game_type_proto_goTypes,
		DependencyIndexes: file_multi_v1_game_type_proto_depIdxs,
		MessageInfos:      file_multi_v1_game_type_proto_msgTypes,
	}.Build()
	File_multi_v1_game_type_proto = out.File
	file_multi_v1_game_type_proto_rawDesc = nil
	file_multi_v1_game_type_proto_goTypes = nil
	file_multi_v1_game_type_proto_depIdxs = nil
}

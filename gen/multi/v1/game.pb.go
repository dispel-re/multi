// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: multi/v1/game.proto

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

type CreateGameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId        int64  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	GameName      string `protobuf:"bytes,2,opt,name=game_name,json=gameName,proto3" json:"game_name,omitempty"`
	Password      string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	HostIpAddress string `protobuf:"bytes,4,opt,name=host_ip_address,json=hostIpAddress,proto3" json:"host_ip_address,omitempty"`
	MapId         int64  `protobuf:"varint,5,opt,name=map_id,json=mapId,proto3" json:"map_id,omitempty"`
}

func (x *CreateGameRequest) Reset() {
	*x = CreateGameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_multi_v1_game_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGameRequest) ProtoMessage() {}

func (x *CreateGameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_multi_v1_game_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGameRequest.ProtoReflect.Descriptor instead.
func (*CreateGameRequest) Descriptor() ([]byte, []int) {
	return file_multi_v1_game_proto_rawDescGZIP(), []int{0}
}

func (x *CreateGameRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CreateGameRequest) GetGameName() string {
	if x != nil {
		return x.GameName
	}
	return ""
}

func (x *CreateGameRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *CreateGameRequest) GetHostIpAddress() string {
	if x != nil {
		return x.HostIpAddress
	}
	return ""
}

func (x *CreateGameRequest) GetMapId() int64 {
	if x != nil {
		return x.MapId
	}
	return 0
}

type CreateGameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Game *Game `protobuf:"bytes,1,opt,name=game,proto3" json:"game,omitempty"`
}

func (x *CreateGameResponse) Reset() {
	*x = CreateGameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_multi_v1_game_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGameResponse) ProtoMessage() {}

func (x *CreateGameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_multi_v1_game_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGameResponse.ProtoReflect.Descriptor instead.
func (*CreateGameResponse) Descriptor() ([]byte, []int) {
	return file_multi_v1_game_proto_rawDescGZIP(), []int{1}
}

func (x *CreateGameResponse) GetGame() *Game {
	if x != nil {
		return x.Game
	}
	return nil
}

type GetGameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   int64  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	GameName string `protobuf:"bytes,2,opt,name=game_name,json=gameName,proto3" json:"game_name,omitempty"`
}

func (x *GetGameRequest) Reset() {
	*x = GetGameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_multi_v1_game_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGameRequest) ProtoMessage() {}

func (x *GetGameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_multi_v1_game_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGameRequest.ProtoReflect.Descriptor instead.
func (*GetGameRequest) Descriptor() ([]byte, []int) {
	return file_multi_v1_game_proto_rawDescGZIP(), []int{2}
}

func (x *GetGameRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *GetGameRequest) GetGameName() string {
	if x != nil {
		return x.GameName
	}
	return ""
}

type GetGameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Game *Game `protobuf:"bytes,1,opt,name=game,proto3" json:"game,omitempty"`
}

func (x *GetGameResponse) Reset() {
	*x = GetGameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_multi_v1_game_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGameResponse) ProtoMessage() {}

func (x *GetGameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_multi_v1_game_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGameResponse.ProtoReflect.Descriptor instead.
func (*GetGameResponse) Descriptor() ([]byte, []int) {
	return file_multi_v1_game_proto_rawDescGZIP(), []int{3}
}

func (x *GetGameResponse) GetGame() *Game {
	if x != nil {
		return x.Game
	}
	return nil
}

type ListGamesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *ListGamesRequest) Reset() {
	*x = ListGamesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_multi_v1_game_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListGamesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListGamesRequest) ProtoMessage() {}

func (x *ListGamesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_multi_v1_game_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListGamesRequest.ProtoReflect.Descriptor instead.
func (*ListGamesRequest) Descriptor() ([]byte, []int) {
	return file_multi_v1_game_proto_rawDescGZIP(), []int{4}
}

func (x *ListGamesRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type ListGamesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Games []*Game `protobuf:"bytes,1,rep,name=games,proto3" json:"games,omitempty"`
}

func (x *ListGamesResponse) Reset() {
	*x = ListGamesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_multi_v1_game_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListGamesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListGamesResponse) ProtoMessage() {}

func (x *ListGamesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_multi_v1_game_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListGamesResponse.ProtoReflect.Descriptor instead.
func (*ListGamesResponse) Descriptor() ([]byte, []int) {
	return file_multi_v1_game_proto_rawDescGZIP(), []int{5}
}

func (x *ListGamesResponse) GetGames() []*Game {
	if x != nil {
		return x.Games
	}
	return nil
}

type JoinGameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId      int64  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	CharacterId int64  `protobuf:"varint,2,opt,name=character_id,json=characterId,proto3" json:"character_id,omitempty"`
	GameRoomId  int64  `protobuf:"varint,3,opt,name=game_room_id,json=gameRoomId,proto3" json:"game_room_id,omitempty"`
	IpAddress   string `protobuf:"bytes,4,opt,name=ip_address,json=ipAddress,proto3" json:"ip_address,omitempty"`
}

func (x *JoinGameRequest) Reset() {
	*x = JoinGameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_multi_v1_game_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinGameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinGameRequest) ProtoMessage() {}

func (x *JoinGameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_multi_v1_game_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinGameRequest.ProtoReflect.Descriptor instead.
func (*JoinGameRequest) Descriptor() ([]byte, []int) {
	return file_multi_v1_game_proto_rawDescGZIP(), []int{6}
}

func (x *JoinGameRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *JoinGameRequest) GetCharacterId() int64 {
	if x != nil {
		return x.CharacterId
	}
	return 0
}

func (x *JoinGameRequest) GetGameRoomId() int64 {
	if x != nil {
		return x.GameRoomId
	}
	return 0
}

func (x *JoinGameRequest) GetIpAddress() string {
	if x != nil {
		return x.IpAddress
	}
	return ""
}

type JoinGameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *JoinGameResponse) Reset() {
	*x = JoinGameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_multi_v1_game_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinGameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinGameResponse) ProtoMessage() {}

func (x *JoinGameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_multi_v1_game_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinGameResponse.ProtoReflect.Descriptor instead.
func (*JoinGameResponse) Descriptor() ([]byte, []int) {
	return file_multi_v1_game_proto_rawDescGZIP(), []int{7}
}

type ListPlayersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GameRoomId int64 `protobuf:"varint,1,opt,name=game_room_id,json=gameRoomId,proto3" json:"game_room_id,omitempty"`
}

func (x *ListPlayersRequest) Reset() {
	*x = ListPlayersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_multi_v1_game_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPlayersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPlayersRequest) ProtoMessage() {}

func (x *ListPlayersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_multi_v1_game_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPlayersRequest.ProtoReflect.Descriptor instead.
func (*ListPlayersRequest) Descriptor() ([]byte, []int) {
	return file_multi_v1_game_proto_rawDescGZIP(), []int{8}
}

func (x *ListPlayersRequest) GetGameRoomId() int64 {
	if x != nil {
		return x.GameRoomId
	}
	return 0
}

type ListPlayersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Players []*Player `protobuf:"bytes,1,rep,name=players,proto3" json:"players,omitempty"`
}

func (x *ListPlayersResponse) Reset() {
	*x = ListPlayersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_multi_v1_game_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPlayersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPlayersResponse) ProtoMessage() {}

func (x *ListPlayersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_multi_v1_game_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPlayersResponse.ProtoReflect.Descriptor instead.
func (*ListPlayersResponse) Descriptor() ([]byte, []int) {
	return file_multi_v1_game_proto_rawDescGZIP(), []int{9}
}

func (x *ListPlayersResponse) GetPlayers() []*Player {
	if x != nil {
		return x.Players
	}
	return nil
}

var File_multi_v1_game_proto protoreflect.FileDescriptor

var file_multi_v1_game_proto_rawDesc = []byte{
	0x0a, 0x13, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x2e, 0x76, 0x31, 0x1a,
	0x18, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa4, 0x01, 0x0a, 0x11, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x67, 0x61, 0x6d, 0x65,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x67, 0x61, 0x6d,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x12, 0x26, 0x0a, 0x0f, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x69, 0x70, 0x5f, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x68, 0x6f, 0x73, 0x74,
	0x49, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x15, 0x0a, 0x06, 0x6d, 0x61, 0x70,
	0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6d, 0x61, 0x70, 0x49, 0x64,
	0x22, 0x38, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x04, 0x67, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x61, 0x6d, 0x65, 0x52, 0x04, 0x67, 0x61, 0x6d, 0x65, 0x22, 0x46, 0x0a, 0x0e, 0x47, 0x65,
	0x74, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x67, 0x61, 0x6d, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x22, 0x35, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x04, 0x67, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x61, 0x6d, 0x65, 0x52, 0x04, 0x67, 0x61, 0x6d, 0x65, 0x22, 0x2b, 0x0a, 0x10, 0x4c, 0x69, 0x73,
	0x74, 0x47, 0x61, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x39, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x61,
	0x6d, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x05, 0x67,
	0x61, 0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6d, 0x75, 0x6c,
	0x74, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x05, 0x67, 0x61, 0x6d, 0x65,
	0x73, 0x22, 0x8e, 0x01, 0x0a, 0x0f, 0x4a, 0x6f, 0x69, 0x6e, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x21,
	0x0a, 0x0c, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x20, 0x0a, 0x0c, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x67, 0x61, 0x6d, 0x65, 0x52, 0x6f, 0x6f,
	0x6d, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x70, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x22, 0x12, 0x0a, 0x10, 0x4a, 0x6f, 0x69, 0x6e, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x36, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0c,
	0x67, 0x61, 0x6d, 0x65, 0x5f, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0a, 0x67, 0x61, 0x6d, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x22, 0x41,
	0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x73, 0x32, 0xf5, 0x02, 0x0a, 0x0b, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x40, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x2e, 0x6d,
	0x75, 0x6c, 0x74, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x73,
	0x12, 0x1a, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x47, 0x61, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x6d,
	0x75, 0x6c, 0x74, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x61, 0x6d, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x0a, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x2e, 0x6d, 0x75, 0x6c, 0x74,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x61, 0x6d, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x08, 0x4a, 0x6f, 0x69, 0x6e, 0x47, 0x61,
	0x6d, 0x65, 0x12, 0x19, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x6f,
	0x69, 0x6e, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e,
	0x6d, 0x75, 0x6c, 0x74, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x47, 0x61, 0x6d,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4c, 0x0a, 0x0b, 0x4c,
	0x69, 0x73, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x12, 0x1c, 0x2e, 0x6d, 0x75, 0x6c,
	0x74, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x8e, 0x01, 0x0a, 0x0c, 0x63, 0x6f,
	0x6d, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x2e, 0x76, 0x31, 0x42, 0x09, 0x47, 0x61, 0x6d, 0x65,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x6d, 0x73, 0x70, 0x65, 0x6c, 0x6c, 0x2f, 0x67, 0x6c, 0x61,
	0x64, 0x69, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69,
	0x2f, 0x76, 0x31, 0x3b, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x4d, 0x58,
	0x58, 0xaa, 0x02, 0x08, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x08, 0x4d,
	0x75, 0x6c, 0x74, 0x69, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x14, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x5c,
	0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x09, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_multi_v1_game_proto_rawDescOnce sync.Once
	file_multi_v1_game_proto_rawDescData = file_multi_v1_game_proto_rawDesc
)

func file_multi_v1_game_proto_rawDescGZIP() []byte {
	file_multi_v1_game_proto_rawDescOnce.Do(func() {
		file_multi_v1_game_proto_rawDescData = protoimpl.X.CompressGZIP(file_multi_v1_game_proto_rawDescData)
	})
	return file_multi_v1_game_proto_rawDescData
}

var file_multi_v1_game_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_multi_v1_game_proto_goTypes = []any{
	(*CreateGameRequest)(nil),   // 0: multi.v1.CreateGameRequest
	(*CreateGameResponse)(nil),  // 1: multi.v1.CreateGameResponse
	(*GetGameRequest)(nil),      // 2: multi.v1.GetGameRequest
	(*GetGameResponse)(nil),     // 3: multi.v1.GetGameResponse
	(*ListGamesRequest)(nil),    // 4: multi.v1.ListGamesRequest
	(*ListGamesResponse)(nil),   // 5: multi.v1.ListGamesResponse
	(*JoinGameRequest)(nil),     // 6: multi.v1.JoinGameRequest
	(*JoinGameResponse)(nil),    // 7: multi.v1.JoinGameResponse
	(*ListPlayersRequest)(nil),  // 8: multi.v1.ListPlayersRequest
	(*ListPlayersResponse)(nil), // 9: multi.v1.ListPlayersResponse
	(*Game)(nil),                // 10: multi.v1.Game
	(*Player)(nil),              // 11: multi.v1.Player
}
var file_multi_v1_game_proto_depIdxs = []int32{
	10, // 0: multi.v1.CreateGameResponse.game:type_name -> multi.v1.Game
	10, // 1: multi.v1.GetGameResponse.game:type_name -> multi.v1.Game
	10, // 2: multi.v1.ListGamesResponse.games:type_name -> multi.v1.Game
	11, // 3: multi.v1.ListPlayersResponse.players:type_name -> multi.v1.Player
	2,  // 4: multi.v1.GameService.GetGame:input_type -> multi.v1.GetGameRequest
	4,  // 5: multi.v1.GameService.ListGames:input_type -> multi.v1.ListGamesRequest
	0,  // 6: multi.v1.GameService.CreateGame:input_type -> multi.v1.CreateGameRequest
	6,  // 7: multi.v1.GameService.JoinGame:input_type -> multi.v1.JoinGameRequest
	8,  // 8: multi.v1.GameService.ListPlayers:input_type -> multi.v1.ListPlayersRequest
	3,  // 9: multi.v1.GameService.GetGame:output_type -> multi.v1.GetGameResponse
	5,  // 10: multi.v1.GameService.ListGames:output_type -> multi.v1.ListGamesResponse
	1,  // 11: multi.v1.GameService.CreateGame:output_type -> multi.v1.CreateGameResponse
	7,  // 12: multi.v1.GameService.JoinGame:output_type -> multi.v1.JoinGameResponse
	9,  // 13: multi.v1.GameService.ListPlayers:output_type -> multi.v1.ListPlayersResponse
	9,  // [9:14] is the sub-list for method output_type
	4,  // [4:9] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_multi_v1_game_proto_init() }
func file_multi_v1_game_proto_init() {
	if File_multi_v1_game_proto != nil {
		return
	}
	file_multi_v1_game_type_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_multi_v1_game_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CreateGameRequest); i {
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
		file_multi_v1_game_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CreateGameResponse); i {
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
		file_multi_v1_game_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetGameRequest); i {
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
		file_multi_v1_game_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetGameResponse); i {
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
		file_multi_v1_game_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*ListGamesRequest); i {
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
		file_multi_v1_game_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*ListGamesResponse); i {
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
		file_multi_v1_game_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*JoinGameRequest); i {
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
		file_multi_v1_game_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*JoinGameResponse); i {
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
		file_multi_v1_game_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*ListPlayersRequest); i {
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
		file_multi_v1_game_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*ListPlayersResponse); i {
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
			RawDescriptor: file_multi_v1_game_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_multi_v1_game_proto_goTypes,
		DependencyIndexes: file_multi_v1_game_proto_depIdxs,
		MessageInfos:      file_multi_v1_game_proto_msgTypes,
	}.Build()
	File_multi_v1_game_proto = out.File
	file_multi_v1_game_proto_rawDesc = nil
	file_multi_v1_game_proto_goTypes = nil
	file_multi_v1_game_proto_depIdxs = nil
}

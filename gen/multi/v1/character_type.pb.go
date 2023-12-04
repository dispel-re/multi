// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: multi/v1/character_type.proto

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

type Gender int32

const (
	Gender_Male   Gender = 0
	Gender_Female Gender = 1
)

// Enum value maps for Gender.
var (
	Gender_name = map[int32]string{
		0: "Male",
		1: "Female",
	}
	Gender_value = map[string]int32{
		"Male":   0,
		"Female": 1,
	}
)

func (x Gender) Enum() *Gender {
	p := new(Gender)
	*p = x
	return p
}

func (x Gender) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Gender) Descriptor() protoreflect.EnumDescriptor {
	return file_multi_v1_character_type_proto_enumTypes[0].Descriptor()
}

func (Gender) Type() protoreflect.EnumType {
	return &file_multi_v1_character_type_proto_enumTypes[0]
}

func (x Gender) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Gender.Descriptor instead.
func (Gender) EnumDescriptor() ([]byte, []int) {
	return file_multi_v1_character_type_proto_rawDescGZIP(), []int{0}
}

type ClassType int32

const (
	ClassType_Knight  ClassType = 0
	ClassType_Warrior ClassType = 1
	ClassType_Archer  ClassType = 2
	ClassType_Mage    ClassType = 3
)

// Enum value maps for ClassType.
var (
	ClassType_name = map[int32]string{
		0: "Knight",
		1: "Warrior",
		2: "Archer",
		3: "Mage",
	}
	ClassType_value = map[string]int32{
		"Knight":  0,
		"Warrior": 1,
		"Archer":  2,
		"Mage":    3,
	}
)

func (x ClassType) Enum() *ClassType {
	p := new(ClassType)
	*p = x
	return p
}

func (x ClassType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ClassType) Descriptor() protoreflect.EnumDescriptor {
	return file_multi_v1_character_type_proto_enumTypes[1].Descriptor()
}

func (ClassType) Type() protoreflect.EnumType {
	return &file_multi_v1_character_type_proto_enumTypes[1]
}

func (x ClassType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ClassType.Descriptor instead.
func (ClassType) EnumDescriptor() ([]byte, []int) {
	return file_multi_v1_character_type_proto_rawDescGZIP(), []int{1}
}

type SkinCarnation int32

const (
	SkinCarnation_SkinCarnationNone SkinCarnation = 0
	SkinCarnation_MaleBeige         SkinCarnation = 102
	SkinCarnation_MalePink          SkinCarnation = 103
	SkinCarnation_MaleLightBrown    SkinCarnation = 104
	SkinCarnation_MaleBrown         SkinCarnation = 105
	SkinCarnation_MaleGray          SkinCarnation = 106
	SkinCarnation_FemaleBeige       SkinCarnation = 107
	SkinCarnation_FemalePink        SkinCarnation = 108
	SkinCarnation_FemaleLightBrown  SkinCarnation = 109
	SkinCarnation_FemaleBrown       SkinCarnation = 110
	SkinCarnation_FemaleGray        SkinCarnation = 111
)

// Enum value maps for SkinCarnation.
var (
	SkinCarnation_name = map[int32]string{
		0:   "SkinCarnationNone",
		102: "MaleBeige",
		103: "MalePink",
		104: "MaleLightBrown",
		105: "MaleBrown",
		106: "MaleGray",
		107: "FemaleBeige",
		108: "FemalePink",
		109: "FemaleLightBrown",
		110: "FemaleBrown",
		111: "FemaleGray",
	}
	SkinCarnation_value = map[string]int32{
		"SkinCarnationNone": 0,
		"MaleBeige":         102,
		"MalePink":          103,
		"MaleLightBrown":    104,
		"MaleBrown":         105,
		"MaleGray":          106,
		"FemaleBeige":       107,
		"FemalePink":        108,
		"FemaleLightBrown":  109,
		"FemaleBrown":       110,
		"FemaleGray":        111,
	}
)

func (x SkinCarnation) Enum() *SkinCarnation {
	p := new(SkinCarnation)
	*p = x
	return p
}

func (x SkinCarnation) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SkinCarnation) Descriptor() protoreflect.EnumDescriptor {
	return file_multi_v1_character_type_proto_enumTypes[2].Descriptor()
}

func (SkinCarnation) Type() protoreflect.EnumType {
	return &file_multi_v1_character_type_proto_enumTypes[2]
}

func (x SkinCarnation) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SkinCarnation.Descriptor instead.
func (SkinCarnation) EnumDescriptor() ([]byte, []int) {
	return file_multi_v1_character_type_proto_rawDescGZIP(), []int{2}
}

type HairStyle int32

const (
	HairStyle_HairStyleNone       HairStyle = 0
	HairStyle_MaleShortWhite      HairStyle = 112
	HairStyle_MaleShortBrown      HairStyle = 113
	HairStyle_MaleShortBlack      HairStyle = 114
	HairStyle_MaleShortRed        HairStyle = 115
	HairStyle_MaleShortGray       HairStyle = 116
	HairStyle_MaleMediumWhite     HairStyle = 117
	HairStyle_MaleMediumBrown     HairStyle = 118
	HairStyle_MaleMediumBlack     HairStyle = 119
	HairStyle_MaleMediumRed       HairStyle = 120
	HairStyle_MaleMediumGray      HairStyle = 121
	HairStyle_MaleLongWhite       HairStyle = 122
	HairStyle_MaleLongBrown       HairStyle = 123
	HairStyle_MaleLongBlack       HairStyle = 124
	HairStyle_MaleLongRed         HairStyle = 125
	HairStyle_MaleLongGray        HairStyle = 126
	HairStyle_MaleBald            HairStyle = 200
	HairStyle_FemaleShortWhite    HairStyle = 127
	HairStyle_FemaleShortBrown    HairStyle = 128
	HairStyle_FemaleShortBlack    HairStyle = 129
	HairStyle_FemaleShortRed      HairStyle = 130
	HairStyle_FemaleShortGray     HairStyle = 131
	HairStyle_FemaleMediumWhite   HairStyle = 132
	HairStyle_FemaleMediumBrown   HairStyle = 133
	HairStyle_FemaleMediumBlack   HairStyle = 134
	HairStyle_FemaleMediumRed     HairStyle = 135
	HairStyle_FemaleMediumGray    HairStyle = 136
	HairStyle_FemalePonytailWhite HairStyle = 137
	HairStyle_FemalePonytailBrown HairStyle = 138
	HairStyle_FemalePonytailBlack HairStyle = 139
	HairStyle_FemalePonytailRed   HairStyle = 140
	HairStyle_FemalePonytailGray  HairStyle = 141
	HairStyle_FemaleLongWhite     HairStyle = 142
	HairStyle_FemaleLongBrown     HairStyle = 143
	HairStyle_FemaleLongBlack     HairStyle = 144
	HairStyle_FemaleLongRed       HairStyle = 145
	HairStyle_FemaleLongGray      HairStyle = 146
)

// Enum value maps for HairStyle.
var (
	HairStyle_name = map[int32]string{
		0:   "HairStyleNone",
		112: "MaleShortWhite",
		113: "MaleShortBrown",
		114: "MaleShortBlack",
		115: "MaleShortRed",
		116: "MaleShortGray",
		117: "MaleMediumWhite",
		118: "MaleMediumBrown",
		119: "MaleMediumBlack",
		120: "MaleMediumRed",
		121: "MaleMediumGray",
		122: "MaleLongWhite",
		123: "MaleLongBrown",
		124: "MaleLongBlack",
		125: "MaleLongRed",
		126: "MaleLongGray",
		200: "MaleBald",
		127: "FemaleShortWhite",
		128: "FemaleShortBrown",
		129: "FemaleShortBlack",
		130: "FemaleShortRed",
		131: "FemaleShortGray",
		132: "FemaleMediumWhite",
		133: "FemaleMediumBrown",
		134: "FemaleMediumBlack",
		135: "FemaleMediumRed",
		136: "FemaleMediumGray",
		137: "FemalePonytailWhite",
		138: "FemalePonytailBrown",
		139: "FemalePonytailBlack",
		140: "FemalePonytailRed",
		141: "FemalePonytailGray",
		142: "FemaleLongWhite",
		143: "FemaleLongBrown",
		144: "FemaleLongBlack",
		145: "FemaleLongRed",
		146: "FemaleLongGray",
	}
	HairStyle_value = map[string]int32{
		"HairStyleNone":       0,
		"MaleShortWhite":      112,
		"MaleShortBrown":      113,
		"MaleShortBlack":      114,
		"MaleShortRed":        115,
		"MaleShortGray":       116,
		"MaleMediumWhite":     117,
		"MaleMediumBrown":     118,
		"MaleMediumBlack":     119,
		"MaleMediumRed":       120,
		"MaleMediumGray":      121,
		"MaleLongWhite":       122,
		"MaleLongBrown":       123,
		"MaleLongBlack":       124,
		"MaleLongRed":         125,
		"MaleLongGray":        126,
		"MaleBald":            200,
		"FemaleShortWhite":    127,
		"FemaleShortBrown":    128,
		"FemaleShortBlack":    129,
		"FemaleShortRed":      130,
		"FemaleShortGray":     131,
		"FemaleMediumWhite":   132,
		"FemaleMediumBrown":   133,
		"FemaleMediumBlack":   134,
		"FemaleMediumRed":     135,
		"FemaleMediumGray":    136,
		"FemalePonytailWhite": 137,
		"FemalePonytailBrown": 138,
		"FemalePonytailBlack": 139,
		"FemalePonytailRed":   140,
		"FemalePonytailGray":  141,
		"FemaleLongWhite":     142,
		"FemaleLongBrown":     143,
		"FemaleLongBlack":     144,
		"FemaleLongRed":       145,
		"FemaleLongGray":      146,
	}
)

func (x HairStyle) Enum() *HairStyle {
	p := new(HairStyle)
	*p = x
	return p
}

func (x HairStyle) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HairStyle) Descriptor() protoreflect.EnumDescriptor {
	return file_multi_v1_character_type_proto_enumTypes[3].Descriptor()
}

func (HairStyle) Type() protoreflect.EnumType {
	return &file_multi_v1_character_type_proto_enumTypes[3]
}

func (x HairStyle) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HairStyle.Descriptor instead.
func (HairStyle) EnumDescriptor() ([]byte, []int) {
	return file_multi_v1_character_type_proto_rawDescGZIP(), []int{3}
}

type Character struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId        int64  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	CharacterId   int64  `protobuf:"varint,2,opt,name=character_id,json=characterId,proto3" json:"character_id,omitempty"`
	CharacterName int64  `protobuf:"varint,3,opt,name=character_name,json=characterName,proto3" json:"character_name,omitempty"`
	Info          []byte `protobuf:"bytes,4,opt,name=info,proto3" json:"info,omitempty"`
	Inventory     []byte `protobuf:"bytes,5,opt,name=inventory,proto3" json:"inventory,omitempty"`
	Spell         []byte `protobuf:"bytes,6,opt,name=spell,proto3" json:"spell,omitempty"`
}

func (x *Character) Reset() {
	*x = Character{}
	if protoimpl.UnsafeEnabled {
		mi := &file_multi_v1_character_type_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Character) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Character) ProtoMessage() {}

func (x *Character) ProtoReflect() protoreflect.Message {
	mi := &file_multi_v1_character_type_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Character.ProtoReflect.Descriptor instead.
func (*Character) Descriptor() ([]byte, []int) {
	return file_multi_v1_character_type_proto_rawDescGZIP(), []int{0}
}

func (x *Character) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Character) GetCharacterId() int64 {
	if x != nil {
		return x.CharacterId
	}
	return 0
}

func (x *Character) GetCharacterName() int64 {
	if x != nil {
		return x.CharacterName
	}
	return 0
}

func (x *Character) GetInfo() []byte {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *Character) GetInventory() []byte {
	if x != nil {
		return x.Inventory
	}
	return nil
}

func (x *Character) GetSpell() []byte {
	if x != nil {
		return x.Spell
	}
	return nil
}

type CharacterInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Strength             int64  `protobuf:"varint,1,opt,name=strength,proto3" json:"strength,omitempty"`
	Agility              int64  `protobuf:"varint,2,opt,name=agility,proto3" json:"agility,omitempty"`
	Wisdom               int64  `protobuf:"varint,3,opt,name=wisdom,proto3" json:"wisdom,omitempty"`
	Constitution         int64  `protobuf:"varint,4,opt,name=constitution,proto3" json:"constitution,omitempty"`
	HealthPoints         int64  `protobuf:"varint,5,opt,name=health_points,json=healthPoints,proto3" json:"health_points,omitempty"`
	MagicPoints          int64  `protobuf:"varint,6,opt,name=magic_points,json=magicPoints,proto3" json:"magic_points,omitempty"`
	ExperiencePoints     int64  `protobuf:"varint,7,opt,name=experience_points,json=experiencePoints,proto3" json:"experience_points,omitempty"`
	Money                int64  `protobuf:"varint,8,opt,name=money,proto3" json:"money,omitempty"`
	ScorePoints          int64  `protobuf:"varint,9,opt,name=score_points,json=scorePoints,proto3" json:"score_points,omitempty"`
	ClassType            int64  `protobuf:"varint,10,opt,name=class_type,json=classType,proto3" json:"class_type,omitempty"`
	SkinCarnation        int64  `protobuf:"varint,11,opt,name=skin_carnation,json=skinCarnation,proto3" json:"skin_carnation,omitempty"`
	HairStyle            int64  `protobuf:"varint,12,opt,name=hair_style,json=hairStyle,proto3" json:"hair_style,omitempty"`
	LightArmourLegs      int64  `protobuf:"varint,13,opt,name=light_armour_legs,json=lightArmourLegs,proto3" json:"light_armour_legs,omitempty"`
	LightArmourTorso     int64  `protobuf:"varint,14,opt,name=light_armour_torso,json=lightArmourTorso,proto3" json:"light_armour_torso,omitempty"`
	LightArmourHands     int64  `protobuf:"varint,15,opt,name=light_armour_hands,json=lightArmourHands,proto3" json:"light_armour_hands,omitempty"`
	LightArmourBoots     int64  `protobuf:"varint,16,opt,name=light_armour_boots,json=lightArmourBoots,proto3" json:"light_armour_boots,omitempty"`
	FullArmour           int64  `protobuf:"varint,17,opt,name=full_armour,json=fullArmour,proto3" json:"full_armour,omitempty"`
	ArmourEmblem         int64  `protobuf:"varint,18,opt,name=armour_emblem,json=armourEmblem,proto3" json:"armour_emblem,omitempty"`
	Helmet               int64  `protobuf:"varint,19,opt,name=helmet,proto3" json:"helmet,omitempty"`
	SecondaryWeapon      int64  `protobuf:"varint,20,opt,name=secondary_weapon,json=secondaryWeapon,proto3" json:"secondary_weapon,omitempty"`
	PrimaryWeapon        int64  `protobuf:"varint,21,opt,name=primary_weapon,json=primaryWeapon,proto3" json:"primary_weapon,omitempty"`
	Shield               int64  `protobuf:"varint,22,opt,name=shield,proto3" json:"shield,omitempty"`
	UnknownEquipmentSlot int64  `protobuf:"varint,23,opt,name=unknown_equipment_slot,json=unknownEquipmentSlot,proto3" json:"unknown_equipment_slot,omitempty"`
	Gender               int64  `protobuf:"varint,24,opt,name=gender,proto3" json:"gender,omitempty"`
	Level                int64  `protobuf:"varint,25,opt,name=level,proto3" json:"level,omitempty"`
	EdgedWeapons         int64  `protobuf:"varint,26,opt,name=edged_weapons,json=edgedWeapons,proto3" json:"edged_weapons,omitempty"`
	BluntedWeapons       int64  `protobuf:"varint,27,opt,name=blunted_weapons,json=bluntedWeapons,proto3" json:"blunted_weapons,omitempty"`
	Archery              int64  `protobuf:"varint,28,opt,name=archery,proto3" json:"archery,omitempty"`
	Polearms             int64  `protobuf:"varint,29,opt,name=polearms,proto3" json:"polearms,omitempty"`
	Wizardry             int64  `protobuf:"varint,30,opt,name=wizardry,proto3" json:"wizardry,omitempty"`
	Kills                []byte `protobuf:"bytes,31,opt,name=kills,proto3" json:"kills,omitempty"`
}

func (x *CharacterInfo) Reset() {
	*x = CharacterInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_multi_v1_character_type_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CharacterInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CharacterInfo) ProtoMessage() {}

func (x *CharacterInfo) ProtoReflect() protoreflect.Message {
	mi := &file_multi_v1_character_type_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CharacterInfo.ProtoReflect.Descriptor instead.
func (*CharacterInfo) Descriptor() ([]byte, []int) {
	return file_multi_v1_character_type_proto_rawDescGZIP(), []int{1}
}

func (x *CharacterInfo) GetStrength() int64 {
	if x != nil {
		return x.Strength
	}
	return 0
}

func (x *CharacterInfo) GetAgility() int64 {
	if x != nil {
		return x.Agility
	}
	return 0
}

func (x *CharacterInfo) GetWisdom() int64 {
	if x != nil {
		return x.Wisdom
	}
	return 0
}

func (x *CharacterInfo) GetConstitution() int64 {
	if x != nil {
		return x.Constitution
	}
	return 0
}

func (x *CharacterInfo) GetHealthPoints() int64 {
	if x != nil {
		return x.HealthPoints
	}
	return 0
}

func (x *CharacterInfo) GetMagicPoints() int64 {
	if x != nil {
		return x.MagicPoints
	}
	return 0
}

func (x *CharacterInfo) GetExperiencePoints() int64 {
	if x != nil {
		return x.ExperiencePoints
	}
	return 0
}

func (x *CharacterInfo) GetMoney() int64 {
	if x != nil {
		return x.Money
	}
	return 0
}

func (x *CharacterInfo) GetScorePoints() int64 {
	if x != nil {
		return x.ScorePoints
	}
	return 0
}

func (x *CharacterInfo) GetClassType() int64 {
	if x != nil {
		return x.ClassType
	}
	return 0
}

func (x *CharacterInfo) GetSkinCarnation() int64 {
	if x != nil {
		return x.SkinCarnation
	}
	return 0
}

func (x *CharacterInfo) GetHairStyle() int64 {
	if x != nil {
		return x.HairStyle
	}
	return 0
}

func (x *CharacterInfo) GetLightArmourLegs() int64 {
	if x != nil {
		return x.LightArmourLegs
	}
	return 0
}

func (x *CharacterInfo) GetLightArmourTorso() int64 {
	if x != nil {
		return x.LightArmourTorso
	}
	return 0
}

func (x *CharacterInfo) GetLightArmourHands() int64 {
	if x != nil {
		return x.LightArmourHands
	}
	return 0
}

func (x *CharacterInfo) GetLightArmourBoots() int64 {
	if x != nil {
		return x.LightArmourBoots
	}
	return 0
}

func (x *CharacterInfo) GetFullArmour() int64 {
	if x != nil {
		return x.FullArmour
	}
	return 0
}

func (x *CharacterInfo) GetArmourEmblem() int64 {
	if x != nil {
		return x.ArmourEmblem
	}
	return 0
}

func (x *CharacterInfo) GetHelmet() int64 {
	if x != nil {
		return x.Helmet
	}
	return 0
}

func (x *CharacterInfo) GetSecondaryWeapon() int64 {
	if x != nil {
		return x.SecondaryWeapon
	}
	return 0
}

func (x *CharacterInfo) GetPrimaryWeapon() int64 {
	if x != nil {
		return x.PrimaryWeapon
	}
	return 0
}

func (x *CharacterInfo) GetShield() int64 {
	if x != nil {
		return x.Shield
	}
	return 0
}

func (x *CharacterInfo) GetUnknownEquipmentSlot() int64 {
	if x != nil {
		return x.UnknownEquipmentSlot
	}
	return 0
}

func (x *CharacterInfo) GetGender() int64 {
	if x != nil {
		return x.Gender
	}
	return 0
}

func (x *CharacterInfo) GetLevel() int64 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *CharacterInfo) GetEdgedWeapons() int64 {
	if x != nil {
		return x.EdgedWeapons
	}
	return 0
}

func (x *CharacterInfo) GetBluntedWeapons() int64 {
	if x != nil {
		return x.BluntedWeapons
	}
	return 0
}

func (x *CharacterInfo) GetArchery() int64 {
	if x != nil {
		return x.Archery
	}
	return 0
}

func (x *CharacterInfo) GetPolearms() int64 {
	if x != nil {
		return x.Polearms
	}
	return 0
}

func (x *CharacterInfo) GetWizardry() int64 {
	if x != nil {
		return x.Wizardry
	}
	return 0
}

func (x *CharacterInfo) GetKills() []byte {
	if x != nil {
		return x.Kills
	}
	return nil
}

var File_multi_v1_character_type_proto protoreflect.FileDescriptor

var file_multi_v1_character_type_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x68, 0x61, 0x72, 0x61,
	0x63, 0x74, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x08, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x2e, 0x76, 0x31, 0x22, 0xb6, 0x01, 0x0a, 0x09, 0x43, 0x68,
	0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x21, 0x0a, 0x0c, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x63, 0x68, 0x61,
	0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x6e,
	0x66, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x1c,
	0x0a, 0x09, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x09, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x73, 0x70, 0x65, 0x6c, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x73, 0x70, 0x65,
	0x6c, 0x6c, 0x22, 0xac, 0x08, 0x0a, 0x0d, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x74, 0x72, 0x65, 0x6e, 0x67, 0x74, 0x68,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x74, 0x72, 0x65, 0x6e, 0x67, 0x74, 0x68,
	0x12, 0x18, 0x0a, 0x07, 0x61, 0x67, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x61, 0x67, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x69,
	0x73, 0x64, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x77, 0x69, 0x73, 0x64,
	0x6f, 0x6d, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x69, 0x74, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x69,
	0x74, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68,
	0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x68,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x6d,
	0x61, 0x67, 0x69, 0x63, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0b, 0x6d, 0x61, 0x67, 0x69, 0x63, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x2b,
	0x0a, 0x11, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x65, 0x78, 0x70, 0x65, 0x72,
	0x69, 0x65, 0x6e, 0x63, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6d,
	0x6f, 0x6e, 0x65, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6d, 0x6f, 0x6e, 0x65,
	0x79, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x50, 0x6f,
	0x69, 0x6e, 0x74, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x6b, 0x69, 0x6e, 0x5f, 0x63, 0x61, 0x72, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x73, 0x6b, 0x69,
	0x6e, 0x43, 0x61, 0x72, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x68, 0x61,
	0x69, 0x72, 0x5f, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x68, 0x61, 0x69, 0x72, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x6c, 0x69, 0x67,
	0x68, 0x74, 0x5f, 0x61, 0x72, 0x6d, 0x6f, 0x75, 0x72, 0x5f, 0x6c, 0x65, 0x67, 0x73, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x41, 0x72, 0x6d, 0x6f, 0x75,
	0x72, 0x4c, 0x65, 0x67, 0x73, 0x12, 0x2c, 0x0a, 0x12, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x61,
	0x72, 0x6d, 0x6f, 0x75, 0x72, 0x5f, 0x74, 0x6f, 0x72, 0x73, 0x6f, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x10, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x41, 0x72, 0x6d, 0x6f, 0x75, 0x72, 0x54, 0x6f,
	0x72, 0x73, 0x6f, 0x12, 0x2c, 0x0a, 0x12, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x61, 0x72, 0x6d,
	0x6f, 0x75, 0x72, 0x5f, 0x68, 0x61, 0x6e, 0x64, 0x73, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x10, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x41, 0x72, 0x6d, 0x6f, 0x75, 0x72, 0x48, 0x61, 0x6e, 0x64,
	0x73, 0x12, 0x2c, 0x0a, 0x12, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x61, 0x72, 0x6d, 0x6f, 0x75,
	0x72, 0x5f, 0x62, 0x6f, 0x6f, 0x74, 0x73, 0x18, 0x10, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x6c,
	0x69, 0x67, 0x68, 0x74, 0x41, 0x72, 0x6d, 0x6f, 0x75, 0x72, 0x42, 0x6f, 0x6f, 0x74, 0x73, 0x12,
	0x1f, 0x0a, 0x0b, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x61, 0x72, 0x6d, 0x6f, 0x75, 0x72, 0x18, 0x11,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x66, 0x75, 0x6c, 0x6c, 0x41, 0x72, 0x6d, 0x6f, 0x75, 0x72,
	0x12, 0x23, 0x0a, 0x0d, 0x61, 0x72, 0x6d, 0x6f, 0x75, 0x72, 0x5f, 0x65, 0x6d, 0x62, 0x6c, 0x65,
	0x6d, 0x18, 0x12, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x61, 0x72, 0x6d, 0x6f, 0x75, 0x72, 0x45,
	0x6d, 0x62, 0x6c, 0x65, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x6c, 0x6d, 0x65, 0x74, 0x18,
	0x13, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x68, 0x65, 0x6c, 0x6d, 0x65, 0x74, 0x12, 0x29, 0x0a,
	0x10, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x5f, 0x77, 0x65, 0x61, 0x70, 0x6f,
	0x6e, 0x18, 0x14, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x61,
	0x72, 0x79, 0x57, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x72, 0x69, 0x6d,
	0x61, 0x72, 0x79, 0x5f, 0x77, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x18, 0x15, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0d, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x57, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x68, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x16, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x73, 0x68, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x34, 0x0a, 0x16, 0x75, 0x6e, 0x6b, 0x6e, 0x6f,
	0x77, 0x6e, 0x5f, 0x65, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x6c, 0x6f,
	0x74, 0x18, 0x17, 0x20, 0x01, 0x28, 0x03, 0x52, 0x14, 0x75, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e,
	0x45, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x6c, 0x6f, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x18, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x67,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x19,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x23, 0x0a, 0x0d, 0x65,
	0x64, 0x67, 0x65, 0x64, 0x5f, 0x77, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x73, 0x18, 0x1a, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0c, 0x65, 0x64, 0x67, 0x65, 0x64, 0x57, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x73,
	0x12, 0x27, 0x0a, 0x0f, 0x62, 0x6c, 0x75, 0x6e, 0x74, 0x65, 0x64, 0x5f, 0x77, 0x65, 0x61, 0x70,
	0x6f, 0x6e, 0x73, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x62, 0x6c, 0x75, 0x6e, 0x74,
	0x65, 0x64, 0x57, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x72, 0x63,
	0x68, 0x65, 0x72, 0x79, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x61, 0x72, 0x63, 0x68,
	0x65, 0x72, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x6c, 0x65, 0x61, 0x72, 0x6d, 0x73, 0x18,
	0x1d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x6f, 0x6c, 0x65, 0x61, 0x72, 0x6d, 0x73, 0x12,
	0x1a, 0x0a, 0x08, 0x77, 0x69, 0x7a, 0x61, 0x72, 0x64, 0x72, 0x79, 0x18, 0x1e, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x77, 0x69, 0x7a, 0x61, 0x72, 0x64, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x6b,
	0x69, 0x6c, 0x6c, 0x73, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x6b, 0x69, 0x6c, 0x6c,
	0x73, 0x2a, 0x1e, 0x0a, 0x06, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x08, 0x0a, 0x04, 0x4d,
	0x61, 0x6c, 0x65, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x65, 0x6d, 0x61, 0x6c, 0x65, 0x10,
	0x01, 0x2a, 0x3a, 0x0a, 0x09, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0a,
	0x0a, 0x06, 0x4b, 0x6e, 0x69, 0x67, 0x68, 0x74, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x57, 0x61,
	0x72, 0x72, 0x69, 0x6f, 0x72, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x72, 0x63, 0x68, 0x65,
	0x72, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x4d, 0x61, 0x67, 0x65, 0x10, 0x03, 0x2a, 0xcc, 0x01,
	0x0a, 0x0d, 0x53, 0x6b, 0x69, 0x6e, 0x43, 0x61, 0x72, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x15, 0x0a, 0x11, 0x53, 0x6b, 0x69, 0x6e, 0x43, 0x61, 0x72, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x4e, 0x6f, 0x6e, 0x65, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x4d, 0x61, 0x6c, 0x65, 0x42, 0x65,
	0x69, 0x67, 0x65, 0x10, 0x66, 0x12, 0x0c, 0x0a, 0x08, 0x4d, 0x61, 0x6c, 0x65, 0x50, 0x69, 0x6e,
	0x6b, 0x10, 0x67, 0x12, 0x12, 0x0a, 0x0e, 0x4d, 0x61, 0x6c, 0x65, 0x4c, 0x69, 0x67, 0x68, 0x74,
	0x42, 0x72, 0x6f, 0x77, 0x6e, 0x10, 0x68, 0x12, 0x0d, 0x0a, 0x09, 0x4d, 0x61, 0x6c, 0x65, 0x42,
	0x72, 0x6f, 0x77, 0x6e, 0x10, 0x69, 0x12, 0x0c, 0x0a, 0x08, 0x4d, 0x61, 0x6c, 0x65, 0x47, 0x72,
	0x61, 0x79, 0x10, 0x6a, 0x12, 0x0f, 0x0a, 0x0b, 0x46, 0x65, 0x6d, 0x61, 0x6c, 0x65, 0x42, 0x65,
	0x69, 0x67, 0x65, 0x10, 0x6b, 0x12, 0x0e, 0x0a, 0x0a, 0x46, 0x65, 0x6d, 0x61, 0x6c, 0x65, 0x50,
	0x69, 0x6e, 0x6b, 0x10, 0x6c, 0x12, 0x14, 0x0a, 0x10, 0x46, 0x65, 0x6d, 0x61, 0x6c, 0x65, 0x4c,
	0x69, 0x67, 0x68, 0x74, 0x42, 0x72, 0x6f, 0x77, 0x6e, 0x10, 0x6d, 0x12, 0x0f, 0x0a, 0x0b, 0x46,
	0x65, 0x6d, 0x61, 0x6c, 0x65, 0x42, 0x72, 0x6f, 0x77, 0x6e, 0x10, 0x6e, 0x12, 0x0e, 0x0a, 0x0a,
	0x46, 0x65, 0x6d, 0x61, 0x6c, 0x65, 0x47, 0x72, 0x61, 0x79, 0x10, 0x6f, 0x2a, 0x9e, 0x06, 0x0a,
	0x09, 0x48, 0x61, 0x69, 0x72, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x12, 0x11, 0x0a, 0x0d, 0x48, 0x61,
	0x69, 0x72, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x4e, 0x6f, 0x6e, 0x65, 0x10, 0x00, 0x12, 0x12, 0x0a,
	0x0e, 0x4d, 0x61, 0x6c, 0x65, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x57, 0x68, 0x69, 0x74, 0x65, 0x10,
	0x70, 0x12, 0x12, 0x0a, 0x0e, 0x4d, 0x61, 0x6c, 0x65, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x42, 0x72,
	0x6f, 0x77, 0x6e, 0x10, 0x71, 0x12, 0x12, 0x0a, 0x0e, 0x4d, 0x61, 0x6c, 0x65, 0x53, 0x68, 0x6f,
	0x72, 0x74, 0x42, 0x6c, 0x61, 0x63, 0x6b, 0x10, 0x72, 0x12, 0x10, 0x0a, 0x0c, 0x4d, 0x61, 0x6c,
	0x65, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x64, 0x10, 0x73, 0x12, 0x11, 0x0a, 0x0d, 0x4d,
	0x61, 0x6c, 0x65, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x47, 0x72, 0x61, 0x79, 0x10, 0x74, 0x12, 0x13,
	0x0a, 0x0f, 0x4d, 0x61, 0x6c, 0x65, 0x4d, 0x65, 0x64, 0x69, 0x75, 0x6d, 0x57, 0x68, 0x69, 0x74,
	0x65, 0x10, 0x75, 0x12, 0x13, 0x0a, 0x0f, 0x4d, 0x61, 0x6c, 0x65, 0x4d, 0x65, 0x64, 0x69, 0x75,
	0x6d, 0x42, 0x72, 0x6f, 0x77, 0x6e, 0x10, 0x76, 0x12, 0x13, 0x0a, 0x0f, 0x4d, 0x61, 0x6c, 0x65,
	0x4d, 0x65, 0x64, 0x69, 0x75, 0x6d, 0x42, 0x6c, 0x61, 0x63, 0x6b, 0x10, 0x77, 0x12, 0x11, 0x0a,
	0x0d, 0x4d, 0x61, 0x6c, 0x65, 0x4d, 0x65, 0x64, 0x69, 0x75, 0x6d, 0x52, 0x65, 0x64, 0x10, 0x78,
	0x12, 0x12, 0x0a, 0x0e, 0x4d, 0x61, 0x6c, 0x65, 0x4d, 0x65, 0x64, 0x69, 0x75, 0x6d, 0x47, 0x72,
	0x61, 0x79, 0x10, 0x79, 0x12, 0x11, 0x0a, 0x0d, 0x4d, 0x61, 0x6c, 0x65, 0x4c, 0x6f, 0x6e, 0x67,
	0x57, 0x68, 0x69, 0x74, 0x65, 0x10, 0x7a, 0x12, 0x11, 0x0a, 0x0d, 0x4d, 0x61, 0x6c, 0x65, 0x4c,
	0x6f, 0x6e, 0x67, 0x42, 0x72, 0x6f, 0x77, 0x6e, 0x10, 0x7b, 0x12, 0x11, 0x0a, 0x0d, 0x4d, 0x61,
	0x6c, 0x65, 0x4c, 0x6f, 0x6e, 0x67, 0x42, 0x6c, 0x61, 0x63, 0x6b, 0x10, 0x7c, 0x12, 0x0f, 0x0a,
	0x0b, 0x4d, 0x61, 0x6c, 0x65, 0x4c, 0x6f, 0x6e, 0x67, 0x52, 0x65, 0x64, 0x10, 0x7d, 0x12, 0x10,
	0x0a, 0x0c, 0x4d, 0x61, 0x6c, 0x65, 0x4c, 0x6f, 0x6e, 0x67, 0x47, 0x72, 0x61, 0x79, 0x10, 0x7e,
	0x12, 0x0d, 0x0a, 0x08, 0x4d, 0x61, 0x6c, 0x65, 0x42, 0x61, 0x6c, 0x64, 0x10, 0xc8, 0x01, 0x12,
	0x14, 0x0a, 0x10, 0x46, 0x65, 0x6d, 0x61, 0x6c, 0x65, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x57, 0x68,
	0x69, 0x74, 0x65, 0x10, 0x7f, 0x12, 0x15, 0x0a, 0x10, 0x46, 0x65, 0x6d, 0x61, 0x6c, 0x65, 0x53,
	0x68, 0x6f, 0x72, 0x74, 0x42, 0x72, 0x6f, 0x77, 0x6e, 0x10, 0x80, 0x01, 0x12, 0x15, 0x0a, 0x10,
	0x46, 0x65, 0x6d, 0x61, 0x6c, 0x65, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x42, 0x6c, 0x61, 0x63, 0x6b,
	0x10, 0x81, 0x01, 0x12, 0x13, 0x0a, 0x0e, 0x46, 0x65, 0x6d, 0x61, 0x6c, 0x65, 0x53, 0x68, 0x6f,
	0x72, 0x74, 0x52, 0x65, 0x64, 0x10, 0x82, 0x01, 0x12, 0x14, 0x0a, 0x0f, 0x46, 0x65, 0x6d, 0x61,
	0x6c, 0x65, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x47, 0x72, 0x61, 0x79, 0x10, 0x83, 0x01, 0x12, 0x16,
	0x0a, 0x11, 0x46, 0x65, 0x6d, 0x61, 0x6c, 0x65, 0x4d, 0x65, 0x64, 0x69, 0x75, 0x6d, 0x57, 0x68,
	0x69, 0x74, 0x65, 0x10, 0x84, 0x01, 0x12, 0x16, 0x0a, 0x11, 0x46, 0x65, 0x6d, 0x61, 0x6c, 0x65,
	0x4d, 0x65, 0x64, 0x69, 0x75, 0x6d, 0x42, 0x72, 0x6f, 0x77, 0x6e, 0x10, 0x85, 0x01, 0x12, 0x16,
	0x0a, 0x11, 0x46, 0x65, 0x6d, 0x61, 0x6c, 0x65, 0x4d, 0x65, 0x64, 0x69, 0x75, 0x6d, 0x42, 0x6c,
	0x61, 0x63, 0x6b, 0x10, 0x86, 0x01, 0x12, 0x14, 0x0a, 0x0f, 0x46, 0x65, 0x6d, 0x61, 0x6c, 0x65,
	0x4d, 0x65, 0x64, 0x69, 0x75, 0x6d, 0x52, 0x65, 0x64, 0x10, 0x87, 0x01, 0x12, 0x15, 0x0a, 0x10,
	0x46, 0x65, 0x6d, 0x61, 0x6c, 0x65, 0x4d, 0x65, 0x64, 0x69, 0x75, 0x6d, 0x47, 0x72, 0x61, 0x79,
	0x10, 0x88, 0x01, 0x12, 0x18, 0x0a, 0x13, 0x46, 0x65, 0x6d, 0x61, 0x6c, 0x65, 0x50, 0x6f, 0x6e,
	0x79, 0x74, 0x61, 0x69, 0x6c, 0x57, 0x68, 0x69, 0x74, 0x65, 0x10, 0x89, 0x01, 0x12, 0x18, 0x0a,
	0x13, 0x46, 0x65, 0x6d, 0x61, 0x6c, 0x65, 0x50, 0x6f, 0x6e, 0x79, 0x74, 0x61, 0x69, 0x6c, 0x42,
	0x72, 0x6f, 0x77, 0x6e, 0x10, 0x8a, 0x01, 0x12, 0x18, 0x0a, 0x13, 0x46, 0x65, 0x6d, 0x61, 0x6c,
	0x65, 0x50, 0x6f, 0x6e, 0x79, 0x74, 0x61, 0x69, 0x6c, 0x42, 0x6c, 0x61, 0x63, 0x6b, 0x10, 0x8b,
	0x01, 0x12, 0x16, 0x0a, 0x11, 0x46, 0x65, 0x6d, 0x61, 0x6c, 0x65, 0x50, 0x6f, 0x6e, 0x79, 0x74,
	0x61, 0x69, 0x6c, 0x52, 0x65, 0x64, 0x10, 0x8c, 0x01, 0x12, 0x17, 0x0a, 0x12, 0x46, 0x65, 0x6d,
	0x61, 0x6c, 0x65, 0x50, 0x6f, 0x6e, 0x79, 0x74, 0x61, 0x69, 0x6c, 0x47, 0x72, 0x61, 0x79, 0x10,
	0x8d, 0x01, 0x12, 0x14, 0x0a, 0x0f, 0x46, 0x65, 0x6d, 0x61, 0x6c, 0x65, 0x4c, 0x6f, 0x6e, 0x67,
	0x57, 0x68, 0x69, 0x74, 0x65, 0x10, 0x8e, 0x01, 0x12, 0x14, 0x0a, 0x0f, 0x46, 0x65, 0x6d, 0x61,
	0x6c, 0x65, 0x4c, 0x6f, 0x6e, 0x67, 0x42, 0x72, 0x6f, 0x77, 0x6e, 0x10, 0x8f, 0x01, 0x12, 0x14,
	0x0a, 0x0f, 0x46, 0x65, 0x6d, 0x61, 0x6c, 0x65, 0x4c, 0x6f, 0x6e, 0x67, 0x42, 0x6c, 0x61, 0x63,
	0x6b, 0x10, 0x90, 0x01, 0x12, 0x12, 0x0a, 0x0d, 0x46, 0x65, 0x6d, 0x61, 0x6c, 0x65, 0x4c, 0x6f,
	0x6e, 0x67, 0x52, 0x65, 0x64, 0x10, 0x91, 0x01, 0x12, 0x13, 0x0a, 0x0e, 0x46, 0x65, 0x6d, 0x61,
	0x6c, 0x65, 0x4c, 0x6f, 0x6e, 0x67, 0x47, 0x72, 0x61, 0x79, 0x10, 0x92, 0x01, 0x42, 0x9b, 0x01,
	0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x2e, 0x76, 0x31, 0x42, 0x12,
	0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x64, 0x69, 0x73, 0x70, 0x65, 0x6c, 0x2d, 0x72, 0x65, 0x2f, 0x64, 0x69, 0x73, 0x70, 0x65,
	0x6c, 0x2d, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x6d, 0x75, 0x6c, 0x74,
	0x69, 0x2f, 0x76, 0x31, 0x3b, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x4d,
	0x58, 0x58, 0xaa, 0x02, 0x08, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x08,
	0x4d, 0x75, 0x6c, 0x74, 0x69, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x14, 0x4d, 0x75, 0x6c, 0x74, 0x69,
	0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x09, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_multi_v1_character_type_proto_rawDescOnce sync.Once
	file_multi_v1_character_type_proto_rawDescData = file_multi_v1_character_type_proto_rawDesc
)

func file_multi_v1_character_type_proto_rawDescGZIP() []byte {
	file_multi_v1_character_type_proto_rawDescOnce.Do(func() {
		file_multi_v1_character_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_multi_v1_character_type_proto_rawDescData)
	})
	return file_multi_v1_character_type_proto_rawDescData
}

var file_multi_v1_character_type_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_multi_v1_character_type_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_multi_v1_character_type_proto_goTypes = []interface{}{
	(Gender)(0),           // 0: multi.v1.Gender
	(ClassType)(0),        // 1: multi.v1.ClassType
	(SkinCarnation)(0),    // 2: multi.v1.SkinCarnation
	(HairStyle)(0),        // 3: multi.v1.HairStyle
	(*Character)(nil),     // 4: multi.v1.Character
	(*CharacterInfo)(nil), // 5: multi.v1.CharacterInfo
}
var file_multi_v1_character_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_multi_v1_character_type_proto_init() }
func file_multi_v1_character_type_proto_init() {
	if File_multi_v1_character_type_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_multi_v1_character_type_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Character); i {
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
		file_multi_v1_character_type_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CharacterInfo); i {
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
			RawDescriptor: file_multi_v1_character_type_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_multi_v1_character_type_proto_goTypes,
		DependencyIndexes: file_multi_v1_character_type_proto_depIdxs,
		EnumInfos:         file_multi_v1_character_type_proto_enumTypes,
		MessageInfos:      file_multi_v1_character_type_proto_msgTypes,
	}.Build()
	File_multi_v1_character_type_proto = out.File
	file_multi_v1_character_type_proto_rawDesc = nil
	file_multi_v1_character_type_proto_goTypes = nil
	file_multi_v1_character_type_proto_depIdxs = nil
}

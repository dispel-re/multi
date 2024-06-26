package model

import (
	"encoding/binary"
)

type CharacterInfo struct {
	Strength             uint16
	Agility              uint16
	Wisdom               uint16
	Constitution         uint16
	HealthPoints         uint16
	MagicPoints          uint16
	ExperiencePoints     uint32
	Money                uint32
	ScorePoints          uint32
	ClassType            ClassType
	SkinCarnation        SkinCarnation
	HairStyle            HairStyle
	LightArmourLegs      EquipmentSlot
	LightArmourTorso     EquipmentSlot
	LightArmourHands     EquipmentSlot
	LightArmourBoots     EquipmentSlot
	FullArmour           EquipmentSlot
	ArmourEmblem         EquipmentSlot
	Helmet               EquipmentSlot
	SecondaryWeapon      EquipmentSlot // Buggy
	PrimaryWeapon        EquipmentSlot
	Shield               EquipmentSlot
	UnknownEquipmentSlot EquipmentSlot // Unknown
	Gender               Gender
	Level                byte
	EdgedWeapons         uint16 // First byte is stat, the next one is number of kills
	BluntedWeapons       uint16 // After advancing to 100 kills, the stat is increased
	Archery              uint16 // by one point.
	Polearms             uint16 // All stats (EdgedWeapons, BluntedWeapons, Archery, Polearms, Wizardry)
	Wizardry             uint16 // start from 1 point.
	HolyMagic            uint16 // Unused
	DarkMagic            uint16 // Unused
	BonusPoints          uint16
}

func ParseCharacterInfo(buf []byte) CharacterInfo {
	return CharacterInfo{
		Strength:             binary.LittleEndian.Uint16(buf[0:2]),
		Agility:              binary.LittleEndian.Uint16(buf[2:4]),
		Wisdom:               binary.LittleEndian.Uint16(buf[4:6]),
		Constitution:         binary.LittleEndian.Uint16(buf[6:8]),
		HealthPoints:         binary.LittleEndian.Uint16(buf[8:10]),
		MagicPoints:          binary.LittleEndian.Uint16(buf[10:12]),
		ExperiencePoints:     binary.LittleEndian.Uint32(buf[12:16]),
		Money:                binary.LittleEndian.Uint32(buf[16:20]),
		ScorePoints:          binary.LittleEndian.Uint32(buf[20:24]),
		ClassType:            ClassType(buf[24]),
		SkinCarnation:        SkinCarnation(buf[25]),
		HairStyle:            HairStyle(buf[26]),
		LightArmourLegs:      EquipmentSlot(buf[27]),
		LightArmourTorso:     EquipmentSlot(buf[28]),
		LightArmourHands:     EquipmentSlot(buf[29]),
		LightArmourBoots:     EquipmentSlot(buf[30]),
		FullArmour:           EquipmentSlot(buf[31]),
		ArmourEmblem:         EquipmentSlot(buf[32]),
		Helmet:               EquipmentSlot(buf[33]),
		SecondaryWeapon:      EquipmentSlot(buf[34]),
		PrimaryWeapon:        EquipmentSlot(buf[35]),
		Shield:               EquipmentSlot(buf[36]),
		UnknownEquipmentSlot: EquipmentSlot(buf[37]),
		Gender:               Gender(buf[38]),
		Level:                buf[39],
		EdgedWeapons:         binary.LittleEndian.Uint16(buf[40:42]),
		BluntedWeapons:       binary.LittleEndian.Uint16(buf[42:44]),
		Archery:              binary.LittleEndian.Uint16(buf[44:46]),
		Polearms:             binary.LittleEndian.Uint16(buf[46:48]),
		Wizardry:             binary.LittleEndian.Uint16(buf[48:50]),
		HolyMagic:            binary.LittleEndian.Uint16(buf[50:52]),
		DarkMagic:            binary.LittleEndian.Uint16(buf[52:54]),
		BonusPoints:          binary.LittleEndian.Uint16(buf[54:56]),
	}
}

func (c *CharacterInfo) ToBytes() []byte {
	buf := make([]byte, 56)

	binary.LittleEndian.PutUint16(buf[0:2], c.Strength)
	binary.LittleEndian.PutUint16(buf[2:4], c.Agility)
	binary.LittleEndian.PutUint16(buf[4:6], c.Wisdom)
	binary.LittleEndian.PutUint16(buf[6:8], c.Constitution)
	binary.LittleEndian.PutUint16(buf[8:10], c.HealthPoints)
	binary.LittleEndian.PutUint16(buf[10:12], c.MagicPoints)
	binary.LittleEndian.PutUint32(buf[12:16], c.ExperiencePoints)
	binary.LittleEndian.PutUint32(buf[16:20], c.Money)
	binary.LittleEndian.PutUint32(buf[20:24], c.ScorePoints)

	buf[24] = byte(c.ClassType)
	buf[25] = byte(c.SkinCarnation)
	buf[26] = byte(c.HairStyle)

	buf[27] = byte(c.LightArmourLegs)
	buf[28] = byte(c.LightArmourTorso)
	buf[29] = byte(c.LightArmourHands)
	buf[30] = byte(c.LightArmourBoots)
	buf[31] = byte(c.FullArmour)
	buf[32] = byte(c.ArmourEmblem)
	buf[33] = byte(c.Helmet)
	buf[34] = byte(c.SecondaryWeapon)
	buf[35] = byte(c.PrimaryWeapon)
	buf[36] = byte(c.Shield)
	buf[37] = byte(c.UnknownEquipmentSlot)

	buf[38] = byte(c.Gender)
	buf[39] = c.Level

	binary.LittleEndian.PutUint16(buf[40:42], c.EdgedWeapons)
	binary.LittleEndian.PutUint16(buf[42:44], c.BluntedWeapons)
	binary.LittleEndian.PutUint16(buf[44:46], c.Archery)
	binary.LittleEndian.PutUint16(buf[46:48], c.Polearms)
	binary.LittleEndian.PutUint16(buf[48:50], c.Wizardry)
	binary.LittleEndian.PutUint16(buf[50:52], c.HolyMagic)
	binary.LittleEndian.PutUint16(buf[52:54], c.DarkMagic)
	binary.LittleEndian.PutUint16(buf[54:56], c.BonusPoints)

	return buf
}

func (c *CharacterInfo) EdgedWeaponsLevel() uint8 { return uint8(c.EdgedWeapons) }
func (c *CharacterInfo) EdgedWeaponsKills() uint8 { return byte(c.EdgedWeapons >> 8) }

func (c *CharacterInfo) BluntedWeaponsLevel() uint8 { return uint8(c.BluntedWeapons) }
func (c *CharacterInfo) BluntedWeaponsKills() uint8 { return byte(c.BluntedWeapons >> 8) }

func (c *CharacterInfo) ArcheryLevel() uint8 { return uint8(c.Archery) }
func (c *CharacterInfo) ArcheryKills() uint8 { return byte(c.Archery >> 8) }

func (c *CharacterInfo) PolearmsLevel() uint8 { return uint8(c.Polearms) }
func (c *CharacterInfo) PolearmsKills() uint8 { return byte(c.Polearms >> 8) }

func (c *CharacterInfo) WizardryLevel() uint8 { return uint8(c.Wizardry) }
func (c *CharacterInfo) WizardryKills() uint8 { return byte(c.Wizardry >> 8) }

type EquipmentSlot byte

func (slot EquipmentSlot) IsEquipped() bool { return slot != 100 }

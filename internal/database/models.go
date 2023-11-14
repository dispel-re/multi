// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0

package database

import (
	"database/sql"
)

type Character struct {
	ID                   int64
	UserID               int64
	CharacterName        string
	SortOrder            int64
	Strength             int64
	Agility              int64
	Wisdom               int64
	Constitution         int64
	HealthPoints         int64
	MagicPoints          int64
	ExperiencePoints     int64
	Money                int64
	ScorePoints          int64
	ClassType            int64
	SkinCarnation        int64
	HairStyle            int64
	LightArmourLegs      int64
	LightArmourTorso     int64
	LightArmourHands     int64
	LightArmourBoots     int64
	FullArmour           int64
	ArmourEmblem         int64
	Helmet               int64
	SecondaryWeapon      int64
	PrimaryWeapon        int64
	Shield               int64
	UnknownEquipmentSlot int64
	Gender               int64
	Level                int64
	EdgedWeapons         int64
	BluntedWeapons       int64
	Archery              int64
	Polearms             int64
	Wizardry             int64
	Unknown              sql.NullString
	Inventory            sql.NullString
	Spells               sql.NullString
}

type GameRoom struct {
	ID            int64
	Name          string
	Password      sql.NullString
	HostIpAddress string
}

type User struct {
	ID       int64
	Username string
	Password string
}

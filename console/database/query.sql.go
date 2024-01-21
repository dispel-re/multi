// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: query.sql

package database

import (
	"context"
	"database/sql"
)

const addPlayerToRoom = `-- name: AddPlayerToRoom :exec
INSERT INTO game_room_players (game_room_id, character_id, ip_address)
VALUES (?, ?, ?)
`

type AddPlayerToRoomParams struct {
	GameRoomID  int64
	CharacterID int64
	IpAddress   string
}

func (q *Queries) AddPlayerToRoom(ctx context.Context, arg AddPlayerToRoomParams) error {
	_, err := q.exec(ctx, q.addPlayerToRoomStmt, addPlayerToRoom, arg.GameRoomID, arg.CharacterID, arg.IpAddress)
	return err
}

const createCharacter = `-- name: CreateCharacter :one
INSERT INTO characters (strength,
                        agility,
                        wisdom,
                        constitution,
                        health_points,
                        magic_points,
                        experience_points,
                        money,
                        score_points,
                        class_type,
                        skin_carnation,
                        hair_style,
                        light_armour_legs,
                        light_armour_torso,
                        light_armour_hands,
                        light_armour_boots,
                        full_armour,
                        armour_emblem,
                        helmet,
                        secondary_weapon,
                        primary_weapon,
                        shield,
                        unknown_equipment_slot,
                        gender,
                        level,
                        edged_weapons,
                        blunted_weapons,
                        archery,
                        polearms,
                        wizardry,
                        holy_magic,
                        dark_magic,
                        bonus_points,
                        character_name,
                        user_id)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
RETURNING id, user_id, character_name, strength, agility, wisdom, constitution, health_points, magic_points, experience_points, money, score_points, class_type, skin_carnation, hair_style, light_armour_legs, light_armour_torso, light_armour_hands, light_armour_boots, full_armour, armour_emblem, helmet, secondary_weapon, primary_weapon, shield, unknown_equipment_slot, gender, level, edged_weapons, blunted_weapons, archery, polearms, wizardry, holy_magic, dark_magic, bonus_points, inventory, spells
`

type CreateCharacterParams struct {
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
	HolyMagic            int64
	DarkMagic            int64
	BonusPoints          int64
	CharacterName        string
	UserID               int64
}

func (q *Queries) CreateCharacter(ctx context.Context, arg CreateCharacterParams) (Character, error) {
	row := q.queryRow(ctx, q.createCharacterStmt, createCharacter,
		arg.Strength,
		arg.Agility,
		arg.Wisdom,
		arg.Constitution,
		arg.HealthPoints,
		arg.MagicPoints,
		arg.ExperiencePoints,
		arg.Money,
		arg.ScorePoints,
		arg.ClassType,
		arg.SkinCarnation,
		arg.HairStyle,
		arg.LightArmourLegs,
		arg.LightArmourTorso,
		arg.LightArmourHands,
		arg.LightArmourBoots,
		arg.FullArmour,
		arg.ArmourEmblem,
		arg.Helmet,
		arg.SecondaryWeapon,
		arg.PrimaryWeapon,
		arg.Shield,
		arg.UnknownEquipmentSlot,
		arg.Gender,
		arg.Level,
		arg.EdgedWeapons,
		arg.BluntedWeapons,
		arg.Archery,
		arg.Polearms,
		arg.Wizardry,
		arg.HolyMagic,
		arg.DarkMagic,
		arg.BonusPoints,
		arg.CharacterName,
		arg.UserID,
	)
	var i Character
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.CharacterName,
		&i.Strength,
		&i.Agility,
		&i.Wisdom,
		&i.Constitution,
		&i.HealthPoints,
		&i.MagicPoints,
		&i.ExperiencePoints,
		&i.Money,
		&i.ScorePoints,
		&i.ClassType,
		&i.SkinCarnation,
		&i.HairStyle,
		&i.LightArmourLegs,
		&i.LightArmourTorso,
		&i.LightArmourHands,
		&i.LightArmourBoots,
		&i.FullArmour,
		&i.ArmourEmblem,
		&i.Helmet,
		&i.SecondaryWeapon,
		&i.PrimaryWeapon,
		&i.Shield,
		&i.UnknownEquipmentSlot,
		&i.Gender,
		&i.Level,
		&i.EdgedWeapons,
		&i.BluntedWeapons,
		&i.Archery,
		&i.Polearms,
		&i.Wizardry,
		&i.HolyMagic,
		&i.DarkMagic,
		&i.BonusPoints,
		&i.Inventory,
		&i.Spells,
	)
	return i, err
}

const createGameRoom = `-- name: CreateGameRoom :one
INSERT INTO game_rooms (name, password, host_ip_address, map_id)
VALUES (?, ?, ?, ?)
RETURNING id, name, password, host_ip_address, map_id
`

type CreateGameRoomParams struct {
	Name          string
	Password      sql.NullString
	HostIpAddress string
	MapID         int64
}

func (q *Queries) CreateGameRoom(ctx context.Context, arg CreateGameRoomParams) (GameRoom, error) {
	row := q.queryRow(ctx, q.createGameRoomStmt, createGameRoom,
		arg.Name,
		arg.Password,
		arg.HostIpAddress,
		arg.MapID,
	)
	var i GameRoom
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Password,
		&i.HostIpAddress,
		&i.MapID,
	)
	return i, err
}

const createUser = `-- name: CreateUser :one
INSERT INTO users (username, password)
VALUES (?, ?)
RETURNING id, username, password
`

type CreateUserParams struct {
	Username string
	Password string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.queryRow(ctx, q.createUserStmt, createUser, arg.Username, arg.Password)
	var i User
	err := row.Scan(&i.ID, &i.Username, &i.Password)
	return i, err
}

const deleteCharacter = `-- name: DeleteCharacter :exec
DELETE
FROM characters
WHERE character_name = ?
  AND user_id = ?
`

type DeleteCharacterParams struct {
	CharacterName string
	UserID        int64
}

func (q *Queries) DeleteCharacter(ctx context.Context, arg DeleteCharacterParams) error {
	_, err := q.exec(ctx, q.deleteCharacterStmt, deleteCharacter, arg.CharacterName, arg.UserID)
	return err
}

const existPlayerInRoom = `-- name: ExistPlayerInRoom :one
SELECT 1 as exist
FROM game_room_players
WHERE game_room_id = ?
  AND character_id = ?
`

type ExistPlayerInRoomParams struct {
	GameRoomID  int64
	CharacterID int64
}

func (q *Queries) ExistPlayerInRoom(ctx context.Context, arg ExistPlayerInRoomParams) (int64, error) {
	row := q.queryRow(ctx, q.existPlayerInRoomStmt, existPlayerInRoom, arg.GameRoomID, arg.CharacterID)
	var exist int64
	err := row.Scan(&exist)
	return exist, err
}

const findCharacter = `-- name: FindCharacter :one
SELECT id, user_id, character_name, strength, agility, wisdom, constitution, health_points, magic_points, experience_points, money, score_points, class_type, skin_carnation, hair_style, light_armour_legs, light_armour_torso, light_armour_hands, light_armour_boots, full_armour, armour_emblem, helmet, secondary_weapon, primary_weapon, shield, unknown_equipment_slot, gender, level, edged_weapons, blunted_weapons, archery, polearms, wizardry, holy_magic, dark_magic, bonus_points, inventory, spells
FROM characters
WHERE character_name = ?
  AND user_id = ?
`

type FindCharacterParams struct {
	CharacterName string
	UserID        int64
}

func (q *Queries) FindCharacter(ctx context.Context, arg FindCharacterParams) (Character, error) {
	row := q.queryRow(ctx, q.findCharacterStmt, findCharacter, arg.CharacterName, arg.UserID)
	var i Character
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.CharacterName,
		&i.Strength,
		&i.Agility,
		&i.Wisdom,
		&i.Constitution,
		&i.HealthPoints,
		&i.MagicPoints,
		&i.ExperiencePoints,
		&i.Money,
		&i.ScorePoints,
		&i.ClassType,
		&i.SkinCarnation,
		&i.HairStyle,
		&i.LightArmourLegs,
		&i.LightArmourTorso,
		&i.LightArmourHands,
		&i.LightArmourBoots,
		&i.FullArmour,
		&i.ArmourEmblem,
		&i.Helmet,
		&i.SecondaryWeapon,
		&i.PrimaryWeapon,
		&i.Shield,
		&i.UnknownEquipmentSlot,
		&i.Gender,
		&i.Level,
		&i.EdgedWeapons,
		&i.BluntedWeapons,
		&i.Archery,
		&i.Polearms,
		&i.Wizardry,
		&i.HolyMagic,
		&i.DarkMagic,
		&i.BonusPoints,
		&i.Inventory,
		&i.Spells,
	)
	return i, err
}

const getCurrentUser = `-- name: GetCurrentUser :one
SELECT position, cte.score_points, cte.username, cte.character_name
FROM (SELECT ROW_NUMBER() over (ORDER BY score_points) as position,
             score_points,
             username,
             character_name
      FROM characters
               JOIN users ON characters.user_id = users.id
      WHERE users.id = ?
        AND characters.character_name = ?) as cte
LIMIT 1
`

type GetCurrentUserParams struct {
	ID            int64
	CharacterName string
}

type GetCurrentUserRow struct {
	Position      interface{}
	ScorePoints   int64
	Username      string
	CharacterName string
}

func (q *Queries) GetCurrentUser(ctx context.Context, arg GetCurrentUserParams) (GetCurrentUserRow, error) {
	row := q.queryRow(ctx, q.getCurrentUserStmt, getCurrentUser, arg.ID, arg.CharacterName)
	var i GetCurrentUserRow
	err := row.Scan(
		&i.Position,
		&i.ScorePoints,
		&i.Username,
		&i.CharacterName,
	)
	return i, err
}

const getGameRoom = `-- name: GetGameRoom :one
SELECT id,
       name,
       password,
       host_ip_address,
       map_id
FROM game_rooms
WHERE game_rooms.name = ?
LIMIT 1
`

func (q *Queries) GetGameRoom(ctx context.Context, name string) (GameRoom, error) {
	row := q.queryRow(ctx, q.getGameRoomStmt, getGameRoom, name)
	var i GameRoom
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Password,
		&i.HostIpAddress,
		&i.MapID,
	)
	return i, err
}

const getGameRoomPlayers = `-- name: GetGameRoomPlayers :many
SELECT DISTINCT username,
                character_name,
                class_type,
                ip_address
FROM game_rooms
         JOIN game_room_players ON game_rooms.id = game_room_players.game_room_id
         JOIN characters ON game_room_players.character_id = characters.id
         JOIN users on users.id = characters.user_id
WHERE game_rooms.id = ?
`

type GetGameRoomPlayersRow struct {
	Username      string
	CharacterName string
	ClassType     int64
	IpAddress     string
}

func (q *Queries) GetGameRoomPlayers(ctx context.Context, id int64) ([]GetGameRoomPlayersRow, error) {
	rows, err := q.query(ctx, q.getGameRoomPlayersStmt, getGameRoomPlayers, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetGameRoomPlayersRow
	for rows.Next() {
		var i GetGameRoomPlayersRow
		if err := rows.Scan(
			&i.Username,
			&i.CharacterName,
			&i.ClassType,
			&i.IpAddress,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserByID = `-- name: GetUserByID :one
SELECT id, username, password
FROM users
WHERE id = ?
LIMIT 1
`

func (q *Queries) GetUserByID(ctx context.Context, id int64) (User, error) {
	row := q.queryRow(ctx, q.getUserByIDStmt, getUserByID, id)
	var i User
	err := row.Scan(&i.ID, &i.Username, &i.Password)
	return i, err
}

const getUserByName = `-- name: GetUserByName :one
SELECT id, username, password
FROM users
WHERE username = ?
LIMIT 1
`

func (q *Queries) GetUserByName(ctx context.Context, username string) (User, error) {
	row := q.queryRow(ctx, q.getUserByNameStmt, getUserByName, username)
	var i User
	err := row.Scan(&i.ID, &i.Username, &i.Password)
	return i, err
}

const listCharacters = `-- name: ListCharacters :many
SELECT id, user_id, character_name, strength, agility, wisdom, constitution, health_points, magic_points, experience_points, money, score_points, class_type, skin_carnation, hair_style, light_armour_legs, light_armour_torso, light_armour_hands, light_armour_boots, full_armour, armour_emblem, helmet, secondary_weapon, primary_weapon, shield, unknown_equipment_slot, gender, level, edged_weapons, blunted_weapons, archery, polearms, wizardry, holy_magic, dark_magic, bonus_points, inventory, spells
FROM characters
WHERE user_id = ?
`

func (q *Queries) ListCharacters(ctx context.Context, userID int64) ([]Character, error) {
	rows, err := q.query(ctx, q.listCharactersStmt, listCharacters, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Character
	for rows.Next() {
		var i Character
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.CharacterName,
			&i.Strength,
			&i.Agility,
			&i.Wisdom,
			&i.Constitution,
			&i.HealthPoints,
			&i.MagicPoints,
			&i.ExperiencePoints,
			&i.Money,
			&i.ScorePoints,
			&i.ClassType,
			&i.SkinCarnation,
			&i.HairStyle,
			&i.LightArmourLegs,
			&i.LightArmourTorso,
			&i.LightArmourHands,
			&i.LightArmourBoots,
			&i.FullArmour,
			&i.ArmourEmblem,
			&i.Helmet,
			&i.SecondaryWeapon,
			&i.PrimaryWeapon,
			&i.Shield,
			&i.UnknownEquipmentSlot,
			&i.Gender,
			&i.Level,
			&i.EdgedWeapons,
			&i.BluntedWeapons,
			&i.Archery,
			&i.Polearms,
			&i.Wizardry,
			&i.HolyMagic,
			&i.DarkMagic,
			&i.BonusPoints,
			&i.Inventory,
			&i.Spells,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listGameRooms = `-- name: ListGameRooms :many
SELECT id, name, password, host_ip_address, map_id
FROM game_rooms
`

func (q *Queries) ListGameRooms(ctx context.Context) ([]GameRoom, error) {
	rows, err := q.query(ctx, q.listGameRoomsStmt, listGameRooms)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GameRoom
	for rows.Next() {
		var i GameRoom
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Password,
			&i.HostIpAddress,
			&i.MapID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const selectRanking = `-- name: SelectRanking :many
SELECT ROW_NUMBER() over (ORDER BY score_points) as position,
       score_points,
       username,
       character_name
FROM characters
         JOIN users ON characters.user_id = users.id
WHERE class_type = ?
ORDER BY score_points
LIMIT 10 OFFSET ?
`

type SelectRankingParams struct {
	ClassType int64
	Offset    int64
}

type SelectRankingRow struct {
	Position      interface{}
	ScorePoints   int64
	Username      string
	CharacterName string
}

func (q *Queries) SelectRanking(ctx context.Context, arg SelectRankingParams) ([]SelectRankingRow, error) {
	rows, err := q.query(ctx, q.selectRankingStmt, selectRanking, arg.ClassType, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []SelectRankingRow
	for rows.Next() {
		var i SelectRankingRow
		if err := rows.Scan(
			&i.Position,
			&i.ScorePoints,
			&i.Username,
			&i.CharacterName,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateCharacterInventory = `-- name: UpdateCharacterInventory :exec
UPDATE characters
SET inventory = ?
WHERE character_name = ?
  AND user_id = ?
`

type UpdateCharacterInventoryParams struct {
	Inventory     sql.NullString
	CharacterName string
	UserID        int64
}

func (q *Queries) UpdateCharacterInventory(ctx context.Context, arg UpdateCharacterInventoryParams) error {
	_, err := q.exec(ctx, q.updateCharacterInventoryStmt, updateCharacterInventory, arg.Inventory, arg.CharacterName, arg.UserID)
	return err
}

const updateCharacterSpells = `-- name: UpdateCharacterSpells :exec
UPDATE characters
SET spells = ?
WHERE character_name = ?
  AND user_id = ?
`

type UpdateCharacterSpellsParams struct {
	Spells        sql.NullString
	CharacterName string
	UserID        int64
}

func (q *Queries) UpdateCharacterSpells(ctx context.Context, arg UpdateCharacterSpellsParams) error {
	_, err := q.exec(ctx, q.updateCharacterSpellsStmt, updateCharacterSpells, arg.Spells, arg.CharacterName, arg.UserID)
	return err
}

const updateCharacterStats = `-- name: UpdateCharacterStats :exec
UPDATE characters
SET strength               = ?,
    agility                = ?,
    wisdom                 = ?,
    constitution           = ?,
    health_points          = ?,
    magic_points           = ?,
    experience_points      = ?,
    money                  = ?,
    score_points           = ?,
    class_type             = ?,
    skin_carnation         = ?,
    hair_style             = ?,
    light_armour_legs      = ?,
    light_armour_torso     = ?,
    light_armour_hands     = ?,
    light_armour_boots     = ?,
    full_armour            = ?,
    armour_emblem          = ?,
    helmet                 = ?,
    secondary_weapon       = ?,
    primary_weapon         = ?,
    shield                 = ?,
    unknown_equipment_slot = ?,
    gender                 = ?,
    level                  = ?,
    edged_weapons          = ?,
    blunted_weapons        = ?,
    archery                = ?,
    polearms               = ?,
    wizardry               = ?,
    holy_magic             = ?,
    dark_magic             = ?,
    bonus_points           = ?
WHERE character_name = ?
  AND user_id = ?
`

type UpdateCharacterStatsParams struct {
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
	HolyMagic            int64
	DarkMagic            int64
	BonusPoints          int64
	CharacterName        string
	UserID               int64
}

func (q *Queries) UpdateCharacterStats(ctx context.Context, arg UpdateCharacterStatsParams) error {
	_, err := q.exec(ctx, q.updateCharacterStatsStmt, updateCharacterStats,
		arg.Strength,
		arg.Agility,
		arg.Wisdom,
		arg.Constitution,
		arg.HealthPoints,
		arg.MagicPoints,
		arg.ExperiencePoints,
		arg.Money,
		arg.ScorePoints,
		arg.ClassType,
		arg.SkinCarnation,
		arg.HairStyle,
		arg.LightArmourLegs,
		arg.LightArmourTorso,
		arg.LightArmourHands,
		arg.LightArmourBoots,
		arg.FullArmour,
		arg.ArmourEmblem,
		arg.Helmet,
		arg.SecondaryWeapon,
		arg.PrimaryWeapon,
		arg.Shield,
		arg.UnknownEquipmentSlot,
		arg.Gender,
		arg.Level,
		arg.EdgedWeapons,
		arg.BluntedWeapons,
		arg.Archery,
		arg.Polearms,
		arg.Wizardry,
		arg.HolyMagic,
		arg.DarkMagic,
		arg.BonusPoints,
		arg.CharacterName,
		arg.UserID,
	)
	return err
}

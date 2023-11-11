// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: query.sql

package sqlite

import (
	"context"
	"database/sql"
)

const createAuthor = `-- name: CreateAuthor :one
INSERT INTO users (username, password)
VALUES (?, ?) RETURNING id, username, password
`

type CreateAuthorParams struct {
	Username string
	Password string
}

func (q *Queries) CreateAuthor(ctx context.Context, arg CreateAuthorParams) (User, error) {
	row := q.queryRow(ctx, q.createAuthorStmt, createAuthor, arg.Username, arg.Password)
	var i User
	err := row.Scan(&i.ID, &i.Username, &i.Password)
	return i, err
}

const createGameRoom = `-- name: CreateGameRoom :one
INSERT INTO game_rooms (name, password, host_ip_address)
VALUES (?, ?, ?) RETURNING id, name, password, host_ip_address
`

type CreateGameRoomParams struct {
	Name          string
	Password      sql.NullString
	HostIpAddress string
}

func (q *Queries) CreateGameRoom(ctx context.Context, arg CreateGameRoomParams) (GameRoom, error) {
	row := q.queryRow(ctx, q.createGameRoomStmt, createGameRoom, arg.Name, arg.Password, arg.HostIpAddress)
	var i GameRoom
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Password,
		&i.HostIpAddress,
	)
	return i, err
}

const createUser = `-- name: CreateUser :one
INSERT INTO users (username, password)
VALUES (?, ?) RETURNING id, username, password
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

const deleteAuthor = `-- name: DeleteAuthor :exec
DELETE
FROM users
WHERE id = ?
`

func (q *Queries) DeleteAuthor(ctx context.Context, id int64) error {
	_, err := q.exec(ctx, q.deleteAuthorStmt, deleteAuthor, id)
	return err
}

const getAuthor = `-- name: GetAuthor :one
SELECT id, username, password
FROM users
WHERE id = ? LIMIT 1
`

func (q *Queries) GetAuthor(ctx context.Context, id int64) (User, error) {
	row := q.queryRow(ctx, q.getAuthorStmt, getAuthor, id)
	var i User
	err := row.Scan(&i.ID, &i.Username, &i.Password)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT id, username, password
FROM users
WHERE username = ? LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, username string) (User, error) {
	row := q.queryRow(ctx, q.getUserStmt, getUser, username)
	var i User
	err := row.Scan(&i.ID, &i.Username, &i.Password)
	return i, err
}

const listAuthors = `-- name: ListAuthors :many
SELECT id, username, password
FROM users
ORDER BY username
`

func (q *Queries) ListAuthors(ctx context.Context) ([]User, error) {
	rows, err := q.query(ctx, q.listAuthorsStmt, listAuthors)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(&i.ID, &i.Username, &i.Password); err != nil {
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

const listCharacters = `-- name: ListCharacters :many
SELECT id, user_id, character_name, sort_order, strength, agility, wisdom, constitution, health_points, magic_points, experience_points, money, score_points, class_type, skin_carnation, hair_style, light_armour_legs, light_armour_torso, light_armour_hands, light_armour_boots, full_armour, armour_emblem, helmet, secondary_weapon, primary_weapon, shield, unknown_equipment_slot, gender, level, edged_weapons, blunted_weapons, archery, polearms, wizardry, unknown, inventory_base64, spells_base64
FROM characters
WHERE user_id = ?
ORDER BY slot_order
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
			&i.SortOrder,
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
			&i.Unknown,
			&i.InventoryBase64,
			&i.SpellsBase64,
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
SELECT id, name, password, host_ip_address
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

const updateAuthor = `-- name: UpdateAuthor :exec
UPDATE users
set username = ?,
    password  = ?
WHERE id = ?
`

type UpdateAuthorParams struct {
	Username string
	Password string
	ID       int64
}

func (q *Queries) UpdateAuthor(ctx context.Context, arg UpdateAuthorParams) error {
	_, err := q.exec(ctx, q.updateAuthorStmt, updateAuthor, arg.Username, arg.Password, arg.ID)
	return err
}

package backend

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/base64"
	"fmt"

	"github.com/dispel-re/dispel-multi/internal/database/sqlite"
	"github.com/dispel-re/dispel-multi/model"
)

func (b *Backend) HandleUpdateCharacterSpells(session *model.Session, req UpdateCharacterSpellsRequest) error {
	data, err := req.Parse()
	if err != nil {
		return err
	}

	if err := b.DB.UpdateCharacterSpells(context.TODO(), sqlite.UpdateCharacterSpellsParams{
		Spells: sql.NullString{
			String: base64.StdEncoding.EncodeToString(data.Spells),
			Valid:  len(data.Spells) > 0,
		},
		ID: session.CharacterID,
	}); err != nil {
		return fmt.Errorf("packet-73: could not update character spells: %s", err)
	}

	return b.Send(session.Conn, UpdateCharacterSpells, []byte{1, 0, 0, 0})
}

type UpdateCharacterSpellsRequest []byte

type UpdateCharacterSpellsRequestData struct {
	Username      string
	CharacterName string
	Spells        []byte
}

func (r UpdateCharacterSpellsRequest) Parse() (data UpdateCharacterSpellsRequestData, err error) {
	split := bytes.SplitN(r, []byte{0}, 3)
	if len(split) != 3 {
		return data, fmt.Errorf("packet-73: no enough arguments, malformed request payload: %s", base64.StdEncoding.EncodeToString(r))
	}
	if len(split[2]) != 43 {
		return data, fmt.Errorf("packet-73: the spells array has invalid length: %s", base64.StdEncoding.EncodeToString(r))
	}

	data.Username = string(split[0])
	data.CharacterName = string(split[1])
	data.Spells = split[2]
	return data, nil
}

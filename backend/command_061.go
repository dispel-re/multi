package backend

import (
	"bytes"
	"fmt"

	"github.com/dispel-re/dispel-multi/model"
)

func (b *Backend) HandleDeleteCharacter(session *model.Session, req DeleteCharacterRequest) error {
	data, err := req.Parse()
	if err != nil {
		return err
	}

	var characters []model.Character
	for _, ch := range session.User.Characters {
		if ch.CharacterName != data.CharacterName {
			characters = append(characters, ch)
		}
	}
	session.User.Characters = characters

	response := make([]byte, len(data.CharacterName)+1)
	copy(response, data.CharacterName)

	return b.Send(session.Conn, DeleteCharacter, response)
}

type DeleteCharacterRequest []byte

type DeleteCharacterRequestData struct {
	Username      string
	CharacterName string
}

func (r DeleteCharacterRequest) Parse() (data DeleteCharacterRequestData, err error) {
	if bytes.Count(r, []byte{0}) < 2 {
		return data, fmt.Errorf("packet-61: malformed packet, not enough null-terminators")
	}

	split := bytes.SplitN(r, []byte{0}, 3)
	data.Username = string(split[0])
	data.CharacterName = string(split[1])

	return data, nil
}
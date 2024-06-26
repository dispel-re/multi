package backend

import (
	"context"
	"encoding/binary"
	"fmt"
	"log/slog"

	"connectrpc.com/connect"
	"github.com/dimspell/gladiator/backend/packet"
	multiv1 "github.com/dimspell/gladiator/gen/multi/v1"
	"github.com/dimspell/gladiator/model"
)

// HandleCreateGame handles 0x1cff (255-28) command
func (b *Backend) HandleCreateGame(session *model.Session, req CreateGameRequest) error {
	if session.UserID == 0 {
		return fmt.Errorf("packet-28: user is not logged in")
	}

	data, err := req.Parse()
	if err != nil {
		return err
	}

	response := make([]byte, 4)

	hostIPAddress, err := b.Proxy.Create(session.LocalIpAddress, session.Username)
	if err != nil {
		return fmt.Errorf("packet-28: incorrect host address %w", err)
	}

	switch data.State {
	case uint32(model.GameStateNone):
		respGame, err := b.gameClient.CreateGame(context.TODO(), connect.NewRequest(&multiv1.CreateGameRequest{
			UserId:        session.UserID,
			GameName:      data.RoomName,
			Password:      data.Password,
			HostIpAddress: hostIPAddress.String(),
			MapId:         int64(data.MapID),
		}))
		if err != nil {
			return err
		}
		slog.Info("packet-28: created game room", "id", respGame.Msg.Game.GameId, "name", respGame.Msg.Game.Name)

		_, err = b.gameClient.JoinGame(context.TODO(), connect.NewRequest(&multiv1.JoinGameRequest{
			UserId:      session.UserID,
			CharacterId: session.CharacterID,
			GameRoomId:  respGame.Msg.Game.GetGameId(),
			IpAddress:   hostIPAddress.String(),
		}))
		if err != nil {
			return err
		}
		binary.LittleEndian.PutUint32(response[0:4], uint32(model.GameStateCreating)) // Game state
		break
	case uint32(model.GameStateCreating):
		// b.EventChan <- EventHostGame
		binary.LittleEndian.PutUint32(response[0:4], uint32(model.GameStateStarted)) // Game state
		break
	}

	return b.Send(session.Conn, CreateGame, response)
}

type CreateGameRequest []byte

type CreateGameRequestData struct {
	State    uint32
	MapID    uint32
	RoomName string
	Password string
}

func (r CreateGameRequest) Parse() (data CreateGameRequestData, err error) {
	rd := packet.NewReader(r)

	data.State, err = rd.ReadUint32()
	if err != nil {
		return data, fmt.Errorf("packet-28: malformed state %w", err)
	}
	data.MapID, err = rd.ReadUint32()
	if err != nil {
		return data, fmt.Errorf("packet-28: malformed map id %w", err)
	}
	if data.MapID > 5 {
		return data, fmt.Errorf("packet-28: incorrect map id %w", err)
	}
	data.RoomName, err = rd.ReadString()
	if err != nil {
		return data, fmt.Errorf("packet-28: malformed room name %w", err)
	}
	data.Password, err = rd.ReadString()
	if err != nil {
		return data, fmt.Errorf("packet-28: malformed password %w", err)
	}

	return data, rd.Close()
}

package backend

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"log/slog"

	"connectrpc.com/connect"
	multiv1 "github.com/dispel-re/dispel-multi/gen/multi/v1"
	"github.com/dispel-re/dispel-multi/model"
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

	// hostIPAddress := "192.168.121.212"
	hostIPAddress := "127.0.1.28"
	// hostIPAddress := session.Conn.RemoteAddr().(*net.TCPAddr).IP.String()

	switch data.State {
	case uint32(0):
		respGame, err := b.GameClient.CreateGame(context.TODO(), connect.NewRequest(&multiv1.CreateGameRequest{
			UserId:   session.UserID,
			GameName: data.RoomName,
			// Password:      data.Password,
			Password:      "",
			HostIpAddress: hostIPAddress,
			MapId:         int64(data.MapID),
		}))
		if err != nil {
			return err
		}
		slog.Info("packet-28: created game room",
			"id", respGame.Msg.Game.GameId,
			"name", respGame.Msg.Game.Name)

		_, err = b.GameClient.JoinGame(context.TODO(), connect.NewRequest(&multiv1.JoinGameRequest{
			UserId:      session.UserID,
			CharacterId: session.CharacterID,
			GameRoomId:  respGame.Msg.Game.GetGameId(),
			IpAddress:   hostIPAddress,
		}))
		if err != nil {
			return err
		}
		binary.LittleEndian.PutUint32(response[0:4], 1)
		break
	case uint32(1):
		// b.EventChan <- EventHostGame
		binary.LittleEndian.PutUint32(response[0:4], 2)
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
	data.State = binary.LittleEndian.Uint32(r[0:4])
	data.MapID = binary.LittleEndian.Uint32(r[4:8])

	split := bytes.Split(r[8:], []byte{0})
	data.RoomName = string(split[0])
	data.Password = string(split[1])
	return data, nil
}

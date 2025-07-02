package server

import (
	"time"

	"github.com/gorilla/websocket"
)

type Lobby struct {
	HostId  string
	LobbyId string

	GameState string //????
	CreatedAt time.Time
}

type GameState struct {
	Status  string
	Players []Player
}

type Player struct {
	ID     string
	Conn   *websocket.Conn
	Score  int
	Status string
}

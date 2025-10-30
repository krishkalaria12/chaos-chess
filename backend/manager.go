package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var (
	websocketUpgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

type Manager struct {
	Matches       map[string]*Match
	WaitingPlayer *Player
}

func NewManager() *Manager {
	return &Manager{
		Matches: make(map[string]*Match),
	}
}

func (m *Manager) serveWs(w http.ResponseWriter, r *http.Request) {
	log.Println("New connection")

	conn, err := websocketUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	player := NewPlayer(conn, nil)
	m.addPlayer(player)
}

func (m *Manager) addPlayer(player *Player) {
	if m.WaitingPlayer != nil {
		match := CreateMatch(m.WaitingPlayer, player)
		m.Matches[match.ID] = match
	} else {
		m.WaitingPlayer = player
	}
}

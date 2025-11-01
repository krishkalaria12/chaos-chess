package main

import (
	"errors"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var (
	websocketUpgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     checkOrigin,
	}
)

type Manager struct {
	Matches       map[string]*Match
	WaitingPlayer *Player
	sync.RWMutex

	handlers map[string]EventHandler
}

func NewManager() *Manager {
	manager := &Manager{
		Matches:  make(map[string]*Match),
		handlers: make(map[string]EventHandler),
	}

	manager.setupHandlers()
	return manager
}

func (m *Manager) setupHandlers() {
	m.handlers[EventSendPlayMove] = PlayMoveHandler
}

func (m *Manager) routeEvent(event Event, match *Match, player *Player) error {
	if handler, ok := m.handlers[event.Type]; !ok {
		return errors.New("there is no such event")
	} else {
		if err := handler(event, match, player); err != nil {
			return err
		}
		return nil
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

	// player services
	go player.ReadMessages()
	go player.WriteMessages()
}

func (m *Manager) addPlayer(player *Player) {
	m.Lock()
	defer m.Unlock()

	if m.WaitingPlayer != nil {
		match := CreateMatch(m, m.WaitingPlayer, player)
		m.Matches[match.ID] = match
		m.WaitingPlayer = nil
	} else {
		m.WaitingPlayer = player
	}
}

func checkOrigin(req *http.Request) bool {
	origin := req.Header.Get("Origin")

	switch origin {
	case "http://localhost:3000":
		return true
	default:
		return false
	}
}

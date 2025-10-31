package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

var (
	pongWait     = 10 * time.Second
	pingInterval = (9 * pongWait) / 10
)

type Player struct {
	Connection *websocket.Conn
	Match      *Match
	Color      string
	Egress     chan []byte // for outgoing messages
}

type PlayerList map[*Player]bool

func NewPlayer(conn *websocket.Conn, match *Match) *Player {
	return &Player{
		Connection: conn,
		Match:      match,
		Egress:     make(chan []byte),
	}
}

func (player *Player) ReadMessages() {
	defer func() {
		// remove the player from the match
	}()

	if err := player.Connection.SetReadDeadline(time.Now().Add(pongWait)); err != nil {
		log.Println("error receiving the pong: ", err)
		return
	}

	player.Connection.SetPongHandler(player.pongHandler)

	for {
		_, payload, err := player.Connection.ReadMessage()

		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseAbnormalClosure, websocket.CloseGoingAway) {
				log.Printf("error reading message: %v", err)
			}
			break
		}

		var request Event

		if err := json.Unmarshal(payload, &request); err != nil {
			log.Printf("Error marshalling event: %v", err)
			break
		}

		if err := player.Match.Manager.routeEvent(request, player.Match, player); err != nil {
			log.Println("error handling event: ", err)
		}
	}
}

func (player *Player) WriteMessages() {
	defer func() {
		// remove the player from the match
	}()

	ticker := time.NewTicker(pingInterval)

	for {
		select {
		case message, ok := <-player.Egress:
			if !ok {
				if err := player.Connection.WriteMessage(websocket.CloseMessage, nil); err != nil {
					log.Println("Connection closed", err)
				}
				return
			}

			data, err := json.Marshal(message)
			if err != nil {
				log.Println(err)
				return
			}

			if err := player.Connection.WriteMessage(websocket.TextMessage, data); err != nil {
				log.Printf("error writing message: %v", err)
			}

			log.Println("Message sent")

		case <-ticker.C:
			log.Println("ping")

			if err := player.Connection.WriteMessage(websocket.PingMessage, []byte(``)); err != nil {
				log.Println("error writing ping message: ", err)
				return
			}
		}
	}
}

func (player *Player) pongHandler(pongMsg string) error {
	log.Println("pong")

	return player.Connection.SetReadDeadline(time.Now().Add(pongWait))
}

package main

import "github.com/gorilla/websocket"

type Player struct {
	Connection *websocket.Conn
	Match      *Match
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

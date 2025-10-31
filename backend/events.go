package main

import "encoding/json"

type Event struct {
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

type EventHandler func(event Event, match *Match, player *Player) error

const (
	EventPlayMove = "play_move"
)

type PlayMoveEvent struct {
	From string `json:"from"`
	To   string `json:"to"`
}

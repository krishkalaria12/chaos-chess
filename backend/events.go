package main

import (
	"encoding/json"
	"log"
)

type Event struct {
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

type EventHandler func(event Event, match *Match, player *Player) error

const (
	EventPlayMove      = "play_move"
	EventError         = "error"
	EventMatchComplete = "match_complete"
)

type PlayMoveEvent struct {
	From string `json:"from"`
	To   string `json:"to"`
}

type ErrorEvent struct {
	Message string `json:"message"`
}

type MatchCompleteEvent struct {
	Message string `json:"message"`
}

func (player *Player) sendError(message string) {
	errorEvent := ErrorEvent{
		Message: message,
	}

	payload, err := json.Marshal(errorEvent)
	if err != nil {
		log.Printf("error marshalling error event: %v", err)
		return
	}

	event := Event{
		Type:    EventError,
		Payload: payload,
	}

	eventData, err := json.Marshal(event)
	if err != nil {
		log.Printf("error marshalling error event wrapper: %v", err)
		return
	}

	select {
	case player.Egress <- eventData:
	default:
		log.Println("client egress channel full, dropping error message")
	}
}

func (player *Player) MatchCompleteHandler(message string) {
	matchCompleteEvent := MatchCompleteEvent{
		Message: message,
	}

	payload, err := json.Marshal(matchCompleteEvent)
	if err != nil {
		log.Printf("error marshalling match complete event: %v", err)
		return
	}

	event := Event{
		Type:    EventMatchComplete,
		Payload: payload,
	}

	eventData, err := json.Marshal(event)
	if err != nil {
		log.Printf("error marshalling match complete event wrapper: %v", err)
		return
	}

	select {
	case player.Egress <- eventData:
	default:
		log.Println("client egress channel full, dropping error message")
	}
}

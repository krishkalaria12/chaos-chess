package main

import (
	"encoding/json"
	"log"
)

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

	select {
	case player.Egress <- event:
	default:
		log.Println("client egress channel full, dropping error message")
	}
}

func (player *Player) sendMatchComplete(message string) {
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

	select {
	case player.Egress <- event:
	default:
		log.Println("client egress channel full, dropping error message")
	}
}

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/notnil/chess"
)

type Event struct {
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

type EventHandler func(event Event, match *Match, player *Player) error

const (
	EventSendPlayMove    = "send_play_move"
	EventReceivePlayMove = "receive_play_move"
	EventMatchStart      = "match_start"
	EventError           = "error"
	EventMatchComplete   = "match_complete"
)

type SendPlayMoveEvent struct {
	From string `json:"from"`
	To   string `json:"to"`
}

type ReceivePlayMoveEvent struct {
	SendPlayMoveEvent
	Sent time.Time `json:"sent"`
}

type ErrorEvent struct {
	Message string `json:"message"`
}

type MatchCompleteEvent struct {
	Message string `json:"message"`
}

type MatchStartEvent struct {
	Color       string `json:"color"`
	Orientation string `json:"orientation"`
	Position    string `json:"position"`
}

func PlayMoveHandler(event Event, match *Match, player *Player) error {
	var move SendPlayMoveEvent

	if err := json.Unmarshal(event.Payload, &move); err != nil {
		return fmt.Errorf("bad payload in request: %v", err)
	}

	if player.Color != match.Turn {
		return fmt.Errorf("not your turn")
	}

	// Find the valid move from the list of legal moves
	var validMove *chess.Move
	for _, m := range match.State.ValidMoves() {
		if m.S1().String() == move.From && m.S2().String() == move.To {
			validMove = m
			break
		}
	}

	if validMove == nil {
		return fmt.Errorf("invalid move: no valid move from %s to %s", move.From, move.To)
	}

	// Execute the move
	if err := match.State.Move(validMove); err != nil {
		return fmt.Errorf("invalid move: %v", err)
	}

	if match.Turn == "white" {
		match.Turn = "black"
	} else {
		match.Turn = "white"
	}

	var request ReceivePlayMoveEvent
	request.From = move.From
	request.To = move.To
	request.Sent = time.Now()

	data, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("failed to marshal broadcast message: %v", err)
	}

	var outgoingMove Event
	outgoingMove.Payload = data
	outgoingMove.Type = EventReceivePlayMove

	var otherPlayer *Player
	if match.Player1 == player {
		otherPlayer = match.Player2
	} else {
		otherPlayer = match.Player1
	}

	select {
	case otherPlayer.Egress <- outgoingMove:
	default:
		log.Println("client egress channel full, dropping move message")
	}

	log.Printf("Player %s moved from %s to %s", player.Color, move.From, move.To)

	// Check if game is over
	outcome := match.State.Outcome()
	if outcome != "*" {
		match.Won = player
		player.sendMatchComplete("You won!")
		otherPlayer.sendMatchComplete("You lost!")

		delete(match.Manager.Matches, match.ID)
		player.Match = nil
		otherPlayer.Match = nil
	}

	return nil
}

package main

import (
	"math/rand"

	"github.com/google/uuid"
	"github.com/notnil/chess"
)

type Match struct {
	ID      string
	Player1 *Player
	Player2 *Player
	Turn    string
	Won     *Player
	State   *chess.Game
}

func CreateMatch(player1 *Player, player2 *Player) *Match {
	if rand.Intn(2) == 0 {
		player1.Color = "white"
		player2.Color = "black"
	} else {
		player1.Color = "black"
		player2.Color = "white"
	}

	match := &Match{
		ID:      uuid.NewString(),
		Player1: player1,
		Player2: player2,
		Turn:    "white",
	}

	player1.Match = match
	player2.Match = match

	// start the match
	// go match.run()

	return match
}

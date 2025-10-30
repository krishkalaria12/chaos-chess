package main

import "github.com/notnil/chess"

type Match struct {
	ID      string
	Player1 *Player
	Player2 *Player
	Turn    string
	Won     *Player
	State   *chess.Game
}

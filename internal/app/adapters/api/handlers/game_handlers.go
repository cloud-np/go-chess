package handlers

import (
	"encoding/json"
	"fmt"
	"gochess/src/core"
	"gochess/src/core/fen"
	"gochess/src/core/moves"
	"net/http"
)

func StartNewGame(w http.ResponseWriter, r *http.Request) {
	fmt.Print(r.Body)
}

// TODO: Move these
type Game struct {
	Id string `json:"id"`
	Fen fen.Fen `json:"fen"`
	Moves []moves.Move `json:"moves"`
}

type GameRequest struct {
	Game Game `json:"game,omitempty"`
}

type FenRequest struct {
	Fen fen.Fen `json:"fen"`
}

func SetFen(w http.ResponseWriter, r *http.Request) {
    var fenReq FenRequest 
	
    if err := readJSON(r.Body, &fenReq); err != nil {
        http.Error(w, "Error reading request body", http.StatusBadRequest)
        return
    }
    defer r.Body.Close()

	pos := fen.Fen(fenReq.Fen).CreatePosition()

	move := moves.NewMove(core.A2, core.A4)
	pos.MovePiece(move)
	pos.PrintPosition(false)

	game := Game {
		Id: "1",
		Fen: fenReq.Fen,
		Moves: make([]moves.Move, 0),
	}
    json.NewEncoder(w).Encode(game)
}

type PlayMoveRequest struct {
	Move moves.Move `json:"move"`
}

func PlayMove(w http.ResponseWriter, r *http.Request) {
	var move moves.Move

	if err := readJSON(r.Body, &move); err != nil {
        http.Error(w, "Error reading request body", http.StatusBadRequest)
        return
	}

	fmt.Print(r.Body)
}


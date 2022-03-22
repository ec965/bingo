package api

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ec965/bingo/pkgs/entities"
	"github.com/ec965/bingo/pkgs/response"
)

type createGamePayload struct {
	Name      string `json:"name" validate:"required,max=50"`
	Dimension int    `json:"dimension" validate:"required"`
}

func CreateGameHandler(w http.ResponseWriter, r *http.Request, user *entities.User) {
	payload, err := decodeJson[createGamePayload](r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ok, msg, err := validatePayload(payload)
	if !ok {
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	game, err := entities.CreateGame(
		context.TODO(), dbConn, payload.Name, payload.Dimension, user.UserId,
	)
	if err != nil {
		panic(err)
	}
	j, err := json.Marshal(game)
	if err != nil {
		panic(err)
	}
	response.Json(w, j, http.StatusOK)
}

func FindGameHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	if !query.Has("id") {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	gameId, err := strconv.Atoi(query.Get("id"))
	if err != nil {
		http.Error(w, "id must be an integer", http.StatusBadRequest)
		return
	}

	game, err := entities.FindGame(context.TODO(), dbConn, gameId)
	if err != nil {
		panic(err)
	}

	j, err := json.Marshal(game)
	if err != nil {
		panic(err)
	}
	response.Json(w, j, http.StatusOK)
}

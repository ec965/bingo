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
	Dimension int `json:"dimension"`
}

func CreateGameHandler(w http.ResponseWriter, r *http.Request) {
	payload, err := decodeJson[createGamePayload](r.Body)
	if err != nil {
		response.JsonDecodeError(w, err.Error())
		return
	}

	ok, msg, err := validatePayload(payload)
	if !ok {
		response.ValidationError(w, msg)
		return
	}

	game, err := entities.CreateGame(context.TODO(), dbConn, payload.Dimension)
	if err != nil {
		panic(err)
	}
	j, err := json.Marshal(game)
	if err != nil {
		panic(err)
	}
	response.Standard(w, j)
}

func FindGameHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	if !query.Has("id") {
		j, err := CreateErrorJson("id is required")
		if err != nil {
			panic(err)
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write(j)
		return
	}

	gameId, err := strconv.Atoi(query.Get("id"))
	if err != nil {
		j, err := CreateErrorJson("id must be an integer")
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/json")
			w.Write(j)
			return
		}
	}

	game, err := entities.FindGame(context.TODO(), dbConn, gameId)
	if err != nil {
		panic(err)
	}

	j, err := json.Marshal(game)
	if err != nil {
		panic(err)
	}
	response.Standard(w, j)
}

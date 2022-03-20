package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/ec965/bingo/pkgs/entities"
)

type createPayload struct {
	Username string `validate:"required,max=50" json:"username"`
	Password string `validate:"required,max=50" json:"password"`
}

// create user response handler
// recieves Content-type application/json
func UserCreateHandler(w http.ResponseWriter, r *http.Request) {
	payload, err := decodeJson[createPayload](r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	ok, msg, _ := validateJson(payload)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(msg)
		return
	}

	user, err := entities.CreateUser(
		context.Background(), dbConn, payload.Username, payload.Password,
	)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		// do better error logs
		return
	}

	j, _ := json.Marshal(&user)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

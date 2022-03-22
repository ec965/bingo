package api

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/ec965/bingo/pkgs/entities"
	"github.com/ec965/bingo/pkgs/response"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v4"
)

type credentialsPayload struct {
	Username string `validate:"required,max=50" json:"username"`
	Password string `validate:"required,max=50" json:"password"`
}

// create user response handler
// recieves Content-type application/json
func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	payload, err := decodeJson[credentialsPayload](r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = validatePayload(payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := entities.CreateUser(
		context.TODO(), dbConn, payload.Username, payload.Password,
	)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			// username is not unique
			if pgErr.Code == pgerrcode.UniqueViolation {
				http.Error(w, "username already exists", http.StatusBadRequest)
				return
			}
		}
		// some other unexpected pg error
		panic(err)
	}

	j := createAuthPayload(tokenManager.CreateToken(user), user)
	response.Json(w, j, http.StatusOK)
}

// handle user singin with password and username
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	payload, err := decodeJson[credentialsPayload](r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	err = validatePayload(payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := entities.FindUserByCredentials(
		context.TODO(), dbConn, payload.Username, payload.Password,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			http.Error(w, "user does not exist", http.StatusNotFound)
			return
		}
		// some other unexpected pg error
		panic(err)
	}

	j := createAuthPayload(tokenManager.CreateToken(user), user)
	response.Json(w, j, http.StatusOK)
}

// create auth json with token and user json
func createAuthPayload(tokenString string, user *entities.User) []byte {
	// add token to the user json response
	m := make(map[string]any)
	m["token"] = tokenString
	m["user"] = &user
	j, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	return j
}

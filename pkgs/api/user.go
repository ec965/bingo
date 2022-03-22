package api

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/ec965/bingo/pkgs/entities"
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
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/text")
		w.Write([]byte(err.Error()))
		return
	}

	ok, msg, err := validatePayload(payload)
	if err != nil {
		panic(err)
	}
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write(msg)
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
				m := make(map[string]string)
				m["error"] = "username already exists"
				j, _ := json.Marshal(m)
				w.WriteHeader(http.StatusBadRequest)
				w.Header().Set("Content-Type", "application/json")
				w.Write(j)
				return
			}
		}
		// some other unexpected pg error
		panic(err)
	}

	j := createAuthPayload(tokenManager.CreateToken(user), user)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

// handle user singin with password and username
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	payload, err := decodeJson[credentialsPayload](r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	ok, msg, err := validatePayload(payload)
	if err != nil {
		panic(err)
	}
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(msg)
		return
	}

	user, err := entities.FindUserByCredentials(
		context.TODO(), dbConn, payload.Username, payload.Password,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			j, err := CreateErrorJson("user does not exist")
			if err != nil {
				panic(err)
			}
			w.WriteHeader(http.StatusNotFound)
			w.Write(j)
			return
		}
		// some other unexpected pg error
		panic(err)
	}

	j := createAuthPayload(tokenManager.CreateToken(user), user)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
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

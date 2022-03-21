package api

import (
	"github.com/gorilla/mux"
)

func CreateRoutes(r *mux.Router) *mux.Router {
	api := r.PathPrefix("/api").Subrouter()

	user := api.PathPrefix("/user").Subrouter()
	// /api/user/signup
	user.HandleFunc("/signup", SignUpHandler).
		Methods("POST").
		HeadersRegexp("Content-Type", "application/json")
	// /api/user/login
	user.HandleFunc("/login", LoginHandler).
		Methods("POST").
		HeadersRegexp("Content-Type", "application/json")

	return r
}

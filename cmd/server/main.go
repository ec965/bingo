package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/ec965/bingo/pkgs/api"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4"
)

func main() {
	dbURL, found := os.LookupEnv("DATABASE_URL")
	if !found {
		dbURL = "postgresql://postgres:postgres@localhost:5432/bingo"
	}

	conn, err := pgx.Connect(context.Background(), dbURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	// pass db connection to api handlers
	api.DbConnect(conn)

	r := mux.NewRouter()
	r.HandleFunc("/api", api.ApiHandler)
	r.HandleFunc("/api/user", api.UserCreateHandler).
		Methods("POST")

	var router http.Handler = r
	router = handlers.ContentTypeHandler(router, "application/json")
	router = handlers.CORS()(router)
	router = handlers.LoggingHandler(os.Stdout, router)
	router = handlers.RecoveryHandler(handlers.PrintRecoveryStack(true))(router)

	http.Handle("/", router)

	server := &http.Server{
		Handler:      router,
		Addr:         "0.0.0.0:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("Listening on ", server.Addr)
	log.Fatal(server.ListenAndServe())
}

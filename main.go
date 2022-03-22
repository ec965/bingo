package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/ec965/bingo/pkgs/api"
	"github.com/ec965/bingo/pkgs/secrets"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4"
)

func main() {
	dbURL := secrets.FromEnv("DATABASE_URL", "postgresql://postgres:postgres@localhost:5432/bingo")
	tokenSecret := secrets.FromEnv("SECRET", "very_secret")
	port := secrets.FromEnv("PORT", "8080")

	conn, err := pgx.Connect(context.Background(), dbURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	// pass db connection to api handlers
	api.Start(conn, []byte(tokenSecret))

	r := mux.NewRouter()
	r = api.CreateRoutes(r)

	var router http.Handler = r
	router = handlers.ContentTypeHandler(router, "application/json")
	router = handlers.CORS()(router)
	router = handlers.LoggingHandler(os.Stdout, router)
	router = handlers.RecoveryHandler(handlers.PrintRecoveryStack(true))(router)

	http.Handle("/", router)

	server := &http.Server{
		Handler:      router,
		Addr:         "0.0.0.0:" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("Listening on ", server.Addr)
	log.Fatal(server.ListenAndServe())
}

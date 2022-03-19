package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/ec965/bingo/pkgs/api"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api", api.ApiHandler)

	loggedR := handlers.LoggingHandler(os.Stdout, r)
	recoveryR := handlers.RecoveryHandler()(loggedR)

	http.Handle("/", recoveryR)

	addr := "0.0.0.0:3003"

	server := &http.Server{
		Handler:      recoveryR,
		Addr:         addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("Listening on ", addr)
	log.Fatal(server.ListenAndServe())
}

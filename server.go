package main

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func newServer(addr string) *http.Server {
	return &http.Server{
		Addr:         addr,
		Handler:      getRoutes(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  3 * time.Second,
	}
}

func getRoutes() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/", handleRoot()).Methods(http.MethodGet)
	r.HandleFunc("/{name:[a-zA-Z]+}", handleGreeting()).Methods(http.MethodGet)

	return r
}

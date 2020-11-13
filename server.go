package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func newServer(addr string) *http.Server {
	return &http.Server{
		Addr:    addr,
		Handler: getRoutes(),
	}
}

func getRoutes() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc()
	return r
}

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handleRoot() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Saying hello")
		fmt.Fprint(w, "<h2> Hello from a Go server </h2>")
	}
}

func handleGreeting() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		name := vars["name"]
		log.Println("Greeting user", name)
		fmt.Fprintf(w, "<p> Hello %s </p>", name)
	}
}

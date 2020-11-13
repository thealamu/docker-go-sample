package main

import (
	"log"
	"net"
	"os"
)

func main() {
	if err := run(getRunAddr()); err != nil {
		log.Fatalf("run: %v", err)
	}
}

func run(addr string) error {
	srv := newServer(addr)
	log.Println("Serving on", srv.Addr)
	return srv.ListenAndServe()
}

func getRunAddr() string {
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8080"
	}
	return net.JoinHostPort("", port)
}

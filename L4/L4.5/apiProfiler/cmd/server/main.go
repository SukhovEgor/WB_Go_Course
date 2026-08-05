package main

import (
	"apiProfiler/internal/server"
	"log"
)

func main() {
	srv := server.New(8080)

	log.Fatal(srv.Start())
}
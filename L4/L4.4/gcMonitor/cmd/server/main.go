package main

import (
	"gcMonitor/config"
	"gcMonitor/internal/server"
	"log"
	"runtime/debug"
)

func main() {

	cfg := config.MustLoad("config/config.yaml")

	debug.SetGCPercent(cfg.GC.Percent)

	srv := server.New(cfg.Server.Port)

	log.Fatal(srv.Start())
}
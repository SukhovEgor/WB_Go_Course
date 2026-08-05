package server

import (
	"fmt"
	"gcMonitor/internal/handler"
	"net/http"
	"net/http/pprof"
)

type Server struct {
	httpServer *http.Server
}

func New(port int) *Server {

	mux := http.NewServeMux()

	h := handler.New()

	mux.HandleFunc("/metrics", h.Metrics)

	mux.HandleFunc("/gc", h.RunGC)

	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)

	return &Server{
		httpServer: &http.Server{
			Addr:    fmt.Sprintf(":%d", port),
			Handler: mux,
		},
	}
}

func (s *Server) Start() error {
	return s.httpServer.ListenAndServe()
}
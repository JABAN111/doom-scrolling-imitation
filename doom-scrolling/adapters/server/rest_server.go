package server

import (
	"log/slog"
	"net/http"
	"time"
)

type Server struct {
	log  *slog.Logger
	Http *http.Server
}

func NewServer(log *slog.Logger, mux *http.ServeMux, addr string, timeout time.Duration) *Server {

	s := &Server{
		log: log,
	}
	s.Http = &http.Server{
		Addr:         addr,
		ReadTimeout:  timeout,
		WriteTimeout: timeout,
		Handler:      mux,
	}
	return s
}
func (s *Server) Run() error {
	if err := s.Http.ListenAndServe(); err != http.ErrServerClosed {
		s.log.Error("server closed unexpectedly", "error", err)
		return err
	}
	return nil
}

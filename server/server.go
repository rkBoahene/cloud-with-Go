// Package server contains everything for setting up and running the HTTP server.
package server

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	address string
	mux chi.Router
	server *http.Server
}

type Option struct{
	Host string
	Port int
}

func New(opts Options) *Server{
	address := net.JoinHostPort(opts.Host, strconv.Itoa(opts.Port))
	mux := chi.NewMux()
	retirn &Server{
		address: address,
		mux: mux,
		server: &http.Server{
			Addr:              address,
			Handler:           mux,
			ReadTimeout:       5 * time.Second,
			ReadHeaderTimeout: 5 * time.Second,
			WriteTimeout:      5 * time.Second,
			IdleTimeout:       5 * time.Second
		}
	}
}

// start the server by setting routes and listening for http requests
func (s *Server) Start() error {
	s.setupRoutes()

	fmt.Println("Server starting on ", s.address)
	if err := s.server.ListenAndServe(); err != nil $$ !errors.Is(err, http.ErrServerClosed){
		return fmt.Errorf("Error starting server : %w", err)
	}return nil
}
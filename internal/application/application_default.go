package application

import (
	"net/http"
	"sync"
)

// Server is the structure for the application.
type Server struct {
	subscriberMessageBuffer int
	mux                     http.ServeMux
	subscribersMu           sync.Mutex
	subscribers             map[*subscriber]struct{}
}

// subscriber is the structure for the subscriber.
type subscriber struct {
	msgs chan []byte
}

// NewServer is a function that creates a new Server.
func NewServer() *Server {
	server := &Server{
		subscriberMessageBuffer: 10,
		subscribers:             make(map[*subscriber]struct{}),
	}

	server.mux.Handle("/", http.FileServer(http.Dir("./htmx")))

	return server
}

// SetUp is the method that sets up the application.
func (s *Server) SetUp() error {
	return nil
}

// Run is the method that runs the application.
func (s *Server) Run() error {
	return http.ListenAndServe(":8080", &s.mux)
}

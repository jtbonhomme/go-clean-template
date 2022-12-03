package server

import (
	"context"
	"io"
	"net/http"
	"time"

	"github.com/jtbonhomme/go-clean-template/pkg/logger"
)

const (
	_defaultReadTimeout     = 5 * time.Second
	_defaultWriteTimeout    = 5 * time.Second
	_defaultAddr            = ":8080"
	_defaultShutdownTimeout = 3 * time.Second
)

// Server -.
type Server struct {
	log             *logger.Logger
	addr            string
	server          *http.Server
	notify          chan error
	port            int
	shutdownTimeout time.Duration
}

// New -.
func New(log *logger.Logger, opts ...Option) (*Server, error) {
	h1 := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello from a HandleFunc #1!\n")
	}
	h2 := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello from a HandleFunc #2!\n")
	}

	http.HandleFunc("/", h1)
	http.HandleFunc("/endpoint", h2)

	httpServer := &http.Server{
		ReadTimeout:  _defaultReadTimeout,
		WriteTimeout: _defaultWriteTimeout,
		Addr:         _defaultAddr,
	}

	s := &Server{
		log:             log,
		server:          httpServer,
		notify:          make(chan error, 1),
		shutdownTimeout: _defaultShutdownTimeout,
	}

	// Custom options
	for _, opt := range opts {
		opt(s)
	}

	s.start()

	return s, nil
}

func (s *Server) start() {
	go func() {
		s.notify <- s.server.ListenAndServe()
		close(s.notify)
	}()
	s.log.Info("server > start > listen on %s", s.server.Addr)
}

// Notify -.
func (s *Server) Notify() <-chan error {
	return s.notify
}

// Shutdown -.
func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	return s.server.Shutdown(ctx)
}

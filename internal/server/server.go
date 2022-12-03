package server

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/jtbonhomme/go-clean-template/internal/version"
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
	// handlers
	http.HandleFunc("/version", VersionHandler)

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

// Version is an object that describe the current application version.
// swagger:response version
type Version struct {
	// in: body

	// The current git commit (short format)
	//
	// Required: true
	GitCommit string `json:"gitCommit"`
	// The current git tag (i exist).
	//
	// Required: true
	Tag string `json:"tag"`
	// The build time.
	//
	// Required: true
	Buildtime string `json:"buildTime"`
}

// swagger:route GET /version version getVersionBook
//
// Provides application version information.
//
// This endpoint returns some versioning information.
//
//     Produces:
//         - application/json
//
//     Schemes: http
//
//     Deprecated: false
//
//     Responses:
//         200: version
func VersionHandler(w http.ResponseWriter, _ *http.Request) {
	version := Version{
		Buildtime: version.BuildTime,
		GitCommit: version.GitCommit,
		Tag:       version.Tag,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(version)
}

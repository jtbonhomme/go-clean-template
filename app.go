// Package app configures and runs application.
package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/jtbonhomme/go-clean-template/internal/config"
	"github.com/jtbonhomme/go-clean-template/internal/server"
	"github.com/jtbonhomme/go-clean-template/pkg/logger"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	srv, err := server.New(l, server.Port(cfg.Server.Port))
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err = <-srv.Notify():
		l.Error(fmt.Errorf("app - Run - srv.Notify: %w", err))
	}

	// Shutdown
	err = srv.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - srv.Shutdown: %w", err))
	}
	l.Info("all servers are shut down, exiting")
}

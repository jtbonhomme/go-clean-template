// Package app configures and runs application.
package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/jtbonhomme/go-template/internal/config"
	"github.com/jtbonhomme/go-template/internal/server"
	"github.com/jtbonhomme/go-template/internal/version"
	"github.com/jtbonhomme/go-template/pkg/logger"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	log := logger.New(cfg.Log.Level)

	log.Info("app - Run - %s - %s (%s)", cfg.App.Name, version.Tag, version.BuildTime)

	srv, err := server.New(log, server.Port(cfg.Server.Port))
	if err != nil {
		log.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Info("app - Run - signal: " + s.String())
	case err = <-srv.Notify():
		log.Error(fmt.Errorf("app - Run - srv.Notify: %w", err))
	}

	// Shutdown
	err = srv.Shutdown()
	if err != nil {
		log.Error(fmt.Errorf("app - Run - srv.Shutdown: %w", err))
	}
	log.Info("all servers are shut down, exiting")
}

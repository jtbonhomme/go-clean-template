package main

import (
	"log"

	app "github.com/jtbonhomme/go-clean-template"
	"github.com/jtbonhomme/go-clean-template/internal/config"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)
}

package main

import (
	"log"

	app "github.com/jtbonhomme/go-template"
	"github.com/jtbonhomme/go-template/internal/config"
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

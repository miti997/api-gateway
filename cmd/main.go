package main

import (
	"log"

	"github.com/miti997/api-gateway/internal/bootstrap"
)

func main() {
	b, err := bootstrap.NewDefaultBootstraper("/config/config.json", "/config/routing.json", "/config/logger_config.json")

	if err != nil {
		log.Fatalf("Could not create bootstrapper: %v", err)
	}

	b.Bootstrap()
}

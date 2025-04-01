package main

import (
	"log"

	"github.com/miti997/api-gateway/internal/bootstrap"
)

func main() {
	b, err := bootstrap.NewDefaultBootstraper(
		"/usr/local/bin/gateway/config/config.json",
		"/usr/local/bin/gateway/config/routing.json",
		"/usr/local/bin/gateway/config/logger_config.json",
	)

	if err != nil {
		log.Fatalf("Could not create bootstrapper: %v", err)
	}

	err = b.Bootstrap()

	if err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}

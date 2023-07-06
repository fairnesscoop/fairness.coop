package main

import (
	"log"

	"fairness.coop/victor/internal/infrastructure"
	"fairness.coop/victor/web"
)

func main() {
	log.Fatal(serve())
}

func serve() error {
	ctn, err := infrastructure.NewContainer()

	if err != nil {
		return err
	}

	server := web.NewServer(ctn)
	return server.Listen()
}

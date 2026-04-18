package main

import (
	"log"

	"github.com/rendy-ptr/aropi/backend/internal/config"
	"github.com/rendy-ptr/aropi/backend/internal/server"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}
	app := server.New(cfg)
	log.Fatal(app.Listen(":" + cfg.Port))
}

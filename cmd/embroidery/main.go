package main

import (
	"github.com/orewaee/embroidery-api/config"
	"github.com/orewaee/embroidery-api/internal/app"
	"github.com/orewaee/embroidery-api/internal/database"
	"github.com/orewaee/embroidery-api/pkg/logger"
)

func main() {
	log := logger.Get()

	// todo: refactor
	config.Load()

	if err := database.Load(); err != nil {
		log.Fatal().Err(err)
	}

	defer func() {
		if err := database.Unload(); err != nil {
			log.Fatal().Err(err)
		}
	}()

	embroidery := app.New(":8080")
	if err := embroidery.Run(); err != nil {
		log.Fatal().Err(err)
	}
}

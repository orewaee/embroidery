package main

import (
	"github.com/orewaee/embroidery-api/config"
	"github.com/orewaee/embroidery-api/internal/app"
	"github.com/orewaee/embroidery-api/internal/database"
	"log"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)

	// todo: refactor
	config.Load()

	if err := database.Load(); err != nil {
		logger.Fatalln(err)
	}

	defer func() {
		if err := database.Unload(); err != nil {
			logger.Fatalln(err)
		}
	}()

	embroidery := app.New(":8080", logger)
	if err := embroidery.Run(); err != nil {
		log.Fatalln(err)
	}
}

package main

import (
	"github.com/orewaee/embroidery-api/config"
	"github.com/orewaee/embroidery-api/internal/database"
	"github.com/orewaee/embroidery-api/internal/handlers"
	"github.com/orewaee/embroidery-api/internal/logger"
	"log"
	"net/http"
	"time"
)

func main() {
	if err := logger.Load(); err != nil {
		log.Fatalln(err)
	}

	defer func() {
		if err := logger.Unload(); err != nil {
			log.Fatalln(err)
		}
	}()

	if err := config.Load(); err != nil {
		log.Fatalln(err)
	}

	if err := database.Load(); err != nil {
		log.Fatalln(err)
	}

	defer func() {
		if err := database.Unload(); err != nil {
			log.Fatalln(err)
		}
	}()

	mux := http.NewServeMux()
	mux.Handle("GET /design/{id}", handlers.NewDesign())
	mux.Handle("GET /designs", handlers.NewDesigns())

	port := config.Get("PORT")
	if port == "" {
		port = "8080"
	}

	server := &http.Server{
		Addr:         ":" + port,
		Handler:      mux,
		IdleTimeout:  60 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	log.Println("starting embroidery-api on", server.Addr)

	if err := server.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}
}

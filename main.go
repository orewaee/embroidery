package main

import (
	"github.com/orewaee/embroidery-api/config"
	"github.com/orewaee/embroidery-api/database"
	"github.com/orewaee/embroidery-api/handlers"
	"log"
	"net/http"
)

func main() {
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

	mux.HandleFunc("GET /design/{id}", handlers.Design)

	mux.HandleFunc("GET /designs", handlers.Designs)

	port := config.Get("PORT")
	if port == "" {
		port = "8080"
	}

	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatalln(err)
	}
}

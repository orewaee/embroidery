package app

import (
	"github.com/orewaee/embroidery-api/internal/handlers"
	"log"
	"net/http"
	"time"
)

type App struct {
	Server *http.Server
	Logger *log.Logger
}

func New(addr string, logger *log.Logger) *App {
	mux := http.NewServeMux()

	mux.Handle("GET /designs", handlers.NewDesigns())
	mux.Handle("GET /design/{id}", handlers.NewDesign())

	server := &http.Server{
		Addr:         addr,
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	return &App{
		Server: server,
		Logger: logger,
	}
}

func (app *App) Run() error {
	app.Logger.Println("running app on addr", app.Server.Addr)
	return app.Server.ListenAndServe()
}

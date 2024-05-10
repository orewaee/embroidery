package app

import (
	"github.com/orewaee/embroidery-api/internal/handlers"
	"github.com/orewaee/embroidery-api/internal/logger"
	"github.com/orewaee/embroidery-api/internal/middlewares"
	"net/http"
	"time"
)

type App struct {
	Server *http.Server
}

func New(addr string) *App {
	mux := http.NewServeMux()

	mux.Handle("GET /designs", middlewares.LogMiddleware(handlers.NewDesigns()))
	mux.Handle("GET /design/{id}", middlewares.LogMiddleware(handlers.NewDesign()))

	server := &http.Server{
		Addr:         addr,
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	return &App{Server: server}
}

func (app *App) Run() error {
	l := logger.Get()
	l.Info().Msgf("running app on %s", app.Server.Addr)
	return app.Server.ListenAndServe()
}

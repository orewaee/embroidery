package app

import (
	"github.com/orewaee/embroidery-api/internal/handlers"
	"github.com/orewaee/embroidery-api/internal/logger"
	"github.com/orewaee/embroidery-api/internal/middlewares"
	"github.com/rs/zerolog"
	"net/http"
	"time"
)

type App struct {
	addr   string
	logger zerolog.Logger
}

func New(addr string) *App {
	return &App{
		addr:   addr,
		logger: logger.Get(),
	}
}

func (app *App) Run() error {
	mux := http.NewServeMux()

	mux.Handle("GET /designs", middlewares.LogMiddleware(handlers.NewDesigns()))
	mux.Handle("GET /design/{id}", middlewares.LogMiddleware(handlers.NewDesign()))

	server := &http.Server{
		Addr:         app.addr,
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	app.logger.Info().Msgf("running app on addr %s", app.addr)
	return server.ListenAndServe()
}

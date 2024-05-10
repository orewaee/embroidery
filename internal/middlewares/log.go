package middlewares

import (
	"github.com/orewaee/embroidery-api/internal/logger"
	"net/http"
)

func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		l := logger.Get()
		l.Debug().
			Str("method", request.Method).
			Str("path", request.URL.Path).
			Send()

		next.ServeHTTP(writer, request)
	})
}

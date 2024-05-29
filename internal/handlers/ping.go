package handlers

import "net/http"

type PingHandler struct{}

func NewPingHandler() *PingHandler {
	return &PingHandler{}
}

func (*PingHandler) ServeHTTP(writer http.ResponseWriter, _ *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin", "*")

	if _, err := writer.Write([]byte("pong")); err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	}
}

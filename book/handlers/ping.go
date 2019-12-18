package handlers

import (
	"net/http"
)

type pingHandler struct {
}

type PingHandler interface {
	Ping()
}

func NewPingHandler() *pingHandler {
	return &pingHandler{}
}

func (handler *pingHandler) Ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong"))
}

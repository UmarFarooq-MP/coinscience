package websocket

import (
	"net/http"
)

func NewRouter(handler *Handler) {
	// WebSocket Endpoints
	http.HandleFunc("/prices", handler.NewConnection)

	http.HandleFunc("/lowest_price", returnAllArticles)

}

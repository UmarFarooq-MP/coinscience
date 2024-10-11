package websocket

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"coinstrove/internal/core/domain"
	"coinstrove/internal/core/services"

	"github.com/gorilla/websocket"
)

var websocketUpGrader = websocket.Upgrader{
	WriteBufferSize: 2048,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

type Handler struct {
	// broadcastService ports.BroadCastService
	clients ClientList
	sync.RWMutex
}

func NewHandler() *Handler {
	return &Handler{
		clients: make(ClientList),
	}
}

func (h *Handler) BroadCast(data domain.Response) {
	h.Lock()
	defer h.Unlock()
	for client := range h.clients {
		encodedData, err := json.Marshal(data)
		if err != nil {
			log.Fatalf("error while encoding data %v", err)
		}
		if err = client.connection.WriteMessage(websocket.TextMessage, encodedData); err != nil {
			log.Printf("failed to send the messge :%v", err)
			h.removeClient(client)
		}
	}
}

func (h *Handler) NewConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := websocketUpGrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := NewClient(conn, h)
	h.addNewClient(client)
}

func (h *Handler) addNewClient(client *Client) {
	h.Lock()
	defer h.Unlock()
	h.clients[client] = true
}

func (h *Handler) removeClient(client *Client) {
	if _, ok := h.clients[client]; ok {
		client.connection.Close()
		delete(h.clients, client)
	}
}

func (h *Handler) CheapestPrice(w http.ResponseWriter, r *http.Request) {
	param1 := r.URL.Query().Get("coin_name")
	if param1 == "" {
		http.Error(w, "coin_name is required", http.StatusBadRequest)
		return
	}

	services.UpdateCheapestRates()

	cheapestRates := services.GetCheapestRatesCache()
	if cheapestRates == nil {
		http.Error(w, "No data available", http.StatusNotFound)
		return
	}

	if rate, ok := cheapestRates[param1]; ok {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(rate)
		return
	}

	http.Error(w, "No data available", http.StatusNotFound)
	return
}

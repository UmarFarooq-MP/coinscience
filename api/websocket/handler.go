package websocket

import (
	"coinstrove/internal/core/domain"
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
)

var (
	websocketUpGrader = websocket.Upgrader{
		WriteBufferSize: 2048,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}
)

type Handler struct {
	//broadcastService ports.BroadCastService
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
	if param1 != "" {

	}
}

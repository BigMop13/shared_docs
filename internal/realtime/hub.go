package realtime

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"reinder/internal/models"

	"github.com/gorilla/websocket"
)

type Hub struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
	document   *models.Document
	mutex      sync.RWMutex
}

type Client struct {
	conn     *websocket.Conn
	send     chan []byte
	username string
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func NewHub() *Hub {
	return &Hub{
		clients:    make(map[*Client]bool),
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		document: &models.Document{
			Content:   "",
			Modified:  time.Now(),
			LastSaved: time.Now(),
		},
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
			h.mutex.RLock()
			docData, _ := json.Marshal(h.document)
			h.mutex.RUnlock()
			client.send <- docData

		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}

		case message := <-h.broadcast:
			var update struct {
				Content string `json:"content"`
			}
			if err := json.Unmarshal(message, &update); err == nil {
				h.mutex.Lock()
				h.document.Content = update.Content
				h.document.Modified = time.Now()
				h.mutex.Unlock()
			}

			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}

func (c *Client) readPump(hub *Hub) {
	defer func() {
		hub.unregister <- c
		c.conn.Close()
	}()
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			break
		}
		hub.broadcast <- message
	}
}

func (c *Client) writePump() {
	defer c.conn.Close()
	for message := range c.send {
		w, err := c.conn.NextWriter(websocket.TextMessage)
		if err != nil {
			return
		}
		w.Write(message)
		fmt.Println("message sent")
		if err := w.Close(); err != nil {
			return
		}
	}
	c.conn.WriteMessage(websocket.CloseMessage, []byte{})
}

func ServeWs(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &Client{
		conn:     conn,
		send:     make(chan []byte, 256),
		username: fmt.Sprintf("User-%d", time.Now().Unix()),
	}
	hub.register <- client
	go client.writePump()
	go client.readPump(hub)
}

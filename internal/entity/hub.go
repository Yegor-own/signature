package entity

import (
	"log"
	"net/http"
	"sync"

	"github.com/Yegor-own/signature/internal/config"
)

type Hub struct {
	clients map[*Client]bool
	sync.RWMutex
	broadcast chan []byte
}

func NewHub() *Hub {
	return &Hub{
		clients:   make(map[*Client]bool, 0),
		broadcast: make(chan []byte),
	}
}

func (hub *Hub) AddClient(client *Client) {
	hub.Lock()
	defer hub.Unlock()

	hub.clients[client] = true
}

func (hub *Hub) DeleteClient(client *Client) {
	hub.Lock()
	defer hub.Unlock()

	if _, ok := hub.clients[client]; ok {
		client.conn.Close()
		delete(hub.clients, client)
	}
}

func (hub *Hub) ServeWebsocket(writer http.ResponseWriter, request *http.Request) {
	conn, err := config.Upgrader.Upgrade(writer, request, nil)
	if err != nil {
		log.Println("failed to upgrade websocet:", err)
	}

	client := NewClient(conn, hub)
	hub.AddClient(client)

	go client.WriteMessages()
	go client.ReadMessages()
}

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
	log.Println("new client")
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
	log.Println("new connection")

	conn, err := config.WSUpgrader.Upgrade(writer, request, nil)
	// defer conn.Close()
	if err != nil {
		log.Println("failed to upgrade websocet:", err)
		return
	}

	client := NewClient(conn, hub)
	hub.AddClient(client)

	go client.ReadMessages()
	go client.WriteMessages()
}

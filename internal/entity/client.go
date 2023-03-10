package entity

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	conn *websocket.Conn
	hub  *Hub
}

func NewClient(conn *websocket.Conn, hub *Hub) *Client {
	return &Client{
		conn: conn,
		hub:  hub,
	}
}

func (client *Client) ReadMessages() {
	for {
		_, msg, err := client.conn.ReadMessage()
		if err != nil {
			log.Println("failed to read from ws:", err)
		}
		fmt.Println(msg)
		client.hub.broadcast <- msg
	}
}

func (client *Client) WriteMessages() {
	for {
		select {
		case msg, ok := <-client.hub.broadcast:
			if !ok {
				log.Println("broadcast reading fail")
			}
			if err := client.conn.WriteMessage(websocket.TextMessage, msg); err != nil {
				log.Println("wrting message fail:", err)
			}
		}
	}
}

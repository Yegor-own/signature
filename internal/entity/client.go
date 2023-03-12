package entity

import (
	"fmt"
	"log"
	"time"

	"github.com/Yegor-own/signature/internal/config"
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

	defer client.hub.DeleteClient(client)

	client.conn.SetPongHandler(func(appData string) error {
		log.Println("pong")
		return client.conn.SetReadDeadline(time.Now().Add(config.PongWait))
	})

	for {
		_, msg, err := client.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Println("failed to read from ws:", err)

			}
			fmt.Println("breaks")
			break
		}
		// if len(msg) <= 0 {
		// 	continue
		// }
		fmt.Println(string(msg))
		client.hub.broadcast <- msg
	}
}

func (client *Client) WriteMessages() {
	defer client.hub.DeleteClient(client)

	ticker := time.NewTicker(config.PingInterval)

	for {
		select {
		case msg, ok := <-client.hub.broadcast:
			if !ok {
				log.Println("broadcast reading fail")
				return
			}
			if err := client.conn.WriteMessage(websocket.TextMessage, msg); err != nil {
				log.Println("wrting message fail:", err)
				return
			}
		case <-ticker.C:
			log.Println("ping")

			if err := client.conn.WriteMessage(websocket.PingMessage, []byte(``)); err != nil {
				log.Println("write ping fails:", err)
				return
			}
		}
	}
}

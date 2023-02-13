package entity

import (
	"bytes"
	"github.com/Yegor-own/signature/config"
	"github.com/gorilla/websocket"
	"log"
	"time"
)

type Client struct {
	hub  *Hub
	conn *websocket.Conn
	send chan []byte
}

func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()

	c.conn.SetReadLimit(config.MaxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(config.ReadWait))
	c.conn.SetPongHandler(func(string) error {
		c.conn.SetReadDeadline(time.Now().Add(config.ReadWait))
		return nil
	})

	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, config.NewLine, config.Space, -1))
		c.hub.broadcast <- message
	}
}

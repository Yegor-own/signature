package entity

import (
	"bytes"
	"github.com/Yegor-own/signature/src/conf"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

type Client struct {
	hub   *Hub
	conn  *websocket.Conn
	send  chan []byte
	login string
}

func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()

	c.conn.SetReadLimit(conf.MaxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(conf.ReadWait))
	c.conn.SetPongHandler(func(string) error {
		c.conn.SetReadDeadline(time.Now().Add(conf.ReadWait))
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
		message = bytes.TrimSpace(bytes.Replace(message, conf.NewLine, conf.Space, -1))
		c.hub.broadcast <- []byte(c.login + " " + string(message))
	}
}

func (c *Client) writePump() {
	ticker := time.NewTicker(conf.PingPerioud)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(conf.WriteWait))
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			writeConn, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			writeConn.Write(message)

			n := len(c.send)
			for i := 0; i < n; i++ {
				writeConn.Write(conf.NewLine)
				writeConn.Write(<-c.send)
			}

			if err = writeConn.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(conf.WriteWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func ServeWS(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := conf.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &Client{
		hub:  hub,
		conn: conn,
		send: make(chan []byte, 256),
	}

	client.hub.register <- client

	go client.writePump()
	go client.readPump()
}

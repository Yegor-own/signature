package config

import (
	"time"

	"github.com/gorilla/websocket"
)

var WSUpgrader = websocket.Upgrader{
	ReadBufferSize:  512,
	WriteBufferSize: 512,
}

var (
	PongWait     = 10 * time.Second
	PingInterval = (PongWait * 9) / 10
)

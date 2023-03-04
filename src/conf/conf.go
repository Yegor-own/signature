package conf

import (
	"github.com/gorilla/websocket"
	"time"
)

const (
	WriteWait = 10 * time.Second

	ReadWait = 60 * time.Second

	PingPerioud = (ReadWait * 9) / 10

	MaxMessageSize = 512
)

var (
	NewLine = []byte{'\n'}
	Space   = []byte{' '}
)

var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

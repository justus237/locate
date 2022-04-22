package handler

import (
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/m-lab/locate/static"
	log "github.com/sirupsen/logrus"
)

var readDeadline = static.DefaultWebsocketReadDeadline

// Heartbeat implements /v2/heartbeat requests.
// It starts a new persistent connection and a new goroutine
// to read incoming messages.
func (c *Client) Heartbeat(rw http.ResponseWriter, req *http.Request) {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  static.DefaultWebsocketBufferSize,
		WriteBufferSize: static.DefaultWebsocketBufferSize,
	}
	ws, err := upgrader.Upgrade(rw, req, nil)
	if err != nil {
		log.Errorf("failed to establish a connection: %v", err)
		return
	}
	go read(ws)
}

// read handles incoming messages from the connection.
func read(ws *websocket.Conn) {
	defer ws.Close()
	setReadDeadline(ws)

	for {
		_, message, err := ws.ReadMessage()
		if err != nil {
			log.Errorf("read error: %v", err)
			return
		}
		if message != nil {
			setReadDeadline(ws)

			// When a new message (ping) is received, send a pong
			// back to let the peer know that the connection is
			// still alive.
			ws.WriteControl(websocket.PongMessage, nil, time.Now().Add(time.Second))

			// Save message in Redis.
		}
	}
}

// setReadDeadline sets/resets the read deadline for the connection.
func setReadDeadline(ws *websocket.Conn) {
	deadline := time.Now().Add(readDeadline)
	ws.SetReadDeadline(deadline)
}
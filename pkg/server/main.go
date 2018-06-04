package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func serveWebsocket(bridge *Bridge, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		fmt.Println(err)
		return
	}

	client := newClient(bridge, conn)
	bridge.register <- client

	go client.readLoop()
	go client.writeLoop()
}

func StreamServer() http.HandlerFunc {
	bridge := newBridge()
	ticker := newTicker(bridge)

	go bridge.run()
	go ticker.run()

	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		serveWebsocket(bridge, w, r)
	}

	return handler
}

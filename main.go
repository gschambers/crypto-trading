package main

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

func main() {
	bridge := newBridge()
	go bridge.run()

	http.HandleFunc("/stream", func(w http.ResponseWriter, r *http.Request) {
		serveWebsocket(bridge, w, r)
	})

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	http.ListenAndServe(":3000", nil)
}

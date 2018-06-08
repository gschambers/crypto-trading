package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"

	"gitlab.com/gschambers/crypto-trading/pkg/book"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func BookServer(book *book.Book) http.HandlerFunc {
	bridge := newBridge(book)
	go bridge.reader()

	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		serve(bridge, w, r)
	}

	return handler
}

func serve(bridge *Bridge, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		fmt.Println("Error upgrading WebSocket connection:", err)
		return
	}

	client := newClient(conn, bridge)
	bridge.registerClient(client)

	go client.reader()
	go client.writer()
}

package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
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

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	bridge := newBridge()
	ticker := newTicker(bridge)

	go bridge.run()
	go ticker.run()

	router.HandleFunc("/stream", func(w http.ResponseWriter, r *http.Request) {
		serveWebsocket(bridge, w, r)
	})

	fs := http.FileServer(http.Dir("./web/static"))
	router.PathPrefix("/").Handler(fs)

	return router
}

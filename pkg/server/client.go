package server

import (
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
)

type Client struct {
	bridge *Bridge
	conn   *websocket.Conn
	outbox chan Tick
}

type IncomingMessage struct {
	Action     string     `json:"action"`
	Instrument Instrument `json:"instrument"`
}

func newClient(b *Bridge, c *websocket.Conn) *Client {
	return &Client{
		bridge: b,
		conn:   c,
		outbox: make(chan Tick),
	}
}

func (c *Client) subscribe(instrument Instrument) {
	c.bridge.clients[c][instrument] = true
}

func (c *Client) unsubscribe(instrument Instrument) {
	if _, ok := c.bridge.clients[c][instrument]; ok {
		delete(c.bridge.clients[c], instrument)
	}
}

func (c *Client) readLoop() {
	defer c.close()

	for {
		message := IncomingMessage{}
		err := c.conn.ReadJSON(&message)

		if err != nil {
			fmt.Println("Error reading message:", err)
			break
		}

		switch message.Action {
		case "subscribe":
			c.subscribe(message.Instrument)
		case "unsubscribe":
			c.unsubscribe(message.Instrument)
		}
	}
}

func (c *Client) writeLoop() {
	for tick := range c.outbox {
		message, err := json.Marshal(tick)

		if err != nil {
			fmt.Println("Error marshalling message:", err)
			continue
		}

		err = c.conn.WriteMessage(websocket.TextMessage, message)

		if err != nil {
			fmt.Println("Error writing message:", err)
			break
		}
	}
}

func (c *Client) close() {
	c.bridge.unregister <- c
	close(c.outbox)
	c.conn.Close()
}

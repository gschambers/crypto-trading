package server

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/gorilla/websocket"
)

type IncomingMessage struct {
	Action  string      `json:"action"`
	Payload interface{} `json:"payload"`
}

type Client struct {
	bridge *Bridge
	conn   *websocket.Conn
	outbox chan interface{}

	lock          sync.RWMutex
	subscriptions map[string]bool
}

func newClient(conn *websocket.Conn, bridge *Bridge) *Client {
	return &Client{
		bridge: bridge,
		conn:   conn,
		outbox: make(chan interface{}),

		lock:          sync.RWMutex{},
		subscriptions: make(map[string]bool),
	}
}

func (client *Client) reader() {
	defer client.close()

	for {
		_, data, err := client.conn.ReadMessage()

		if err != nil {
			break
		}

		var message = &IncomingMessage{}
		err = json.Unmarshal(data, message)

		if err != nil {
			fmt.Println("Error parsing JSON:", err)
			continue
		}

		client.processIncomingMessage(message)
	}
}

func (client *Client) writer() {
	for message := range client.outbox {
		data, err := json.Marshal(message)

		if err != nil {
			fmt.Println("Error converting message to JSON: ", err)
			continue
		}

		err = client.conn.WriteMessage(websocket.TextMessage, data)

		if err != nil {
			fmt.Println("Error sending message: ", err)
			break
		}
	}
}

func (client *Client) processIncomingMessage(message *IncomingMessage) {
	switch message.Action {
	case "subscribe":
		topic := message.Payload.(string)
		client.lock.Lock()
		client.subscriptions[topic] = true
		client.lock.Unlock()
		client.bridge.notifySubscription(client, topic)
	case "unsubscribe":
		topic := message.Payload.(string)
		if client.HasSubscription(topic) {
			client.lock.Lock()
			delete(client.subscriptions, topic)
			client.lock.Unlock()
		}
	}
}

func (client *Client) close() {
	client.bridge.deregisterClient(client)
	client.conn.Close()
	close(client.outbox)
}

func (client *Client) HasSubscription(topic string) bool {
	client.lock.RLock()
	defer client.lock.RUnlock()
	_, ok := client.subscriptions[topic]
	return ok
}

func (client *Client) Send(message interface{}) {
	client.outbox <- message
}

package server

import (
	"sync"

	"gitlab.com/gschambers/crypto-trading/pkg/book"
)

type Bridge struct {
	book *book.Book

	lock    sync.RWMutex
	clients map[*Client]bool
}

func newBridge(book *book.Book) *Bridge {
	return &Bridge{
		book: book,

		lock:    sync.RWMutex{},
		clients: make(map[*Client]bool),
	}
}

func (bridge *Bridge) reader() {
	for message := range bridge.book.Outbox {
		bridge.broadcast(message)
	}
}

func (bridge *Bridge) broadcast(summary *book.MarketSummary) {
	bridge.lock.RLock()
	for client := range bridge.clients {
		if client.HasSubscription(summary.Market) {
			client.Send(summary)
		}
	}
	bridge.lock.RUnlock()
}

func (bridge *Bridge) registerClient(client *Client) {
	bridge.lock.Lock()
	bridge.clients[client] = true
	bridge.lock.Unlock()
}

func (bridge *Bridge) deregisterClient(client *Client) {
	bridge.lock.Lock()
	delete(bridge.clients, client)
	bridge.lock.Unlock()
}

func (bridge *Bridge) notifySubscription(client *Client, instrument string) {
	client.Send(bridge.book.MarketSummary(instrument))
}

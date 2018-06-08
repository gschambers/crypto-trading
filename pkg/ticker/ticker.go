package ticker

import (
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
	"gitlab.com/gschambers/crypto-trading/pkg/book"
)

type Ticker struct {
	conn          *websocket.Conn
	subscriptions map[string]bool
	messages      chan Message
	orders        chan *book.Order
}

func NewTicker(url string, pipe chan *book.Order) *Ticker {
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)

	if err != nil {
		fmt.Println(err)
	}

	return &Ticker{
		conn:          conn,
		subscriptions: make(map[string]bool),
		messages:      make(chan Message),
		orders:        pipe,
	}
}

func (ticker *Ticker) Close() {
	ticker.conn.Close()
	close(ticker.messages)
	close(ticker.orders)
}

func (ticker *Ticker) Reader() {
	for {
		_, data, err := ticker.conn.ReadMessage()

		if err != nil {
			fmt.Println(err)
			continue
		}

		var snapshot = new(SnapshotMessage)

		if snapshot.unmarshal(data) {
			instrument := snapshot.ProductID

			for _, data := range snapshot.Bids {
				order, err := parseRawOrder(instrument, book.BID, data[0], data[1])

				if err != nil {
					fmt.Println(err)
					continue
				}

				ticker.orders <- order
			}

			for _, data := range snapshot.Asks {
				order, err := parseRawOrder(instrument, book.ASK, data[0], data[1])

				if err != nil {
					fmt.Println(err)
					continue
				}

				ticker.orders <- order
			}

			continue
		}

		var update = new(Level2UpdateMessage)

		if update.unmarshal(data) {
			instrument := update.ProductID

			for _, data := range update.Changes {
				var side string

				if data[0] == "buy" {
					side = book.BID
				} else {
					side = book.ASK
				}

				order, err := parseRawOrder(instrument, side, data[1], data[2])

				if err != nil {
					fmt.Println(err)
					continue
				}

				ticker.orders <- order
			}

			continue
		}
	}
}

func (ticker *Ticker) Writer() {
	for message := range ticker.messages {
		switch message.messageType {
		case SUBSCRIBE:
			data, err := json.Marshal(SubscriptionMessage{
				Type:       "subscribe",
				ProductIds: []string{message.data.(string)},
				Channels:   []string{"level2"},
			})

			if err != nil {
				fmt.Println(err)
				continue
			}

			ticker.conn.WriteMessage(websocket.TextMessage, data)
		}
	}
}

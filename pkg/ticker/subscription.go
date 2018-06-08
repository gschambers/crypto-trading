package ticker

const (
	SUBSCRIBE   = "subscribe"
	UNSUBSCRIBE = "unsubscribe"
)

type SubscriptionMessage struct {
	Type       string   `json:"type"`
	ProductIds []string `json:"product_ids"`
	Channels   []string `json:"channels"`
}

func (ticker *Ticker) GetSubscriptions() []string {
	keys := make([]string, 0, len(ticker.subscriptions))
	for key := range ticker.subscriptions {
		keys = append(keys, key)
	}
	return keys
}

func (ticker *Ticker) Subscribe(instrument string) {
	if _, ok := ticker.subscriptions[instrument]; !ok {
		ticker.messages <- Message{
			messageType: SUBSCRIBE,
			data:        instrument,
		}
	}
}

func (ticker *Ticker) Unsubscribe(instrument string) {
	if _, ok := ticker.subscriptions[instrument]; ok {
		ticker.messages <- Message{
			messageType: UNSUBSCRIBE,
			data:        instrument,
		}
	}
}

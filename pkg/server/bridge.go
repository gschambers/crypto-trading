package server

// Bridge between clients and price stream ticks
type Bridge struct {
	clients    map[*Client]map[Instrument]bool
	register   chan *Client
	unregister chan *Client
	ticks      chan Tick
}

func newBridge() *Bridge {
	return &Bridge{
		clients:    make(map[*Client]map[Instrument]bool),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		ticks:      make(chan Tick),
	}
}

func (bridge *Bridge) broadcast(tick Tick) {
	for client, instruments := range bridge.clients {
		if _, ok := instruments[tick.Instrument]; ok {
			client.outbox <- tick
		}
	}
}

func (bridge *Bridge) run() {
	for {
		select {
		case client := <-bridge.register:
			bridge.clients[client] = make(map[Instrument]bool)
		case client := <-bridge.unregister:
			if _, ok := bridge.clients[client]; ok {
				delete(bridge.clients, client)
			}
		case tick := <-bridge.ticks:
			bridge.broadcast(tick)
		}
	}
}

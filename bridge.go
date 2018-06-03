package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Bridge between clients and price stream ticks
type Bridge struct {
	clients    map[*Client]map[Instrument]bool
	register   chan *Client
	unregister chan *Client
	ticker     chan Tick
}

func newBridge() *Bridge {
	return &Bridge{
		clients:    make(map[*Client]map[Instrument]bool),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		ticker:     make(chan Tick),
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
	go startTickStream(bridge)

	for {
		select {
		case client := <-bridge.register:
			bridge.clients[client] = make(map[Instrument]bool)
		case client := <-bridge.unregister:
			if _, ok := bridge.clients[client]; ok {
				delete(bridge.clients, client)
			}
		case tick := <-bridge.ticker:
			bridge.broadcast(tick)
		}
	}
}

func getNextTick() (tick Tick, err interface{}) {
	tick = Tick{
		Instrument: Instrument{From: "BTC", To: "USD"},
		Price:      (rand.Int31n(5) + 100),
	}

	return tick, nil
}

func startTickStream(bridge *Bridge) {
	for {
		tick, err := getNextTick()

		if err != nil {
			fmt.Println("Error receiving tick:", err)
			continue
		}

		bridge.ticker <- tick

		time.Sleep(50 * time.Millisecond)
	}
}

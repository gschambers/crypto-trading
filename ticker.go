package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Ticker struct {
	bridge *Bridge
}

type Instrument struct {
	From string `json:"from"`
	To   string `json:"to"`
}

type Tick struct {
	Instrument Instrument `json:"instrument"`
	Price      int32      `json:"price"`
}

func newTicker(bridge *Bridge) *Ticker {
	return &Ticker{
		bridge: bridge,
	}
}

func (ticker *Ticker) getNextTick() (tick Tick, err interface{}) {
	tick = Tick{
		Instrument: Instrument{From: "BTC", To: "USD"},
		Price:      (rand.Int31n(5) + 100),
	}

	return tick, nil
}

func (ticker *Ticker) run() {
	for {
		tick, err := ticker.getNextTick()

		if err != nil {
			fmt.Println("Error receiving tick:", err)
			continue
		}

		ticker.bridge.ticks <- tick

		time.Sleep(50 * time.Millisecond)
	}
}

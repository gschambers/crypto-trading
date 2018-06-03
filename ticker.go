package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Instrument struct {
	From string `json:"from"`
	To   string `json:"to"`
}

type Tick struct {
	Instrument Instrument `json:"instrument"`
	Price      int32      `json:"price"`
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

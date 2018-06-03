package main

type Instrument struct {
	From string `json:"from"`
	To   string `json:"to"`
}

type Tick struct {
	Instrument Instrument `json:"instrument"`
	Price      int32      `json:"price"`
}

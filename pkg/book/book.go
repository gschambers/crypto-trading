package book

const (
	ASK = "ask"
	BID = "bid"
)

type Order struct {
	Instrument string
	Side       string
	Price      float64
	Size       float64
}

type LimitOrderBook struct {
	Ask map[float64]float64
	Bid map[float64]float64
}

type PriceSummary struct {
	Price  float64 `json:"price"`
	Volume float64 `json:"volume"`
}

type MarketSummary struct {
	Market string       `json:"market"`
	Ask    PriceSummary `json:"ask"`
	Bid    PriceSummary `json:"bid"`
}

type Market struct {
	AskPrice    float64
	BidPrice    float64
	LimitOrders *LimitOrderBook
}

type Book struct {
	markets map[string]*Market
	inbox   chan *Order

	Outbox chan *MarketSummary
}

func NewBook(inbox chan *Order) *Book {
	return &Book{
		markets: make(map[string]*Market),
		inbox:   inbox,

		Outbox: make(chan *MarketSummary),
	}
}

func (book *Book) Reader() {
	for order := range book.inbox {
		book.processOrder(order)
	}
}

func (book *Book) MarketSummary(instrument string) *MarketSummary {
	market := book.markets[instrument]

	return &MarketSummary{
		Market: instrument,

		Ask: PriceSummary{
			Price:  market.AskPrice,
			Volume: market.LimitOrders.Ask[market.AskPrice],
		},

		Bid: PriceSummary{
			Price:  market.BidPrice,
			Volume: market.LimitOrders.Bid[market.BidPrice],
		},
	}
}

package book

func (book *Book) processOrder(order *Order) {
	if _, ok := book.markets[order.Instrument]; !ok {
		book.markets[order.Instrument] = &Market{
			LimitOrders: &LimitOrderBook{
				Ask: make(map[float64]float64),
				Bid: make(map[float64]float64),
			},
		}
	}

	market := book.markets[order.Instrument]

	if order.Side == ASK {
		if order.Size == 0 {
			delete(market.LimitOrders.Ask, order.Price)

			if market.AskPrice == order.Price {
				book.updateMarketPrices(order.Instrument)
				book.Outbox <- book.MarketSummary(order.Instrument)
			}

			return
		}

		market.LimitOrders.Ask[order.Price] = order.Size

		if order.Size > 0 && (market.AskPrice == 0 || order.Price <= market.AskPrice) {
			market.AskPrice = order.Price
			book.Outbox <- book.MarketSummary(order.Instrument)
		}
	} else if order.Side == BID {
		if order.Size == 0 {
			delete(market.LimitOrders.Bid, order.Price)

			if market.BidPrice == order.Price {
				book.updateMarketPrices(order.Instrument)
				book.Outbox <- book.MarketSummary(order.Instrument)
			}

			return
		}

		market.LimitOrders.Bid[order.Price] = order.Size

		if order.Size > 0 && order.Price >= market.BidPrice {
			market.BidPrice = order.Price
			book.Outbox <- book.MarketSummary(order.Instrument)
		}
	}
}

func (book *Book) updateMarketPrices(instrument string) {
	market := book.markets[instrument]

	var askPrice float64
	var bidPrice float64

	for price := range market.LimitOrders.Ask {
		if askPrice == 0 || price < askPrice {
			askPrice = price
		}
	}

	for price := range market.LimitOrders.Bid {
		if price > bidPrice {
			bidPrice = price
		}
	}

	market.AskPrice = askPrice
	market.BidPrice = bidPrice
}

package ticker

import (
	"strconv"

	"gitlab.com/gschambers/crypto-trading/pkg/book"
)

func parseRawOrder(instrument string, side string, rawPrice string, rawSize string) (*book.Order, error) {
	price, err := strconv.ParseFloat(rawPrice, 64)

	if err != nil {
		return nil, err
	}

	size, err := strconv.ParseFloat(rawSize, 64)

	if err != nil {
		return nil, err
	}

	return &book.Order{
		Instrument: instrument,
		Side:       side,
		Price:      price,
		Size:       size,
	}, nil
}

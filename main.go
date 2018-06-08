package main

import (
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
	"gitlab.com/gschambers/crypto-trading/pkg/book"
	"gitlab.com/gschambers/crypto-trading/pkg/server"
	"gitlab.com/gschambers/crypto-trading/pkg/ticker"
)

func initGDAXTicker(pipe chan *book.Order) *ticker.Ticker {
	url := url.URL{Scheme: "wss", Host: "ws-feed.gdax.com", Path: "/"}
	ticker := ticker.NewTicker(url.String(), pipe)

	go ticker.Reader()
	go ticker.Writer()

	ticker.Subscribe("BTC-USD")

	return ticker
}

func initBook(pipe chan *book.Order) *book.Book {
	book := book.NewBook(pipe)
	go book.Reader()
	return book
}

func main() {
	pipe := make(chan *book.Order)
	book := initBook(pipe)
	initGDAXTicker(pipe)

	router := mux.NewRouter()
	router.HandleFunc("/stream", server.BookServer(book))

	fs := http.FileServer(http.Dir("web/static"))
	router.PathPrefix("/").Handler(fs)

	http.Handle("/", router)
	http.ListenAndServe(":3000", nil)
}

package main

import (
	"net/http"

	"gitlab.com/gschambers/crypto-trading/pkg/server"
)

func main() {
	http.Handle("/", server.NewRouter())
	http.ListenAndServe(":3000", nil)
}

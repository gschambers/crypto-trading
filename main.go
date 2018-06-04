package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"gitlab.com/gschambers/crypto-trading/pkg/server"
)

func main() {
	router := mux.NewRouter()
	router.Handle("/stream", server.StreamServer())

	fs := http.FileServer(http.Dir("web/static"))
	router.PathPrefix("/").Handler(fs)

	http.Handle("/", router)
	http.ListenAndServe(":3000", nil)
}

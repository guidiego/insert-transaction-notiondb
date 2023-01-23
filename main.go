package main;

import (
	"net/http"
	"github.com/guidiego/insert-transaction-notiondb/api"
)

func main() {
	http.HandleFunc("/api", handler.Handler)
	http.ListenAndServe(":3000", nil)
}
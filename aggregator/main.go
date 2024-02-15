package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"

	"github.com/umizu/tollcalc/types"
)

func main() {
	listenAddr := flag.String("listenAddr", ":3000", "http server port")
	flag.Parse()
	store := NewInMemStore()
	var (
		svc = NewInvoiceAggregator(store)
	)
	
	startHttpTransport(*listenAddr, svc)
}

func startHttpTransport(listenAddr string, svc Aggregator) {
	fmt.Println("aggregator: starting http server on port", listenAddr)
	http.HandleFunc("/aggregate", handleAggregate(svc))
	http.ListenAndServe(listenAddr, nil)
}

func handleAggregate(svc Aggregator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var dist types.Distance
		if err := json.NewDecoder(r.Body).Decode(&dist); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
}

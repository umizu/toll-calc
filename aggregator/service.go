package main

import (
	"fmt"

	"github.com/umizu/tollcalc/types"
)

type Aggregator interface {
	AggregateDistance(types.Distance) error
}

type Storer interface {
	Insert(types.Distance) error
}

type InvoiceAggregator struct {
	storer Storer
}

func NewInvoiceAggregator(storer Storer) *InvoiceAggregator {
	return &InvoiceAggregator{storer: storer}
}

func (ia *InvoiceAggregator) AggregateDistance(dist types.Distance) error {
	fmt.Println("processing and storing distance")
	return ia.storer.Insert(dist)
}

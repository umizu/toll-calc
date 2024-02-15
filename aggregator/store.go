package main

import "github.com/umizu/tollcalc/types"

type InMemStore struct {
}

func (ims *InMemStore) Insert(dist types.Distance) error {
	return nil
}

func NewInMemStore() *InMemStore {
	return &InMemStore{}
}

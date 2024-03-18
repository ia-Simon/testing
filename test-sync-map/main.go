package main

import "sync"

var sm sync.Map

type Sample struct {
	ID    string
	Value float64
	Color string
}

func main() {
	item1 := Sample{
		ID:    "sample1",
		Value: 2.3,
		Color: "black",
	}

	item2 := Sample{
		ID:    "sample1",
		Value: 2.3,
		Color: "black",
	}

	sm.Store("key1", item1)

	sm.Co
}

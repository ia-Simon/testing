package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Event struct {
	Name string    `json:"name"`
	Type EventType `json:"type"`
	At   time.Time `json:"at"`
}

type EventType string

// Enum values for EventType
const (
	EventTypeLoad  EventType = "LOAD"
	EventTypeStore EventType = "STORE"
	EventTypeReset EventType = "RESET"
)

func main() {
	jsonBody := []byte(`{
		"name": "Event1",
		"type": "LOAD",
		"at": "2022-10-17T10:45:22.123456-03:00"
	}`)

	var event Event
	err := json.Unmarshal(jsonBody, &event)
	if err != nil {
		panic(err)
	}

	fmt.Println(event.At, event.Type, event.Name)
}

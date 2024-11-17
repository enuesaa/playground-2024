package main

import (
	"errors"
	"fmt"
)

func NewEvent(g Greeter) (Event, error) {
	if g.Grumpy {
		return Event{}, errors.New("could not create event: event greeter is grumpy")
	}
	return Event{Greeter: g}, nil
}

type Event struct {
	Greeter Greeter
}

func (e *Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}

package main

import (
	"fmt"

	"go.uber.org/fx/fxevent"
)

func NewLogger() fxevent.Logger {
	return &Logger{}
}

type Logger struct {}

func (l *Logger) LogEvent(event fxevent.Event) {
	switch e := event.(type) {
	case *fxevent.Invoked:
		if e.Err != nil {
			fmt.Printf("Error: %s\n", e.Err.Error())
		}
	}
}
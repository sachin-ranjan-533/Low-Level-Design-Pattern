package main

import "fmt"

type ErrorHandler struct {
	nextHandler HandlerInterface
}

func (eh *ErrorHandler) Handle(level string, message string) {
	if level == "Error" {
		fmt.Printf("Error: %s\n", message)
	} else if eh.nextHandler != nil {
		eh.nextHandler.Handle(level, message)
	}
}

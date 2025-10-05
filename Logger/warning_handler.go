package main

import "fmt"

type WarningHandler struct {
	nextHandler HandlerInterface
}

func (wh *WarningHandler) Handle(level string, message string) {
	if level == "Warning" {
		fmt.Printf("Warning: %s\n", message)
	} else if wh.nextHandler != nil {
		wh.nextHandler.Handle(level, message)
	}
}

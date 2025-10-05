package main

import "fmt"

type InfoHandler struct {
	nextHandler HandlerInterface
}

func (ih *InfoHandler) Handle(level string, message string) {
	if level == "Info" {
		fmt.Printf("Info: %s\n", message)
	} else if ih.nextHandler != nil {
		ih.nextHandler.Handle(level, message)
	}
}

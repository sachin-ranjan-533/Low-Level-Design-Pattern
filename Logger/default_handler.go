package main

import "fmt"

type DefaultHandler struct {
	nextHandler HandlerInterface
}

func (dh *DefaultHandler) Handle(level string, message string) {
	fmt.Printf("%s: %s\n", level, message)
}

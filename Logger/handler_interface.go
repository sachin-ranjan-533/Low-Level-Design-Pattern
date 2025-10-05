package main

type HandlerInterface interface {
	Handle(level string, message string)
}

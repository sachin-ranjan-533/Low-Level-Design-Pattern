package main

func main() {
	handler := &InfoHandler{}
	warningHandler := &WarningHandler{}
	errorHandler := &ErrorHandler{}
	defaultHandler := &DefaultHandler{}

	handler.nextHandler = warningHandler
	warningHandler.nextHandler = errorHandler
	errorHandler.nextHandler = defaultHandler

	handler.Handle("Info", "This is an info message.")
	handler.Handle("Warning", "This is a warning message.")
	handler.Handle("Error", "This is an error message.")
	handler.Handle("Debug", "This is a debug message.")
}

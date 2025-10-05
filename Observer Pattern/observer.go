package main

// Observer interface defines methods that observers must implement
type Observer interface {
	update()    // Called when observable changes
	getId() int // Unique ID of the observer
}

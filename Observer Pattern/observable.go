package main

// Observable interface defines methods for managing observers
type Observable interface {
	add(observer Observer)    // Add an observer
	remove(observer Observer) // Remove an observer
	notifyAll()               // Notify all observers of a change
	setData(data int)         // Update the state
	getData() int             // Get the current state
}

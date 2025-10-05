package main

// Concrete Observable: iPhone stock
type iPhoneObservable struct {
	observers []Observer // List of observers
	count     int        // Current stock count
}

// Add an observer to the list
func (ipo *iPhoneObservable) add(o Observer) {
	ipo.observers = append(ipo.observers, o)
}

// Remove an observer from the list
func (ipo *iPhoneObservable) remove(o Observer) {
	for i, observer := range ipo.observers {
		if observer.getId() == o.getId() {
			ipo.observers = append(ipo.observers[:i], ipo.observers[i+1:]...)
			return
		}
	}
}

// Notify all observers about the current state
func (ipo *iPhoneObservable) notifyAll() {
	for _, observer := range ipo.observers {
		observer.update()
	}
}

// Update the state and notify observers
func (ipo *iPhoneObservable) setData(data int) {
	ipo.count = ipo.count + data
	ipo.notifyAll()
}

// Get the current state
func (ipo *iPhoneObservable) getData() int {
	return ipo.count
}

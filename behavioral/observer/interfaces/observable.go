package interfaces

// Observable ...
type Observable interface {
	// AddObserver subscribes observers
	AddObserver(name string, o Observer)
	// RemoveObserver unsubscribes observers
	RemoveObserver(name string)
	// NotifyObservers notifies when the state changes
	NotifyObservers()
}

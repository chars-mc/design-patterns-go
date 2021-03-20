package interfaces

// Observer ...
type Observer interface {
	// Notify sends the data to the observers
	Notify(data string)
}

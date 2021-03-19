package decorator

import "fmt"

// LogRegistry dontains a Decorator
// it's a decorator for the routes
type LogRegistry struct {
	Handler Decorator
}

// NewLogRegistry creates a new LogRegistry
func NewLogRegistry(handler Decorator) *LogRegistry {
	return &LogRegistry{handler}
}

// Process prints "request"
func (lr *LogRegistry) Process() error {
	defer fmt.Println("request log") // execute after sending the response
	return lr.Handler.Process()
}

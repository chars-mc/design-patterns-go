package decorator

import "fmt"

type HandlerSayHello struct{}

// NewHandlerSayHello creates a new HandlerSayHello
func NewHandlerSayHello() *HandlerSayHello {
	return &HandlerSayHello{}
}

// Process just prints "Hello"
func (h *HandlerSayHello) Process() error {
	fmt.Println("Hello")
	return nil
}

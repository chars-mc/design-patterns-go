package decorator

import "fmt"

type HandlerSayGoodbye struct{}

// NewHandlerSayGoodbye creates a new HandlerSayGoodbye
func NewHandlerSayGoodbye() *HandlerSayGoodbye {
	return &HandlerSayGoodbye{}
}

// Process just prints "Goodbye"
func (h *HandlerSayGoodbye) Process() error {
	fmt.Println("Goodbye")
	return nil
}

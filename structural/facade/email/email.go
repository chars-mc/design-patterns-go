package email

import "fmt"

// Email contains the email and message
type Email struct {
	to      string
	message string
}

// NewEmail creates and Email
func NewEmail() Email {
	return Email{}
}

// Send send the email
func (e *Email) Send(to, message string) {
	e.to = to
	e.message = message
	fmt.Printf("to %s: \tmessage: %s", e.to, e.message)
}

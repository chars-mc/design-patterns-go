package adapter

import "github.com/chars-mc/design-patterns-go/structural/adapter/paypal"

// PaypalAdapter adapts the payment method
type PaypalAdapter struct{}

// Transfer transfers an amount from one account to another
func (p *PaypalAdapter) Transfer(from, to string, amount int) {
	paypal.Send(from, to, amount)
}

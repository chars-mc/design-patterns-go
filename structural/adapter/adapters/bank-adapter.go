package adapter

import "github.com/chars-mc/design-patterns-go/structural/adapter/bank"

// BankAdapter adapts the payment method
type BankAdapter struct{}

// Transfer transfers an amount from one account to another
func (b *BankAdapter) Transfer(from, to string, amount int) {
	bank.Transfer(from, to, amount)
}

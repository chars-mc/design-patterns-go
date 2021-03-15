package adapter

// PaymentMethod allows to use any payment method
type PaymentMethod interface {
	Transfer(from, to string, amount int)
}

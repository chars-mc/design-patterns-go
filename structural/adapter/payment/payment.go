package payment

import adapter "github.com/chars-mc/design-patterns-go/structural/adapter/adapters"

// Payment contains the transaction information
type Payment struct {
	From   string
	To     string
	Amount int
	Method adapter.PaymentMethod
}

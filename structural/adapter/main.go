package main

import (
	adapter "github.com/chars-mc/design-patterns-go/structural/adapter/adapters"
	"github.com/chars-mc/design-patterns-go/structural/adapter/payment"
)

func main() {
	p1 := &payment.Payment{
		From:   "carlos@mail.com",
		To:     "manuel@mail.com",
		Amount: 3000,
		Method: &adapter.PaypalAdapter{},
	}

	p2 := &payment.Payment{
		From:   "245634573457",
		To:     "124365468970",
		Amount: 3500,
		Method: &adapter.BankAdapter{},
	}

	p1.Method.Transfer(p1.From, p1.To, p1.Amount)
	p2.Method.Transfer(p2.From, p2.To, p2.Amount)
}

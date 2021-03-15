package bank

import (
	"fmt"
)

// Transfer transfers an amount from one account to another
func Transfer(fromAccount, toAccount string, amount int) {
	fmt.Printf("Transfered %d  from %s to %s via bank transfer\n", amount, fromAccount, toAccount)
}

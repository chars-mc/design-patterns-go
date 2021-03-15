package paypal

import (
	"fmt"
)

// Send transfers an amount from one account to another
func Send(fromMobile, toMobile string, amount int) {
	fmt.Printf("Send %d  from %s to %s via paypal transfer\n", amount, fromMobile, toMobile)
}

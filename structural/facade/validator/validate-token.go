package validator

import (
	"errors"
	"fmt"
)

// ErrTokenNotValid user not logged in
var ErrTokenNotValid = errors.New("User not logged in")

// ValidatorToken contains the token
type ValidatorToken struct {
	token string
}

// NewValidatorToken creates a ValidatorToken
func NewValidatorToken(t string) ValidatorToken {
	return ValidatorToken{token: t}
}

// Validate checks if the token is valid
func (vt *ValidatorToken) Validate() error {
	if vt.token != "valid-token" {
		return ErrTokenNotValid
	}

	fmt.Println(" * valid token")
	return nil
}

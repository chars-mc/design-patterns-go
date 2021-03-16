package validator

import (
	"errors"
	"fmt"
)

// ErrPermissionNotValid user not authorized
var ErrPermissionNotValid = errors.New("User not authorized")

// ValidatorPermission contains the user id
type ValidatorPermission struct {
	userId string
}

// NewValidatorPermission creates a ValidatorPermission
func NewValidatorPermission(ID string) ValidatorPermission {
	return ValidatorPermission{ID}
}

// Validate checks if the user has permission
func (vp *ValidatorPermission) Validate() error {
	if vp.userId != "user-blog" {
		return ErrPermissionNotValid
	}

	fmt.Println(" * user has permission")
	return nil
}

package facade

import (
	"github.com/chars-mc/design-patterns-go/structural/facade/email"
	"github.com/chars-mc/design-patterns-go/structural/facade/storage"
	"github.com/chars-mc/design-patterns-go/structural/facade/validator"
)

// Facade contains all information about the email
// the structure must know all the subsystems that are going to be a property of the struct
type Facade struct {
	to                  string
	comment             string
	validatorToken      validator.ValidatorToken
	validatorPermission validator.ValidatorPermission
	storage             storage.Storage
	notificator         email.Email
}

// New creates a Facade
func New(to, comment, token, user string) Facade {
	return Facade{
		to:                  to,
		comment:             comment,
		validatorToken:      validator.NewValidatorToken(token),
		validatorPermission: validator.NewValidatorPermission(user),
		storage:             storage.NewStorage("mysql"),
		notificator:         email.NewEmail(),
	}
}

// Comment validates, saves and notify for a comment
func (f *Facade) Comment() error {
	if err := f.validatorToken.Validate(); err != nil {
		return err
	}

	if err := f.validatorPermission.Validate(); err != nil {
		return err
	}

	f.storage.Save(f.comment)
	f.Notify()

	return nil
}

// Notify sends an mail
func (f *Facade) Notify() {
	f.notificator.Send(f.to, f.comment)
}

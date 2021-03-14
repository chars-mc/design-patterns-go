package builder

// MessageBuilder message
type MessageBuilder interface {
	SetRecipient(recipient string) MessageBuilder
	SetMessage(message string) MessageBuilder
	Build() (*MessageRepresented, error)
}

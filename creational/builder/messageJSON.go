package builder

import "encoding/json"

// JSONMessageBuilder = ConcreteBuilder
type JSONMessageBuilder struct {
	message Message
}

// SetRecipient sets message.Recipient
func (b *JSONMessageBuilder) SetRecipient(recipient string) MessageBuilder {
	b.message.Recipient = recipient
	return b
}

// SetMessage sets message.Text
func (b *JSONMessageBuilder) SetMessage(text string) MessageBuilder {
	b.message.Text = text
	return b
}

// Build builds the object and represents it on JSON
func (b *JSONMessageBuilder) Build() (*MessageRepresented, error) {
	data, err := json.Marshal(b.message)
	if err != nil {
		return nil, err
	}
	return &MessageRepresented{data, "JSON"}, nil
}

package builder

import "encoding/xml"

// XMLMessageBuilder = ConcreteBuilder
type XMLMessageBuilder struct {
	message Message
}

// SetRecipient sets message.Recipient
func (b *XMLMessageBuilder) SetRecipient(recipient string) MessageBuilder {
	b.message.Recipient = recipient
	return b
}

// SetMessage sets message.Text
func (b *XMLMessageBuilder) SetMessage(text string) MessageBuilder {
	b.message.Text = text
	return b
}

// Build builds the object and represents it on XML
func (b *XMLMessageBuilder) Build() (*MessageRepresented, error) {
	data, err := xml.Marshal(b.message)
	if err != nil {
		return nil, err
	}
	return &MessageRepresented{data, "XML"}, nil
}

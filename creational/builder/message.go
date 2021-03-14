package builder

// Message = Product
type Message struct {
	Recipient string `json:"recipient" xml:"recipient"`
	Text      string `json:"text" xml:"text"`
}

// MessageRepresented = Object representation
type MessageRepresented struct {
	Body   []byte
	Format string
}

package event

// Message is a generic type agnostic of actual payload
type Message struct {
	// Id uniquely identifies a message from others
	// Typically it's a system generated ID that application has no control over it
	Id string

	// Payload is application-specific
	Payload *string

	// Additional metadata that is not part of Payload, support string only for simplicity
	Metadata map[string]*string
}

// IMessageHandler defines API for receipient of Message
type IMessageHandler interface {
	Handle(m *Message) error
}

// IQueueApi defines generic operations available by a queue
type IQueueApi interface {
	// Get n messages from queue, it should return no more than n messages
	GetMessages(n int) ([]Message, error)
	SendMessage(m *Message) error
	DeleteMessage(m *Message) error
}

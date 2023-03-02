package message

type Message struct {
	Message string
}

func New(m string) *Message {
	return &Message{
		Message: m,
	}
}

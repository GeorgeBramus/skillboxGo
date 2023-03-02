package storage

import "pro/pkg/message"

type Storage []message.Message

func New() Storage {
	var s []message.Message
	return s
}

func (s *Storage) Put(mes *message.Message) {
	*s = append(*s, *mes)
}

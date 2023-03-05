package main

import (
	"pro/pkg/message"
	"pro/pkg/storage"
)

func main() {
	sms := message.NewSMS("сообщение SMS")
	mms := message.NewMMS("сообщение MMS")

	s := storage.New()
	s.Put(sms)
	s.Put(mms)

	s.Get()
}

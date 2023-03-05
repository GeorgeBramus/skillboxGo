package message

type SMS struct {
	Text string
}

func NewSMS(text string) *SMS {
	return &SMS{
		Text: text,
	}
}

type MMS struct {
	Text string
}

func NewMMS(text string) *MMS {
	return &MMS{
		Text: text,
	}
}

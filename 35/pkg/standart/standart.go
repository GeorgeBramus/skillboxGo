package standart

type Codes struct {
	Code    string
	Country string
}

func New(code, country string) *Codes {
	return &Codes{
		Code:    code,
		Country: country,
	}
}

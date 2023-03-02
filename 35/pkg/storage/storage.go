package storage

import (
	"fmt"

	"code/pkg/standart"
)

type Codes map[string]*standart.Codes

func New() Codes {
	return make(map[string]*standart.Codes)
}

func (c Codes) Put(code *standart.Codes) {
	c[code.Code] = code
}

func (c Codes) Get(codeF string) (*standart.Codes, error) {
	code, ok := c[codeF]
	if !ok {
		return nil, fmt.Errorf("no such user")
	}
	return code, nil
}

// ***
// Эксперименты

type SMSData struct {
	User    string
	Message string
}

type SMS []
package surname

import (
	"fmt"
)

var (
	MaxLength      = 100
	ErrWrongLength = fmt.Errorf("Surname must be less or equal to %d chatacters", MaxLength)
)

type Surname string

func (n Surname) String() string {
	return string(n)
}

func New(surname string) (*Surname, error) {
	if len([]rune(surname)) > MaxLength {
		return nil, ErrWrongLength
	}
	n := Surname(surname)

	return &n, nil
}

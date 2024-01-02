package name

import (
	"fmt"
)

var (
	MaxLength      = 50
	ErrWrongLength = fmt.Errorf("Name must be less or equal to %d chatacters", MaxLength)
)

type Name string

func (n Name) String() string {
	return string(n)
}

func New(name string) (*Name, error) {
	if len([]rune(name)) > MaxLength {
		return nil, ErrWrongLength
	}
	n := Name(name)

	return &n, nil
}

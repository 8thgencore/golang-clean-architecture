package patronymic

import (
	"fmt"
)

var (
	MaxLength      = 50
	ErrWrongLength = fmt.Errorf("Name must be less or equal to %d chatacters", MaxLength)
)

type Patronymic string

func (n Patronymic) String() string {
	return string(n)
}

func New(patronymic string) (*Patronymic, error) {
	if len([]rune(patronymic)) > MaxLength {
		return nil, ErrWrongLength
	}
	n := Patronymic(patronymic)

	return &n, nil
}

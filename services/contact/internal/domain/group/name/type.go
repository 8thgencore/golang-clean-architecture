package name

import "fmt"

var (
	MaxLength      = 250
	ErrWrongLength = fmt.Errorf("name must be less than or equal to %d characters", MaxLength)
)

type Name struct {
	value string
}

func New(name string) (Name, error) {
	if len([]rune(name)) > MaxLength {
		return Name{}, ErrWrongLength
	}
	return Name{value: name}, nil
}

func (d Name) Value() string {
	return d.value
}

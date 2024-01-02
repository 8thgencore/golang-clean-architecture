package grpc

import (
	"app/services/contact/internal/useCase"

	contact "app/services/contact/internal/delivery/grpc/interface"
	"time"
)

type Delivery struct {
	contact.UnimplementedContactServiceServer
	ucContact useCase.Contact
	ucGroup   useCase.Group
}

type options struct {
	Timeout time.Duration
}

type Option func(*options)

func WithTimeout(timeout time.Duration) Option {
	return func(args *options) {
		args.Timeout = timeout
	}
}

func (d *Delivery) SetOptions(setters ...Option) {
	args := &options{
		Timeout: time.Second * 30,
	}

	for _, setter := range setters {
		setter(args)
	}
}

func New(ucContact useCase.Contact, ucGroup useCase.Group, setters ...Option) *Delivery {
	var d = &Delivery{
		ucContact: ucContact,
		ucGroup:   ucGroup,
	}

	d.SetOptions(setters...)

	return d
}

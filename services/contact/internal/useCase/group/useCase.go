package group

import (
	"app/services/contact/internal/useCase/adapters/storage"
	"time"
)

type UseCase struct {
	adapterStorage storage.Group
	options        options
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

func (uc *UseCase) SetOptions(setters ...Option) {
	args := &options{
		Timeout: time.Second * 30,
	}

	for _, setter := range setters {
		setter(args)
	}

	uc.options = *args
}

func New(storage storage.Storage, setters ...Option) *UseCase {
	var uc = &UseCase{
		adapterStorage: storage,
	}
	uc.SetOptions(setters...)

	return uc
}

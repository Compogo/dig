package dig

import (
	"fmt"

	"go.uber.org/dig"
)

type Decorator struct {
	dig *dig.Container
}

func NewDecorator(dig *dig.Container) *Decorator {
	return &Decorator{dig: dig}
}

func (container *Decorator) Provide(constructor interface{}) error {
	return container.dig.Provide(constructor)
}

func (container *Decorator) Provides(constructors ...interface{}) error {
	var errs error

	for _, constructor := range constructors {
		if err := container.Provide(constructor); err != nil {
			errs = fmt.Errorf("%w\n%w", errs, err)
		}
	}

	return errs
}

func (container *Decorator) Invoke(function interface{}) error {
	return container.dig.Invoke(function)
}

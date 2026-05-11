package dig

import (
	"fmt"

	"go.uber.org/dig"
)

type Container struct {
	dig *dig.Container
}

func NewContainer(dig *dig.Container) *Container {
	return &Container{dig: dig}
}

func (container *Container) Provide(constructor interface{}) error {
	return container.dig.Provide(constructor)
}

func (container *Container) Provides(constructors ...interface{}) error {
	var errs error

	for _, constructor := range constructors {
		if err := container.Provide(constructor); err != nil {
			errs = fmt.Errorf("%w\n%w", errs, err)
		}
	}

	return errs
}

func (container *Container) Invoke(function any) error {
	return container.dig.Invoke(function)
}

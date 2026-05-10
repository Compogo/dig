package dig

import (
	"github.com/Compogo/compogo"
	"github.com/Compogo/compogo/component"
	"github.com/Compogo/compogo/container"
	uberDig "go.uber.org/dig"
)

func WithDig() compogo.Option {
	dig := uberDig.New()
	digContainer := NewContainer(dig)

	return compogo.WithContainer(digContainer, &component.Component{
		Name: "container.Dig",
		Init: component.StepFunc(func(c container.Container) error {
			return c.Provides(
				func() *uberDig.Container { return dig },
				func() *Container { return digContainer },
				func(decorator *Container) container.Container { return decorator },
			)
		}),
	})
}

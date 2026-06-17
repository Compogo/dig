package dig

import (
	"github.com/Compogo/compogo"
	uberDig "go.uber.org/dig"
)

// WithDig возвращает опцию для подключения Dig-контейнера к приложению Compogo.
//
// Опция выполняет:
//   - Создание нового экземпляра Dig-контейнера
//   - Создание обёртки Container, реализующей интерфейс compogo.Container
//   - Регистрацию системного компонента "container.Dig", который:
//   - Регистрирует Dig-контейнер в DI (чтобы его можно было получить)
//   - Регистрирует обёртку Container
//   - Регистрирует обёртку как интерфейс compogo.Container
//
// Пример использования:
//
//	app := compogo.NewApp("myapp",
//	    compogo.WithConfigurator(configurator, configuratorCmp),
//	    compogo.WithLogger(logger, loggerCmp),
//	    compogo.WithCloser(closer, closerCmp),
//	    dig.WithDig(), // подключаем Dig
//	)
//
//	// Теперь все компоненты могут регистрировать зависимости через DI:
//	app.AddComponents(myComponent)
func WithDig() compogo.Option {
	dig := uberDig.New()
	digContainer := NewContainer(dig)

	return compogo.WithContainer(digContainer, &compogo.Component{
		Name: "container.Dig",
		Init: compogo.StepFunc(func(container compogo.Container) error {
			return container.Provides(
				func() *uberDig.Container { return dig },
				func() *Container { return digContainer },
				func(container *Container) compogo.Container { return container },
			)
		}),
	})
}

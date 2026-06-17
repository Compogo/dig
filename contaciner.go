package dig

import (
	"fmt"

	"go.uber.org/dig"
)

// Container реализует интерфейс compogo.Container на основе Uber Dig.
// Предоставляет удобную обёртку над dig.Container с поддержкой массовой
// регистрации конструкторов через метод Provides.
//
// Используется в Compogo как стандартная реализация DI-контейнера.
// Позволяет регистрировать зависимости и вызывать функции с внедрёнными
// аргументами.
//
// Пример:
//
//	container := NewContainer(dig.New())
//	container.Provide(func() *Config { return &Config{Port: 8080} })
//	container.Provide(func(cfg *Config) *http.Server {
//	    return &http.Server{Addr: fmt.Sprintf(":%d", cfg.Port)}
//	})
//	container.Invoke(func(srv *http.Server) {
//	    srv.ListenAndServe()
//	})
type Container struct {
	dig *dig.Container
}

// NewContainer создаёт новый Container на основе переданного Dig-контейнера.
// Обычно используется внутри WithDig() и не требует ручного вызова.
func NewContainer(dig *dig.Container) *Container {
	return &Container{dig: dig}
}

// Provide регистрирует конструктор в контейнере.
// Конструктор должен быть функцией, возвращающей один или несколько сервисов.
// Реализует интерфейс compogo.Container.
//
// Пример:
//
//	container.Provide(func() *Database { return &Database{DSN: "..."} })
func (container *Container) Provide(constructor interface{}) error {
	return container.dig.Provide(constructor)
}

// Provides регистрирует несколько конструкторов за один вызов.
// Если хотя бы один конструктор вызывает ошибку, она возвращается.
// Реализует интерфейс compogo.Container.
//
// Пример:
//
//	container.Provides(
//	    func() *Config { return &Config{} },
//	    func(cfg *Config) *Database { return NewDatabase(cfg) },
//	)
func (container *Container) Provides(constructors ...interface{}) error {
	var errs error

	for _, constructor := range constructors {
		if err := container.Provide(constructor); err != nil {
			errs = fmt.Errorf("%w\n%w", errs, err)
		}
	}

	return errs
}

// Invoke выполняет переданную функцию, внедряя в неё все необходимые
// зависимости из контейнера. Реализует интерфейс compogo.Container.
//
// Пример:
//
//	container.Invoke(func(db *Database) {
//	    db.Connect()
//	})
func (container *Container) Invoke(function any) error {
	return container.dig.Invoke(function)
}

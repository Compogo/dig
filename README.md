# Compogo Dig 🧩

**Compogo Dig** — это готовая интеграция [uber-go/dig](https://github.com/uber-go/dig) с фреймворком [Compogo](https://github.com/Compogo/compogo). Добавляется одной строкой и автоматически регистрирует контейнер в DI, делая его доступным для всех компонентов.

## 🚀 Установка

```bash
go get github.com/Compogo/dig
```

## 📦 Использование

```go
package main

import (
    "github.com/Compogo/compogo"
    "github.com/Compogo/dig"
    "github.com/Compogo/logrus"
    "github.com/Compogo/myapp/service"
)

func main() {
    app := compogo.NewApp("myapp",
        compogo.WithOsSignalCloser(),
        dig.WithDig(),                          // ← одна строка
        logrus.WithLogrus(),
        compogo.WithComponents(
            service.Component,
        ),
    )

    if err := app.Serve(); err != nil {
        panic(err)
    }
}
```

## ✨ Возможности

### 🔧 Три способа получить контейнер

В любом компоненте можно запросить:

```go
// 1. Как интерфейс container.Container
type Service struct {
    container container.Container
}

// 2. Как *dig.Decorator (с методами Provide/Provides/Invoke)
type Service struct {
    dig *dig.Decorator
}

// 3. Как чистый *dig.Container
type Service struct {
    rawDig *uberDig.Container
}
```

### 📦 Пакетная регистрация

```go
container.Provides(
    NewConfig,
    NewService,
    NewRepository,
) // возвращает одну ошибку со всеми проблемами
```

## 🧪 Пример компонента, использующего контейнер

```go
var ServiceComponent = &component.Component{
    Init: component.StepFunc(func(c container.Container) error {
        return c.Provides(
            NewService,
            func() *Repository { return NewRepository() },
        )
    }),
    Run: component.StepFunc(func(c container.Container) error {
        return c.Invoke(func(s *Service) {
            s.Start()
        })
    }),
}
```

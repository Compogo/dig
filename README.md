# Compogo Dig

[![Go Reference](https://pkg.go.dev/badge/github.com/Compogo/dig.svg)](https://pkg.go.dev/github.com/Compogo/dig)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Адаптер [Uber Dig](https://github.com/uber-go/dig) для фреймворка [Compogo](https://github.com/Compogo/compogo).

Реализует интерфейс `compogo.Container` через Dig, позволяя использовать мощный DI-контейнер для управления зависимостями в приложениях на Compogo.

## Установка

```shell
go get github.com/Compogo/compogo-dig
```

## Использование

Подключение Dig-контейнера к приложению Compogo:

```go
package main

import (
    "github.com/Compogo/compogo"
    "github.com/Compogo/dig"
)

func main() {
    // Создание приложения с Dig
    app := compogo.NewApp("myapp",
        // Подключаем Dig как DI-контейнер
        dig.WithDig(),
        
        // Другие опции
        compogo.WithConfigurator(configurator, configuratorCmp),
        compogo.WithLogger(logger, loggerCmp),
        compogo.WithOsSignalCloser(),
    )

    // Добавление компонентов
    app.AddComponents(
        HTTPServerComponent,
        DatabaseComponent,
    )

    // Запуск
    if err := app.Serve(); err != nil {
        panic(err)
    }
}
```

## Что происходит под капотом

При вызове `dig.WithDig()` происходит следующее:

* Создаётся экземпляр dig.Container
* Создаётся обёртка Container, реализующая compogo.Container
* Регистрируется системный компонент container.Dig, который:
* * Регистрирует оригинальный `*dig.Container`
* * Регистрирует обёртку `*Container`
* * Регистрирует обёртку как `compogo.Container`

После этого все компоненты могут использовать container.Provide() и container.Invoke() через Dig.

## Методы

### Container

Реализует интерфейс `compogo.Container`:

```go
type Container struct {
    dig *dig.Container
}

// Provide регистрирует конструктор
func (c *Container) Provide(constructor interface{}) error

// Provides регистрирует несколько конструкторов
func (c *Container) Provides(constructors ...interface{}) error

// Invoke вызывает функцию с внедрёнными зависимостями
func (c *Container) Invoke(function interface{}) error
```

### WithDig

```go
func WithDig() compogo.Option
```

Возвращает опцию для подключения Dig к приложению Compogo.

## Зависимости

* [Compogo](https://github.com/Compogo/compogo) — основной фреймворк
* [Uber Dig](https://github.com/uber-go/dig) — DI-контейнер

## Лицензия

```plantuml
MIT License

Copyright (c) 2026 Compogo

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

```

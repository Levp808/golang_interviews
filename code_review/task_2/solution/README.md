## Реализовать паттерн "Functional Options" (Функциональные опции).

Паттерн «Функциональные опции» в Go — это идиоматический способ конфигурирования структур и функций, особенно полезный, когда требуется множество необязательных параметров. Он позволяет избежать перегрузки конструкторов и делает код более читаемым и расширяемым.

### 🔧 Суть паттерна
Вместо передачи большого количества аргументов в конструктор, вы определяете функции-опции, каждая из которых изменяет конкретное поле структуры. Затем в конструкторе применяете эти функции к объекту.

### Вариант решения:

```
package main

import (
  "fmt"
  "time"
)


type Config struct {
  Host    string
  Port    int
  UseSSL  bool
  Timeout time.Duration
}

type ConfigOption func(*Config)

func WithHost(host string) *Config{
    return func(cfg *Config){
        cfg.Host = host
    }
}

func WithPort(port int) *Config{
    return func(cfg *Config){
        cfg.Port = port
    }
}

func WithUseSSL(ssl bool){
    return func(cfg *Config){
        cfg.UseSSL = ssl
    }
}

func WithTimeout(timeout time.Duration){
    return func(cfg *Config){
        cfg.Timeout = timeout
    }
}



func NewConfig(opts ...ConfigOption) *Config{
    const (
        defaultHost = "example.com"
        defaultPort = 8080
        defaultUseSSL = false
        defaultTimeout = 1*time.Second
    )
    
    cfg := &Config{
        Host: defaultHost,
        Port: defaultPort,
        UseSSL: defaultUseSSL,
        Timeout: defaultTimeout,
    }

    for _, opt := range opts{
        opt(cfg)
    }

    return cfg
}

```
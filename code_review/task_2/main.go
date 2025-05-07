package main

import (
	"time"
)

type Config struct {
	Host    string
	Port    int
	UseSSL  bool
	Timeout time.Duration
}

// Реализовать паттерн "Functional Options" (Функциональные опции).

package main

import (
	"log"
	"time"
)

func LongRunningFunction() int {
	// Имитация долгой работы
	time.Sleep(3 * time.Second)
	return 42
}

func RunWithTimeout(f func() int, timeout time.Duration) (int, error) {
	//TODO: реализовать функцию с таймаутом
	return 0, nil
}

func main() {
	startTime := time.Now()
	result, err := RunWithTimeout(LongRunningFunction, 3*time.Second)
	executionTime := time.Since(startTime)

	if err != nil {
		// Выводим ошибку
		log.Println("Error:", err)
	} else {
		// Выводим результат
		log.Println("Result:", result)
	}

	// Выводим время выполнения
	log.Println("Execution Time:", executionTime)
}

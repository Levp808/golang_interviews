package main

import (
	"fmt"
	"sync"
)

const maxThreads = 3

type Task struct {
	Text string
}

type Result struct {
	ThreadNumber int
	Reverted     string
}

func Revert(str string) string {
	runes := []rune(str)
	for i := 0; i < len(runes)/2; i++ {
		runes[i], runes[len(runes)-1-i] = runes[len(runes)-1-i], runes[i]
	}
	return string(runes)
}

func main() {
	texts := []string{"Hello", "qwerty", "Golang", "platypus", "тест", "level", "generics"}
	taskCh := make(chan Task, maxThreads)
	resultCh := make(chan Result)
	wg := sync.WaitGroup{}

	// Запускаем воркеры
	wg.Add(maxThreads)
	for i := 0; i < maxThreads; i++ {
		go func(threadID int) {
			defer wg.Done()
			for task := range taskCh {
				reverted := Revert(task.Text)
				resultCh <- Result{
					ThreadNumber: threadID + 1,
					Reverted:     reverted,
				}
			}
		}(i)
	}

	// Отправляем задачи
	go func() {
		for _, text := range texts {
			taskCh <- Task{Text: text}
		}
		close(taskCh)
	}()

	// Ждём завершения воркеров и закрываем канал результатов
	go func() {
		wg.Wait()
		close(resultCh)
	}()

	// Печатаем результаты
	lineNumber := 1
	for res := range resultCh {
		fmt.Printf("line %d, thread %d: \"%s\"\n", lineNumber, res.ThreadNumber, res.Reverted)
		lineNumber++
	}
	fmt.Println("done")
}

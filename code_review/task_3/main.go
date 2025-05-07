package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Task struct {
	ID int
}

func (t Task) Execute(ctx context.Context) error {
	delay := time.Duration(rand.Intn(5)+1) * time.Second

	select {
	case <-time.After(delay):
		fmt.Printf("✅ Task %d completed in %v\n", t.ID, delay)
		return nil
	case <-ctx.Done():
		fmt.Printf("⛔️ Task %d canceled (timeout)\n", t.ID)
		return ctx.Err()
	}
}

func taskGenerator(total int) <-chan Task {
	out := make(chan Task)

	go func() {
		defer close(out)
		for i := 1; i <= total; i++ {
			out <- Task{ID: i}
		}
	}()

	return out
}

// worker получает задачи из канала и обрабатывает их
func worker(ctx context.Context, tasks <-chan Task, wg *sync.WaitGroup, results chan<- error) {
	defer wg.Done()

	for task := range tasks {
		results <- task.Execute(ctx)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	const (
		totalTasks  = 20
		workerCount = 5
		timeout     = 10 * time.Second
	)

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	tasks := taskGenerator(totalTasks)

	taskChan := make(chan Task)
	results := make(chan error, totalTasks)

	var wg sync.WaitGroup

	wg.Add(workerCount)

	for i := 0; i < workerCount; i++ {
		go worker(ctx, taskChan, &wg, results)
	}

	go func() {
		defer close(taskChan)
		for task := range tasks {
			select {
			case <-ctx.Done():
				return
			case taskChan <- task:
			}
		}
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	var success, failed int

	for err := range results {
		if err == nil {
			success++
		} else {
			failed++
		}
	}

	fmt.Println("---------------")
	fmt.Printf("✅ Success: %d\n", success)
	fmt.Printf("⛔️ Failed (timeout): %d\n", failed)
}

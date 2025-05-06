package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Processor struct {
	results map[int]int
}

func NewProcessor() *Processor {
	return &Processor{
		results: make(map[int]int),
	}
}

func (p *Processor) Process(ctx context.Context, data []int) error {
	var wg sync.WaitGroup
	errCh := make(chan error, 1)

	for _, v := range data {
		wg.Add(1)
		go func() {
			select {
			case <-time.After(time.Duration(v) * 100 * time.Millisecond):
				p.results[v] = v * 2
			case <-ctx.Done():
				select {
				case errCh <- fmt.Errorf("operation for %d canceled: %v", v, ctx.Err()):
				default:
				}

				wg.Done()
				return
			}
		}()
	}

	wg.Wait()

	for err := range errCh {
		return err
	}
	return nil
}

func (p *Processor) GetResults() map[int]int {
	return p.results
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	processor := NewProcessor()

	if err := processor.Process(ctx, data); err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("Results:", processor.GetResults())
}

// Провести код ревью

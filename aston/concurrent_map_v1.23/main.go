package main

import (
	"fmt"
	"sync"
)

func main() {

	data := make(map[int]struct{})

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			data[i] = struct{}{}
		}
	}()

	go func() {
		defer wg.Done()
		for i := 10; i < 20; i++ {
			data[i] = struct{}{}
		}
	}()

	wg.Wait()

	fmt.Println(data)
}

// Устранить потенциальные проблемы

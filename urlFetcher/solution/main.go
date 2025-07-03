package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
)

// Написать функцию, которая запрашивает URL из списка и в случае положительного кода 200 выводит
// в stdout в отдельной строке url: <url>, code: <statusCode>
// В случае ошибки выводит в отдельной строке url: <url>, code: <statusCode>
// Функция должна завершаться при отмене контекста.
// Доп. задание: реализовать ограничение количества одновременно запущенных горутин.

func fetchParallel(ctx context.Context, urls []string) {
	const maxWorkers int = 3
	wg := sync.WaitGroup{}
	ch := make(chan string, len(urls))
	for _, url := range urls {
		ch <- url
	}
	close(ch)

	for i := 0; i < maxWorkers; i++ {
		wg.Add(1)
		go func(ch <-chan string) {
			defer wg.Done()
			for url := range ch {
				select {
				case <-ctx.Done():
					return
				default:
					req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
					if err != nil {
						fmt.Printf("url: %s, code: %d\n", url, http.StatusBadRequest)
						continue
					}
					resp, err := http.DefaultClient.Do(req)
					if err != nil {
						fmt.Printf("url: %s, code: %d\n", url, http.StatusBadRequest)
						continue
					}
					defer resp.Body.Close()
					fmt.Printf("url: %s, code: %d\n", url, resp.StatusCode)
				}
			}
		}(ch)
	}

	wg.Wait()
}

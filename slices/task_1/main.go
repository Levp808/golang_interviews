package main

import "fmt"

func main() {
	a := []int{
		1, 2, 3, 4, 5,
	}

	for _, i := range a {
		go func() {
			fmt.Println(i)
		}()
	}
}

// Что будет выведено?

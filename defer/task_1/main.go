package main

import "fmt"

func main() {
	fmt.Println(someFunc())
}

func someFunc() int {
	i := 6
	defer fmt.Println(i)
	defer func() {
		i++
		fmt.Println(i)
	}()

	i++
	return i
}

// Что будет выведено?

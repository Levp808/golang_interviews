package main

import (
	"fmt"
)

func ConcateSlices(slice1 []int, slice2 []int) []int {
	//TODO
}

func main() {
	s1 := []int{1, 2, 2, 3, 4, 5}
	s2 := []int{2, 3, 3, 6}

	fmt.Println(ConcateSlices(s1, s2))
}

// На вход подаются два неупорядоченных слайса любой длины
// Надо написать функцию, которая возвращает их пересечение

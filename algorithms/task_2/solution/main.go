package main

import (
	"fmt"
)

func ConcateSlices(slice1 []int, slice2 []int) []int {
	res := []int{}

	nums := make(map[int](bool))

	for _, v := range slice1 {
		nums[v] = true
	}

	for _, v := range slice2 {
		if nums[v] == true {
			res = append(res, v)
			nums[v] = false
		}
	}
	return res
}

func main() {
	s1 := []int{1, 2, 2, 3, 4, 5}
	s2 := []int{2, 3, 3, 6}

	fmt.Println(ConcateSlices(s1, s2))
}

// На вход подаются два неупорядоченных слайса любой длины
// Надо написать функцию, которая возвращает их пересечение

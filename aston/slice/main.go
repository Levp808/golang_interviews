package main

import (
	"fmt"
)

func main() {
	nums := make([]int, 1, 2)
	fmt.Println(nums) // <- what's the output?

	appendSlice(nums, 1024)
	fmt.Println(nums) // <- what's the output?

	mutateSlice(nums, 1, 512)
	fmt.Println(nums) // <- what's the output?

} 

func appendSlice(sl []int, val int) {
	sl = append(sl, val)
}

func mutateSlice(sl []int, idx, val int) {
	sl[idx] = val
}

// 1. Что выведет программа?
// 2. Как можно исправить второй вывод? 
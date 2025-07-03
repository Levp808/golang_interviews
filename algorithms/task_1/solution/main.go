package main

import "fmt"

func main() {
	max_count_substr("aafbaaaaffc") // a:4 b:1 f:2 c:1
}

func max_count_substr(str string) {
	storage := make(map[rune]int)
	curr_char := rune(str[0]) // byte
	curr_count := 0
	for _, v := range str { // rune
		if v == curr_char {
			curr_count++
			continue
		} else {
			if storage[curr_char] < curr_count {
				storage[curr_char] = curr_count
			}
			curr_char = v
			curr_count = 1

		}
	}
	if storage[curr_char] < curr_count {
		storage[curr_char] = curr_count
	}
	fmt.Println(storage)
}

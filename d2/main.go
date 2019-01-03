package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

func product(l []int64) []int64 {
	var (
		p      int64 = 1
		result       = make([]int64, 0, len(l))
	)
	for _, v := range l {
		p *= v
	}

	for _, v := range l {
		result = append(result, p/v)
	}
	return result
}

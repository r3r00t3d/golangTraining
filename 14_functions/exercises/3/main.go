package main

import "fmt"

func max(numbers ...int) int {
	var largest int
	for _, v := range numbers {
		if v > largest {
			largest = v
		}
	}
	return largest
}

func main() {
	largest := max(1, 10, 25, 654, 23, 1)
	fmt.Println(largest)
}

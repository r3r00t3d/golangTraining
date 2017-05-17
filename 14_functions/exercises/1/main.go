package main

import "fmt"

func half(num int) (int, bool) {
	return num / 2, num%2 == 0
}

func main() {
	h, even := half(5)
	fmt.Println(h, even)
}

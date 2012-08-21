package main

import (
	"fmt"
)

func main() {
	// Testing
	// fmt.Println(Problem1(0, 0) == 0)
	// fmt.Println(Problem1(0, 10) == 23)

	fmt.Println(Problem1(0, 1000))
}

func Problem1(start, finish int) int {
	var sum int
	for ; start < finish; start++ {
		if start%3 == 0 || start%5 == 0 {
			sum += start
		}
	}
	return sum
}

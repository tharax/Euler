package main

import "fmt"

var m map[int]int

func main() {
	Problem2(4000000)
}

func Problem2(maximum int) {
	m = make(map[int]int)
	var total int
	fmt.Println("maximum =", maximum)
	for i := 0; fib(i) < maximum; i++ {
		m[i] = fib(i)
	}
	for i := 0; fib(i) < maximum; i++ {
		if m[i]%2 == 0 {
			fmt.Println(m[i])
			total = total + m[i]
		}
	}
	fmt.Println("total =", total)
}

func fib(x int) int {
	if x < 1 {
		return 1
	}
	return m[x-1] + m[x-2]
}

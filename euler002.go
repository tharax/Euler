package main

import "fmt"

func main() {
	Problem2(4000000)
}

func Problem2(maximum int) {
	var total int
	fmt.Println("maximum =", maximum)
	m := createFibSequence(maximum)
	for i := 0; fib(i, m) < maximum; i++ {
		if m[i]%2 == 0 {
			fmt.Println(m[i])
			total += m[i]
		}
	}
	fmt.Println("total =", total)
}

func createFibSequence(maximum int) map[int]int {
	m := make(map[int]int)
	for i := 0; fib(i, m) < maximum; i++ {
		m[i] = fib(i, m)
	}
	return m
}

func fib(x int, m map[int]int) int {
	if x < 1 {
		return 1
	}
	return m[x-1] + m[x-2]
}

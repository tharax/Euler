package main

import "fmt"

var m map[int]uint64

func main() {
	m = make(map[int]uint64)
	var maximum, total uint64
	maximum = 4000000

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

func fib(x int) uint64 {
	if x < 1 {
		return 1
	}
	return m[x-1] + m[x-2]
}

// This section not needed - it's my first attempt at Fibonacci in go, but it's got a Big-Oh of 2^n
// I wrote this just to check that I can in Go, and because it's been awhile since i wrote something with recursion

func printRecursiveFib(start, finish int) {
	var x int
	for x = start; x <= finish; x++ {
		fmt.Println("fib of", x, "is", recursiveFib(x))
	}
}

func recursiveFib(x int) int {
	if x < 0 {
		return 0
	}
	if x <= 1 {
		return 1
	}
	return recursiveFib(x-1) + recursiveFib(x-2)
}

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

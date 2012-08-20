package main

import "fmt"

func main() {
	var x, y, z int
	y = 1000

	fmt.Println(x, y, z)

	for ; x < y; x++ {
		if x%3 == 0 {
			z = z + x
		} else if x%5 == 0 {
			z = z + x
		}
	}

	fmt.Println(x, y, z)
}

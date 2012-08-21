// A Pythagorean triplet is a set of three natural numbers, a  b  c, for which,

// a2 + b2 = c2
// For example, 32 + 42 = 9 + 16 = 25 = 52.

// There exists exactly one Pythagorean triplet for which a + b + c = 1000.
// Find the product abc.

package main

import "fmt"

func main() {

	total := 1000
	a, b, c := 1, 1, total
	long, short, hypotenuse := a*a, b*b, c*c
	fmt.Println("start: ", a, b, c, long, short, hypotenuse)

L:
	for ; a <= total; a++ {
		for b = 1; a+b <= total; b++ {
			c = total - a - b
			long, short, hypotenuse = a*a, b*b, c*c
			if long+short == hypotenuse {
				fmt.Println("start: ", a, b, c, long, short, hypotenuse)
				break L
			}
		}
	}

	fmt.Println("check: a + b + c =", a+b+c)
	fmt.Println("check: long + short =", long+short, "hypotenuse = ", hypotenuse)
	fmt.Println("answer: a * b * c =", a*b*c)
}

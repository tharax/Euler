/*The sum of the squares of the first ten natural numbers is,

12 + 22 + ... + 102 = 385
The square of the sum of the first ten natural numbers is,

(1 + 2 + ... + 10)2 = 552 = 3025
Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025  385 = 2640.

Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.

*/
package main

import "fmt"

func main() {
	var numbers int
	var sumOfSquares, squareOfSums int

	numbers = 100

	for i := 0; i <= numbers; i++ {
		sumOfSquares = sumOfSquares + i*i
	}
	fmt.Println("sumOfSquares", sumOfSquares)
	for i := 0; i <= numbers; i++ {
		squareOfSums = squareOfSums + i
	}
	squareOfSums = squareOfSums * squareOfSums
	fmt.Println("squareOfSums", squareOfSums)

	fmt.Println("difference =", squareOfSums-sumOfSquares)
}

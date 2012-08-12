/*

2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
*/

//I feel I should mention that after about 10 seconds of thinking on this, I realise it's just
// "multiple all the prime numbers below this number" - because anything divisble by a prime (eg 2,3)
// can be divided by a non-prime as well (eg 4, 6, 9).
// actually now i am not sure - i guess 15 is divisble by 3 but not by 9. so i may have to rethink this as i go.
// but i am pretty sure the key lies in multiplying numbers that are primes. ( 2vs4 is still a valid argument, although it is
// a square. (2^2) so possible that its just a special case.

//turns out i needed a different approach entirely.

package main

import "fmt"

var rangeMin, rangeMax, evenlyDivisble int = 1, 10, 1

func main() {

	// for every number in the range
	for i := rangeMin; i <= rangeMax; i++ {
		// if the current total not divisble by the number
		if evenlyDivisble%i != 0 {
			// multiply by lowest number needed
			evenlyDivisble = evenlyDivisble * lcm(evenlyDivisble, i)
		}
		fmt.Println(evenlyDivisble, "/", i, "=", evenlyDivisble/i)
	}
}

//lowest common multiple - this probably needs to be renamed
func lcm(x, y int) int {
	for i := rangeMin; i <= rangeMax; i++ {
		if x*i%y == 0 {
			return i
		}
	}
	return y
}

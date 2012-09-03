// Find the number of non-empty subsets of {1^1, 2^2, 3^3,..., 250250^250250},
// the sum of whose elements is divisible by 250. Enter the rightmost 16 digits as your answer.

// 25^25 = 8.8817842 × 10^34
// It's clear there is no simple number-crunching method.
// There are 250250 Numbers in the largest set.
// There are 1 Numbers in the 250250 possible smallest sets.

// I actually wonder if this is a prime number question - the primes of 250 are 2, 5, 5, 5.
// If the number is divisible by those primes, then it should be ok.

// Ok so this is the case - I'm hacking this code from previous questions I've answered (3 and 7)
// The idea is to check if a number is divisble by any of the primes of 250, and if it is, multiply it.

// the theory is, for any number in the set, it has a set number of primes as it's Lowest common
// divisors. and for each time you multiply it, you multiply the set of prime divisors.
// so for the specific example of 250, i just have to figure out if n is divisble by 2 or 5, 
// and if that is the case, is n > 3 (because there are 3 "5"s in the 250.)

// this is going to give me all the numbers in the super set that % 250 == 0

// ahh you say, the problem is bigger than this - perhaps, but this is my starting block. 

// 25 july - ok so the trick is possibly something like "multiply by the number, divide by 250, 
// repeat y times (x^y)
// 

// 21 August - Has it really been a month?

// 27 August - I think this could be a breadth/depth first problem. you have the root level - 1 to n, then every level under each of those is the same until n = 1 where you return

// 3 September - 250 Groupings

// Length | Count | Combinations
//      1 |     1 | 0
//      2 |     2 | 0, 1+1
//      3 |     4 | 0, 2+2+2, 2+1, 1+1+1
//      4 |     7 | 0, 3+3+3+3, 3+3+2, 3+1, 2+2, 2+1+1, 1+1+1+1
//      5 |    15 | 0, 4+4+4+4+4, 4+4+4+3, 4+4+2, 4+1, 4+3+3, 4+2+2+2, 3+3+3+1, 3+3+3+3+3, 3+2, 3+1+1, 2+2+2+2+2, 2+2+1, 2+1+1+1, 1+1+1+1+1
//      6 |    19 | 0, 5+5+5+5+5+5, 5+5+5+5+4,5+5+5+5+4,5+5+5+3, 5+3+2+2, 5+5+2, 5+1,4+4+3+1, 4+4+4, 4+1+1, 4+2, 3+3, 2+2+2, 3+2+1, 1+1+1, 2+2+1+1, 2+1+1+1+1, 1+1+1+1+1+1
//      7 |    59 |

// My expectation is that the algorithm is the sum of the previous amounts:
// combos 0 = 1
// combos n = sum of combos n-1 to 0

// So far this problem has a solution for the first two numbers.

package main

import (
	"fmt"
	"math"
	// "time"
)

func main() {
	Problem250(250250, 250)
}

func Problem250(number, divisor int) {

	// Calculate the remainder for each number in the set
	// startTime := time.Now()
	remainders := arrayOfRemainders(number)
	// endTime := time.Now()
	// fmt.Printf("Creating remainder array took %v to run.\n\n", endTime.Sub(startTime))

	// Count each remainder, returns a smaller set (of length 250)
	// startTime = time.Now()
	countForEachRemainder := countEachRemainder(remainders)
	// fmt.Println(countForEachRemainder)
	// endTime = time.Now()
	// fmt.Printf("Counting each remainder took %v to run.\n\n", endTime.Sub(startTime))

	//Returns all the combinations that add up to 
	// startTime = time.Now()
	fmt.Println(countForEachRemainder[0:2])
	combos := combinations(divisor, countForEachRemainder[0:2])
	fmt.Println(combos)
	// endTime = time.Now()
	// fmt.Printf("Creating every combination took %v to run.\n\n", endTime.Sub(startTime))

}

func combinations(x int, slice []int) map[int]uint64 {
	s := createFibSequence(x)
	var number uint64
	number = uint64(slice[0]) * uint64(slice[1]) * uint64(slice[1]-1)
	fmt.Println(slice[0], slice[1], (slice[1] - 1), "number=", number)
	// fmt.Println("slice length = ", len(s))
	// fmt.Println("maximum uint64 = 18446744073709551615")
	// fmt.Println("highest prime  =", s[len(s)-1])

	var sum uint64
	for i, _ := range s {
		sum += s[i]
		// fmt.Println("sum=", sum)
	}
	// fmt.Println("sum of all the=", sum)

	return s
}

func makePrimeSlice(lowerBottom, upperBottom uint64) []uint64 {
	isPrime, primeSlice := true, make([]uint64, 0)
	// K:
	for i := lowerBottom; i <= upperBottom; i++ {
		isPrime = true
	L:
		for j := i - lowerBottom; j > lowerBottom; j-- {
			if i%j == 0 {
				isPrime = false
				break L
			}
		}
		if isPrime && i > 1 {
			primeSlice = append(primeSlice, i)
		}

		// if len(primeSlice) == 250 {
		// 	fmt.Println(primeSlice[len(primeSlice)-1])
		// 	break K
		// }

	}
	// fmt.Println("Slice created")
	return primeSlice
}

// func createComboSequence(maximum int) map[int]uint64 {
// 	m := make(map[int]uint64, maximum)
// 	for i := 0; i < maximum; i++ {
// 		m[i] = fib(i, m)
// 	}
// 	return m
// }

func createFibSequence(maximum int) map[int]uint64 {
	m := make(map[int]uint64, maximum)
	for i := 0; i < maximum; i++ {
		m[i] = fib(i, m)
	}
	return m
}

func fib(x int, m map[int]uint64) uint64 {
	if x < 1 {
		return 1
	}
	return m[x-1] + m[x-2]
}

// Returns the remainder of (x^x / 250)  
func arrayOfRemainders(x int) []int {
	remainderSlice := make([]int, 0)
	for i := 1; i <= x; i++ {
		steps, remainder := make([]int, 0), i
	J:
		for j := 1; j < i; j++ {
			remainder = remainder * i % 250
			// steps = append(steps, remainder)
			for k := 0; k < len(steps); k++ {
				if remainder == steps[k] {
					remainder = steps[i%len(steps)]
					// fmt.Println("i", i, "j", j, "k", k,  "remainder", remainder)
					break J
				}
			}
			steps = append(steps, remainder)
		}
		remainderSlice = append(remainderSlice, remainder)
	}
	return remainderSlice
}

// Returns the count of each remainder
func countEachRemainder(array []int) []int {
	count := make([]int, 250)
	for i := 0; i < len(array); i++ {
		count[int(math.Abs(float64(array[i])))]++
	}
	return count
}

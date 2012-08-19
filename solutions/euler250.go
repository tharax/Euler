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

package main

import (
	"fmt"
	"math"
	"time"
)

func main() {

	// divisors := nonRepeatingDivisorsOf(250)
	// fmt.Println(divisors)

	// allDivisors := allDivisorsOf(250)
	// fmt.Println(allDivisors)

	// numbers := make([]int, 0)

	// x, y := 250250, 250250
	// for i := 1; i <= 250250; i++ {
	// 	thisIsNotADivisor := true
	// 	remainder := i % 250
	// 	//fmt.Print(remainder, " ")
	// 	for j := 0; j < len(allDivisors); j++ {
	// 		if remainder % allDivisors[j] == 0 {
	// 			remainder = remainder / allDivisors[j]
	// 		} else {
	// 			thisIsNotADivisor = false
	// 		}
	// 	}
	// 	if thisIsNotADivisor {
	// 		numbers = append(numbers, int(i))
	// 		// fmt.Println(allDivisors, len(allDivisors), i)
	// 	}
	// }
	// fmt.Println(numbers)

	// t0 := time.Now()
	// remainders := arrayOfRemainders(25)
	// fmt.Println(remainders)
	// t1 := time.Now()
	// fmt.Printf("The call took %v to run.\n", t1.Sub(t0))

	// t0 = time.Now()
	// remainders = arrayOfRemainders(1000)
	// fmt.Println(remainders)
	// t1 = time.Now()
	// fmt.Printf("The call took %v to run.\n", t1.Sub(t0))

	// t0 = time.Now()
	// remainders = arrayOfRemainders(10000)
	// fmt.Println(remainders)
	// t1 = time.Now()
	// fmt.Printf("The call took %v to run.\n", t1.Sub(t0))

	// t0 = time.Now()
	// remainders = arrayOfRemainders(20000)
	// fmt.Println(remainders)
	// t1 = time.Now()
	// fmt.Printf("The call took %v to run.\n", t1.Sub(t0))

	// t0 = time.Now()
	// remainders = arrayOfRemainders(40000)
	// fmt.Println(remainders)
	// t1 = time.Now()
	// fmt.Printf("The call took %v to run.\n", t1.Sub(t0))

	// t0 = time.Now()
	// remainders = arrayOfRemainders(50000)
	// fmt.Println(remainders)
	// t1 = time.Now()
	// fmt.Printf("The call took %v to run.\n", t1.Sub(t0))

	// t0 = time.Now()
	// remainders = arrayOfRemainders(100000)
	// fmt.Println(remainders)
	// t1 = time.Now()
	// fmt.Printf("The call took %v to run.\n", t1.Sub(t0))

	// t0 = time.Now()
	// remainders = arrayOfRemainders(200000)
	// fmt.Println(remainders)
	// t1 = time.Now()
	// fmt.Printf("The call took %v to run.\n", t1.Sub(t0))

	// t0 := time.Now()
	// remainders := arrayOfRemainders(250250)
	// //fmt.Println(remainders)
	// t1 := time.Now()
	// fmt.Printf("creating the array took %v to run.\n", t1.Sub(t0))
	// countForEachRemainder := countEachRemainder(remainders)
	// fmt.Println(countForEachRemainder[0:4])
	t2 := time.Now()
	// fmt.Printf("counting each remainder took %v to run.\n", t2.Sub(t1))
	testArray := []int{1, 2, 3, 4, 5}
	countOfCombinations := countCombinations(testArray)
	fmt.Println(countOfCombinations)
	t3 := time.Now()
	fmt.Printf("counting each combination took %v to run.\n", t3.Sub(t2))

}

/* Takes a number and returns all the prime divisors, without repeating previous numbers.
 * eg. 10 returns [2 5], 100 returns [2 5], 8 returns [2].
 */
func nonRepeatingDivisorsOf(upper int64) []int {
	var primeSlice = make([]int, 0)
	for i := int64(2); i <= upper; i++ {
		if upper%i == 0 {
			primeSlice = append(primeSlice, int(i))
			for upper%i == 0 {
				upper = upper / i
			}
		}
	}
	return primeSlice
}

/* Takes a number and returns all the prime divisors, with repeating previous numbers.
 * eg. 10 returns [2 5], 100 returns [2 5 5], 8 returns [2 2 2].
 */
func allDivisorsOf(upper int64) []int {
	var primeSlice = make([]int, 0)
	for i := int64(2); i <= upper; i++ {
		for upper%i == 0 {
			primeSlice = append(primeSlice, int(i))
			upper = upper / i

		}
	}
	return primeSlice
}

/* Returns the remainder of (x^x / 250) 
 * 
 */
func arrayOfRemainders(x int) []int {
	remainderSlice := make([]int, 0)
	for i := 1; i <= x; i++ {
		steps, remainder := make([]int, 0), i
	J:
		for j := 1; j < i; j++ {
			remainder = remainder * i % 250
			//steps = append(steps, remainder)
			for k := 0; k < len(steps); k++ {
				if remainder == steps[k] {
					remainder = steps[i%len(steps)]
					// fmt.Println("i", i, "j", j, "k", k,  "remainder", remainder)
					break J
				}
			}
			steps = append(steps, remainder)
			// fmt.Println("i", i, "j", j, "remainder", remainder)
			// if remainder == remainder * i % 250 {
			// 	break
			// }
		}
		remainderSlice = append(remainderSlice, remainder)
	}
	return remainderSlice
}

func countEachRemainder(array []int) []int {
	count := make([]int, 250)
	//fmt.Println(len(array))
	for i := 0; i < len(array); i++ {
		count[int(math.Abs(float64(array[i])))]++ // = count[n] + 1
	}
	return count
}

func countCombinations(array []int) int {
	fmt.Println(array, len(array))
	count := array[0]

	for i := len(array) - 1; i >= 1; i++ {
		fmt.Println("count + array[", i, "] * array[", len(array)-i, "]")
		fmt.Println(count, "+", array[i], "*", array[len(array)-i], "=")
		count = count + array[i]*countCombinations(array[i:len(array)-i+1])
	}

	// if len(array) <= 0 {
	// 	fmt.Println("len less than 0", array, len(array))
	// 	return 1
	// } else if len(array) <= 1 {
	// 	fmt.Println("len less than 1", array, len(array))
	// 	count = uint64(array[0])
	// } else {
	// 	for i := len(array) ; i >= 0 ; i-- {
	// 		count = uint64(array[i-1]) * countCombinations(array[0:len(array)-i])
	// 	}
	// }
	return count
}

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

package main

import (
	"fmt"
	"math"
	"time"
)

func main() {

	Problem250(250250) // Modulus = 250 is also a special number, this should probably be abstracted to a variable.

	// len(array) == 3, ten of each number - should get a result of 1100
	// 10			=  10 {0}
	// 10 * 10		= 100 {1, 2}
	// 10 * 10 * 9	= 900 {1, 1, 1}
	testAmounts := []int{10, 10, 10}

	combos := combinations(testAmounts)
	fmt.Println(combos)
}

func Problem250(number int) {

	// Calculate the remainder for each number in the set
	startTime := time.Now()
	remainders := arrayOfRemainders(number)
	// fmt.Println(remainders)
	endTime := time.Now()
	fmt.Printf("Creating remainder array took %v to run.\n\n", endTime.Sub(startTime))

	// Count each remainder, returns a smaller set (of length 250)
	startTime = time.Now()
	countForEachRemainder := countEachRemainder(remainders)
	fmt.Println(countForEachRemainder)
	endTime = time.Now()
	fmt.Printf("Counting each remainder took %v to run.\n\n", endTime.Sub(startTime))

}

// Given an array, add all the numbers in the array that add to len array 
// eg. for [10, 10, 10, 10, 10]  is an array of len 5, remainders 0-4.

// for array len 5

// 5			10							   10
// 4+1			10 * 10						  100
// 3+2			10 * 10						  100
// 3+1+1		10 * 10 *  9				  900
// 2+2+1		10 *  9	* 10				  900
// 2+1+1+1		10 * 10 *  9  *  8			 7200
// 1+1+1+1+1	10 *  9 *  8  *  7 *  6		30240

// 17 635 968 000 000 000 000

// function(x) 
// if x = 1, return the current count of x 
// if else split into x-y and y
// otherwise recursive call of this over y

// 11 ways to add 6. 7 ways to add 5. 5 ways to add 4. 3 ways to add 3. definitely a prime sequence.
// dont know how that will come in handy yet.
// len(array) == 6. There are at least 3 algorithms that I can see, listed below.

// 1. Exhaust all of the combos where array[0] is highest, before reducing array[0] by 1...
// 2. All combos of len 1, then 2, then 3...
// 3. A wierd mix of trying to spread all numbers (I wrote this one ages ago and could be wrong about the order) (more notes below)

// 3. ok so the pattern seems to be "take a number from the first one, shift it 
// 3. accross to the next, if that hasnt been done, calculate, if it has,
// 3. keep shifting until either its a new combo or its +1 on the end.

// 6				6				6
// 5+1				5+1				5+1
// 4+2				4+2				4+1+1
// 4+1+1			3+3				3+1+1+1
// 3+3				4+1+1			2+1+1+1+1
// 3+2+1			3+2+1			1+1+1+1+1+1
// 3+1+1+1			2+2+2			4+2
// 2+2+2			3+1+1+1			3+3
// 2+2+1+1			2+2+1+1			3+2+1
// 2+1+1+1+1		2+1+1+1+1		2+2+2
// 1+1+1+1+1+1		1+1+1+1+1+1		2+2+1+1
func combinations(x []int) int {
	// len(array) == 3, ten of each number - should get a result of 1100
	// 10			=  10 {0}
	// 10 * 10		= 100 {1, 2}
	// 10 * 10 * 9	= 900 {1, 1, 1}	
	sum := 0
	for i := 0; i < len(x); i++ {

	}
	return sum
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

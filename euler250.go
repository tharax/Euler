// Find the number of non-empty subsets of {1^1, 2^2, 3^3,..., 250250^250250},
// the sum of whose elements is divisible by 250. Enter the rightmost 16 digits as your answer.

// Currently I have an array length(250) showing frequency of each remainder.
// My next step is to calculate the number of combinations possible.

// eg. Say you had {1, 1, 1, 2, 2, 2, 3, 3, 3} and you were adding to
// Possible Combinations
// 	1	3*2*1			6
// 	2	3*2*1			6
// 	3	3			3
// 	1&2	3*2*1 * 3*2*1		36
// 	1&3	3*2*1 * 3		18
// 	2&3	3*2*1 * 3		18
// 	1&2&3	3*2*1 * 3*2*1 * 3	108

//195 possible cominbations. and that's just for matching sets, not crossover like 1+1+2+2. ugh. I have more thinking to do.

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

	// Count the occurence of each remainder in the set.
	remainders := countOfEachRemainder(number, divisor)
	fmt.Println(remainders)

	//Returns all the combinations that add up to
	// startTime = time.Now()
	// fmt.Println(countForEachRemainder[0:2])
	// combos := combinations(divisor, countForEachRemainder[0:2])
	// fmt.Println(combos)
	// endTime = time.Now()
	// fmt.Printf("Creating every combination took %v to run.\n\n", endTime.Sub(startTime))

}

// Return map of len(divisor), where each value is the frequency of the remainder in the larger array.
func countOfEachRemainder(number, divisor int) []int {
	remainders := arrayOfRemainders(number, divisor)
	remain := make([]int, divisor)
	for _, value := range remainders {
		remain[value] += 1
	}
	return remain
}

// Returns an array of number^number % divisor
func arrayOfRemainders(number, divisor int) []int {
	array := make([]int, 0)
	for i := 1; i <= number; i++ {
		array = append(array, remainder(i, divisor))
	}
	return array
}

// Returns number^number % divisor
func remainder(number, divisor int) int {
	total, remainderChain := number, make([]int, 0)
	total = total % divisor
	remainderChain = append(remainderChain, total)
J:
	for i := 1; i < number; i++ {
		total = (total * number) % divisor
		if checkRemainderAlreadyFound(remainderChain, total) {
			total = remainderChain[(number-1)%len(remainderChain)]
			break J
		}
		remainderChain = append(remainderChain, total)
	}
	return total
}

// If total already exists in remainderChain array, return true.
func checkRemainderAlreadyFound(remainderChain []int, total int) bool {
	for k := 0; k < len(remainderChain); k++ {
		if total == remainderChain[k] {
			return true
		}
	}
	return false
}

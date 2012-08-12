// The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

// Find the sum of all the primes below two million.

package main

import "fmt"

func main() {
	low, high := uint64(0), uint64(2000000) //2 000 000
	s := sumOfPrimesBetween(low, high)

	fmt.Println("sum of primes between", low, "and", high, "is", s)
}

func sumOfPrimesBetween(lowerBottom, upperBottom uint64) uint64 {
	if lowerBottom < 2 {
		lowerBottom = 2
	}
	isPrime, sumOfPrimes := true, uint64(0)
	var primes = make([]uint64, 0)
	for i := lowerBottom; i <= upperBottom; i++ {
		isPrime = true
	L:
		for j := 0; j < len(primes); j++ {
			if i%primes[j] == 0 {
				isPrime = false
				break L
			}
		}
		if isPrime {
			sumOfPrimes += i
			primes = append(primes, i)
			//fmt.Println(i, sumOfPrimes)
		}
	}
	//fmt.Println("Slice created")
	return sumOfPrimes
}

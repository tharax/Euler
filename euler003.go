//The prime factors of 13195 are 5, 7, 13 and 29.
//What is the largest prime factor of the number 600851475143 ?

// Ok first I am going to produce a Slice of prime numbers <30
// now divide the large number by members of the primeSlice

// ok now that the program works - the reason that it worked before was because
// i created a map of all the possible primes, and then iterated through them all.
// I changed it so that it finds the next prime, and then checks if it can be divided by that

package main

import "fmt"

func main() {
	startTime = time.Now()
	fmt.Println(Problem3(1, 600851475143))
	endTime = time.Now()
	fmt.Printf("Slice took %v to make.\n\n", endTime.Sub(startTime))
}

func Problem3(lowerBottom, upperBottom uint64) []uint64 {
	isPrime := true
	primeSlice := make([]uint64, 0)
K:
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
			for upperBottom%i == 0 {
				upperBottom = upperBottom / i
				primeSlice = append(primeSlice, i)
				if upperBottom == i {
					break K
				}
			}
		}
	}
	return primeSlice
}

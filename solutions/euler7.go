// rebuilt euler 3 - my first program for primes - 
// because thats a hell of a lot quicker.

package main

import "fmt"


func main(){
	largeNumber := uint64(250)e
	s := makePrimeSlice(1, largeNumber)
	//fmt.Println(s)

	if len(s) == 10001 {
		fmt.Println("10001 prime =", s[len(s) - 1])
	} else {
		fmt.Println("slice length = ", len(s))
		fmt.Println("highest prime =", s[len(s) - 1])
	}run
	
	
	/*for i := range s  {
		if s[i] != 0 {
			for ;largeNumber % s[i] == 0; {
				largeNumber = largeNumber / s[i]
				fmt.Printf("%d ",s[i])
			}
		}
	}*/	
}

func makePrimeSlice(lowerBottom, upperBottom uint64) []uint64{
	var isPrime bool = true
	var primeSlice = make([]uint64, 0)
K:	for i := lowerBottom; i <= upperBottom; i++ {
		isPrime = true		
L:		for j := i-lowerBottom; j>lowerBottom; j--{
			if i % j == 0 {
				isPrime = false
				break L
			}
		}
		if isPrime && i > 1 {
			//for ;upperBottom % i == 0; {
			//	upperBottom = upperBottom / i
				primeSlice = append(primeSlice, i)
				fmt.Println(i)
			//	if upperBottom == i {
					// break K
			//	}
			//}
		}

		if len(primeSlice) == 10001 {
			fmt.Println(primeSlice[len(primeSlice) - 1])
			break K
		}

	}
	fmt.Println("Slice created")
	return primeSlice
}
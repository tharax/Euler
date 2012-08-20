//A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 99.
//Find the largest palindrome made from the product of two 3-digit numbers.

//ok the quickest way to do this, reusing the code I have written before
//is probably to work backwards from 999*999 (the highest number) and then test
//for the first palindrome that can be made from multiplying two 3-digits together.

package main

import (
	"fmt"
	"strconv"
)

var x, y int

func main() {
	upper := 999
	biggestPalindrome := upper * upper
L:
	for ; biggestPalindrome > 1; biggestPalindrome-- {
		if numIsPalindrome(biggestPalindrome) {
			m := makePrimeSlice(1, biggestPalindrome)
			if checkSlice(upper, m) {
				fmt.Println("Use", biggestPalindrome, m)
				break L
			} else {
				fmt.Println("Don't use", biggestPalindrome, m)
			}

		}
	}
}

//this algorithm is dirty, I don't like it.
func checkSlice(upper int, m []int) bool {
	x = int(m[len(m)-1])
	y = 1

	for i := len(m) - 1; i > 0; i-- {
		if x*m[i-1] < upper {
			x = x * m[i-1]
		} else {
			y = y * m[i-1]
		}
		if x > upper || y > upper {
			return false
		}
	}
	return true
}

func numIsPalindrome(x int) bool {
	s := strconv.Itoa(x)
	return s == Reverse(s)
}

//pinched from http://stackoverflow.com/questions/1752414/how-to-reverse-a-string-in-go
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func makePrimeSlice(lowerBottom, upperBottom int) []int {
	var isPrime bool = true
	var primeSlice = make([]int, 0)
	//var unusedPrimes = make([]int, 0)
	var i, j int
K:
	for i = lowerBottom; i <= upperBottom; i++ {
		isPrime = true
	L:
		for j = i - lowerBottom; j > lowerBottom; j-- {
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

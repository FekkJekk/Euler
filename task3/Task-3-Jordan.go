package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(largestPrimeFactor(600851475143))
}

func largestPrimeFactor(n int) int {
	var possiblePrime int
	debug := 0
OuterLoop:
	for i := 2; i < n; i++ {
		if n%i == 0 {
			possiblePrime = n / i
			for j := 2; j < int(math.Sqrt(float64(possiblePrime))); j++ {
				if possiblePrime%j == 0 {
					continue OuterLoop
				}
			}
			return possiblePrime
		}
	}
	return debug
}

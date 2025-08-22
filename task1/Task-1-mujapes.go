package main

import (
	"fmt"
)

func main() {
	var sum int
	//Check and sum multiples by brute force
	for i := 1.0; i < 1000; i++ {
		if i/3 == float64(int(i/3)) || i/5 == float64(int(i/5)) {
			sum += int(i)
		}
	}
	//sum of multiples of 3 and 5 are added together, common multiples of 3 and 5 = multiples of 15, sum of which are subtracted
	fmt.Println(sum, sumOfMultiples(3, 1000)+sumOfMultiples(5, 1000)-sumOfMultiples(15, 1000))
}

// Divide maxValue by base and round down to find largest multiplicand, calculate sum of all smaller or equal multiplicands,
// multiply this value by base for sum of multiples less than maxValue
// sum of multiples = 3*1 + 3*2 ... + 3*333 = 3*(1 + 2 ... + 333) = 3*((333 * 334) / 2)
func sumOfMultiples(base, maxValue int) int {
	maxMultiplicand := (maxValue - 1) / base
	sumOfMultiplicands := (maxMultiplicand * (maxMultiplicand + 1)) / 2
	return sumOfMultiplicands * base
}

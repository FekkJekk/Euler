package main

import "fmt"

func main() {
	fmt.Println(sumOfEvens(fibSlice(4000000)), sumOfEvenFibs(4000000))
}

func fibSlice(maxValue int) []int {
	var fibSequence []int
	fibSequence = append(fibSequence, 1)
	fibSequence = append(fibSequence, 2)
	for i := 2; fibSequence[i-1] <= maxValue; i++ {
		fibSequence = append(fibSequence, fibSequence[i-1]+fibSequence[i-2])
	}
	return fibSequence
}

func sumOfEvens(fibSlice []int) int {
	sum := 0
	for i := 0; i < len(fibSlice)-1; i++ {
		if fibSlice[i]%2 == 0 {
			sum += fibSlice[i]
		}
	}
	return sum
}

// More efficient solution, 4 variables less space than slice containing all fibonacci numbers calculated, 
//summation and modulo 2 maybe more efficient than append
func sumOfEvenFibs(maxValue int) int {
	fibPrevious := 1
	fibCurrent := 2
	fibNext := 0
	sum := 2
	for fibNext < maxValue {
		fibNext = fibPrevious + fibCurrent
		fibPrevious = fibCurrent
		fibCurrent = fibNext
		if fibCurrent%2 == 0 {
			sum += fibCurrent
		}
	}
	return sum
}

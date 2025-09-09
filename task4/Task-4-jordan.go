package main

import (
	"fmt"
	"strconv"
)

func palindrome(num int) bool {
	numConv := strconv.Itoa(num)
	j := len(numConv) - 1
	for i := 0; i < len(numConv)/2; i++ {
		if numConv[i] != numConv[j] {
			return false
		}
		j--
	}
	return true
}

func main() {
	maxPal := 0
	next := 0
	checked := make(map[int]struct{})
	for i := 999; i > 99; i-- {
		for j := 999; j > 99; j-- {
			next = i * j
			if next > maxPal {
				if _, ok := checked[next]; !ok {
					if palindrome(next) {
						maxPal = next
					}
					checked[next] = struct{}{}
				}
			}
		}
	}
	fmt.Println(maxPal)
}

//Time: O(n^2)
//Space: O(n^2)

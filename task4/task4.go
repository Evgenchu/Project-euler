package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(num int) bool {
	numStr := strconv.Itoa(num)
	for i, j := 0, len(numStr)-1; i < j; i, j = i+1, j-1 {
		if numStr[i] != numStr[j] {
			return false
		}
	}
	return true
}

func maxPal() int {
	maxPali := 0
	for i := 100; i < 1000; i++ {
		for j := 100; j < 1000; j++ {
			product := i * j
			if isPalindrome(product) && product > maxPali {
				maxPali = product
			}
		}
	}
	return maxPali
}
func main() {
	result := maxPal()
	fmt.Println(result)
}

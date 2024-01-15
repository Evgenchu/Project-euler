package main

import "fmt"

func collatzSequenceLength(n int) int {
	length := 1
	for n != 1 {
		if n%2 == 0 {
			n /= 2
		} else {
			n = 3*n + 1
		}
		length++
	}
	return length
}

func findLongestCollatzChain(limit int) int {
	maxLength := 0
	startingNumber := int(0)

	for i := int(1); i < limit; i++ {
		currentLength := collatzSequenceLength(i)
		if currentLength > maxLength {
			maxLength = currentLength
			startingNumber = i
		}
	}

	return startingNumber
}

func main() {
	limit := int(1000000)
	result := findLongestCollatzChain(limit)
	fmt.Println(result)
}

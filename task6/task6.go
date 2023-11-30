package main

import "fmt"

func squareSum(num int) int {
	result := 0
	for i := 1; i <= num; i++ {
		result += i * i
	}
	return result
}

func sumSquare(num int) int {
	result := 0
	sum := 0
	for i := 1; i <= num; i++ {
		sum += i
	}
	result = sum * sum
	return result
}

func main() {
	number := 100
	fmt.Println(sumSquare(number) - squareSum(number))
}

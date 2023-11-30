package main

import "fmt"

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}

	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func findLargestPrimeFactor(number int) int {
	result := 0
	for i := 2; i <= number; i++ {
		if isPrime(i) && number%i == 0 {
			result = i
			number /= i
		}
	}
	return result
}

func main() {
	number := 600851475143
	fmt.Println(findLargestPrimeFactor(number))
}

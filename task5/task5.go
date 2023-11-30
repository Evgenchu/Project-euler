package main

import "fmt"

func smallestMult(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		if result%i != 0 {
			result *= i / commondiv(result, i)
		}
	}
	return result
}

func commondiv(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	result := smallestMult(20)
	fmt.Println(result)
}

package main

import "fmt"

func fib(limit int) int {
	a, b := 1, 2
	sum := 0
	for a <= limit {
		if a%2 == 0 {
			sum += a
		}
		a, b = b, a+b
	}
	return sum
}

func main() {
	limit := 4000000
	fmt.Println(fib(limit))
}

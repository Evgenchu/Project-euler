package main

import "fmt"

func checkIfPitagorean(a, b, c int) bool {
	return a*a+b*b == c*c
}

func main() {
	for a := 1; a <= 1000; a++ {
		for b := a + 1; b <= 1000; b++ {
			c := 1000 - a - b
			if checkIfPitagorean(a, b, c) && a+b+c == 1000 {
				fmt.Printf("The numbers are %d, %d, %d\n", a, b, c)
				product := a * b * c
				fmt.Printf("The product is %d\n", product)
				return
			}
		}
	}
}

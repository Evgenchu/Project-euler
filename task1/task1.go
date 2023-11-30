package main

import "fmt"

func main() {
	var answer int
	for number := 1; number < 1000; number++ {
		if number%3 == 0 || number%5 == 0 {
			answer = answer + number
			fmt.Println(answer)
		}
	}
}

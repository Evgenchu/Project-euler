package main

import "fmt"

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}

	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func ShowPrimeList(num int) int {
	primearr := []int{}
	i := 2
	for {
		if isPrime(i) {
			primearr = append(primearr, i)
			if len(primearr) == num {
				break
			}
		}
		i++
	}
	return primearr[num-1]
}

func main() {
	number := 10001
	fmt.Println(ShowPrimeList(number))
}

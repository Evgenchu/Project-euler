package main

import "fmt"

func sieveOfEratosthenes(limit int) []int {
	isPrime := make([]bool, limit+1)
	primes := []int{}

	for i := 2; i*i <= limit; i++ {
		if !isPrime[i] {
			for j := i * i; j <= limit; j += i {
				isPrime[j] = true
			}
		}
	}

	for i := 2; i <= limit; i++ {
		if !isPrime[i] {
			primes = append(primes, i)
		}
	}

	return primes
}

func sumOfPrimesBelow(limit int) int {
	primes := sieveOfEratosthenes(limit)
	sum := 0
	for _, prime := range primes {
		sum += prime
	}
	return sum
}

func main() {
	limit := 2000000
	result := sumOfPrimesBelow(limit)
	fmt.Printf("The sum of all primes below %d is: %d\n", limit, result)
}

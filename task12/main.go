package main

import (
	"fmt"
	"math"
)

func triangleNumber(n int) int {
	return (n * (n + 1)) / 2
}

func listFactors(num int) int {
	count := 0
	sqrtNum := int(math.Sqrt(float64(num)))

	for i := 1; i <= sqrtNum; i++ {
		if num%i == 0 {
			count += 2
		}
	}

	if sqrtNum*sqrtNum == num {
		count--
	}

	return count
}

func main() {
	i := 1
	for {
		result := triangleNumber(i)
		if listFactors(result) > 500 {
			fmt.Println(result)
			break
		}
		i++
	}
}

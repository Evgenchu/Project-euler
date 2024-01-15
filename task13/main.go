package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

func readNumbersFromFile(filename string) ([]*big.Int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var numbers []*big.Int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var number big.Int
		_, success := number.SetString(line, 10)
		if !success {
			return nil, fmt.Errorf("failed to parse number: %s", line)
		}
		numbers = append(numbers, &number)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return numbers, nil
}

func main() {
	numbers, err := readNumbersFromFile("input.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	var sum big.Int
	for _, num := range numbers {
		sum.Add(&sum, num)
	}
	sumStr := sum.String()
	fmt.Println(sumStr[:10])
}

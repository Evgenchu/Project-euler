package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findGreatestProduct(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])
	maxProduct := 0

	for i := 0; i < rows; i++ {
		for j := 0; j <= cols-4; j++ {
			product := grid[i][j] * grid[i][j+1] * grid[i][j+2] * grid[i][j+3]
			maxProduct = max(maxProduct, product)
		}
	}

	for i := 0; i <= rows-4; i++ {
		for j := 0; j < cols; j++ {
			product := grid[i][j] * grid[i+1][j] * grid[i+2][j] * grid[i+3][j]
			maxProduct = max(maxProduct, product)
		}
	}

	for i := 0; i <= rows-4; i++ {
		for j := 0; j <= cols-4; j++ {
			product := grid[i][j] * grid[i+1][j+1] * grid[i+2][j+2] * grid[i+3][j+3]
			maxProduct = max(maxProduct, product)
		}
	}

	for i := 3; i < rows; i++ {
		for j := 0; j <= cols-4; j++ {
			product := grid[i][j] * grid[i-1][j+1] * grid[i-2][j+2] * grid[i-3][j+3]
			maxProduct = max(maxProduct, product)
		}
	}

	return maxProduct
}

func readFile(filename string) ([][]int, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	contentStr := string(content)
	contentStr = strings.TrimSpace(contentStr)
	rowsStr := strings.Split(contentStr, "\n")

	var grid [][]int
	for _, rowStr := range rowsStr {
		numStr := strings.Fields(rowStr)
		var row []int

		for _, num := range numStr {
			n, err := strconv.Atoi(num)
			if err != nil {
				return nil, err
			}
			row = append(row, n)
		}

		grid = append(grid, row)
	}

	return grid, nil
}

func main() {
	filename := "input.txt"
	grid, err := readFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Print(grid)

	result := findGreatestProduct(grid)
	fmt.Println("The greatest product of four adjacent numbers:", result)
}

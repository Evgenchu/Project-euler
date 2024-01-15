package main

func routes(size int) int {
	grid := make([][]int, size+1)

	for i := 0; i < size+1; i++ {
		grid[i] = make([]int, size+1)
	}

	for i := 0; i < size; i++ {
		grid[size][i] = 1
		grid[i][size] = 1
	}
	grid[size][size] = 0

	for x := size - 1; x >= 0; x-- {
		for y := size - 1; y >= 0; y-- {
			grid[x][y] = grid[x+1][y] + grid[x][y+1]
		}
	}

	return grid[0][0]

}

func main() {
	println(routes(20))
}

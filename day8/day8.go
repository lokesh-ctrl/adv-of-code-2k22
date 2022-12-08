package adv22

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func initalizeInput(inputFile string) ([][]int, int) {
	lines, _ := parseAndReadLines(inputFile)

	grid := make([][]int, 0)

	for _, line := range lines {
		temp := make([]int, 0)
		for _, char := range line {
			value, _ := strconv.Atoi(string(char))
			temp = append(temp, value)
		}
		grid = append(grid, temp)
	}

	return grid, len(lines)
}

func checkVisiblity(grid [][]int, x int, y int, height int) bool {
	isVisible := 0
	width := len(grid[0])
	currentValue := grid[x][y]
	// fmt.Printf("Coming here \n")
loop1:
	for i := x - 1; i >= 0; i-- {
		if currentValue <= grid[i][y] {
			// fmt.Printf("Not visible top %d %d %d\n", x, y, grid[x][y])
			isVisible++
			break loop1
		}
	}
	currentValue = grid[x][y]
loop2:
	for i := x + 1; i <= height-1; i++ {
		if currentValue <= grid[i][y] {
			// fmt.Printf("Not visible bottom %d %d %d\n", x, y, grid[x][y])
			isVisible++
			break loop2
		}
	}
	currentValue = grid[x][y]
loop3:
	for j := y - 1; j >= 0; j-- {
		if currentValue <= grid[x][j] {
			// fmt.Printf("Not visible left %d %d %d\n", x, y, grid[x][y])
			isVisible++
			break loop3
		}
	}
	currentValue = grid[x][y]
loop4:
	for j := y + 1; j <= width-1; j++ {
		if currentValue <= grid[x][j] {
			// fmt.Printf("Not visible right %d %d %d\n", x, y, grid[x][y])
			isVisible++
			break loop4
		}
	}
	if isVisible == 4 {
		fmt.Printf("\n this number is not visiblie %d %d %d\n", x, y, grid[x][y])
		return false
	} else {
		return true
	}
}

func findVisibliScore(grid [][]int, x int, y int, height int) int {
	currentScore := 1
	width := len(grid[0])
	currentValue := grid[x][y]
	localScore := 0
loop1:
	for i := x - 1; i >= 0; i-- {
		if currentValue <= grid[i][y] {
			// fmt.Printf("Not visible top %d %d %d\n", x, y, grid[x][y])
			localScore++
			break loop1
		} else {
			localScore++
		}
	}
	currentScore = currentScore * localScore
	localScore = 0
loop2:
	for i := x + 1; i <= height-1; i++ {
		if currentValue <= grid[i][y] {
			// fmt.Printf("Not visible bottom %d %d %d\n", x, y, grid[x][y])
			localScore++
			break loop2
		} else {
			localScore++
		}
	}
	currentScore = currentScore * localScore
	localScore = 0
loop3:
	for j := y - 1; j >= 0; j-- {
		if currentValue <= grid[x][j] {
			// fmt.Printf("Not visible left %d %d %d\n", x, y, grid[x][y])
			localScore++
			break loop3
		} else {
			localScore++
		}
	}
	currentScore = currentScore * localScore
	localScore = 0
loop4:
	for j := y + 1; j <= width-1; j++ {
		if currentValue <= grid[x][j] {
			// fmt.Printf("Not visible right %d %d %d\n", x, y, grid[x][y])
			localScore++
			break loop4
		} else {
			localScore++
		}
	}
	currentScore = currentScore * localScore
	return currentScore
}

func Day8P1(inputFile string) {
	grid, height := initalizeInput(inputFile)
	fmt.Printf("\n Outer visibilt %d \n", 2*(len(grid[0]))+2*(height-2))
	totalVisible := 2*(len(grid[0])) + 2*(height-2)
	for i := 1; i < len(grid[0])-1; i++ {
		for j := 1; j < height-1; j++ {
			if checkVisiblity(grid, i, j, height) {
				totalVisible++
			}
		}
	}

	fmt.Printf("\n Answer is total visibile %d", totalVisible)
}

func Day8P2(inputFile string) {
	var totalVisibleScores []int
	highestPossible := 0
	grid, height := initalizeInput(inputFile)
	for i := 0; i < len(grid[0]); i++ {
		for j := 0; j < height; j++ {
			currentScore := findVisibliScore(grid, i, j, height)
			totalVisibleScores = append(totalVisibleScores, currentScore)
			if currentScore > highestPossible {
				highestPossible = currentScore
			}
		}
	}

	fmt.Printf("\n Answer highest visibile %d", highestPossible)
	fmt.Println("\n total scores", totalVisibleScores)
}

func parseAndReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

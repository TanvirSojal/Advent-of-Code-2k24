// problem: https://adventofcode.com/2024/day/6
package main

import (
	"fmt"
	"os"
	"strings"
)

type position struct {
	x int
	y int
}

const up int = 1 << 0
const right int = 1 << 1
const down int = 1 << 2
const left int = 1 << 3

func main() {
	file, err := os.ReadFile("input.txt")

	if err != nil {
		return
	}

	text := string(file)
	inputGrid := strings.Split(text, "\n")

	height := len(inputGrid)
	width := len(inputGrid[0])

	startX, startY := getGuardPosition(inputGrid)

	grid := make([][]int, height)

	for i := range grid {
		grid[i] = make([]int, width)
	}

	for i := range inputGrid {
		for j := range inputGrid[i] {
			grid[i][j] = int(inputGrid[i][j])
		}
	}

	solveFirstHalf(grid, height, width, position{startX, startY})
	solveSecondHalf(grid, height, width, position{startX, startY})
}

func solveFirstHalf(grid [][]int, height int, width int, pos position) {
	direction := up
	dx, dy := -1, 0

	visited := make(map[position]bool)

	count := 0

	for {
		if !visited[pos] {
			visited[pos] = true
			count++
		}

		nextX := pos.x + dx
		nextY := pos.y + dy

		if isOut(height, width, nextX, nextY) {
			break
		}

		if isObstacle(grid, nextX, nextY) {
			dx, dy, direction = getNewDirection(direction)
			continue
		}

		pos.x = nextX
		pos.y = nextY
	}

	fmt.Println(count)
}

func solveSecondHalf(grid [][]int, height int, width int, pos position) {
	count := 0

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == int('.') {
				newGrid := getCopy(grid)
				newGrid[i][j] = int('#')

				if checkLoop(newGrid, height, width, pos) {
					count++
				}
			}
		}
	}

	fmt.Println(count)
}

func checkLoop(grid [][]int, height int, width int, pos position) bool {
	direction := up
	dx, dy := -1, 0

	visited := make(map[position]int)

	spot := make([][]int, height)

	for i := range spot {
		spot[i] = make([]int, width)
	}

	for {
		// bitwise operation to check
		// if the same spot was visited again from the same direction!
		if (visited[pos] & direction) == direction {
			return true
		}

		// mark only the direction bits (to keep up to 4 directions in a single int)
		visited[pos] = visited[pos] | direction

		spot[pos.x][pos.y]++

		if spot[pos.x][pos.y] > 1000 {
			break
		}

		nextX := pos.x + dx
		nextY := pos.y + dy

		if isOut(height, width, nextX, nextY) {
			break
		}

		if isObstacle(grid, nextX, nextY) {
			dx, dy, direction = getNewDirection(direction)
			continue
		}

		pos.x = nextX
		pos.y = nextY
	}

	return false
}

func getGuardPosition(grid []string) (int, int) {
	for i := range grid {
		for j := range grid[i] {
			if rune(grid[i][j]) == '^' {
				return i, j
			}
		}
	}

	return 0, 0
}

func isObstacle(grid [][]int, x int, y int) bool {
	return grid[x][y] == '#'
}

func isOut(height int, width int, x int, y int) bool {
	return x < 0 || x >= height || y < 0 || y >= width
}

func getNewDirection(direction int) (int, int, int) {
	var dx, dy int
	var newDirection int

	switch direction {
	case up:
		dx = 0
		dy = 1
		newDirection = right
	case right:
		dx = 1
		dy = 0
		newDirection = down
	case down:
		dx = 0
		dy = -1
		newDirection = left

	case left:
		dx = -1
		dy = 0
		newDirection = up
	}

	return dx, dy, newDirection
}

func getCopy(grid [][]int) [][]int {
	duplicate := make([][]int, len(grid))

	for i := range grid {
		duplicate[i] = make([]int, len(grid[0]))
		copy(duplicate[i], grid[i])
	}

	return duplicate
}

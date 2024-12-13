// problem: https://adventofcode.com/2024/day/9
package main

import (
	"fmt"
	"os"
	"strings"
)

var grid [][]int
var visited [][]bool
var height int
var width int

var dx = []int{0, 0, -1, 1}
var dy = []int{-1, 1, 0, 0}

var dp [][]int // only used in second half

func main() {
	file, err := os.ReadFile("input.txt")

	if err != nil {
		return
	}

	lines := strings.Split(string(file), "\n")

	grid, height, width = parseGrid(lines)

	solveFirstHalf()
	solveSecondHalf()
}

func solveFirstHalf() {
	answer := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 0 {
				resetVisitedMap()
				score := hike(i, j)
				answer += score
			}
		}
	}

	fmt.Println(answer)
}

func solveSecondHalf() {
	resetVisitedMap()
	initializeDp()

	answer := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 0 {
				score := hikeAllTrails(i, j)
				answer += score
			}
		}
	}

	fmt.Println(answer)
}

func hike(x int, y int) int {
	visited[x][y] = true

	if grid[x][y] == 9 {
		return 1
	}

	count := 0

	for i := range dx {
		_x := x + dx[i]
		_y := y + dy[i]

		if isValid(_x, _y, grid[x][y]) {
			count += hike(_x, _y)
		}
	}

	return count
}

func hikeAllTrails(x int, y int) int {
	// we will find all trails leading to one peak (9)
	// therefore, not marking the cell as visited
	if grid[x][y] == 9 {
		return 1
	}

	if dp[x][y] != -1 {
		return dp[x][y]
	}

	count := 0

	for i := range dx {
		_x := x + dx[i]
		_y := y + dy[i]

		if isValid(_x, _y, grid[x][y]) {
			count += hikeAllTrails(_x, _y)
		}
	}

	dp[x][y] = count

	return count
}

func parseGrid(lines []string) ([][]int, int, int) {
	grid := make([][]int, len(lines))
	for i := range lines {
		grid[i] = make([]int, len(strings.TrimSpace(lines[i])))
		for j := range grid[i] {
			grid[i][j] = int(lines[i][j] - '0')
		}
	}

	return grid, len(grid), len(grid[0])
}

func isValid(x int, y int, v int) bool {
	return x >= 0 && x < height && y >= 0 && y < width && !visited[x][y] && grid[x][y]-v == 1
}

func resetVisitedMap() {
	visited = make([][]bool, len(grid))
	for i := range grid {
		visited[i] = make([]bool, len(grid[i]))
	}
}

func initializeDp() {
	dp = make([][]int, len(grid))
	for i := range dp {
		dp[i] = make([]int, len(grid[i]))
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
}

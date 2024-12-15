// problem: https://adventofcode.com/2024/day/12
package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strings"
)

const up int = 1
const down int = 2
const left int = 3
const right int = 4

type cell struct {
	x         int
	y         int
	direction int
}

var height int
var width int
var grid []string
var visited [][]bool

var dx = []int{-1, 1, 0, 0}
var dy = []int{0, 0, -1, 1}
var dir = []int{up, down, left, right}

var sideMap map[byte][]cell

func main() {
	file, err := os.ReadFile("input.txt")

	if err != nil {
		return
	}

	input := strings.Split(string(file), "\n")

	grid, height, width = parseGrid(input)

	calculateCostBasedOnAreaAndPerimeter()
	calculateCostBasedOnAreaAndSide()
}

func calculateCostBasedOnAreaAndPerimeter() {
	resetVisited()

	cost := 0

	for i := range grid {
		for j := range grid[i] {
			if !visited[i][j] {
				area, perimeter := dfs(i, j)
				cost += (area * perimeter)
			}
		}
	}

	fmt.Println(cost)
}

func calculateCostBasedOnAreaAndSide() {
	resetVisited()

	sideMap = make(map[byte][]cell)

	cost := 0

	for i := range grid {
		for j := range grid[i] {
			if !visited[i][j] {
				// reset side cells because same plant can
				// have multiple disjoint plots in a grid
				sideMap[grid[i][j]] = make([]cell, 0)

				area := dfs2(i, j)
				sides := countSides(sideMap[grid[i][j]])

				cost += (area * sides)
			}
		}
	}

	fmt.Println(cost)
}

func dfs(x int, y int) (int, int) {
	visited[x][y] = true

	perimeter := 0
	area := 1

	for i := 0; i < len(dx); i++ {
		_x := x + dx[i]
		_y := y + dy[i]

		// outer borders or plot of different plant
		if isOut(_x, _y) || grid[_x][_y] != grid[x][y] {
			perimeter++
			continue
		}

		if isVisited(_x, _y) { // plot of same plant but already visited
			continue
		}

		_area, _perimeter := dfs(_x, _y) // same plant but undiscovered plot, go there
		perimeter += _perimeter
		area += _area
	}

	return area, perimeter
}

func dfs2(x int, y int) int {
	visited[x][y] = true

	area := 1

	for i := 0; i < len(dx); i++ {
		_x := x + dx[i]
		_y := y + dy[i]

		if isOut(_x, _y) || grid[_x][_y] != grid[x][y] {
			memo := sideMap[grid[x][y]]
			memo = append(memo, cell{_x, _y, dir[i]})
			sideMap[grid[x][y]] = memo
			continue
		}

		if isVisited(_x, _y) {
			continue
		}

		area += dfs2(_x, _y)
	}

	return area
}

func countSides(list []cell) int {
	sides := 0

	for _, direction := range dir {
		sides += countSidesForDirection(list, direction)
	}

	return sides
}

func countSidesForDirection(list []cell, direction int) int {
	sameDirList := getCellsOnTheSameDirection(list, direction)

	if len(sameDirList) == 0 {
		return 0
	}

	sides := 1

	for i := 1; i < len(sameDirList); i++ {
		if isAdjacent(sameDirList[i], sameDirList[i-1]) {
			continue
		}
		sides++
	}

	return sides
}

func getCellsOnTheSameDirection(list []cell, direction int) []cell {
	result := make([]cell, 0)

	for _, c := range list {
		if c.direction == direction {
			result = append(result, c)
		}
	}

	if direction == up || direction == down {
		sortByRow(result) // since we are trying to find adjacent x
	} else {
		sortByColumn(result) // since we are trying to find adjacent y
	}

	return result
}

func sortByRow(list []cell) {
	sort.Slice(list, func(i, j int) bool {
		if list[i].x == list[j].x {
			return list[i].y < list[j].y
		}

		return list[i].x < list[j].x
	})
}

func sortByColumn(list []cell) {
	sort.Slice(list, func(i, j int) bool {
		if list[i].y == list[j].y {
			return list[i].x < list[j].x
		}

		return list[i].y < list[j].y
	})
}

func isAdjacent(a cell, b cell) bool {
	if a.x == b.x {
		return int(math.Abs(float64((b.y - a.y)))) == 1
	}

	if a.y == b.y {
		return int(math.Abs(float64((b.x - a.x)))) == 1
	}

	return false
}

func resetVisited() {
	visited = make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[i]))
	}
}

func isVisited(x int, y int) bool {
	return visited[x][y]
}

func isOut(x int, y int) bool {
	return x < 0 || x >= height || y < 0 || y >= width
}

func parseGrid(input []string) ([]string, int, int) {
	grid = make([]string, len(input))

	for i := range input {
		grid[i] = strings.TrimSpace(input[i])
	}

	return grid, len(grid), len(grid[0])
}

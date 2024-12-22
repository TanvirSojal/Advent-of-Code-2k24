// problem: https://adventofcode.com/2024/day/15
package main

import (
	"bufio"
	"fmt"
	"os"
)

type position struct {
	x, y int
}

type robot struct {
	position position
}

const up = 1
const down = 2
const left = 3
const right = 4

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		return
	}

	defer file.Close()

	grid, moves := parseInput(file)

	robot := getRobotPosition(grid)

	moveRobot(robot, grid, moves)

	printGrid(grid)

	calculateGps(grid)
}

func calculateGps(grid [][]rune) {
	ans := 0

	for i := range(grid) {
		for j := range(grid[i]) {
			if (grid[i][j] == 'O'){
				ans += (100 * i) + j
			}
		}
	}

	fmt.Println(ans)
}

func moveRobot(robot position, grid [][]rune, moves []int) {
	for _, direction := range(moves) {
		freeSpace := findSpace(grid, robot, direction)
		robot = moveAndPush(grid, robot, freeSpace, direction)
	} 
}

func findSpace(grid [][]rune,  start position, direction int) position {
	dx, dy := getDelta(direction)

	initial := start

	for grid[start.x + dx][start.y + dy] != '#' {
		start.x += dx
		start.y += dy

		if grid[start.x][start.y] == '.' {
			return start
		}
	}

	return initial
}

func moveAndPush(grid [][]rune, from position, to position, direction int) position {
	dx, dy := getDelta(direction)

	robot := from

	for to.x != from.x || to.y != from.y{
		grid[to.x][to.y] = grid[to.x - dx][to.y - dy]

		if grid[to.x][to.y] == '@' {
			robot = to
		}

		to.x -= dx
		to.y -=dy

		grid[to.x][to.y] = '.'
	}

	return robot
}

func getDelta(move int) (int, int){
	switch move {
		case up: 
			return -1, 0
	case down:
		return 1, 0
	case left:
		return 0, -1
	case right:
		return 0, 1
	}
	return 0, 0
}

func getRobotPosition(grid [][]rune) position{
	for i := range(grid) {
		for j := range(grid[i]) {
			if grid[i][j] == '@'{
				return position{i, j}
			}
		}
	}

	return position{0, 0}
}

func parseInput(file *os.File) ([][]rune, []int) {
	scanner := bufio.NewScanner(file)

	grid := make([][]rune, 0)
	moves := make([]int, 0)

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			break
		}

		grid = append(grid, parseRow(line))
	}

	for scanner.Scan() {
		line := scanner.Text()
		for _, ch := range(line) {
			moves = append(moves, getMove(ch))
		}
	}

	return grid, moves
}

func parseRow(line string) []rune {
	row := make([]rune, 0)
	for _, ch := range(line) {
		row = append(row, ch)
	}
	return row
}

func getMove(ch rune) int {
	switch ch {
	case '^':
		return up
	case 'v':
		return down
	case '<':
		return left
	case '>':
		return right
	}
	return 0
}

func printGrid(grid [][]rune) {
	for i := range(grid) {
		for j := range(grid[i]) {
			fmt.Printf("%c", grid[i][j])
		}
		fmt.Println()
	}
}
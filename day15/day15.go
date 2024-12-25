// problem: https://adventofcode.com/2024/day/15
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type position struct {
	x, y int
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

	warehouse := getWareHouseMap(grid)

	robot := getRobotPosition(grid)

	moveRobot(robot, grid, moves)

	warehouseRobot := getRobotPosition(warehouse)

	moveWarehouseRobot(warehouseRobot, warehouse, moves)

	printGrid(grid)
	calculateGps(grid)

	printGrid(warehouse)
	calculateGps(warehouse)
}

func moveWarehouseRobot(robot position, grid [][]rune, moves []int) {
	for _, direction := range moves {
		possible, movables := findMovables(grid, robot, direction)
		if possible {
			robot = moveMovables(grid, movables, direction)
		}
	}
}

func findMovables(grid [][]rune, start position, direction int) (bool, []position) {
	dx, dy := getDelta(direction)
	possible, movables := canMove(grid, start.x, start.y, dx, dy, direction)
	return possible, movables
}

func canMove(grid [][]rune, x int, y int, dx int, dy int, direction int) (bool, []position) {
	if grid[x][y] == '#' {
		return false, make([]position, 0)
	}
	if grid[x][y] == '.' {
		return true, make([]position, 0)
	}

	possible := false
	movables := make([]position, 0)

	ch := grid[x][y]

	if ch == '@' {
		_possible, _movables := canMove(grid, x+dx, y+dy, dx, dy, direction)

		possible = _possible

		if _possible {
			movables = append(movables, _movables...)
			movables = append(movables, position{x, y})
		}
	} else if ch == '[' || ch == ']' {
		if direction == left || direction == right {
			_possible, _movables := canMove(grid, x+dx, y+dy, dx, dy, direction)
			possible = _possible

			if _possible {
				movables = append(movables, _movables...)
				movables = append(movables, position{x, y})
			}
		} else {
			_left, _right := 0, 0

			if ch == '[' {
				_left = y
				_right = y + 1
			} else {
				_left = y - 1
				_right = y
			}

			possibleLeft, movablesLeft := canMove(grid, x+dx, _left+dy, dx, dy, direction)
			possibleRight, movablesRight := canMove(grid, x+dx, _right+dy, dx, dy, direction)
			possible = possibleLeft && possibleRight

			if possibleLeft && possibleRight {
				movables = append(movables, movablesLeft...)
				movables = append(movables, movablesRight...)
				movables = append(movables, position{x, _left})
				movables = append(movables, position{x, _right})
			}
		}

		return possible, movables
	}
	return possible, movables
}

func moveMovables(grid [][]rune, movables []position, direction int) position {
	newRobotPosition := position{}

	mp := make(map[position]rune)

	// sort the movable cells according to direction
	// so that no cell overrides another
	if direction == up {
		sortAscending(movables)
	} else if direction == down {
		sortDescending(movables)
	}

	for _, pos := range movables {
		mp[pos] = grid[pos.x][pos.y]
	}

	dx, dy := getDelta(direction)

	for _, pos := range movables {
		ch := mp[pos]
		grid[pos.x+dx][pos.y+dy] = ch
		grid[pos.x][pos.y] = '.'

		if ch == '@' {
			newRobotPosition.x = pos.x + dx
			newRobotPosition.y = pos.y + dy
		}
	}

	return newRobotPosition
}

func getWareHouseMap(grid [][]rune) [][]rune {
	warehouse := make([][]rune, len(grid))

	for i := range grid {
		row := make([]rune, 0)
		for j := range grid[i] {
			switch grid[i][j] {
			case '#':
				row = append(row, '#')
				row = append(row, '#')
			case 'O':
				row = append(row, '[')
				row = append(row, ']')
			case '.':
				row = append(row, '.')
				row = append(row, '.')
			case '@':
				row = append(row, '@')
				row = append(row, '.')
			}
		}
		warehouse[i] = row
	}

	return warehouse
}

func calculateGps(grid [][]rune) {
	ans := 0

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 'O' || grid[i][j] == '[' {
				ans += (100 * i) + j
			}
		}
	}

	fmt.Println(ans)
}

func moveRobot(robot position, grid [][]rune, moves []int) {
	for _, direction := range moves {
		freeSpace := findSpace(grid, robot, direction)
		robot = moveAndPush(grid, robot, freeSpace, direction)
	}
}

func findSpace(grid [][]rune, start position, direction int) position {
	dx, dy := getDelta(direction)

	initial := start

	for grid[start.x+dx][start.y+dy] != '#' {
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

	for to.x != from.x || to.y != from.y {
		grid[to.x][to.y] = grid[to.x-dx][to.y-dy]

		if grid[to.x][to.y] == '@' {
			robot = to
		}

		to.x -= dx
		to.y -= dy

		grid[to.x][to.y] = '.'
	}

	return robot
}

func getDelta(move int) (int, int) {
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

func getRobotPosition(grid [][]rune) position {
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '@' {
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
		for _, ch := range line {
			moves = append(moves, getMove(ch))
		}
	}

	return grid, moves
}

func parseRow(line string) []rune {
	row := make([]rune, 0)
	for _, ch := range line {
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

func sortAscending(list []position) {
	sort.Slice(list, func(i, j int) bool {
		if list[i].x == list[j].x {
			return list[i].y < list[j].y
		}
		return list[i].x < list[j].x
	})
}

func sortDescending(list []position) {
	sort.Slice(list, func(i, j int) bool {
		if list[i].x == list[j].x {
			return list[i].y > list[j].y
		}
		return list[i].x > list[j].x
	})
}

func printGrid(grid [][]rune) {
	for i := range grid {
		for j := range grid[i] {
			fmt.Printf("%c", grid[i][j])
		}
		fmt.Println()
	}
}

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.ReadFile("input.txt")

	if err != nil {
		return
	}

	lines := strings.Split(string(file), "\n")

	var grid [][]rune = make([][]rune, len(lines))

	for i, line := range lines {
		grid[i] = []rune(strings.TrimSpace(line))
	}

	solveFirstHalf(grid)
	solveSecondHalf(grid)
}

func solveFirstHalf(grid [][]rune) {
	count := 0

	for i, row := range grid {
		for j := range row {
			count += countXmasOccurrenceFromIndex(grid, i, j)
		}
	}

	fmt.Println(count)
}

func solveSecondHalf(grid [][]rune) {
	count := 0

	for i, row := range grid {
		for j, ch := range row {
			if ch == 'A' && isCrossMasOccurrenceFromIndex(grid, i, j) {
				count++
			}
		}
	}

	fmt.Println(count)
}

func isCrossMasOccurrenceFromIndex(grid [][]rune, x int, y int) bool {
	word := []rune("MAS")

	diagonalCount := 0

	// top-left to bottom-right
	if isWord(grid, x-1, y-1, word, []int{0, 1, 2}, []int{0, 1, 2}) {
		diagonalCount++
	}

	// bottom-left to top-right
	if isWord(grid, x+1, y-1, word, []int{0, -1, -2}, []int{0, 1, 2}) {
		diagonalCount++
	}

	// top-right to bottom-left
	if isWord(grid, x-1, y+1, word, []int{0, 1, 2}, []int{0, -1, -2}) {
		diagonalCount++
	}

	// bottom-right to top-left
	if isWord(grid, x+1, y+1, word, []int{0, -1, -2}, []int{0, -1, -2}) {
		diagonalCount++
	}

	// fmt.Println(diagonalCount)

	return diagonalCount >= 2
}

func countXmasOccurrenceFromIndex(grid [][]rune, x int, y int) int {
	count := 0

	wordToFind := []rune("XMAS")

	// right
	if isWord(grid, x, y, wordToFind, []int{0, 0, 0, 0}, []int{0, 1, 2, 3}) {
		count++
	}

	// left
	if isWord(grid, x, y, wordToFind, []int{0, 0, 0, 0}, []int{0, -1, -2, -3}) {
		count++
	}

	// top
	if isWord(grid, x, y, wordToFind, []int{0, -1, -2, -3}, []int{0, 0, 0, 0}) {
		count++
	}

	// bottom
	if isWord(grid, x, y, wordToFind, []int{0, 1, 2, 3}, []int{0, 0, 0, 0}) {
		count++
	}

	// top left
	if isWord(grid, x, y, wordToFind, []int{0, -1, -2, -3}, []int{0, -1, -2, -3}) {
		count++
	}

	// top right
	if isWord(grid, x, y, wordToFind, []int{0, -1, -2, -3}, []int{0, 1, 2, 3}) {
		count++
	}

	// bottom left
	if isWord(grid, x, y, wordToFind, []int{0, 1, 2, 3}, []int{0, -1, -2, -3}) {
		count++
	}

	// bottom right
	if isWord(grid, x, y, wordToFind, []int{0, 1, 2, 3}, []int{0, 1, 2, 3}) {
		count++
	}

	return count
}

func isWord(grid [][]rune, x int, y int, word []rune, deltaX []int, deltaY []int) bool {
	if len(word) != len(deltaX) || len(word) != len(deltaY) {
		return false
	}

	for index := range deltaX {
		_x := x + deltaX[index]
		_y := y + deltaY[index]
		if !isLetter(grid, _x, _y, word[index]) {
			return false
		}
	}

	return true
}

func isLetter(grid [][]rune, x int, y int, ch rune) bool {
	if isValidIndex(grid, x, y) && grid[x][y] == ch {
		return true
	}

	return false
}

func isValidIndex(grid [][]rune, x int, y int) bool {
	h := len(grid)
	w := len(grid[0])

	return x >= 0 && x < h && y >= 0 && y < w
}

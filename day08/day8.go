// problem: https://adventofcode.com/2024/day/8
package main

import (
	"fmt"
	"os"
	"strings"
)

type position struct {
	x         int
	y         int
}

func main() {
	file, err := os.ReadFile("input.txt")

	if err != nil {
		return
	}

	grid := strings.Split(string(file), "\n")

	height := len(grid)
	width := len(strings.TrimSpace(grid[0]))

	antiNodeGrid := make([][]int, height)

	frequencyMap := getFrequencyMap(antiNodeGrid, grid, width)

	uniqueAntiNodeCount := 0

	for _, antennas := range frequencyMap {
		for i := 0; i < len(antennas); i++ {
			for j := i + 1; j < len(antennas); j++ {
				uniqueAntiNodeCount += markAntiNodes(antiNodeGrid, height, width, antennas[i], antennas[j])
			}
		}
	}

	fmt.Println(uniqueAntiNodeCount) // first half answer

	for _, antennas := range frequencyMap {
		for i := 0; i < len(antennas); i++ {
			for j := i + 1; j < len(antennas); j++ {
				uniqueAntiNodeCount += markAntiNodesWithResonateHarmonics(antiNodeGrid, height, width, antennas[i], antennas[j])
			}
		}
	}

	fmt.Println(uniqueAntiNodeCount) // second half answer
}

func markAntiNodes(antiNodeGrid [][]int, height int, width int, a position, b position) int {
	node1X := a.x + (a.x - b.x)
	node1Y := a.y + (a.y - b.y)

	node2X := b.x + (b.x - a.x)
	node2Y := b.y + (b.y - a.y)

	newLocationCount := 0

	if isValid(height, width, node1X, node1Y) && antiNodeGrid[node1X][node1Y] != int('#') {
		newLocationCount++
		antiNodeGrid[node1X][node1Y] = int('#')
	}
	if isValid(height, width, node2X, node2Y) && antiNodeGrid[node2X][node2Y] != int('#') {
		newLocationCount++
		antiNodeGrid[node2X][node2Y] = int('#')
	}

	return newLocationCount
}

func markAntiNodesWithResonateHarmonics(antiNodeGrid [][]int, height int, width int, a position, b position) int {
	positions := getAntiNodePositionsWithResonateHarmonics(height, width, a, (a.x - b.x), (a.y - b.y))
	positions = append(positions, getAntiNodePositionsWithResonateHarmonics(height, width, a, (b.x - a.x), (b.y - a.y))...)

	newLocationCount := 0

	for _, pos := range positions {
		if antiNodeGrid[pos.x][pos.y] != int('#'){
			newLocationCount++
			antiNodeGrid[pos.x][pos.y] = int('#')
		}
	}

	return newLocationCount
}

func getAntiNodePositionsWithResonateHarmonics(height int, width int, pos position, xDelta int, yDelta int) []position {
	positions := make([]position, 0)

	for {
		if !isValid(height, width, pos.x, pos.y) {
			break
		}

		positions = append(positions, position{x: pos.x, y: pos.y})

		pos.x += xDelta
		pos.y += yDelta
	}

	return positions
}

func isValid(height int, width int, x int, y int) bool {
	return x >= 0 && x < height && y >= 0 && y < width
}

func getFrequencyMap(antiNodeGrid [][]int, grid []string, width int) map[int][]position {
	frequencyMap := make(map[int][]position)

	for i := range antiNodeGrid {
		antiNodeGrid[i] = make([]int, width)

		for j := range antiNodeGrid[i] {
			if grid[i][j] != '.' {
				freq := int(grid[i][j])

				_, ok := frequencyMap[freq]

				if !ok {
					frequencyMap[freq] = make([]position, 0)
				}

				frequencyMap[freq] = append(frequencyMap[freq], position{x: i, y: j})
			}
		}
	}

	return frequencyMap
}
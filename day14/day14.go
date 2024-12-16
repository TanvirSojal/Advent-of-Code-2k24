// problem: https://adventofcode.com/2024/day/13
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
	vx, vy   int
}

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		return
	}

	defer file.Close()

	robots := parseInput(file)

	simulate(robots, 103, 101, 100)
}

func simulate(robots []robot, height int, width int, count int) {
	mp := make(map[position]int)

	for _, robot := range robots {
		_x := (robot.position.x + robot.vx*count) % width
		_y := (robot.position.y + robot.vy*count) % height

		if _x < 0 {
			_x = width + _x
		}
		if _y < 0 {
			_y = height + _y
		}

		newPosition := position{_x, _y}

		mp[newPosition]++
	}

	quads := [4]int{0, 0, 0, 0}

	for k, v := range mp {
		quadrant, ok := getQuadrant(k, height, width)

		if ok {
			quads[quadrant] += v
		}
	}

	answer := quads[0] * quads[1] * quads[2] * quads[3]

	for j := 0; j < height; j++ {

		for i := 0; i < width; i++ {
			count := mp[position{i, j}]
			if count > 0 {
				fmt.Print(rune(mp[position{i, j}]))
			} else {
				fmt.Print(".")
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")

	fmt.Println("seconds", count, answer)
}

func getQuadrant(p position, height int, width int) (int, bool) {
	midX := width / 2
	midY := height / 2

	if p.x < midX && p.y < midY {
		return 0, true
	}

	if p.x > midX && p.y < midY {
		return 1, true
	}

	if p.x < midX && p.y > midY {
		return 2, true
	}

	if p.x > midX && p.y > midY {
		return 3, true
	}

	return 0, false
}

func parseInput(file *os.File) []robot {
	scanner := bufio.NewScanner(file)

	robots := make([]robot, 0)

	for scanner.Scan() {
		input := scanner.Text()

		var robot robot

		fmt.Sscanf(input, "p=%v,%v v=%v,%v", &robot.position.x, &robot.position.y, &robot.vx, &robot.vy)

		robots = append(robots, robot)
	}

	return robots
}

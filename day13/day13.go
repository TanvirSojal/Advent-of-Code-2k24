// problem: https://adventofcode.com/2024/day/13
package main

import (
	"bufio"
	"fmt"
	"os"
)

type game struct {
	ax, ay int
	bx, by int
	px, py int
}

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		return
	}

	defer file.Close()

	games := parseGames(file)

	calculateMinimumTokens(games)

	// modify prize location based on 2nd half requirement
	for i := range games {
		games[i].px += 10000000000000
		games[i].py += 10000000000000
	}

	calculateMinimumTokens(games)
}

func calculateMinimumTokens(games []game) {
	answer := 0

	for _, game := range games {
		answer += play(game)
	}

	fmt.Println(answer)
}

func play(game game) int {
	ax, ay, bx, by, px, py := game.ax, game.ay, game.bx, game.by, game.px, game.py

	numerator := ((bx + by) * px) - (bx * (px + py))

	denominator := ((bx + by) * ax) - (bx * (ax + ay))

	if numerator%denominator != 0 {
		return 0
	}

	a := numerator / denominator

	b := (px - (ax * a)) / bx

	return a*3 + b
}

func parseGames(file *os.File) []game {
	scanner := bufio.NewScanner(file)

	games := make([]game, 0)

	for scanner.Scan() {
		a := scanner.Text()
		scanner.Scan()
		b := scanner.Text()
		scanner.Scan()
		p := scanner.Text()
		scanner.Scan()

		var ax, ay, bx, by, px, py int

		fmt.Sscanf(a, "Button A: X+%v, Y+%v", &ax, &ay)
		fmt.Sscanf(b, "Button B: X+%v, Y+%v", &bx, &by)
		fmt.Sscanf(p, "Prize: X=%v, Y=%v", &px, &py)

		games = append(games, game{ax, ay, bx, by, px, py})
	}

	return games
}

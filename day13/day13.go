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

	calculateMinimumTokens(games, 100)

	// modify prize location based on 2nd half requirement
	for i := range games {
		games[i].px += 10000000000000
		games[i].py += 10000000000000
	}

	// calculateMinimumTokens(games, 10000000000000)
}

func calculateMinimumTokens(games []game, limit int) {
	answer := 0

	for _, game := range games {
		answer += play(game, limit)
	}

	fmt.Println(answer)
}

func play(game game, limit int) int {
	best := 0
	possible := false

	for i := 0; i <= limit; i++ {
		_x := game.px - (game.ax * i)
		_y := game.py - (game.ay * i)

		if _x%game.bx == 0 && _y%game.by == 0 && _x/game.bx == _y/game.by && _x/game.bx <= 100 {
			spent := 3*i + (_x / game.bx)

			if !possible || spent < best {
				best = spent
				possible = true
			}
		}
	}

	return best
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

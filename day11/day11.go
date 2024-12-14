// problem: https://adventofcode.com/2024/day/9
package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

var dp map[int][75]int

func main() {
	file, err := os.ReadFile("input.txt")

	if err != nil {
		return
	}

	input := strings.Fields(string(file))

	stones := make([]int, len(input))

	for i, str := range input {
		stones[i], _ = strconv.Atoi(str)
	}

	dp = make(map[int][75]int)

	// iteration 25 answer 220722 time 1.0952ms
	performBlinks(copyOf(stones), 25)
	// iteration 75 answer 261952051690787 time 21.4174ms
	performBlinks(copyOf(stones), 75)
}

func performBlinks(stones []int, count int) {
	start := time.Now()

	answer := 0

	for i := range stones {
		answer += dfs(stones[i], count-1) // count-1 to iterate over [0 - range)
	}

	fmt.Println("iteration", count, "answer", answer, "time", time.Since(start))
}

func dfs(stone int, depth int) int {
	if depth < 0 {
		return 1
	}

	if dp[stone][depth] > 0 {
		return dp[stone][depth]
	}

	var memo = dp[stone]

	if stone == 0 {
		memo[depth] = dfs(1, depth-1)
		dp[stone] = memo
		return dp[stone][depth]
	}

	digits := getDigitCount(stone)

	if digits%2 == 0 {
		div := int(math.Pow(10, float64(digits/2)))
		memo[depth] = dfs(stone/div, depth-1) + dfs(stone%div, depth-1)
		dp[stone] = memo
		return dp[stone][depth]
	}

	memo[depth] = dfs(stone*2024, depth-1)
	dp[stone] = memo

	return dp[stone][depth]
}

func getDigitCount(num int) int {
	count := 0
	for num != 0 {
		num /= 10
		count++
	}
	return count
}

func copyOf(arr []int) []int {
	newArr := make([]int, len(arr))
	copy(newArr, arr)
	return newArr
}

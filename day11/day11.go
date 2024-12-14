// problem: https://adventofcode.com/2024/day/9
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

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

	performBlinks(copyOf(stones), 25)
}

func performBlinks(stones []int, count int) {
	start := time.Now()

	answers := make([]int, len(stones))

	var wg sync.WaitGroup

	for i := range stones {
		wg.Add(1)

		// parallel simulations for every stone
		go func(index, stone, count int) {
			defer wg.Done()
			answers[index] = simulateBlinks(stone, count)
		}(i, stones[i], count)
	}

	wg.Wait()

	sum := 0

	for i := range answers {
		sum += answers[i]
	}

	fmt.Println("iteration", count, "parallel", sum, "time", time.Since(start))
}

func simulateBlinks(stone int, count int) int {
	list := []int{stone}

	for blink := 1; blink <= count; blink++ {
		fmt.Println("for", stone, "iteration", blink, "list size", len(list))
		list = applyRules(list)
	}

	return len(list)
}

func applyRules(stones []int) []int {
	newStones := make([]int, 0)

	size := len(stones)

	for i := 0; i < size; i++ {
		newStones = append(newStones, applyRule(stones[i])...)
	}

	return newStones
}

func applyRule(stone int) []int {
	result := make([]int, 0)

	if stone == 0 {
		result = append(result, 1)
		return result
	}

	digits := getDigits(stone)

	if len(digits)%2 == 0 {

		mid := len(digits) / 2

		result = append(result, buildNumber(digits[:mid]))
		result = append(result, buildNumber(digits[mid:]))

		return result
	}

	result = append(result, stone*2024)

	return result
}

func getDigits(num int) []int {
	digits := make([]int, 0)

	for num != 0 {
		digits = append(digits, num%10)
		num /= 10
	}

	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}

	return digits
}

func buildNumber(digits []int) int {
	d := 1

	num := 0

	for i := len(digits) - 1; i >= 0; i-- {
		num += (digits[i] * d)
		d *= 10
	}

	return num
}

func copyOf(arr []int) []int {
	newArr := make([]int, len(arr))
	copy(newArr, arr)
	return newArr
}

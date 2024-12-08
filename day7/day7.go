// problem: https://adventofcode.com/2024/day/6
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type equation struct {
	result   int
	operands []int
}

func main() {
	file, err := os.ReadFile("input.txt")

	if err != nil {
		return
	}

	inputs := strings.Split(string(file), "\n")

	equations := make([]equation, len(inputs))

	for i, input := range inputs {
		equations[i] = parseEquation(input)
	}

	solveFirstHalf(equations)
	solveSecondHalf(equations)
}

func solveFirstHalf(equations []equation) {
	var sum int = 0

	for _, eq := range equations {
		if isPossible(eq, 1, eq.operands[0]) {
			sum += eq.result
		}
	}

	fmt.Println(sum)
}

func solveSecondHalf(equations []equation) {
	var sum int = 0

	for _, eq := range equations {
		if isPossibleWithConcat(eq, 1, eq.operands[0]) {
			sum += eq.result
		}
	}

	fmt.Println(sum)
}

func isPossible(eq equation, pos int, total int) bool {
	if pos == len(eq.operands) {
		return total == eq.result
	}

	if total > eq.result {
		return false
	}

	branch1 := isPossible(eq, pos+1, total+eq.operands[pos])
	branch2 := isPossible(eq, pos+1, total*eq.operands[pos])

	return branch1 || branch2
}

func isPossibleWithConcat(eq equation, pos int, total int) bool {
	if pos == len(eq.operands) {
		return total == eq.result
	}

	if total > eq.result {
		return false
	}

	branch1 := isPossibleWithConcat(eq, pos+1, total+eq.operands[pos])
	branch2 := isPossibleWithConcat(eq, pos+1, total*eq.operands[pos])
	branch3 := isPossibleWithConcat(eq, pos+1, concatInts(total, eq.operands[pos]))

	return branch1 || branch2 || branch3
}

func parseEquation(input string) equation {
	var eq equation

	index := strings.Index(input, ":")

	eq.result, _ = strconv.Atoi(input[:index])

	numbers := strings.Fields(input[index+1:])

	eq.operands = make([]int, len(numbers))

	for i, n := range numbers {
		eq.operands[i], _ = strconv.Atoi(n)
	}

	return eq
}

func concatInts(a int, b int) int {
	concat := strconv.Itoa(a) + strconv.Itoa(b)
	number, _ := strconv.Atoi(concat)
	return number
}

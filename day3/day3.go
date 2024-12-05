package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {
	file, err := os.ReadFile("input.txt")

	if err != nil {
		return
	}

	text := string(file)

	solveFirstHalf(text)
	solveSecondHalf(text)
}

func solveFirstHalf(text string) {
	pattern := regexp.MustCompile(`mul\([0-9]+,[0-9]+\)`)

	ops := pattern.FindAllString(text, -1)

	var sum int64 = 0

	for _, op := range ops {
		var a, b int64
		fmt.Sscanf(op, "mul(%v,%v)", &a, &b)

		sum += a * b
	}

	fmt.Println(sum)
}

func solveSecondHalf(text string) {
	pattern := regexp.MustCompile(`mul\([0-9]+,[0-9]+\)|do\(\)|don't\(\)`)

	ops := pattern.FindAllString(text, -1)

	var sum int64 = 0

	opEnabled := true

	for _, op := range ops {
		if op == "do()" {
			opEnabled = true
			continue
		} else if op == "don't()" {
			opEnabled = false
			continue
		}

		if !opEnabled {
			continue
		}

		var a, b int64
		fmt.Sscanf(op, "mul(%v,%v)", &a, &b)

		sum += a * b
	}

	fmt.Println(sum)
}

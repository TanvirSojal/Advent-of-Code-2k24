// problem: https://adventofcode.com/2024/day/2
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("input.txt")

	if err != nil {
		return
	}

	var lines = strings.Split(string(file), "\n")

	safeCount := 0
	safeWithDampenerCount := 0

	for _, line := range lines {
		arr := parseInts(strings.Fields(line))

		if isSafe(arr) {
			safeCount++
		}

		isSafe, _ := isSafeWithDampener(arr)

		if isSafe {
			safeWithDampenerCount++
		}
	}

	fmt.Println(safeCount)
	fmt.Println(safeWithDampenerCount)

}

func isSafe(arr []int) bool {
	if len(arr) < 2 {
		return true
	}

	isAsc := (arr[1] - arr[0]) > 0

	for i := range arr {
		if i == 0 {
			continue
		}

		if isAsc {
			diff := arr[i] - arr[i-1]
			if !(diff >= 1 && diff <= 3) {
				return false
			}
		} else {
			diff := arr[i-1] - arr[i]
			if !(diff >= 1 && diff <= 3) {
				return false
			}
		}
	}

	return true
}

func isSafeWithDampener(arr []int) (bool, []int) {
	if len(arr) < 2 {
		return true, arr
	}

	for i := range arr {
		newArr := []int{}

		newArr = append(newArr, arr[:i]...)
		newArr = append(newArr, arr[i+1:]...)

		if isSafe(newArr) {
			return true, newArr
		}
	}

	return false, []int{}
}

func parseInts(arr []string) []int {
	intArr := make([]int, len(arr))

	for i, str := range arr {
		intArr[i], _ = strconv.Atoi(str)
	}

	return intArr
}

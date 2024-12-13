// problem: https://adventofcode.com/2024/day/5
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	childMap := make(map[int][]int)

	for scanner.Scan() {
		input := scanner.Text()

		if len(input) == 0 {
			break
		}

		var n, m int
		fmt.Sscanf(input, "%v|%v", &n, &m)

		children := childMap[n]
		children = append(children, m)
		childMap[n] = children
	}

	var updates [][]int

	for scanner.Scan() {
		update := parseInts(strings.Split(scanner.Text(), ","))
		updates = append(updates, update)
	}

	solveFirstHalf(updates, childMap)
	solveSecondHalf(updates, childMap)
}

func solveFirstHalf(updates [][]int, childMap map[int][]int) {
	sum := 0

	for _, update := range updates {
		isValid, _, _ := isUpdateValid(update, childMap)
		if isValid {
			midIndex := len(update) / 2
			sum += update[midIndex]
		}
	}

	fmt.Println(sum)
}

func solveSecondHalf(updates [][]int, childMap map[int][]int) {
	sum := 0

	for _, update := range updates {
		isValid, parentIndex, childIndex := isUpdateValid(update, childMap)

		if !isValid {
			validUpdate := getValidUpdate(update, parentIndex, childIndex, childMap)
			midIndex := len(validUpdate) / 2
			sum += validUpdate[midIndex]
		}
	}

	fmt.Println(sum)
}

func isUpdateValid(update []int, childMap map[int][]int) (bool, int, int) {
	for parentIndex, parent := range update {
		prefixList := update[:parentIndex]

		isAnyChildFound, childIndex := isAnyChildFound(parent, prefixList, childMap)

		if isAnyChildFound {
			return false, parentIndex, childIndex
		}
	}

	return true, 0, 0
}

func isAnyChildFound(parent int, prefixList []int, childMap map[int][]int) (bool, int) {
	children := childMap[parent]
	for i, p := range prefixList {
		if exists(children, p) { // a value that came before is actually a child
			return true, i
		}
	}

	return false, 0
}

func getValidUpdate(update []int, parentIndex int, childIndex int, childMap map[int][]int) []int {
	isValid := false

	for {
		update = fixParentChildOrdering(update, parentIndex, childIndex)
		isValid, parentIndex, childIndex = isUpdateValid(update, childMap)

		if isValid {
			return update
		}
	}
}

func fixParentChildOrdering(update []int, parentIndex int, childIndex int) []int {
	prefix := update[:childIndex]
	middle := update[childIndex+1 : parentIndex]
	suffix := update[parentIndex+1:]

	arr := make([]int, 0)

	arr = append(arr, prefix...)
	arr = append(arr, middle...)
	arr = append(arr, update[parentIndex])
	arr = append(arr, update[childIndex])
	arr = append(arr, suffix...)

	return arr
}

func parseInts(arr []string) []int {
	intArr := make([]int, len(arr))

	for i, str := range arr {
		intArr[i], _ = strconv.Atoi(str)
	}

	return intArr
}

func exists(arr []int, valueToCheck int) bool {
	for _, v := range arr {
		if valueToCheck == v {
			return true
		}
	}

	return false
}

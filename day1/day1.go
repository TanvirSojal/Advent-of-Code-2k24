// problem: https://adventofcode.com/2024/day/1
package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("input.txt")

	if err != nil {
		return
	}

	var lines = strings.Split(string(file), "\n")

	var list1, list2 []int

	frequency := make(map[int]int)

	for _, line := range lines {
		numbers := strings.Fields(line)

		num1, _ := strconv.Atoi(numbers[0])
		num2, _ := strconv.Atoi(numbers[1])

		list1 = append(list1, num1)
		list2 = append(list2, num2)

		frequency[num2]++
	}

	sort.Ints(list1)
	sort.Ints(list2)

	sum := 0

	for i := range list1 {
		diff := list1[i] - list2[i]

		if diff < 0 {
			diff *= -1
		}

		sum += diff
	}

	fmt.Println(sum)

	var sum2 int64 = 0

	for _, v := range list1 {
		sum2 += int64(v) * int64(frequency[v])
	}

	fmt.Println(sum2)
}

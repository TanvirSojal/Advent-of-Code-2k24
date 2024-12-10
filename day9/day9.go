// problem: https://adventofcode.com/2024/day/9
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type block struct {
	start_index int
	end_index   int
}

type file struct {
	id     int
	blocks []block
}

func main() {
	inputFile, err := os.ReadFile("input.txt")

	if err != nil {
		return
	}

	input := strings.TrimSpace(string(inputFile))

	index := 0
	id := 0

	files := make([]file, 0)
	spaces := make([]block, 0)

	for i, b := range input {
		blockSize, _ := strconv.Atoi(string(b))

		newBlock := block{start_index: index, end_index: index + blockSize - 1}

		index += blockSize

		if i%2 == 0 { // file block
			files = append(files, file{id: id, blocks: []block{newBlock}})
			id++
		} else { // free space block
			spaces = append(spaces, newBlock)
		}
	}

	printMemory(index, files, spaces)
}

func printMemory(size int, files []file, spaces []block) {
	mem := make([]rune, size)

	for i := range mem {
		mem[i] = '.'
	}

	for _, f := range files {
		for _, b := range f.blocks {
			for i := b.start_index; i <= b.end_index; i++ {
				mem[i] = rune(f.id + 48)
			}
		}
	}

	for i := range mem {
		fmt.Printf("%c", mem[i])
	}
	fmt.Println()
}

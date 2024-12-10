// problem: https://adventofcode.com/2024/day/9
package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type block struct {
	start_index int
	end_index   int
}

func (b block) getSize() int {
	return b.end_index - b.start_index + 1
}

type file struct {
	id     int
	blocks []block
}

func (f file) getLastBlock() block {
	return f.blocks[len(f.blocks)-1]
}

func (f file) setLastBlock(b block) {
	f.blocks[len(f.blocks)-1] = b
}

func main() {
	inputFile, err := os.ReadFile("input.txt")

	if err != nil {
		return
	}

	input := strings.TrimSpace(string(inputFile))

	files, spaces, size := processInputMemory(input)

	spaceIndex := 0

	for i := len(files) - 1; i >= 0; i-- {
		for ; spaceIndex < len(spaces); spaceIndex++ {
			// if space is on the right side, do not move file
			if spaces[spaceIndex].start_index > files[i].getLastBlock().start_index {
				continue
			}

			lastBlock := files[i].getLastBlock()
			lastBlockSize := lastBlock.getSize()

			spaceSize := spaces[spaceIndex].getSize()

			availableSpace := int(math.Min(float64(lastBlockSize), float64(spaceSize)))

			if availableSpace < lastBlockSize { // split the last block and store part of it
				// create new block with leftover file part (from the left) 0 1 2 3 4 5
				leftoverBlock := block{start_index: lastBlock.start_index, end_index: lastBlock.end_index - availableSpace}

				// store part of file (from the right)
				lastBlock.start_index = spaces[spaceIndex].start_index
				lastBlock.end_index = lastBlock.start_index + availableSpace - 1

				files[i].setLastBlock(lastBlock)

				files[i].blocks = append(files[i].blocks, leftoverBlock)

			} else { // store last block and exit loop
				fmt.Println("Store entire block", files[i].id)
				// store file
				lastBlock.start_index = spaces[spaceIndex].start_index
				lastBlock.end_index = lastBlock.start_index + lastBlockSize - 1

				files[i].setLastBlock(lastBlock)

				// update space info
				spaces[spaceIndex].start_index = lastBlock.end_index + 1

				break
			}
		}
	}

	fmt.Println(files)

	printMemory(size, files)
}

func processInputMemory(input string) ([]file, []block, int) {
	files := make([]file, 0)
	spaces := make([]block, 0)

	index := 0
	id := 0

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

	return files, spaces, index
}

func printMemory(size int, files []file) {
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

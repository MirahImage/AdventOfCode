package main

import (
	"fmt"

	fileParsing "github.com/MirahImage/AdventOfCode2017/FileParsing"
)

const input = "input.txt"

func JumpsToOutside(instructions []int) int {
	jumps := 0
	for pos := 0; pos >= 0 && pos < len(instructions); jumps++ {
		instructions[pos]++
		pos = pos + instructions[pos] - 1
	}
	return jumps
}

func NewJumpsToOutside(instructions []int) int {
	jumps := 0
	for pos := 0; pos >= 0 && pos < len(instructions); jumps++ {
		if instructions[pos] < 3 {
			instructions[pos]++
			pos = pos + instructions[pos] - 1
		} else {
			instructions[pos]--
			pos = pos + instructions[pos] + 1
		}
	}
	return jumps
}

func ReadInput(file string) []int {
	lines := fileParsing.ReadFileToLines(file)
	instructions, e := fileParsing.SliceAtoi(lines)
	if e != nil {
		panic(e)
	}
	return instructions
}

func main() {
	instructions := ReadInput(input)
	newInstructions := ReadInput(input)
	fmt.Println("It took", JumpsToOutside(instructions), "steps to escape")
	fmt.Println("It took", NewJumpsToOutside(newInstructions), "new jumps to escape")
}

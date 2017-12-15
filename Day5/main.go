package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
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

func sliceAtoi(a []string) ([]int, error) {
	var nums []int
	for _, s := range a {
		i, err := strconv.Atoi(s)
		if err != nil {
			return nums, err
		}
		nums = append(nums, i)
	}
	return nums, nil
}

func ReadInput(file string) []int {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(b), "\n")
	lines = lines[:len(lines)-1] //handle trailing newline
	instructions, e := sliceAtoi(lines)
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

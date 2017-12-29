package main

import (
	"fmt"
	"strconv"
	"strings"

	fileParsing "github.com/MirahImage/AdventOfCode2017/FileParsing"
)

const input = "input.txt"

func Compare(comparator string, register string, value int, registers map[string]int) bool {
	switch comparator {
	case ">":
		return registers[register] > value
	case "<":
		return registers[register] < value
	case "==":
		return registers[register] == value
	case "!=":
		return registers[register] != value
	case ">=":
		return registers[register] >= value
	case "<=":
		return registers[register] <= value
	}
	return false
}

func Execute(instruction string, registers map[string]int) {
	parts := strings.Split(instruction, " ")
	val, _ := strconv.Atoi(parts[6])
	incr, _ := strconv.Atoi(parts[2])
	if Compare(parts[5], parts[4], val, registers) {
		if parts[1] == "inc" {
			registers[parts[0]] += incr
		} else if parts[1] == "dec" {
			registers[parts[0]] -= incr
		}
	}
}

func FindLargest(registers map[string]int) (string, int) {
	var (
		largestKey string
		largestVal int
	)
	largestKey = ""
	largestVal = 0
	for key, val := range registers {
		if largestKey == "" || val > largestVal {
			largestKey = key
			largestVal = val
		}
	}
	return largestKey, largestVal
}

func main() {
	registers := make(map[string]int)
	instructions := fileParsing.ReadFileToLines(input)
	for _, instruction := range instructions {
		Execute(instruction, registers)
	}
	largestKey, largestVal := FindLargest(registers)
	fmt.Println(largestKey, "contains the largest value", largestVal)
}

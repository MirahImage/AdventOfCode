package main

import (
	"fmt"

	. "github.com/MirahImage/AdventOfCode2017/Day6/History"
	. "github.com/MirahImage/AdventOfCode2017/Day6/Memory"
	fileParsing "github.com/MirahImage/AdventOfCode2017/FileParsing"
)

const input = "input.txt"

func CountReallocations(m Memory) (int, int) {
	var h History
	for h.Log(m) == -1 {
		m.Reallocate()
	}
	return len(h.TimeSteps), len(h.TimeSteps) - h.Log(m)
}

func main() {
	var m Memory
	lines := fileParsing.ReadFileToLines(input)
	m.Banks = fileParsing.LineToInts(lines[0])
	fmt.Println("Initial state", m.Banks)
	hist, loopSize := CountReallocations(m)
	fmt.Println(hist, "reallocations to reach a repeated state")
	fmt.Println("Loop is of size", loopSize)
}

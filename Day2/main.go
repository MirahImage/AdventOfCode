package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	. "github.com/MirahImage/AdventOfCode2017/Day2/row"
	. "github.com/MirahImage/AdventOfCode2017/Day2/spreadsheet"
)

const input = "input.txt"

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

func LineToInts(line string) []int {
	if strings.Compare(line, "") == 0 {
		return []int{}
	}
	numberStrings := strings.Split(line, "\t")
	numbers, err := sliceAtoi(numberStrings)
	if err != nil {
		panic(err)
	}
	return numbers
}

func BytesToRows(b []byte) []Row {
	lines := strings.Split(string(b), "\n")
	var rows = make([]Row, len(lines))
	for i, line := range lines {
		rows[i] = Row{Entries: LineToInts(line)}
	}
	return rows
}

func ReadInput(file string) Spreadsheet {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	return Spreadsheet{Rows: BytesToRows(b)}
}

func main() {
	spreadsheet := ReadInput(input)
	fmt.Println("Checksum for spreadsheet", input, "is", spreadsheet.Checksum())
	fmt.Println("Divisible checksum for spreadsheet", input, "is", spreadsheet.DivChecksum())
}

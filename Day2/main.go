package main

import (
	"fmt"
	"strings"

	. "github.com/MirahImage/AdventOfCode2017/Day2/row"
	. "github.com/MirahImage/AdventOfCode2017/Day2/spreadsheet"
	fileParsing "github.com/MirahImage/AdventOfCode2017/FileParsing"
)

const input = "input.txt"

func LineToInts(line string) []int {
	numberStrings := strings.Split(line, "\t")
	numbers, err := fileParsing.SliceAtoi(numberStrings)
	if err != nil {
		panic(err)
	}
	return numbers
}

func LinesToRows(lines []string) []Row {
	var rows = make([]Row, len(lines))
	for i, line := range lines {
		rows[i] = Row{Entries: LineToInts(line)}
	}
	return rows
}

func ReadInput(file string) Spreadsheet {
	lines := fileParsing.ReadFileToLines(file)
	return Spreadsheet{Rows: LinesToRows(lines)}
}

func main() {
	spreadsheet := ReadInput(input)
	fmt.Println("Checksum for spreadsheet", input, "is", spreadsheet.Checksum())
	fmt.Println("Divisible checksum for spreadsheet", input, "is", spreadsheet.DivChecksum())
}

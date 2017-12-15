package main

import (
	"fmt"

	. "github.com/MirahImage/AdventOfCode2017/Day2/row"
	. "github.com/MirahImage/AdventOfCode2017/Day2/spreadsheet"
	fileParsing "github.com/MirahImage/AdventOfCode2017/FileParsing"
)

const input = "input.txt"

func LinesToRows(lines []string) []Row {
	var rows = make([]Row, len(lines))
	for i, line := range lines {
		rows[i] = Row{Entries: fileParsing.LineToInts(line)}
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

package spreadsheet

import (
	. "github.com/MirahImage/AdventOfCode2017/Day2/row"
)

type Spreadsheet struct {
	Rows []Row
}

func (s *Spreadsheet) Checksum() int {
	checksum := 0
	for _, row := range s.Rows {
		checksum += row.Difference()
	}
	return checksum
}

func (s *Spreadsheet) DivChecksum() int {
	checksum := 0
	for _, row := range s.Rows {
		checksum += row.Resultant()
	}
	return checksum
}

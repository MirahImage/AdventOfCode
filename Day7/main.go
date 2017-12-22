package main

import (
	"fmt"
	"strconv"
	"strings"

	. "github.com/MirahImage/AdventOfCode2017/Day7/ProgramData"
	. "github.com/MirahImage/AdventOfCode2017/Day7/Tree"
	fileParsing "github.com/MirahImage/AdventOfCode2017/FileParsing"
)

const input = "input.txt"

func LinesToPrograms(lines []string) []ProgramData {
	progs := make([]ProgramData, len(lines))
	for i, line := range lines {
		words := strings.Split(line, " ")
		for i, word := range words {
			words[i] = strings.Trim(word, ",")
		}
		progs[i].Name = words[0]
		progs[i].Weight, _ = strconv.Atoi(strings.Trim(words[1], "()"))
		if len(words) > 2 {
			progs[i].ChildNames = words[3:]
		} else {
			progs[i].ChildNames = []string{}
		}
	}
	return progs
}

func main() {
	lines := fileParsing.ReadFileToLines(input)
	progs := LinesToPrograms(lines)
	var t Tree
	t.AddPrograms(progs)
	fmt.Println("The root program is named", t.Root.Data.Name)
}

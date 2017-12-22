package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	. "github.com/MirahImage/AdventOfCode2017/Day7/Node"
	. "github.com/MirahImage/AdventOfCode2017/Day7/ProgramData"
	. "github.com/MirahImage/AdventOfCode2017/Day7/Tree"
	fileParsing "github.com/MirahImage/AdventOfCode2017/FileParsing"
)

const input = "input.txt"

var writer = os.Stdout

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

func PrintUnbalanced(w io.Writer, unbalanced []*Node) {
	for _, node := range unbalanced {
		fmt.Fprintln(w, node.Data.Name, "is unbalanced, it's children are of weight")
		for _, child := range node.Children {
			fmt.Fprintln(w, child.Data.Name, "is of total weight", child.Weight(), "and local weight", child.Data.Weight)
		}
	}
}

func main() {
	lines := fileParsing.ReadFileToLines(input)
	progs := LinesToPrograms(lines)
	var t Tree
	t.AddPrograms(progs)
	fmt.Fprintln(writer, "The root program is named", t.Root.Data.Name)
	unbalanced := t.FindUnbalanced()
	PrintUnbalanced(writer, unbalanced)
}

package main

import (
	"fmt"
	"strings"

	. "github.com/MirahImage/AdventOfCode2017/Day4/passphrase"
	fileParsing "github.com/MirahImage/AdventOfCode2017/FileParsing"
)

const input = "input.txt"

func LinesToPassphrases(lines []string) []Passphrase {
	var passphrases = make([]Passphrase, len(lines))
	for i, line := range lines {
		passphrases[i] = Passphrase{Words: strings.Split(line, " ")}
	}
	return passphrases
}

func ReadInput(input string) []Passphrase {
	lines := fileParsing.ReadFileToLines(input)
	return LinesToPassphrases(lines)
}

func NumValid(passphrases []Passphrase, validator func(*Passphrase) bool) int {
	valid := 0
	for _, passphrase := range passphrases {
		if validator(&passphrase) {
			valid++
		}
	}
	return valid
}

func main() {
	passphrases := ReadInput(input)
	fmt.Println("There are", NumValid(passphrases, (*Passphrase).ValidRepeated), "valid part 1 passphrases in", input)
	fmt.Println("There are", NumValid(passphrases, (*Passphrase).ValidAnagram), "valid part 2 passphrases in", input)
}

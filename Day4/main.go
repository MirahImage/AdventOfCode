package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	. "github.com/MirahImage/AdventOfCode2017/Day4/passphrase"
)

const input = "input.txt"

func BytesToPassphrases(bytes []byte) []Passphrase {
	lines := strings.Split(string(bytes), "\n")
	var passphrases = make([]Passphrase, len(lines)-1)
	for i, _ := range passphrases {
		passphrases[i] = Passphrase{Words: strings.Split(lines[i], " ")}
	}
	return passphrases
}

func ReadInput(input string) []Passphrase {
	b, err := ioutil.ReadFile(input)
	if err != nil {
		panic(err)
	}
	return BytesToPassphrases(b)
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

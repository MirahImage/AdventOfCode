package main

import (
	"fmt"
	"strings"

	. "github.com/MirahImage/AdventOfCode2017/Day1/captcha"
	fileParsing "github.com/MirahImage/AdventOfCode2017/FileParsing"
)

const input = "input.txt"

func linesToDigits(lines []string) []int {
	digitStrings := strings.Split(lines[0], "")
	digits, err := fileParsing.SliceAtoi(digitStrings)
	if err != nil {
		panic(err)
	}
	return digits
}

func ReadInput(f string) []int {
	lines := fileParsing.ReadFileToLines(f)
	return linesToDigits(lines)
}

func main() {
	c := Captcha{Digits: ReadInput(input)}
	fmt.Println("The sum for the captcha", c.Digits, "is", c.Sum())
	fmt.Println("The halfway matched sum for the captcha", c.Digits, "is", c.SumHalfway())
}

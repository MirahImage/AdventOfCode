package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	. "github.com/MirahImage/AdventOfCode2017/Day1/captcha"
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

func bytesToDigits(b []byte) []int {
	strs := strings.Split(string(b), "\n")
	digitStrings := strings.Split(strs[0], "")
	digits, err := sliceAtoi(digitStrings)
	if err != nil {
		panic(err)
	}
	return digits
}

func ReadInput(f string) []int {
	b, err := ioutil.ReadFile(f)
	if err != nil {
		panic(err)
	}
	return bytesToDigits(b)
}

func main() {
	c := Captcha{Digits: ReadInput(input)}
	fmt.Println("The sum for the captcha", c.Digits, "is", c.Sum())
	fmt.Println("The halfway matched sum for the captcha", c.Digits, "is", c.SumHalfway())
}

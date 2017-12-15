package fileParsing

import (
  "io/ioutil"
  "strconv"
  "strings"
)

func SliceAtoi(a []string) ([]int, error) {
	nums := make([]int, len(a))
	for i, s := range a {
		num, err := strconv.Atoi(s)
		if err != nil {
			return nums, err
		}
		nums[i] = num
	}
	return nums, nil
}

func ReadFileToLines(file string) []string {
  b, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(b), "\n")
	lines = lines[:len(lines)-1] //handle trailing newline
  return lines
}

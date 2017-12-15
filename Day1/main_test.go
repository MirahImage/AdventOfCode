package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/MirahImage/AdventOfCode2017/Day1"
)

var _ = Describe("Main", func() {
	var (
		fileName     string
		shortFile    string
		digitStrings []string
	)

	BeforeEach(func() {
		fileName = "input.txt"
		shortFile = "shortInput.txt"
		digitStrings = []string{"1", "1"}
	})

	Describe("readInput", func() {
		It("should return a slice of ints", func() {
			Expect(ReadInput(fileName)).To(BeAssignableToTypeOf([]int{}))
		})

		It("should return a non-empty slice if the file is non-empty", func() {
			Expect(ReadInput(fileName)).To(Not(BeNil()))
		})

		It("should return [1, 1, 2, 2] for shortInput", func() {
			Expect(ReadInput(shortFile)).To(Equal([]int{1, 1, 2, 2}))
		})
	})
})

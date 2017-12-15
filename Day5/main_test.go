package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/MirahImage/AdventOfCode2017/Day5"
)

var _ = Describe("Main", func() {

	var (
		jumps    []int
		testFile string
	)

	BeforeEach(func() {
		jumps = []int{0, 3, 0, 1, -3}
		testFile = "test.txt"
	})

	Describe("JumpsToOutside", func() {
		It("Should count the jumps until we are outside", func() {
			Expect(JumpsToOutside(jumps)).To(Equal(5))
		})
	})

	Describe("NewJumpsToOutside", func() {
		It("Should count the jumps until we are outside", func() {
			Expect(NewJumpsToOutside(jumps)).To(Equal(10))
		})
	})

	Describe("ReadInput", func() {
		It("Should parse file into a slice of ints", func() {
			Expect(ReadInput(testFile)).To(Equal(jumps))
		})
	})

})

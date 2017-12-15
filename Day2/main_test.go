package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/MirahImage/AdventOfCode2017/Day2"
	. "github.com/MirahImage/AdventOfCode2017/Day2/row"
)

var _ = Describe("Main", func() {

	Describe("LineToInts", func() {
		var (
			tabSeparatedIntString string
		)

		BeforeEach(func() {
			tabSeparatedIntString = "5\t1\t9\t5"
		})

		It("should return a slice of ints", func() {
			Expect(LineToInts(tabSeparatedIntString)).To(Equal([]int{5, 1, 9, 5}))
		})
	})

	Describe("BytesToRows", func() {
		var (
			bytes        []byte
			expectedRows []Row
		)

		BeforeEach(func() {
			bytes = []byte("5\t1\t9\t5\n7\t5\t3\n2\t4\t6\t8\n")
			expectedRows = []Row{
				Row{Entries: []int{5, 1, 9, 5}},
				Row{Entries: []int{7, 5, 3}},
				Row{Entries: []int{2, 4, 6, 8}},
				Row{Entries: []int{}},
			}
		})

		It("should return a slice of rows", func() {
			Expect(BytesToRows(bytes)).To(Equal(expectedRows))
		})
	})

})

package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/MirahImage/AdventOfCode2017/Day2"
	. "github.com/MirahImage/AdventOfCode2017/Day2/row"
)

var _ = Describe("Main", func() {

	Describe("LinesToRows", func() {
		var (
			lines        []string
			expectedRows []Row
		)

		BeforeEach(func() {
			lines = []string{"5\t1\t9\t5", "7\t5\t3", "2\t4\t6\t8"}
			expectedRows = []Row{
				Row{Entries: []int{5, 1, 9, 5}},
				Row{Entries: []int{7, 5, 3}},
				Row{Entries: []int{2, 4, 6, 8}},
			}
		})

		It("should return a slice of rows", func() {
			Expect(LinesToRows(lines)).To(Equal(expectedRows))
		})
	})

})

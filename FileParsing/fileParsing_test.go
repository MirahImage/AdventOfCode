package fileParsing_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/MirahImage/AdventOfCode2017/FileParsing"
)

var _ = Describe("FileParsing", func() {

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

	Describe("SliceAtoi", func() {
		var (
			a         []string
			iExpected []int
		)

		Context("Input is Valid", func() {
			BeforeEach(func() {
				a = []string{"12", "1", "-2", "0"}
				iExpected = []int{12, 1, -2, 0}
			})

			It("should return integer versions of a slice of strings", func() {
				nums, err := SliceAtoi(a)
				Expect(nums).To(Equal(iExpected))
				Expect(err).ShouldNot(HaveOccurred())
			})
		})

		Context("Input is invalid", func() {
			BeforeEach(func() {
				a = []string{"", "zero"}
			})

			It("should return an error", func() {
				_, err := SliceAtoi(a)
				Expect(err).Should(HaveOccurred())
			})
		})
	})

	Describe("ReadFileToLines", func() {
		var (
			fileName      string
			expectedLines []string
		)

		BeforeEach(func() {
			fileName = "test.txt"
			expectedLines = []string{"12", "1", "-2", "0"}
		})

		It("should return the file contents as a slice of strings, one for each line", func() {
			Expect(ReadFileToLines(fileName)).To(Equal(expectedLines))
		})
	})

})

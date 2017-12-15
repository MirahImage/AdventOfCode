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

	/*Describe("sliceAtoi", func() {
		It("should return a slice of ints and no error", func() {
			Expect(sliceAtoi(digitStrings)).To(Equal([]int{1, 1}, nil))
		})

		Context("When a string contains a non-digit", func() {
			BeforeEach(func() {
				var (
					nondigit  string
					nondigitStrings []string
					err error
				)

			})

			It("should return an error if a string doesn't contain digits", func() {
				Expect(sliceAtoi()).To()
			})
		})
	})*/

})

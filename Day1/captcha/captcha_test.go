package captcha_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/MirahImage/AdventOfCode2017/Day1/captcha"
)

var _ = Describe("Captcha", func() {
	var (
		exampleRepeating   Captcha
		exampleNoRepeating Captcha
		exampleLooping     Captcha
		exampleAllOnes     Captcha
	)

	BeforeEach(func() {
		exampleRepeating = Captcha{
			Digits: []int{1, 1, 2, 2},
		}
		exampleNoRepeating = Captcha{
			Digits: []int{1, 2, 3, 4},
		}
		exampleLooping = Captcha{
			Digits: []int{9, 1, 2, 1, 2, 1, 2, 9},
		}
		exampleAllOnes = Captcha{
			Digits: []int{1, 1, 1, 1},
		}
	})

	Describe("Sum", func() {
		It("should be non-negative", func() {
			Expect(exampleRepeating.Sum()).To(BeNumerically(">=", 0))
		})

		It("should be 0 for an input with no repeating numbers", func() {
			Expect(exampleNoRepeating.Sum()).To(BeNumerically("==", 0))
		})

		It("should be positive for an input with repeating numbers", func() {
			Expect(exampleRepeating.Sum()).To(BeNumerically(">", 0))
		})

		It("should be the sum of all digits that match the next digit in the list", func() {
			Expect(exampleRepeating.Sum()).To(BeNumerically("==", 3))
		})

		It("should loop the digits to check if the first and last match", func() {
			Expect(exampleLooping.Sum()).To(BeNumerically("==", 9))
		})

		It("should accurately count the repeted digits", func() {
			Expect(exampleAllOnes.Sum()).To(BeNumerically("==", 4))
		})
	})

	Describe("MatchPrev", func() {
		It("should return true iff the digit matches the previous digit", func() {
			Expect(exampleRepeating.MatchPrev(1)).To(BeTrue())
			Expect(exampleNoRepeating.MatchPrev(1)).To(BeFalse())
		})

		It("should compare the 0th digit to the last digit", func() {
			Expect(exampleRepeating.MatchPrev(0)).To(BeFalse())
			Expect(exampleAllOnes.MatchPrev(0)).To(BeTrue())
		})
	})

	Describe("MatchHalfway", func() {
		It("should compare the zeroth digit to the digit halfway", func() {
			Expect(exampleRepeating.MatchHalfway(0)).To(BeFalse())
			Expect(exampleAllOnes.MatchHalfway(0)).To(BeTrue())
		})

		It("should compare the digit halfway to the zeroth digit", func() {
			Expect(exampleRepeating.MatchHalfway(2)).To(BeFalse())
			Expect(exampleAllOnes.MatchHalfway(2)).To(BeTrue())
		})
	})

	Describe("SumHalfway", func() {
		It("should sum matching halfway", func() {
			Expect(exampleAllOnes.SumHalfway()).To(Equal(4))
			Expect(exampleRepeating.SumHalfway()).To(Equal(0))
		})
	})

})

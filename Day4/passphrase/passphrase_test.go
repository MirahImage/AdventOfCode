package passphrase_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/MirahImage/AdventOfCode2017/Day4/passphrase"
)

var _ = Describe("Passphrase", func() {

	var (
		ex1             Passphrase
		ex2             Passphrase
		ex3             Passphrase
		unsortedStrings []string
		sortedStrings   []string
	)

	BeforeEach(func() {
		ex1 = Passphrase{
			Words: []string{"aa", "bb", "cc", "dd", "ee"},
		}
		ex2 = Passphrase{
			Words: []string{"aa", "bb", "cc", "dd", "aa"},
		}
		ex3 = Passphrase{
			Words: []string{"aa", "bb", "cc", "dd", "aaa"},
		}
		unsortedStrings = []string{"ab", "ba"}
		sortedStrings = []string{"ab", "ab"}
	})

	Describe("RepeatedString", func() {
		It("Should return false if no words are repeated", func() {
			Expect(ex1.RepeatedString()).To(BeFalse())
			Expect(ex3.RepeatedString()).To(BeFalse())
		})

		It("Should return true if a word is repeated", func() {
			Expect(ex2.RepeatedString()).To(BeTrue())
		})
	})

	Describe("SortStrings", func() {
		It("Should return a list of strings, each sorted by rune", func() {
			Expect(SortStrings(unsortedStrings)).To(Equal(sortedStrings))
		})
	})

	Describe("ValidRepeated", func() {
		It("Should return !RepeatedString", func() {
			Expect(ex1.ValidRepeated()).To(Equal(!ex1.RepeatedString()))
			Expect(ex2.ValidRepeated()).To(Equal(!ex2.RepeatedString()))
			Expect(ex3.ValidRepeated()).To(Equal(!ex3.RepeatedString()))
		})
	})

	Describe("ValidAnagram", func() {
		It("Should return !RepeatedAnagram", func() {
			Expect(ex1.ValidAnagram()).To(Equal(!ex1.RepeatedAnagram()))
			Expect(ex2.ValidAnagram()).To(Equal(!ex2.RepeatedAnagram()))
			Expect(ex3.ValidAnagram()).To(Equal(!ex3.RepeatedAnagram()))
		})
	})

})

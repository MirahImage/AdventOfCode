package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/MirahImage/AdventOfCode2017/Day4"
	. "github.com/MirahImage/AdventOfCode2017/Day4/passphrase"
)

var _ = Describe("Main", func() {
	var (
		testInput  string
		testOutput []Passphrase
		ex1        Passphrase
		ex2        Passphrase
		ex3        Passphrase
		testLines  []string
	)

	BeforeEach(func() {
		testInput = "test.txt"
		ex1 = Passphrase{
			Words: []string{"aa", "bb", "cc", "dd", "ee"},
		}
		ex2 = Passphrase{
			Words: []string{"aa", "bb", "cc", "dd", "aa"},
		}
		ex3 = Passphrase{
			Words: []string{"aa", "bb", "cc", "dd", "aaa"},
		}
		testOutput = []Passphrase{ex1, ex2, ex3}
		testLines = []string{"aa bb cc dd ee", "aa bb cc dd aa", "aa bb cc dd aaa"}
	})

	Describe("ReadInput", func() {
		It("should return a slice of Passphrase", func() {
			Expect(ReadInput(testInput)).To(BeAssignableToTypeOf([]Passphrase{}))
		})

		It("should parse the input into Passphrases", func() {
			Expect(ReadInput(testInput)).To(Equal(testOutput))
		})
	})

	Describe("LinesToPassphrases", func() {
		It("should return a slice of Passphrase", func() {
			Expect(LinesToPassphrases(testLines)).To(BeAssignableToTypeOf([]Passphrase{}))
		})

		It("should parse the test phrase into the passphrases", func() {
			Expect(LinesToPassphrases(testLines)).To(Equal(testOutput))
		})

	})

})

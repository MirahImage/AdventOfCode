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
		testBytes  []byte
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
		testBytes = []byte("aa bb cc dd ee\naa bb cc dd aa\naa bb cc dd aaa\n")
	})

	Describe("ReadInput", func() {
		It("should return a slice of Passphrase", func() {
			Expect(ReadInput(testInput)).To(BeAssignableToTypeOf([]Passphrase{}))
		})

		It("should parse the input into Passphrases", func() {
			Expect(ReadInput(testInput)).To(Equal(testOutput))
		})
	})

	Describe("BytesToPassphrases", func() {
		It("should return a slice of Passphrase", func() {
			Expect(BytesToPassphrases(testBytes)).To(BeAssignableToTypeOf([]Passphrase{}))
		})

		It("should parse the test phrase into the passphrases", func() {
			Expect(BytesToPassphrases(testBytes)).To(Equal(testOutput))
		})

	})

})

package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/MirahImage/AdventOfCode2017/Day8"
)

var _ = Describe("Main", func() {
	var (
		instructions []string
		registers    map[string]int
	)

	Describe("Execute", func() {
		BeforeEach(func() {
			instructions = []string{
				"b inc 5 if a > 1",
				"a inc 1 if b < 5",
				"c dec -10 if a >= 1",
			}
			registers = make(map[string]int)
			for _, instruction := range instructions {
				Execute(instruction, registers)
			}
		})
		It("should not increment b", func() {
			Expect(registers["b"]).To(Equal(0))
		})
		It("should increment a", func() {
			Expect(registers["a"]).To(Equal(1))
		})
		It("should decrement c", func() {
			Expect(registers["c"]).To(Equal(10))
		})
	})

	Describe("Compare", func() {
		var (
			comparator string
			register   string
			value      int
		)
		BeforeEach(func() {
			registers = make(map[string]int)
			register = "a"
		})
		Context("Comparator is >", func() {
			BeforeEach(func() {
				comparator = ">"
			})
			It("should return true if the register is greater than the value", func() {
				value = -1
				Expect(Compare(comparator, register, value, registers)).To(BeTrue())
			})
			It("should return false if the register is less than the value", func() {
				value = 1
				Expect(Compare(comparator, register, value, registers)).To(BeFalse())
			})
		})

		Context("Comparator is <", func() {
			BeforeEach(func() {
				comparator = "<"
			})
			It("should return false if the register is greater than the value", func() {
				value = -1
				Expect(Compare(comparator, register, value, registers)).To(BeFalse())
			})
			It("should return true if the register is less than the value", func() {
				value = 1
				Expect(Compare(comparator, register, value, registers)).To(BeTrue())
			})
		})

		Context("Comparator is ==", func() {
			BeforeEach(func() {
				comparator = "=="
			})
			It("should return true if the register is equal to the value", func() {
				value = 0
				Expect(Compare(comparator, register, value, registers)).To(BeTrue())
			})
			It("should return false if the register is not equal to the value", func() {
				value = 1
				Expect(Compare(comparator, register, value, registers)).To(BeFalse())
			})
		})

		Context("Comparator is !=", func() {
			BeforeEach(func() {
				comparator = "!="
			})
			It("should return false if the register is equal to the value", func() {
				value = 0
				Expect(Compare(comparator, register, value, registers)).To(BeFalse())
			})
			It("should return true if the register is not equal to the value", func() {
				value = 1
				Expect(Compare(comparator, register, value, registers)).To(BeTrue())
			})
		})

		Context("Comparator is >=", func() {
			BeforeEach(func() {
				comparator = ">="
			})
			It("should return true if the register is greater than the value", func() {
				value = -1
				Expect(Compare(comparator, register, value, registers)).To(BeTrue())
			})
			It("should return true if the register equals the value", func() {
				value = 0
				Expect(Compare(comparator, register, value, registers)).To(BeTrue())
			})
			It("should return false if the register is less than the value", func() {
				value = 1
				Expect(Compare(comparator, register, value, registers)).To(BeFalse())
			})
		})

		Context("Comparator is <=", func() {
			BeforeEach(func() {
				comparator = "<="
			})
			It("should return false if the register is greater than the value", func() {
				value = -1
				Expect(Compare(comparator, register, value, registers)).To(BeFalse())
			})
			It("should return true if the register equals the value", func() {
				value = 0
				Expect(Compare(comparator, register, value, registers)).To(BeTrue())
			})
			It("should return true if the register is less than the value", func() {
				value = 1
				Expect(Compare(comparator, register, value, registers)).To(BeTrue())
			})
		})

	})

	Describe("FindLargest", func() {
		var (
			largestKey string
			largestVal int
		)
		BeforeEach(func() {
			instructions = []string{
				"b inc 5 if a > 1",
				"a inc 1 if b < 5",
			}
			registers = make(map[string]int)
			for _, instruction := range instructions {
				Execute(instruction, registers)
			}
			largestKey, largestVal = FindLargest(registers)
		})
		It("should return the key to the largest value", func() {
			Expect(largestKey).To(Equal("a"))
		})
		It("should return the largest value", func() {
			Expect(largestVal).To(Equal(1))
		})
	})

})

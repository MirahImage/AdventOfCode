package memory_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/MirahImage/AdventOfCode2017/Day6/Memory"
)

var _ = Describe("Memory", func() {

	var (
		mem Memory
	)

	BeforeEach(func() {
		mem = Memory{
			Banks: []int{0, 2, 7, 0},
		}
	})

	Describe("largestBank", func() {
		It("should find the largest bank", func() {
			bank, _ := mem.LargestBank()
			Expect(bank).To(Equal(2))
		})

		It("should return the size of the largest bank", func() {
			_, max := mem.LargestBank()
			Expect(max).To(Equal(7))
		})
	})

	Describe("Reallocate", func() {
		It("should reallocate the bank with the most blocks", func() {
			mem.Reallocate()
			Expect(mem.Banks[2]).To(BeNumerically("<", 7))
		})

		It("should distribute the reallocated bank across the banks", func() {
			mem.Reallocate()
			Expect(mem.Banks).To(Equal([]int{2, 4, 1, 2}))
		})
	})

})

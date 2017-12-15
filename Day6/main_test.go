package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/MirahImage/AdventOfCode2017/Day6"
	. "github.com/MirahImage/AdventOfCode2017/Day6/Memory"
)

var _ = Describe("Main", func() {
	var (
		m Memory
	)

	BeforeEach(func() {
		m.Banks = []int{0, 2, 7, 0}
	})

	Describe("CountReallocations", func() {
		It("should return the number of reallocations required to reach a repeated state", func() {
			hist, loopSize := CountReallocations(m)
			Expect(hist).To(Equal(5))
			Expect(loopSize).To(Equal(4))
		})
	})

})

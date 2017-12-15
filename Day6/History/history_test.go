package history_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/MirahImage/AdventOfCode2017/Day6/History"
	. "github.com/MirahImage/AdventOfCode2017/Day6/Memory"
)

var _ = Describe("History", func() {
	var (
		h  History
		m  Memory
		b0 Bank
		b1 Bank
	)

	BeforeEach(func() {
		m.Banks = []int{0, 2, 7, 0}
		b0 = Bank{
			Banks: []int{0, 2, 7, 0},
		}
		b1 = Bank{
			Banks: []int{2, 4, 1, 2},
		}
	})

	Describe("Log", func() {
		It("should add the state to TimeSteps if it has not been seen logged before", func() {
			h.Log(m)
			Expect(h.TimeSteps).To(Equal([]Bank{b0}))
		})

		It("should return the index if the state has been seen before", func() {
			h.Log(m)
			Expect(h.Log(m)).To(Equal(0))
		})

		It("should add a new state to TimeSteps when logged after reallocation", func() {
			h.Log(m)
			m.Reallocate()
			Expect(h.Log(m)).To(Equal(-1))
			Expect(len(h.TimeSteps)).To(Equal(2))
			m.Reallocate()
			Expect(h.Log(m)).To(Equal(-1))
			Expect(len(h.TimeSteps)).To(Equal(3))
		})
	})

})

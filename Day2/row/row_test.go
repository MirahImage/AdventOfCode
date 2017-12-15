package row_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/MirahImage/AdventOfCode2017/Day2/row"
)

var _ = Describe("Row", func() {
	var (
		row Row
		ex1 Row
		ex2 Row
		ex3 Row
	)

	BeforeEach(func() {
		row = Row{
			Entries: []int{5, 2, 9, 5},
		}
		ex1 = Row{
			Entries: []int{5, 9, 2, 8},
		}
		ex2 = Row{
			Entries: []int{9, 4, 7, 3},
		}
		ex3 = Row{
			Entries: []int{3, 8, 6, 5},
		}
	})

	Describe("Difference", func() {
		It("should return the difference between the largest and smallest entries in the Row", func() {
			Expect(row.Difference()).Should(Equal(7))
		})

		Context("Row is empty", func() {
			BeforeEach(func() {
				row = Row{
					Entries: []int{},
				}
			})

			It("should return 0", func() {
				Expect(row.Difference()).Should(Equal(0))
			})
		})
	})

	Describe("Divisible", func() {
		It("should return 0 if the entry at index is not divisible by any other entries in the Row", func() {
			Expect(row.Divisible(1)).To(Equal(0))
			Expect(row.Divisible(2)).To(Equal(0))
		})

		It("should return the resultant if the entry at index is divisble by another entry in the row", func() {
			Expect(row.Divisible(0)).To(Equal(1))
			Expect(row.Divisible(3)).To(Equal(1))
		})

	})

	Describe("Resultant", func() {
		It("should return the resultant for the one divisible entry in the row", func() {
			Expect(ex1.Resultant()).To(Equal(4))
			Expect(ex2.Resultant()).To(Equal(3))
			Expect(ex3.Resultant()).To(Equal(2))
		})
	})

})

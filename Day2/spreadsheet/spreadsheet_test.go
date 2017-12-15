package spreadsheet_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/MirahImage/AdventOfCode2017/Day2/row"
	. "github.com/MirahImage/AdventOfCode2017/Day2/spreadsheet"
)

var _ = Describe("Spreadsheet", func() {
	var (
		spreadsheet Spreadsheet
		row1        Row
		row2        Row
		row3        Row
	)

	BeforeEach(func() {
		row1 = Row{
			Entries: []int{5, 1, 9, 5},
		}
		row2 = Row{
			Entries: []int{7, 5, 3},
		}
		row3 = Row{
			Entries: []int{2, 4, 6, 8},
		}
		spreadsheet = Spreadsheet{
			Rows: []Row{row1, row2, row3},
		}
	})

	Describe("Checksum", func() {
		It("should return the sum of the row differences", func() {
			Expect(spreadsheet.Checksum()).To(Equal(18))
		})
	})

	Context("a row is empty", func() {
		BeforeEach(func() {
			row2 = Row{
				Entries: []int{},
			}
			spreadsheet = Spreadsheet{
				Rows: []Row{row1, row2, row3},
			}
		})

		It("should treat the empty row as difference 0", func() {
			Expect(spreadsheet.Checksum()).To(Equal(14))
		})
	})

	Context("the spreadsheet is empty", func() {
		BeforeEach(func() {
			spreadsheet = Spreadsheet{
				Rows: []Row{},
			}
		})

		It("should give checksum 0", func() {
			Expect(spreadsheet.Checksum()).To(Equal(0))
		})
	})

})

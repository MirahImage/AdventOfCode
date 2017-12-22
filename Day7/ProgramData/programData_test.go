package programData_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/MirahImage/AdventOfCode2017/Day7/ProgramData"
)

var _ = Describe("ProgramData", func() {

	var (
		parent ProgramData
		child  ProgramData
	)

	BeforeEach(func() {
		parent = ProgramData{
			Name:       "ugml",
			Weight:     68,
			ChildNames: []string{"gyxo"},
		}
		child = ProgramData{
			Name:       "gyxo",
			Weight:     61,
			ChildNames: []string{},
		}
	})

	Describe("IsParentOf", func() {
		It("should return true if the program is a parent of the data", func() {
			Expect(parent.IsParentOf(child)).To(BeTrue())
		})
		It("should return false if the program is not a parent of the data", func() {
			Expect(child.IsParentOf(parent)).To(BeFalse())
		})
	})

})

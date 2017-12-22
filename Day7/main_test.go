package main_test

import (
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/MirahImage/AdventOfCode2017/Day7"
	. "github.com/MirahImage/AdventOfCode2017/Day7/ProgramData"
)

var _ = Describe("Main", func() {
	Describe("LinesToPrograms", func() {
		var (
			lines       []string
			grandparent ProgramData
			parent      ProgramData
			child1      ProgramData
			child2      ProgramData
			progs       []ProgramData
			expected    []ProgramData
		)
		BeforeEach(func() {
			lines = []string{"tknk (41) -> ugml", "ugml (68) -> gyxo, ebii", "gyxo (61)", "ebii (61)"}
			grandparent = ProgramData{
				Name:       "tknk",
				Weight:     41,
				ChildNames: []string{"ugml"},
			}
			child1 = ProgramData{
				Name:       "gyxo",
				Weight:     61,
				ChildNames: []string{},
			}
			child2 = ProgramData{
				Name:       "ebii",
				Weight:     61,
				ChildNames: []string{},
			}
			parent = ProgramData{
				Name:       "ugml",
				Weight:     68,
				ChildNames: []string{"gyxo", "ebii"},
			}
			expected = []ProgramData{grandparent, parent, child1, child2}
			progs = LinesToPrograms(lines)
		})
		It("should return all the programs", func() {
			Expect(len(progs)).To(Equal(len(expected)))
		})
		It("should parse the names into ProgramData.Name", func() {
			for i, prog := range progs {
				Expect(strings.Compare(prog.Name, expected[i].Name)).To(Equal(0))
			}
		})
		It("should parse the weights into Weight", func() {
			for i, prog := range progs {
				Expect(prog.Weight).To(Equal(expected[i].Weight))
			}
		})
		It("should parse the children into ChildNames", func() {
			for i, prog := range progs {
				Expect(prog.ChildNames).To(Equal(expected[i].ChildNames))
			}
		})
	})

})

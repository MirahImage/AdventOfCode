package tree_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/MirahImage/AdventOfCode2017/Day7/Node"
	. "github.com/MirahImage/AdventOfCode2017/Day7/ProgramData"
	. "github.com/MirahImage/AdventOfCode2017/Day7/Tree"
)

var _ = Describe("Tree", func() {

	Describe("AddPrograms", func() {
		var (
			progs       []ProgramData
			t           Tree
			grandparent ProgramData
			parent      ProgramData
			child       ProgramData
		)

		BeforeEach(func() {
			grandparent = ProgramData{
				Name:       "tknk",
				Weight:     41,
				ChildNames: []string{"ugml"},
			}
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
			progs = []ProgramData{grandparent, parent, child}
			t.AddPrograms(progs)
		})

		It("should contain Programs", func() {
			Expect(len(t.Programs)).To(BeNumerically(">", 0))
		})
		It("should have identified children", func() {
			Expect(len(t.Programs[0].Children)).To(BeNumerically(">", 0))
			Expect(len(t.Programs[1].Children)).To(BeNumerically(">", 0))
			Expect(len(t.Programs[2].Children)).To(Equal(0))
		})
		It("should have identified parents", func() {
			Expect(t.Programs[0].Parent).To(BeNil())
			Expect(t.Programs[1].Parent).To(Not(BeNil()))
			Expect(t.Programs[2].Parent).To(Not(BeNil()))
		})
		It("should have identified the root", func() {
			Expect(t.Root).To(Not(BeNil()))
			Expect(t.Root.Data.Name).To(Equal("tknk"))
		})
	})

	Describe("FundUnbalanced", func() {
		var (
			t          Tree
			parentData ProgramData
			child1Data ProgramData
			child2Data ProgramData
			unbalanced []*Node
		)
		Context("The tree is balanced", func() {
			BeforeEach(func() {
				parentData = ProgramData{
					Name:       "ugml",
					Weight:     68,
					ChildNames: []string{"gyxo", "ebii"},
				}
				child1Data = ProgramData{
					Name:       "gyxo",
					Weight:     61,
					ChildNames: []string{},
				}
				child2Data = ProgramData{
					Name:       "ebii",
					Weight:     61,
					ChildNames: []string{},
				}
				t.AddPrograms([]ProgramData{parentData, child1Data, child2Data})
				unbalanced = t.FindUnbalanced()
			})
			It("should return an empty slice", func() {
				Expect(len(unbalanced)).To(Equal(0))
			})
		})
		Context("The tree is unbalanced", func() {
			BeforeEach(func() {
				parentData = ProgramData{
					Name:       "ugml",
					Weight:     68,
					ChildNames: []string{"gyxo", "ebii"},
				}
				child1Data = ProgramData{
					Name:       "gyxo",
					Weight:     61,
					ChildNames: []string{},
				}
				child2Data = ProgramData{
					Name:       "ebii",
					Weight:     69,
					ChildNames: []string{},
				}
				t.AddPrograms([]ProgramData{parentData, child1Data, child2Data})
				unbalanced = t.FindUnbalanced()
			})
			It("should return a non-empty slice", func() {
				Expect(len(unbalanced)).To(BeNumerically(">", 0))
			})
		})
	})
})

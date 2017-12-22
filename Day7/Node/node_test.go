package Node_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/MirahImage/AdventOfCode2017/Day7/Node"
	. "github.com/MirahImage/AdventOfCode2017/Day7/ProgramData"
)

var _ = Describe("Node", func() {

	var (
		parentData ProgramData
		child1Data ProgramData
		child2Data ProgramData
	)

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
	})

	Describe("AddChild", func() {
		Context("There are no children yet", func() {
			var (
				parent Node
				child1 Node
			)
			BeforeEach(func() {
				parent.Data = parentData
				child1.Data = child1Data
				parent.AddChild(&child1)
			})
			It("should have a child after AddChild", func() {
				Expect(len(parent.Children)).Should(BeNumerically(">", 0))
			})
		})
		Context("There is already a child", func() {
			var (
				parent Node
				child1 Node
				child2 Node
			)
			BeforeEach(func() {
				parent.Data = parentData
				child1.Data = child1Data
				child2.Data = child2Data
				parent.AddChild(&child1)
				parent.AddChild(&child2)
			})
			It("should have 2 children", func() {
				Expect(len(parent.Children)).To(Equal(2))
			})
		})
	})

	Describe("Weight", func() {
		Context("There are no children", func() {
			var (
				child1 Node
			)
			BeforeEach(func() {
				child1.Data = child1Data
			})

			It("should have weight equal to the stored Weight", func() {
				Expect(child1.Weight()).To(Equal(child1Data.Weight))
			})
		})
		Context("There are children", func() {
			var (
				parent Node
				child1 Node
				child2 Node
			)
			BeforeEach(func() {
				parent.Data = parentData
				child1.Data = child1Data
				child2.Data = child2Data
				parent.AddChild(&child1)
				parent.AddChild(&child2)
			})
			It("should have weight equal to the sum of the parent and children", func() {
				Expect(parent.Weight()).To(Equal(parentData.Weight + child1Data.Weight + child2Data.Weight))
			})
		})
	})

	Describe("Balanced", func() {
		Context("There are no children", func() {
			var (
				child1 Node
			)
			BeforeEach(func() {
				child1.Data = child1Data
			})
			It("should be balanced", func() {
				Expect(child1.Balanced()).To(BeTrue())
			})
		})
		Context("The children are of unequal weight", func() {
			var (
				parent      Node
				child1      Node
				child2      Node
				child1Heavy ProgramData
			)
			BeforeEach(func() {
				child1Heavy = ProgramData{
					Name:       "gyxo",
					Weight:     69,
					ChildNames: []string{},
				}
				parent.Data = parentData
				child1.Data = child1Heavy
				child2.Data = child2Data
				parent.AddChild(&child1)
				parent.AddChild(&child2)
			})
			It("should not be balanced", func() {
				Expect(parent.Balanced()).To(BeFalse())
			})
		})
		Context("A child is unbalanced", func() {
			var (
				grandparent     Node
				parent1         Node
				child1          Node
				child2          Node
				parent2         Node
				child1Heavy     ProgramData
				grandparentData ProgramData
				parent2data     ProgramData
			)
			BeforeEach(func() {
				child1Heavy = ProgramData{
					Name:       "gyxo",
					Weight:     69,
					ChildNames: []string{},
				}
				parent2data = ProgramData{
					Name:       "padx",
					Weight:     198,
					ChildNames: []string{},
				}
				grandparentData = ProgramData{
					Name:       "tknk",
					Weight:     41,
					ChildNames: []string{"ugml", "padx"},
				}
				grandparent.Data = grandparentData
				parent1.Data = parentData
				parent2.Data = parent2data
				child1.Data = child1Heavy
				child2.Data = child2Data
				grandparent.AddChild(&parent1)
				grandparent.AddChild(&parent2)
				parent1.AddChild(&child1)
				parent1.AddChild(&child2)
			})
			It("Should not be balanced", func() {
				Expect(grandparent.Balanced()).To(BeFalse())
			})
		})
		Context("The children are of equal weight and balanced", func() {
			var (
				parent Node
				child1 Node
				child2 Node
			)
			BeforeEach(func() {
				parent.Data = parentData
				child1.Data = child1Data
				child2.Data = child2Data
				parent.AddChild(&child1)
				parent.AddChild(&child2)
			})
			It("should be balanced", func() {
				Expect(parent.Balanced()).To(BeTrue())
			})
		})
	})

})

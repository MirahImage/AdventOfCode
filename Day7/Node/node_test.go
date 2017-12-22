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

})

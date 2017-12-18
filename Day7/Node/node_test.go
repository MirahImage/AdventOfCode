package Node_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/MirahImage/AdventOfCode2017/Day7/Node"
	. "github.com/MirahImage/AdventOfCode2017/Day7/ProgramData"
)

var _ = Describe("Node", func() {

	Describe("NewNode", func() {
		var (
			program ProgramData
		)

		BeforeEach(func() {
			program = ProgramData{
				Name:       "pbga",
				Weight:     66,
				ChildNames: []string{},
			}
		})

		It("should have the given ProgramData as Data", func() {
			n := NewNode(program)
			Expect(n.Data).To(Equal(program))
		})
	})

	Describe("Insert", func() {
		var (
			parent ProgramData
			child  ProgramData
			node   *Node
			err    error
		)

		BeforeEach(func() {
			child = ProgramData{
				Name:       "gyxo",
				Weight:     61,
				ChildNames: []string{},
			}
			parent = ProgramData{
				Name:       "ugml",
				Weight:     68,
				ChildNames: []string{"gyxo"},
			}
		})

		Context("The inserted node is the parent", func() {
			BeforeEach(func() {
				node = NewNode(child)
				node, err = node.Insert(parent)
			})
			It("should return the parent node", func() {
				Expect(node.Data).To(Equal(parent))
			})
			It("should have the child node in Children", func() {
				Expect(len(node.Children)).To(BeNumerically(">", 0))
				Expect(node.Children[0].Data).To(Equal(child))
			})
			It("should not produce an error", func() {
				Expect(err).To(BeNil())
			})
		})
		Context("The inserted node is the child", func() {
			BeforeEach(func() {
				node = NewNode(parent)
				node, err = node.Insert(child)
			})
			It("should return the parent node", func() {
				Expect(node.Data).To(Equal(parent))
			})
			It("should have the child node in Children", func() {
				Expect(len(node.Children)).To(BeNumerically(">", 0))
				Expect(node.Children[0].Data).To(Equal(child))
			})
			It("should not produce an error", func() {
				Expect(err).To(BeNil())
			})
		})
		Context("The inserted node is a grandchild", func() {
			var (
				grandparent ProgramData
			)
			BeforeEach(func() {
				grandparent = ProgramData{
					Name:       "tknk",
					Weight:     41,
					ChildNames: []string{"ugml"},
				}
				node = NewNode(grandparent)
				node, _ = node.Insert(parent)
				node, err = node.Insert(child)
			})
			It("should have the grandchild in children of children of node", func() {
				Expect(len(node.Children[0].Children)).To(BeNumerically(">", 0))
				Expect(node.Children[0].Children[0].Data).To(Equal(child))
			})
			It("should not produce an error", func() {
				Expect(err).To(BeNil())
			})
		})
		Context("The inserted node does not belong to the tree", func() {
			var (
				notPresent ProgramData
				parent     ProgramData
				child      ProgramData
			)
			BeforeEach(func() {
				notPresent = ProgramData{
					Name:       "tflo",
					Weight:     69,
					ChildNames: []string{},
				}
				child = ProgramData{
					Name:       "gyxo",
					Weight:     61,
					ChildNames: []string{},
				}
				parent = ProgramData{
					Name:       "ugml",
					Weight:     68,
					ChildNames: []string{"gyxo"},
				}
				node = NewNode(child)
				node, _ = node.Insert(parent)
				node, err = node.Insert(notPresent)
			})
			It("should produce an error", func() {
				Expect(err).Should(HaveOccurred())
				Expect(err).To(Not(BeNil()))
			})
			It("should contain the parent", func() {
				Expect(node.IsMember(parent)).To(Not(BeNil()))
			})
			It("should contain the child", func() {
				Expect(node.IsMember(child)).To(Not(BeNil()))
			})
			It("should not contain the new program", func() {
				Expect(node.IsMember(notPresent)).To(BeNil())
			})
			It("should not be nil", func() {
				Expect(node).To(Not(BeNil()))
			})
		})
	})

	Describe("IsMember", func() {
		var (
			parent     ProgramData
			child      ProgramData
			notPresent ProgramData
			node       *Node
		)
		BeforeEach(func() {
			child = ProgramData{
				Name:       "gyxo",
				Weight:     61,
				ChildNames: []string{},
			}
			parent = ProgramData{
				Name:       "ugml",
				Weight:     68,
				ChildNames: []string{"gyxo"},
			}
			notPresent = ProgramData{
				Name:       "notHere",
				Weight:     69,
				ChildNames: []string{},
			}
			node = NewNode(parent)
			node, _ = node.Insert(child)
		})
		It("should return the parent node for parent", func() {
			Expect(node.IsMember(parent).Data).To(Equal(parent))
		})
		It("should return the child node for child", func() {
			Expect(node.IsMember(child).Data).To(Equal(child))
		})
		It("should give nil for a non-member", func() {
			Expect(node.IsMember(notPresent)).To(BeNil())
		})
	})

})

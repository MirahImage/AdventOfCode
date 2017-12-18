package tree_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/MirahImage/AdventOfCode2017/Day7/ProgramData"
	. "github.com/MirahImage/AdventOfCode2017/Day7/Tree"
)

var _ = Describe("Tree", func() {

	Describe("Insert", func() {
		var (
			t   Tree
			err error
		)
		Context("The tree is empty", func() {
			var (
				program ProgramData
			)

			BeforeEach(func() {
				t = Tree{
					Root: nil,
				}
				program = ProgramData{
					Name:       "pbga",
					Weight:     66,
					ChildNames: []string{},
				}
				err = t.Insert(program)
			})

			It("should create a new Root node", func() {
				Expect(t.Root).To(Not(BeNil()))
			})
			It("should create a Root node with the given program data", func() {
				Expect(t.Root.Data).To(Equal(program))
			})
			It("should not produce an error", func() {
				Expect(err).To(BeNil())
			})
		})

		Context("The tree is not empty", func() {
			var (
				parent ProgramData
				child  ProgramData
			)
			BeforeEach(func() {
				t = Tree{
					Root: nil,
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
			})
			Context("The new data is a parent of the root", func() {
				BeforeEach(func() {
					_ = t.Insert(child)
					err = t.Insert(parent)
				})

				It("should contain the old root node", func() {
					Expect(t.Root.IsMember(child)).To(Not(BeNil()))
				})
				It("should contain the new node", func() {
					Expect(t.Root.IsMember(parent)).To(Not(BeNil()))
				})
				It("should make the new node the root node", func() {
					Expect(t.Root.Data).To(Equal(parent))
				})
				It("should have the old root node as a child node", func() {
					Expect(len(t.Root.Children)).To(BeNumerically(">", 0))
					Expect(t.Root.IsMember(child)).To(Not(BeNil()))
					Expect(t.Root.Children[0].Data).To(Equal(child))
				})
				It("should not produce an error", func() {
					Expect(err).To(BeNil())
				})
			})
			Context("The new data is a child", func() {
				BeforeEach(func() {
					_ = t.Insert(parent)
					err = t.Insert(child)
				})
				It("should not produce an error", func() {
					Expect(err).To(BeNil())
				})
				It("should contain the parent node", func() {
					Expect(t.Root.IsMember(parent)).To(Not(BeNil()))
				})
				It("should contain the child mode", func() {
					Expect(t.Root.IsMember(child)).To(Not(BeNil()))
				})
				It("should have the new node as the child mode", func() {
					Expect(t.Root.Children[0].Data).To(Equal(child))
				})
				It("should have the new node as a member", func() {
					Expect(t.Root.IsMember(child)).To(Not(BeNil()))
				})
			})
			Context("The new data does not belong to the tree", func() {
				var (
					t          Tree
					notPresent ProgramData
					parent     ProgramData
					child      ProgramData
					err        error
				)
				BeforeEach(func() {
					notPresent = ProgramData{
						Name:       "notHere",
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
					t = Tree{
						Root: nil,
					}
					_ = t.Insert(parent)
					_ = t.Insert(child)
					err = t.Insert(notPresent)
				})
				It("should not have a nil Root", func() {
					Expect(t.Root).To(Not(BeNil()))
				})
				It("should produce an error", func() {
					Expect(err).Should(HaveOccurred())
					Expect(err).To(Not(BeNil()))
				})
				It("should not have the new node as a member", func() {
					Expect(t.Root.IsMember(notPresent)).To(BeNil())
				})
				It("should have the old nodes as members", func() {
					Expect(t.Root.IsMember(parent)).To(Not(BeNil()))
					Expect(t.Root.IsMember(child)).To(Not(BeNil()))
				})
			})
		})
	})

	Describe("InsertMultiple", func() {
		var (
			t           Tree
			grandparent ProgramData
			parent      ProgramData
			child       ProgramData
			progs       []ProgramData
		)
		BeforeEach(func() {
			t = Tree{
				Root: nil,
			}
			grandparent = ProgramData{
				Name:       "tknk",
				Weight:     41,
				ChildNames: []string{"ugml"},
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
		})
		It("should insert a length 1 slice", func() {
			progs = []ProgramData{parent}
			t.InsertMultiple(progs)
			Expect(t.Root.IsMember(parent)).To(Not(BeNil()))
		})
		It("should insert a length 3 slice with 3 generations in order", func() {
			progs = []ProgramData{grandparent, parent, child}
			t.InsertMultiple(progs)
			for _, prog := range progs {
				Expect(t.Root.IsMember(prog)).To(Not(BeNil()))
			}
		})
		It("should insert a length 3 slice with generations out of order", func() {
			progs = []ProgramData{child, grandparent, parent}
			t.InsertMultiple(progs)
			for _, prog := range progs {
				Expect(t.Root.IsMember(prog)).To(Not(BeNil()))
			}
		})
	})

})

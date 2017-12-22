package main_test

import (
	"bytes"
	"fmt"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/MirahImage/AdventOfCode2017/Day7"
	. "github.com/MirahImage/AdventOfCode2017/Day7/Node"
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

	Describe("PrintUnbalanced", func() {
		var (
			buff  bytes.Buffer
			nodes []*Node
		)
		Context("Not given any nodes", func() {
			BeforeEach(func() {
				nodes = []*Node{}
				PrintUnbalanced(&buff, nodes)
			})
			It("should not print anything", func() {
				Expect(buff.Len()).To(Equal(0))
			})
		})
		Context("Given a node with unbalanced children", func() {
			var (
				parent         Node
				child1         Node
				child2         Node
				parentData     ProgramData
				child1Data     ProgramData
				child2Data     ProgramData
				expectedOutput bytes.Buffer
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
					Weight:     69,
					ChildNames: []string{},
				}
				parent.Data = parentData
				child1.Data = child1Data
				child2.Data = child2Data
				parent.AddChild(&child1)
				parent.AddChild(&child2)
				nodes = []*Node{&parent}
				PrintUnbalanced(&buff, nodes)
				fmt.Fprintln(&expectedOutput, parent.Data.Name, "is unbalanced, it's children are of weight")
				fmt.Fprintln(&expectedOutput, child1.Data.Name, "is of total weight", child1.Weight(), "and local weight", child1.Data.Weight)
				fmt.Fprintln(&expectedOutput, child2.Data.Name, "is of total weight", child2.Weight(), "and local weight", child2.Data.Weight)
			})
			It("should print the unbalanced parent name and children's weights", func() {
				Expect(buff.String()).To(Equal(expectedOutput.String()))
			})
		})
	})

})

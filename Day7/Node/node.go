package Node

import (
	. "github.com/MirahImage/AdventOfCode2017/Day7/ProgramData"
)

type Node struct {
	Data     ProgramData
	Children []*Node
	Parent   *Node
}

func (n *Node) AddChild(child *Node) {
	n.Children = append(n.Children, child)
}

func (n *Node) Weight() int {
	weight := n.Data.Weight
	for _, child := range n.Children {
		weight += child.Weight()
	}
	return weight
}

func (n *Node) Balanced() bool {
	if len(n.Children) == 0 {
		return true
	}
	child1weight := n.Children[0].Weight()
	for _, child := range n.Children {
		if child.Weight() != child1weight || !child.Balanced() {
			return false
		}
	}
	return true
}

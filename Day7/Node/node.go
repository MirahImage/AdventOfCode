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

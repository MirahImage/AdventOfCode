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

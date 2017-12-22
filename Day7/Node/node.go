package Node

import (
	. "github.com/MirahImage/AdventOfCode2017/Day7/ProgramData"
)

type Node struct {
	Data     ProgramData
	Children []*Node
	Parent   *Node
}

func (n *Node) IsParentOf(prog ProgramData) bool {
	return n.Data.IsParent(prog)
}

func (n *Node) IsChildOf(prog ProgramData) bool {
	return prog.IsParent(n.Data)
}

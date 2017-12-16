package Node

import (
	"errors"
	"strings"

	. "github.com/MirahImage/AdventOfCode2017/Day7/ProgramData"
)

type Node struct {
	Data     ProgramData
	Children []*Node
}

func NewNode(program ProgramData) *Node {
	node := new(Node)
	node.Data = program
	return node
}

func (n *Node) Insert(program ProgramData) (*Node, error) {
	newNode := NewNode(program)
	if program.IsParent(n.Data) {
		newNode.Children = append(newNode.Children, n)
		return newNode, nil
	}
	if n.Data.IsParent(program) {
		n.Children = append(n.Children, newNode)
		return n, nil
	}
	for i, child := range n.Children {
		var err error
		n.Children[i], err = child.Insert(program)
		if err == nil {
			return n, nil
		}
	}
	return n, errors.New("program does not belong to tree")
}

func (n *Node) IsMember(program ProgramData) *Node {
	if strings.Compare(n.Data.Name, program.Name) == 0 {
		return n
	}
	for _, node := range n.Children {
		location := node.IsMember(program)
		if location != nil {
			return location
		}
	}
	return nil
}

package Node

import (
	. "github.com/MirahImage/AdventOfCode2017/Day7/ProgramData"
)

type Node struct {
	Data     ProgramData
	Children []*Node
	Parent   *Node
}

package tree

import (
	. "github.com/MirahImage/AdventOfCode2017/Day7/Node"
	. "github.com/MirahImage/AdventOfCode2017/Day7/ProgramData"
)

type Tree struct {
	Root     *Node
	Programs []Node
}

func (t *Tree) AddPrograms(progs []ProgramData) {
	t.Programs = make([]Node, len(progs))
	for i, prog := range progs {
		t.Programs[i].Data = prog
	}
	for i, nodei := range t.Programs {
		for j, nodej := range t.Programs {
			if nodei.Data.IsParentOf(nodej.Data) {
				t.Programs[i].AddChild(&t.Programs[j])
			}
			if nodej.Data.IsParentOf(nodei.Data) {
				t.Programs[i].Parent = &t.Programs[j]
			}
		}
		if t.Programs[i].Parent == nil {
			t.Root = &t.Programs[i]
		}
	}
}

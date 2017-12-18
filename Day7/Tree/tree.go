package tree

import (
	node "github.com/MirahImage/AdventOfCode2017/Day7/Node"
	. "github.com/MirahImage/AdventOfCode2017/Day7/ProgramData"
)

type Tree struct {
	Root *node.Node
}

func (t *Tree) Insert(program ProgramData) error {
	if t.Root == nil {
		t.Root = node.NewNode(program)
		return nil
	}
	var err error
	t.Root, err = t.Root.Insert(program)
	return err
}

//If given a list that won't create a valid tree, this is an infinite loop
func (t *Tree) InsertMultiple(progs []ProgramData) {
	for len(progs) > 0 {
		for i, prog := range progs {
			err := t.Insert(prog)
			if err == nil {
				progs[i] = progs[len(progs)-1]
				progs = progs[:len(progs)-1]
				break
			}
		}
	}
}

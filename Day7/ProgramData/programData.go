package programData

import (
	"strings"
)

type ProgramData struct {
	Name       string
	Weight     int
	ChildNames []string
}

func (p *ProgramData) IsParent(data ProgramData) bool {
	for _, child := range p.ChildNames {
		if strings.Compare(child, data.Name) == 0 {
			return true
		}
	}
	return false
}

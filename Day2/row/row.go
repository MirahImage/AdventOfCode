package row

type Row struct {
	Entries []int
}

func (r *Row) Difference() int {
	if len(r.Entries) == 0 {
		return 0
	}
	smallest := r.Entries[0]
	largest := r.Entries[0]
	for _, entry := range r.Entries {
		if entry < smallest {
			smallest = entry
		}
		if entry > largest {
			largest = entry
		}
	}
	return largest - smallest
}

func (r *Row) Divisible(index int) int {
	for i, entry := range r.Entries {
		if i != index && r.Entries[index]%entry == 0 {
			return r.Entries[index] / entry
		}
	}
	return 0
}

func (r *Row) Resultant() int {
	for i, _ := range r.Entries {
		if r.Divisible(i) != 0 {
			return r.Divisible(i)
		}
	}
	return 0
}

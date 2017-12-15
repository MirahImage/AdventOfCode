package memory

type Memory struct {
  Banks []int
}

func (m *Memory) Reallocate() {
  bank, max := 0, 0
  for i, blocks := range m.Banks {
    if blocks > max {
      bank, max = i, blocks
    }
  }
  m.Banks[bank] = 0
  bank++
  for max > 0 {
    m.Banks[bank%len(m.Banks)]++
    max--
    bank++
  }
}

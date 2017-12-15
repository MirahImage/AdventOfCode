package memory

type Memory struct {
  Banks []int
}

//Returns bank with most blocks, number of blocks in said bank
func (m *Memory) LargestBank() (int, int) {
  bank, max := 0, 0
  for i, blocks := range m.Banks {
    if blocks > max {
      bank, max = i, blocks
    }
  }
  return bank, max
}

func (m *Memory) Reallocate() {
  bank, max := m.LargestBank()
  m.Banks[bank] = 0
  bank++
  for max > 0 {
    m.Banks[bank%len(m.Banks)]++
    max--
    bank++
  }
}

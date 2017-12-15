package history

import (
  . "github.com/MirahImage/AdventOfCode2017/Day6/Memory"
  "reflect"
)

type Bank struct {
  Banks []int
}

type History struct {
  TimeSteps []Bank
}

//Logs a memory state, returns true if the memory state has been logged already
func (h *History) Log(m Memory) bool {
  toAppend := Bank{
    Banks: make([]int, len(m.Banks)),
  }
  copy(toAppend.Banks, m.Banks)
  for _, bank := range h.TimeSteps {
    if reflect.DeepEqual(bank, toAppend) {
      return true
    }
  }
  h.TimeSteps = append(h.TimeSteps, toAppend)
  return false
}

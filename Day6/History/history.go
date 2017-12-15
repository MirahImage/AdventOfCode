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

//Logs a memory state, returns index of the repeated memory state, or -1 if it has not been logged already
func (h *History) Log(m Memory) int {
  toAppend := Bank{
    Banks: make([]int, len(m.Banks)),
  }
  copy(toAppend.Banks, m.Banks)
  for i, bank := range h.TimeSteps {
    if reflect.DeepEqual(bank, toAppend) {
      return i
    }
  }
  h.TimeSteps = append(h.TimeSteps, toAppend)
  return -1
}

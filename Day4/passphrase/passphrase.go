package passphrase

import (
	"sort"
	"strings"
)

type Passphrase struct {
	Words []string
}

func (p *Passphrase) RepeatedString() bool {
	for i, word := range p.Words {
		for j := 0; j < i; j++ {
			if strings.Compare(word, p.Words[j]) == 0 {
				return true
			}
		}
	}
	return false
}

type runeSlice []rune

func (r runeSlice) Len() int {
	return len(r)
}

func (r runeSlice) Less(i, j int) bool {
	return r[i] < r[j]
}

func (r runeSlice) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func SortStrings(strings []string) []string {
	sorted := make([]string, len(strings))
	for i, s := range strings {
		runes := runeSlice(s)
		sort.Sort(runes)
		sorted[i] = string(runes)
	}
	return sorted
}

// RepeatedAnagram returns true if an anagram of a word is repeated
func (p *Passphrase) RepeatedAnagram() bool {
	sorted := SortStrings(p.Words)
	for i, word := range sorted {
		for j := 0; j < i; j++ {
			if strings.Compare(word, sorted[j]) == 0 {
				return true
			}
		}
	}
	return false
}

func (p *Passphrase) ValidRepeated() bool {
	return !p.RepeatedString()
}

func (p *Passphrase) ValidAnagram() bool {
	return !p.RepeatedAnagram()
}

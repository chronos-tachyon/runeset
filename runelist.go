package runeset

import (
	"fmt"
	"sort"
)

type RuneList []rune

func (list RuneList) Len() uint {
	return uint(len(list))
}

func (list RuneList) At(index uint) Pair {
	ch := list[index]
	return Pair{ch, ch}
}

func (list RuneList) Contains(ch rune) bool {
	for _, item := range list {
		if ch == item {
			return true
		}
	}
	return false
}

func (list RuneList) Append(out []byte) []byte {
	return appendSource(out, list)
}

func (list RuneList) String() string {
	return toString(list)
}

func (list RuneList) Sort() {
	sort.Slice(list, func(i int, j int) bool {
		a, b := list[i], list[j]
		return a < b
	})
}

var (
	_ Source       = RuneList(nil)
	_ Appender     = RuneList(nil)
	_ fmt.Stringer = RuneList(nil)
)

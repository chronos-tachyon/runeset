package runeset

import (
	"fmt"
	"sort"
)

type PairList []Pair

func (list PairList) Len() uint {
	return uint(len(list))
}

func (list PairList) At(index uint) Pair {
	return list[index]
}

func (list PairList) Contains(ch rune) bool {
	for _, pair := range list {
		if pair.Contains(ch) {
			return true
		}
	}
	return false
}

func (list PairList) Append(out []byte) []byte {
	return appendSource(out, list)
}

func (list PairList) String() string {
	return toString(list)
}

func (list PairList) Sort() {
	sort.Slice(list, func(i int, j int) bool {
		a, b := list[i], list[j]
		return a.LessThan(b)
	})
}

var (
	_ Source       = PairList(nil)
	_ Appender     = PairList(nil)
	_ fmt.Stringer = PairList(nil)
)

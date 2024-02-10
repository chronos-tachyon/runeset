package runeset

import (
	"fmt"
	"unicode"
)

type Set struct{ list []Pair }

func Make(sources ...Source) Set {
	var b Builder
	return b.Reset().Add(sources...).Build()
}

func Empty() Set {
	return Set{}
}

var gFull = [1]Pair{Pair{Lo: 0, Hi: unicode.MaxRune}}

func Full() Set {
	list := gFull[:]
	return Set{list: list}
}

func (set Set) Len() uint {
	return uint(len(set.list))
}

func (set Set) At(index uint) Pair {
	return set.list[index]
}

func (set Set) Search(ch rune) (uint, bool) {
	return searchSource(set, ch)
}

func (set Set) Contains(ch rune) bool {
	_, found := set.Search(ch)
	return found
}

func (set Set) IsEmpty() bool {
	return isEmpty(set)
}

func (set Set) IsFull() bool {
	return isFull(set)
}

func (set Set) Append(out []byte) []byte {
	return appendSource(out, set)
}

func (set Set) String() string {
	return toString(set)
}

func (set Set) Builder() *Builder {
	return NewBuilder().Add(set)
}

var (
	_ Source       = Set{}
	_ Appender     = Set{}
	_ fmt.Stringer = Set{}
)

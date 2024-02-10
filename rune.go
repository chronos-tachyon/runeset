package runeset

import (
	"cmp"
	"fmt"
)

type Rune rune

func (r Rune) IsValid() bool {
	return isValidRune(rune(r))
}

func (r Rune) AssertValid() {
	if r.IsValid() {
		return
	}
	panic(fmt.Errorf("BUG: U+%04X is not a valid Unicode code point", uint32(r)))
}

func (r Rune) Len() uint {
	return 1
}

func (r Rune) At(index uint) Pair {
	if index != 0 {
		panic(fmt.Errorf("index out of range: %d != 0", index))
	}
	r.AssertValid()
	ch := rune(r)
	return Pair{ch, ch}
}

func (r Rune) Contains(ch rune) bool {
	return ch == rune(r)
}

func (r Rune) Append(out []byte) []byte {
	return appendSource(out, r)
}

func (r Rune) String() string {
	return toString(r)
}

func (r Rune) CompareTo(other Rune) int {
	return cmp.Compare(r, other)
}

func (r Rune) LessThan(other Rune) bool {
	return r.CompareTo(other) < 0
}

func (r Rune) EqualTo(other Rune) bool {
	return r == other
}

var (
	_ Source       = Rune(0)
	_ Appender     = Rune(0)
	_ fmt.Stringer = Rune(0)
)

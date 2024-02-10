package runeset

import (
	"cmp"
	"fmt"
)

type Pair struct {
	Lo rune
	Hi rune
}

func MakePair(lo rune, hi rune) Pair {
	pair := Pair{Lo: lo, Hi: hi}
	pair.AssertValid()
	return pair
}

func (pair Pair) IsValid() bool {
	return isValidPair(pair.Lo, pair.Hi)
}

func (pair Pair) AssertValid() {
	if pair.IsValid() {
		return
	}
	if !isValidRune(pair.Lo) {
		panic(fmt.Errorf("BUG: in %v, lower bound is not a valid Unicode code point", pair))
	}
	if !isValidRune(pair.Hi) {
		panic(fmt.Errorf("BUG: in %v, upper bound is not a valid Unicode code point", pair))
	}
	lo := uint32(pair.Lo)
	hi := uint32(pair.Hi)
	panic(fmt.Errorf("BUG: in %v, lower bound U+%04X exceeds upper bound U+%04X", pair, lo, hi))
}

func (pair Pair) Len() uint {
	return 1
}

func (pair Pair) At(index uint) Pair {
	if index != 0 {
		panic(fmt.Errorf("index out of range: %d != 0", index))
	}
	pair.AssertValid()
	return pair
}

func (pair Pair) Contains(ch rune) bool {
	return ch >= pair.Lo && ch <= pair.Hi
}

func (pair Pair) Append(out []byte) []byte {
	return appendSource(out, pair)
}

func (pair Pair) String() string {
	return toString(pair)
}

func (pair Pair) CompareTo(other Pair) int {
	c := cmp.Compare(pair.Lo, other.Lo)
	if c == 0 {
		c = cmp.Compare(pair.Hi, other.Hi)
	}
	return c
}

func (pair Pair) LessThan(other Pair) bool {
	return pair.CompareTo(other) < 0
}

func (pair Pair) EqualTo(other Pair) bool {
	return pair == other
}

var (
	_ Source       = Pair{}
	_ Appender     = Pair{}
	_ fmt.Stringer = Pair{}
)

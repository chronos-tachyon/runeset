package runeset

import (
	"fmt"
	"unicode"
)

type Builder struct {
	l []Pair
	s [16]Pair
}

func NewBuilder() *Builder {
	b := new(Builder)
	return b.Reset()
}

func (b *Builder) assertNotNil() {
	if b == nil {
		panic("Builder is nil")
	}
}

func (b *Builder) Reset() *Builder {
	b.l = b.s[:0]
	clear(b.s[:])
	return b
}

func (b *Builder) Len() uint {
	if b == nil {
		return 0
	}
	return uint(len(b.l))
}

func (b *Builder) At(index uint) Pair {
	return b.l[index]
}

func (b *Builder) Search(ch rune) (uint, bool) {
	return searchSource(b, ch)
}

func (b *Builder) Contains(ch rune) bool {
	_, found := b.Search(ch)
	return found
}

func (b *Builder) IsEmpty() bool {
	return isEmpty(b)
}

func (b *Builder) IsFull() bool {
	return isFull(b)
}

func (b *Builder) Append(out []byte) []byte {
	if b == nil {
		return append(out, "<nil>"...)
	}
	return appendSource(out, b)
}

func (b *Builder) String() string {
	if b == nil {
		return "<nil>"
	}
	return toString(b)
}

func (b *Builder) Negate() *Builder {
	b.assertNotNil()
	in := b.l
	n := uint(len(in))
	out := make([]Pair, n+1)
	out[0].Lo = 0
	if n > 0 {
		out[0].Hi = in[0].Lo - 1
		for i := uint(1); i < n; i++ {
			out[i].Lo = in[i-1].Hi + 1
			out[i].Hi = in[i].Lo - 1
		}
		out[n].Lo = in[n-1].Hi + 1
	}
	out[n].Hi = unicode.MaxRune
	b.update(out)
	return b
}

func (b *Builder) Add(sources ...Source) *Builder {
	b.assertNotNil()
	if len(sources) <= 0 {
		return b
	}

	list := cloneList(b.l)
	for _, src := range sources {
		srcLen := src.Len()
		for i := uint(0); i < srcLen; i++ {
			pair := src.At(i)
			pair.AssertValid()
			list = append(list, pair)
		}
	}
	b.update(list)
	return b
}

func (b *Builder) AddPairs(list []Pair) *Builder {
	return b.Add(PairList(list))
}

func (b *Builder) AddPair(pairs ...Pair) *Builder {
	return b.AddPairs(pairs)
}

func (b *Builder) AddRange(lo rune, hi rune) *Builder {
	return b.Add(Pair{lo, hi})
}

func (b *Builder) AddRune(runes ...rune) *Builder {
	return b.Add(RuneList(runes))
}

func (b *Builder) Remove(sources ...Source) *Builder {
	b.assertNotNil()
	if len(sources) <= 0 {
		return b
	}

	list := b.l
	k := 2 * uint(len(list))
	in := make([]Pair, 0, k)
	out := make([]Pair, 0, k)

	in = append(in, list...)
	fn := func(pair Pair) { out = append(out, pair) }

	for _, src := range sources {
		srcLen := src.Len()
		for i := uint(0); i < srcLen; i++ {
			remove := src.At(i)
			remove.AssertValid()
			for _, pair := range in {
				segment(pair, remove, fn, nil, fn)
			}
			in, out = out, in[:0]
		}
	}

	b.update(in)
	return b
}

func (b *Builder) RemovePairs(list []Pair) *Builder {
	return b.Remove(PairList(list))
}

func (b *Builder) RemovePair(pairs ...Pair) *Builder {
	return b.RemovePairs(pairs)
}

func (b *Builder) RemoveRange(lo rune, hi rune) *Builder {
	return b.Remove(Pair{lo, hi})
}

func (b *Builder) RemoveRune(runes ...rune) *Builder {
	return b.Remove(RuneList(runes))
}

func (b *Builder) Intersect(sources ...Source) *Builder {
	b.assertNotNil()
	if len(sources) <= 0 {
		return b
	}

	list := b.l
	k := 2 * uint(len(list))
	in := make([]Pair, 0, k)
	out := make([]Pair, 0, k)

	in = append(in, list...)
	fn := func(pair Pair) { out = append(out, pair) }

	for _, src := range sources {
		srcLen := src.Len()
		for i := uint(0); i < srcLen; i++ {
			keep := src.At(i)
			keep.AssertValid()
			for _, pair := range in {
				segment(pair, keep, nil, fn, nil)
			}
		}
		in, out = out, in[:0]
	}

	b.update(in)
	return b
}

func (b *Builder) IntersectPairs(list []Pair) *Builder {
	return b.Intersect(PairList(list))
}

func (b *Builder) IntersectPair(pairs ...Pair) *Builder {
	return b.IntersectPairs(pairs)
}

func (b *Builder) IntersectRange(lo rune, hi rune) *Builder {
	return b.Intersect(Pair{lo, hi})
}

func (b *Builder) IntersectRune(runes ...rune) *Builder {
	return b.Intersect(RuneList(runes))
}

func (b *Builder) update(list []Pair) {
	// phase one: remove invalid pairs
	i := uint(0)
	for i < uint(len(list)) {
		pair := list[i]
		if !pair.IsValid() {
			list = splice(list, i, i+1)
			continue
		}
		i++
	}

	// verify
	for i = 0; i < uint(len(list)); i++ {
		list[i].AssertValid()
	}

	// phase two: sort
	PairList(list).Sort()

	// phase three: merge overlapping/adjacent
	i = 1
	for i < uint(len(list)) {
		prev := list[i-1]
		curr := list[i]
		if prev.Hi == unicode.MaxRune {
			n := uint(len(list))
			list = splice(list, i, n)
			continue
		}
		if prev.Hi >= curr.Hi {
			list = splice(list, i, i+1)
			continue
		}
		if (prev.Hi + 1) >= curr.Lo {
			prev.Hi = curr.Hi
			list[i-1] = prev
			list = splice(list, i, i+1)
			continue
		}
		i++
	}

	// verify
	for i = 1; i < uint(len(list)); i++ {
		curr := list[i]
		prev := list[i-1]
		if prev.Lo >= curr.Lo {
			panic(fmt.Errorf("BUG: [%d] %v >= [%d] %v", i-1, prev, i, curr))
		}
		if prev.Hi >= (curr.Lo + 1) {
			panic(fmt.Errorf("BUG: [%d] %v touches [%d] %v and should have been merged with it", i-1, prev, i, curr))
		}
	}

	// commit
	b.l = cloneListInto(b.s[:], list)
}

func (b *Builder) Clone() *Builder {
	return NewBuilder().Add(b)
}

func (b *Builder) Build() Set {
	list := cloneList(b.l)
	return Set{list}
}

var (
	_ Source       = (*Builder)(nil)
	_ Appender     = (*Builder)(nil)
	_ fmt.Stringer = (*Builder)(nil)
)

package runeset

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

type Appender interface {
	Append(out []byte) []byte
}

func toString(a Appender) string {
	var scratch [64]byte
	return string(a.Append(scratch[:0]))
}

func isValidRune(ch rune) bool {
	return ch >= 0 && ch <= unicode.MaxRune
}

func isValidPair(lo rune, hi rune) bool {
	return isValidRune(lo) && isValidRune(hi) && (lo <= hi)
}

func cloneList[T any](in []T) []T {
	if len(in) <= 0 {
		return nil
	}
	out := make([]T, len(in))
	copy(out, in)
	return out
}

func cloneListInto[T any](storage []T, in []T) []T {
	if len(in) <= 0 {
		return nil
	}
	if len(in) <= cap(storage) {
		return append(storage[:0], in...)
	}
	out := make([]T, len(in))
	copy(out, in)
	return out
}

func splice[T any](in []T, i uint, j uint, insert ...T) []T {
	inLen := uint(len(in))
	inCap := uint(cap(in))
	if i > j {
		panic(fmt.Errorf("BUG: i > j [i=%d j=%d n=%d]", i, j, inLen))
	}
	if j > inLen {
		panic(fmt.Errorf("BUG: j > n [i=%d j=%d n=%d]", i, j, inLen))
	}
	head := in[:i]
	tail := in[j:]
	headLen := i
	deleteLen := j - i
	insertLen := uint(len(insert))
	tailLen := inLen - j
	outLen := headLen + insertLen + tailLen
	out := head
	switch {
	case outLen > inCap:
		out = make([]T, headLen, outLen)
		copy(out, head)
	case insertLen > deleteLen && tailLen > 0:
		tmp := make([]T, tailLen)
		copy(tmp, tail)
		tail = tmp
	}
	out = append(out, insert...)
	out = append(out, tail...)
	return out
}

// segment divides p into 3 or fewer pieces:
//
// - the part of p that lies strictly before q
//
// - the part of p that overlaps q
//
// - the part of p that lies strictly after q
func segment(p Pair, q Pair, x, y, z func(Pair)) {
	p.AssertValid()
	q.AssertValid()

	a, b := p.Lo, p.Hi
	c, d := q.Lo, q.Hi

	dummy := func(Pair) {}
	if x == nil {
		x = dummy
	}
	if y == nil {
		y = dummy
	}
	if z == nil {
		z = dummy
	}

	// Of the 24 permutations
	// 18 are eliminated by a <= b && c <= d
	// leaving these 6 cases to handle:
	switch {
	case b < c:
		// a b c d
		x(Pair{a, b})
	case a < c && d >= b:
		// a c b d
		x(Pair{a, c - 1})
		y(Pair{c, b})
	case a < c:
		// a c d b
		x(Pair{a, c - 1})
		y(Pair{c, d})
		z(Pair{d + 1, b})
	case d < a:
		// c d a b
		z(Pair{a, b})
	case d < b:
		// c a d b
		y(Pair{a, d})
		z(Pair{d + 1, b})
	default:
		// c a b d
		y(Pair{a, b})
	}
}

func isEmpty[S Source](src S) bool {
	return src.Len() <= 0
}

func isFull[S Source](src S) bool {
	if src.Len() == 1 {
		p := src.At(0)
		if p.Lo == 0 && p.Hi == unicode.MaxRune {
			return true
		}
	}
	return false
}

func appendSource[S Source](out []byte, src S) []byte {
	switch {
	case isEmpty(src):
		out = append(out, '!', '.')

	case isFull(src):
		out = append(out, '.')

	default:
		srcLen := src.Len()
		out = append(out, '[')
		for i := uint(0); i < srcLen; i++ {
			p := src.At(i)
			out = appendPair(out, p.Lo, p.Hi)
		}
		out = append(out, ']')
	}
	return out
}

func appendPair(out []byte, lo rune, hi rune) []byte {
	out = appendRune(out, lo)
	if lo != hi {
		out = append(out, '-')
		out = appendRune(out, hi)
	}
	return out
}

func appendRune(out []byte, ch rune) []byte {
	u32 := uint32(ch)
	switch {
	case ch >= '0' && ch <= '9':
		fallthrough
	case ch >= 'A' && ch <= 'Z':
		fallthrough
	case ch >= 'a' && ch <= 'z':
		fallthrough
	case unicode.IsLetter(ch):
		fallthrough
	case unicode.IsDigit(ch):
		return utf8.AppendRune(out, ch)

	case ch == 0:
		return append(out, '\\', '0')
	case ch == '\t':
		return append(out, '\\', 't')
	case ch == '\n':
		return append(out, '\\', 'n')
	case ch == '\v':
		return append(out, '\\', 'v')
	case ch == '\f':
		return append(out, '\\', 'f')
	case ch == '\r':
		return append(out, '\\', 'r')
	case ch == unicode.MaxRune:
		return append(out, '\\', 'z')

	case u32 < 0x100:
		return fmt.Appendf(out, "\\x%02x", u32)
	case u32 < 0x10000:
		return fmt.Appendf(out, "\\u%04x", u32)
	default:
		return fmt.Appendf(out, "\\u{%x}", u32)
	}
}

func searchSource[S Source](src S, ch rune) (uint, bool) {
	i := uint(0)
	j := src.Len()
	for i < j {
		k := (i + j) >> 1
		pair := src.At(k)
		switch {
		case ch < pair.Lo:
			j = k
		case ch > pair.Hi:
			i = k + 1
		default:
			return k, true
		}
	}
	return i, false
}

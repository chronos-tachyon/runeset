package runeset

import (
	"fmt"
	"unicode"
)

var classMap map[string]Set

func ForClass(className string) Set {
	if set, found := classMap[className]; found {
		return set
	}
	panic(fmt.Errorf("unknown character class %q", className))
}

func ForTable(table *unicode.RangeTable) Set {
	if table == nil {
		return Empty()
	}
	out := make(PairList, 0, len(table.R16)+len(table.R32))
	fn := func(lo uint32, hi uint32, stride uint32) {
		const maxRune = 0x10ffff
		switch {
		case lo > hi:
			// pass
		case hi > maxRune:
			// pass
		case stride == 1:
			rLo := rune(lo)
			rHi := rune(hi)
			out = append(out, Pair{rLo, rHi})
		default:
			for lo <= hi {
				rLo := rune(lo)
				out = append(out, Pair{rLo, rLo})
				lo += stride
			}
		}
	}
	for _, r16 := range table.R16 {
		fn(uint32(r16.Lo), uint32(r16.Hi), uint32(r16.Stride))
	}
	for _, r32 := range table.R32 {
		fn(r32.Lo, r32.Hi, r32.Stride)
	}
	out.Sort()
	var b Builder
	return b.Reset().Add(out).Build()
}

func init() {
	classMap = make(map[string]Set, 64)
	for name, table := range unicode.Categories {
		classMap[name] = ForTable(table)
	}

	classMap["ascii"] = Make(Pair{0x00, 0x7f})
	classMap["ascii.graph"] = Make(Pair{0x21, 0x7e})
	classMap["ascii.print"] = Make(Pair{0x20, 0x7e})
	classMap["ascii.cntrl"] = Make(Pair{0x00, 0x1f}, Pair{0x7f, 0x7f})
	classMap["ascii.upper"] = Make(Pair{'A', 'Z'})
	classMap["ascii.lower"] = Make(Pair{'a', 'z'})
	classMap["ascii.alpha"] = Make(Pair{'A', 'Z'}, Pair{'a', 'z'})
	classMap["ascii.digit"] = Make(Pair{'0', '9'})
	classMap["ascii.bdigit"] = Make(Pair{'0', '1'})
	classMap["ascii.odigit"] = Make(Pair{'0', '7'})
	classMap["ascii.xdigit"] = Make(Pair{'0', '9'}, Pair{'A', 'F'}, Pair{'a', 'f'})
	classMap["ascii.alnum"] = Make(Pair{'0', '9'}, Pair{'A', 'Z'}, Pair{'a', 'z'})
	classMap["ascii.word"] = Make(ForClass("ascii.alnum"), Pair{'_', '_'})
	classMap["ascii.punct"] = NewBuilder().Add(ForClass("ascii.graph")).Remove(ForClass("ascii.alnum")).Build()
	classMap["ascii.blank"] = Make(Pair{'\t', '\t'}, Pair{' ', ' '})
	classMap["ascii.space"] = Make(Pair{'\t', '\r'}, Pair{' ', ' '})

	classMap["cntrl"] = ForClass("Cc")
	classMap["upper"] = ForClass("Lu")
	classMap["title"] = ForClass("Lt")
	classMap["lower"] = ForClass("Ll")
	classMap["letter"] = ForClass("L")
	classMap["alpha"] = ForClass("L")
	classMap["digit"] = ForClass("Nd")
	classMap["bdigit"] = ForClass("ascii.bdigit")
	classMap["odigit"] = ForClass("ascii.odigit")
	classMap["xdigit"] = ForClass("ascii.xdigit")
	classMap["number"] = ForClass("N")
	classMap["alnum"] = Make(ForClass("L"), ForClass("N"))
	classMap["word"] = Make(ForClass("alnum"), Pair{'_', '_'})
	classMap["punct"] = ForClass("P")
	classMap["symbol"] = ForClass("S")
	classMap["mark"] = ForClass("M")
	classMap["blank"] = ForClass("ascii.blank")
	classMap["space"] = Make(ForClass("ascii.space"), ForClass("Z"))
	classMap["graph"] = Make(ForClass("L"), ForClass("M"), ForClass("N"), ForClass("P"), ForClass("S"))
	classMap["print"] = Make(ForClass("graph"), ForClass("Zs"))
}

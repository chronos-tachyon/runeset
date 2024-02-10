package runeset

import (
	"testing"
	"unicode"
)

func TestSet(t *testing.T) {
	type testRow struct {
		Name   string
		Input  Set
		Expect string
	}

	var b Builder
	testData := [...]testRow{
		{
			Name:   "Empty",
			Input:  Empty(),
			Expect: `!.`,
		},
		{
			Name:   "Full",
			Input:  Full(),
			Expect: `.`,
		},
		{
			Name:   "Base-01",
			Input:  b.Reset().Build(),
			Expect: `!.`,
		},
		{
			Name:   "Base-02",
			Input:  b.Reset().AddRange(0, unicode.MaxRune).Build(),
			Expect: `.`,
		},
		{
			Name:   "Base-03",
			Input:  b.Reset().AddRange('a', 'c').Build(),
			Expect: `[a-c]`,
		},
		{
			Name:   "Base-04",
			Input:  b.Reset().AddRange('a', 'c').AddRange('g', 'i').AddRange('m', 'o').Build(),
			Expect: `[a-cg-im-o]`,
		},
		{
			Name:   "Base-05",
			Input:  b.Reset().AddRange('a', unicode.MaxRune).Build(),
			Expect: `[a-\z]`,
		},
		{
			Name:   "Base-06",
			Input:  b.Reset().AddRange('a', 'c').AddRange('g', 'i').AddRange('m', unicode.MaxRune).Build(),
			Expect: `[a-cg-im-\z]`,
		},

		{
			Name:   "AddRange-AZAZ",
			Input:  b.Reset().AddRange('a', 'z').AddRange('a', 'z').Build(),
			Expect: `[a-z]`,
		},
		{
			Name:   "AddRange-AZNZ",
			Input:  b.Reset().AddRange('a', 'z').AddRange('n', 'z').Build(),
			Expect: `[a-z]`,
		},
		{
			Name:   "AddRange-AMAZ",
			Input:  b.Reset().AddRange('a', 'm').AddRange('a', 'z').Build(),
			Expect: `[a-z]`,
		},
		{
			Name:   "AddRange-AMNZ",
			Input:  b.Reset().AddRange('a', 'm').AddRange('n', 'z').Build(),
			Expect: `[a-z]`,
		},
		{
			Name:   "AddRange-AFNW",
			Input:  b.Reset().AddRange('a', 'f').AddRange('n', 'w').Build(),
			Expect: `[a-fn-w]`,
		},

		{
			Name:   "AddRange-AFIKNW-123",
			Input:  b.Reset().AddRange('a', 'f').AddRange('i', 'k').AddRange('n', 'w').Build(),
			Expect: `[a-fi-kn-w]`,
		},
		{
			Name:   "AddRAnge-AFIKNW-132",
			Input:  b.Reset().AddRange('a', 'f').AddRange('n', 'w').AddRange('i', 'k').Build(),
			Expect: `[a-fi-kn-w]`,
		},
		{
			Name:   "AddRAnge-AFIKNW-213",
			Input:  b.Reset().AddRange('i', 'k').AddRange('a', 'f').AddRange('n', 'w').Build(),
			Expect: `[a-fi-kn-w]`,
		},
		{
			Name:   "AddRAnge-AFIKNW-231",
			Input:  b.Reset().AddRange('i', 'k').AddRange('n', 'w').AddRange('a', 'f').Build(),
			Expect: `[a-fi-kn-w]`,
		},
		{
			Name:   "AddRAnge-AFIKNW-312",
			Input:  b.Reset().AddRange('n', 'w').AddRange('a', 'f').AddRange('i', 'k').Build(),
			Expect: `[a-fi-kn-w]`,
		},
		{
			Name:   "AddRAnge-AFIKNW-321",
			Input:  b.Reset().AddRange('n', 'w').AddRange('i', 'k').AddRange('a', 'f').Build(),
			Expect: `[a-fi-kn-w]`,
		},

		{
			Name:   "AddRange-AFGMNW-123",
			Input:  b.Reset().AddRange('a', 'f').AddRange('g', 'm').AddRange('n', 'w').Build(),
			Expect: `[a-w]`,
		},
		{
			Name:   "AddRange-AFGMNW-132",
			Input:  b.Reset().AddRange('a', 'f').AddRange('n', 'w').AddRange('g', 'm').Build(),
			Expect: `[a-w]`,
		},
		{
			Name:   "AddRange-AFGMNW-213",
			Input:  b.Reset().AddRange('g', 'm').AddRange('a', 'f').AddRange('n', 'w').Build(),
			Expect: `[a-w]`,
		},
		{
			Name:   "AddRange-AFGMNW-231",
			Input:  b.Reset().AddRange('g', 'm').AddRange('n', 'w').AddRange('a', 'f').Build(),
			Expect: `[a-w]`,
		},
		{
			Name:   "AddRange-AFGMNW-312",
			Input:  b.Reset().AddRange('n', 'w').AddRange('a', 'f').AddRange('g', 'm').Build(),
			Expect: `[a-w]`,
		},
		{
			Name:   "AddRange-AFGMNW-321",
			Input:  b.Reset().AddRange('n', 'w').AddRange('g', 'm').AddRange('a', 'f').Build(),
			Expect: `[a-w]`,
		},

		{
			Name:   "AddRange-AFGINW-123",
			Input:  b.Reset().AddRange('a', 'f').AddRange('g', 'i').AddRange('n', 'w').Build(),
			Expect: `[a-in-w]`,
		},
		{
			Name:   "AddRange-AFGINW-132",
			Input:  b.Reset().AddRange('a', 'f').AddRange('n', 'w').AddRange('g', 'i').Build(),
			Expect: `[a-in-w]`,
		},
		{
			Name:   "AddRange-AFGINW-213",
			Input:  b.Reset().AddRange('g', 'i').AddRange('a', 'f').AddRange('n', 'w').Build(),
			Expect: `[a-in-w]`,
		},
		{
			Name:   "AddRange-AFGINW-231",
			Input:  b.Reset().AddRange('g', 'i').AddRange('n', 'w').AddRange('a', 'f').Build(),
			Expect: `[a-in-w]`,
		},
		{
			Name:   "AddRange-AFGINW-312",
			Input:  b.Reset().AddRange('n', 'w').AddRange('a', 'f').AddRange('g', 'i').Build(),
			Expect: `[a-in-w]`,
		},
		{
			Name:   "AddRange-AFGINW-321",
			Input:  b.Reset().AddRange('n', 'w').AddRange('g', 'i').AddRange('a', 'f').Build(),
			Expect: `[a-in-w]`,
		},

		{
			Name:   "AddRange-AFKMNW-123",
			Input:  b.Reset().AddRange('a', 'f').AddRange('k', 'm').AddRange('n', 'w').Build(),
			Expect: `[a-fk-w]`,
		},
		{
			Name:   "AddRange-AFKMNW-132",
			Input:  b.Reset().AddRange('a', 'f').AddRange('n', 'w').AddRange('k', 'm').Build(),
			Expect: `[a-fk-w]`,
		},
		{
			Name:   "AddRange-AFKMNW-213",
			Input:  b.Reset().AddRange('k', 'm').AddRange('a', 'f').AddRange('n', 'w').Build(),
			Expect: `[a-fk-w]`,
		},
		{
			Name:   "AddRange-AFKMNW-231",
			Input:  b.Reset().AddRange('k', 'm').AddRange('n', 'w').AddRange('a', 'f').Build(),
			Expect: `[a-fk-w]`,
		},
		{
			Name:   "AddRange-AFKMNW-312",
			Input:  b.Reset().AddRange('n', 'w').AddRange('a', 'f').AddRange('k', 'm').Build(),
			Expect: `[a-fk-w]`,
		},
		{
			Name:   "AddRange-AFKMNW-321",
			Input:  b.Reset().AddRange('n', 'w').AddRange('k', 'm').AddRange('a', 'f').Build(),
			Expect: `[a-fk-w]`,
		},

		{
			Name:   "AddRange-ACGIMO",
			Input:  b.Reset().AddRange('a', 'c').AddRange('g', 'i').AddRange('m', 'o').Build(),
			Expect: `[a-cg-im-o]`,
		},
		{
			Name:   "AddRange-AXGIM.",
			Input:  b.Reset().AddRange('a', 'c').AddRange('g', 'i').AddRange('m', unicode.MaxRune).Build(),
			Expect: `[a-cg-im-\z]`,
		},

		{
			Name:   "AddRange-ACGIMOG.",
			Input:  b.Reset().AddRange('a', 'c').AddRange('g', 'i').AddRange('m', 'o').AddRange('g', unicode.MaxRune).Build(),
			Expect: `[a-cg-\z]`,
		},
		{
			Name:   "AddRange-ACGIM.G.",
			Input:  b.Reset().AddRange('a', 'c').AddRange('g', 'i').AddRange('m', unicode.MaxRune).AddRange('g', unicode.MaxRune).Build(),
			Expect: `[a-cg-\z]`,
		},
		{
			Name:   "AddRange-ACGIMOJ.",
			Input:  b.Reset().AddRange('a', 'c').AddRange('g', 'i').AddRange('m', 'o').AddRange('j', unicode.MaxRune).Build(),
			Expect: `[a-cg-\z]`,
		},
		{
			Name:   "AddRange-ACGIM.J.",
			Input:  b.Reset().AddRange('a', 'c').AddRange('g', 'i').AddRange('m', unicode.MaxRune).AddRange('j', unicode.MaxRune).Build(),
			Expect: `[a-cg-\z]`,
		},
		{
			Name:   "AddRange-ACGIMOK.",
			Input:  b.Reset().AddRange('a', 'c').AddRange('g', 'i').AddRange('m', 'o').AddRange('k', unicode.MaxRune).Build(),
			Expect: `[a-cg-ik-\z]`,
		},
		{
			Name:   "AddRange-ACGIM.K.",
			Input:  b.Reset().AddRange('a', 'c').AddRange('g', 'i').AddRange('m', unicode.MaxRune).AddRange('k', unicode.MaxRune).Build(),
			Expect: `[a-cg-ik-\z]`,
		},
		{
			Name:   "AddRange-ACGIMOM.",
			Input:  b.Reset().AddRange('a', 'c').AddRange('g', 'i').AddRange('m', 'o').AddRange('m', unicode.MaxRune).Build(),
			Expect: `[a-cg-im-\z]`,
		},
		{
			Name:   "AddRange-ACGIM.M.",
			Input:  b.Reset().AddRange('a', 'c').AddRange('g', 'i').AddRange('m', unicode.MaxRune).AddRange('m', unicode.MaxRune).Build(),
			Expect: `[a-cg-im-\z]`,
		},
		{
			Name:   "AddRange-ACGIMOP.",
			Input:  b.Reset().AddRange('a', 'c').AddRange('g', 'i').AddRange('m', 'o').AddRange('p', unicode.MaxRune).Build(),
			Expect: `[a-cg-im-\z]`,
		},
		{
			Name:   "AddRange-ACGIM.P.",
			Input:  b.Reset().AddRange('a', 'c').AddRange('g', 'i').AddRange('m', unicode.MaxRune).AddRange('p', unicode.MaxRune).Build(),
			Expect: `[a-cg-im-\z]`,
		},
		{
			Name:   "AddRange-ACGIMOQ.",
			Input:  b.Reset().AddRange('a', 'c').AddRange('g', 'i').AddRange('m', 'o').AddRange('q', unicode.MaxRune).Build(),
			Expect: `[a-cg-im-oq-\z]`,
		},
		{
			Name:   "AddRange-ACGIM.Q.",
			Input:  b.Reset().AddRange('a', 'c').AddRange('g', 'i').AddRange('m', unicode.MaxRune).AddRange('q', unicode.MaxRune).Build(),
			Expect: `[a-cg-im-\z]`,
		},

		{
			Name:   "Negate-01",
			Input:  b.Reset().Negate().Build(),
			Expect: `.`,
		},
		{
			Name:   "Negate-02",
			Input:  b.Reset().Add(Full()).Negate().Build(),
			Expect: `!.`,
		},
		{
			Name:   "Negate-03",
			Input:  b.Reset().AddRange('A', 'Z').Negate().Build(),
			Expect: `[\0-\x40\x5b-\z]`,
		},

		{
			Name:   "RemoveRange-ACGIMOG.",
			Input:  b.Reset().AddRange('a', 'c').AddRange('g', 'i').AddRange('m', 'o').RemoveRange('g', unicode.MaxRune).Build(),
			Expect: `[a-c]`,
		},
		{
			Name:   "RemoveRange-ACGIMOH.",
			Input:  b.Reset().AddRange('a', 'c').AddRange('g', 'i').AddRange('m', 'o').RemoveRange('h', unicode.MaxRune).Build(),
			Expect: `[a-cg]`,
		},
		{
			Name:   "RemoveRange-ACGIMOJ.",
			Input:  b.Reset().AddRange('a', 'c').AddRange('g', 'i').AddRange('m', 'o').RemoveRange('j', unicode.MaxRune).Build(),
			Expect: `[a-cg-i]`,
		},
		{
			Name:   "RemoveRange-ACGIMOK.",
			Input:  b.Reset().AddRange('a', 'c').AddRange('g', 'i').AddRange('m', 'o').RemoveRange('k', unicode.MaxRune).Build(),
			Expect: `[a-cg-i]`,
		},
		{
			Name:   "RemoveRange-ACGIMOM.",
			Input:  b.Reset().AddRange('a', 'c').AddRange('g', 'i').AddRange('m', 'o').RemoveRange('m', unicode.MaxRune).Build(),
			Expect: `[a-cg-i]`,
		},
		{
			Name:   "RemoveRange-ACGIMON.",
			Input:  b.Reset().AddRange('a', 'c').AddRange('g', 'i').AddRange('m', 'o').RemoveRange('n', unicode.MaxRune).Build(),
			Expect: `[a-cg-im]`,
		},
		{
			Name:   "RemoveRange-ACGIMOP.",
			Input:  b.Reset().AddRange('a', 'c').AddRange('g', 'i').AddRange('m', 'o').RemoveRange('p', unicode.MaxRune).Build(),
			Expect: `[a-cg-im-o]`,
		},
		{
			Name:   "RemoveRange-ACGIMOQ.",
			Input:  b.Reset().AddRange('a', 'c').AddRange('g', 'i').AddRange('m', 'o').RemoveRange('q', unicode.MaxRune).Build(),
			Expect: `[a-cg-im-o]`,
		},

		{
			Name:   "RemoveRange-0909",
			Input:  b.Reset().AddRange('0', '9').RemoveRange('0', '9').Build(),
			Expect: `!.`,
		},
		{
			Name:   "RemoveRange-0907",
			Input:  b.Reset().AddRange('0', '9').RemoveRange('0', '7').Build(),
			Expect: `[8-9]`,
		},
		{
			Name:   "RemoveRange-0937",
			Input:  b.Reset().AddRange('0', '9').RemoveRange('3', '7').Build(),
			Expect: `[0-28-9]`,
		},
		{
			Name:   "RemoveRange-0989",
			Input:  b.Reset().AddRange('0', '9').RemoveRange('8', '9').Build(),
			Expect: `[0-7]`,
		},

		{
			Name:   "AddRange-09AZaz",
			Input:  b.Reset().AddRange('0', '9').AddRange('A', 'Z').AddRange('a', 'z').Build(),
			Expect: `[0-9A-Za-z]`,
		},
		{
			Name:   "AddRange-azAZ09",
			Input:  b.Reset().AddRange('a', 'z').AddRange('A', 'Z').AddRange('0', '9').Build(),
			Expect: `[0-9A-Za-z]`,
		},
		{
			Name:   "AddRange-AZaz09",
			Input:  b.Reset().AddRange('A', 'Z').AddRange('a', 'z').AddRange('0', '9').Build(),
			Expect: `[0-9A-Za-z]`,
		},
		{
			Name:   "Punct",
			Input:  b.Reset().AddRange(0x20, 0x7e).RemoveRange('0', '9').RemoveRange('A', 'Z').RemoveRange('a', 'z').Build(),
			Expect: `[\x20-\x2f\x3a-\x40\x5b-\x60\x7b-\x7e]`,
		},
		{
			Name:   "AddRange",
			Input:  b.Reset().AddRange('0', '9').Build(),
			Expect: `[0-9]`,
		},
	}

	for _, row := range testData {
		t.Run(row.Name, func(t *testing.T) {
			actual := row.Input.String()
			if expect := row.Expect; actual != expect {
				t.Errorf("wrong set:\n\texpect: %q\n\tactual: %q", expect, actual)
			}
		})
	}
}

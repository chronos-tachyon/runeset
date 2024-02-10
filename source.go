package runeset

type Source interface {
	Len() uint
	At(uint) Pair
	Contains(rune) bool
}

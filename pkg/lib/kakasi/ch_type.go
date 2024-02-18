package kakasi

const (
	chKanji chType = iota + 1
	chKana
	chHiragana
	chSymbol
	chAlpha
)

type chType int

package properties

// Ch is a set of characters.
// It is used to convert symbols to alphabet.
// It is also used to convert Greek characters to alphabet.
// It is also used to convert Zenkaku characters to alphabet.
var Ch = ch{}

type ch struct{}

func (ch) Space() rune                    { return 0x20 }
func (ch) AtMark() rune                   { return 0x40 }
func (ch) AlphabetA() rune                { return 0x41 }
func (ch) AlphabetZ() rune                { return 0x5A }
func (ch) SquareBra() rune                { return 0x5B }
func (ch) BackQuote() rune                { return 0x60 }
func (ch) Alphabet_a() rune               { return 0x61 }
func (ch) Alphabet_z() rune               { return 0x7A }
func (ch) BracketBra() rune               { return 0x7B }
func (ch) Tilda() rune                    { return 0x7E }
func (ch) Delete() rune                   { return 0x7F }
func (ch) Latin1InvertedExclam() rune     { return 0x00A1 }
func (ch) Latin1YDiaeresis() rune         { return 0x00FF }
func (ch) IdeographicSpace() rune         { return 0x3000 }
func (ch) PostalMarkFace() rune           { return 0x3020 }
func (ch) WavyDash() rune                 { return 0x3030 }
func (ch) IdeographicHalfFillSpace() rune { return 0x303F }
func (ch) GreeceAlpha() rune              { return 0x0391 }
func (ch) GreeceRho() rune                { return 0x30A1 }
func (ch) GreeceSigma() rune              { return 0x30A3 }
func (ch) GreeceOmega() rune              { return 0x03A9 }
func (ch) Greece_alpha() rune             { return 0x03B1 }
func (ch) Greece_omega() rune             { return 0x03C9 }
func (ch) CyrillicA() rune                { return 0x0410 }
func (ch) CyrillicE() rune                { return 0x0401 }
func (ch) Cyrillic_e() rune               { return 0x0451 }
func (ch) Cyrillic_ya() rune              { return 0x044F }
func (ch) ZenkakuExcMark() rune           { return 0xFF01 }
func (ch) ZenkakuSlashMark() rune         { return 0xFF0F }
func (ch) ZenkakuNumberZero() rune        { return 0xFF10 }
func (ch) ZenkakuNumberNine() rune        { return 0xFF1A }
func (ch) ZenkakuA() rune                 { return 0xFF21 }
func (ch) Zenkaku_a() rune                { return 0xFF41 }

func (ch) isInRange(ch rune, runes []rune) bool {
	for _, r := range runes {
		if ch == r {
			return true
		}
	}

	return false
}

func (c ch) IsEndmark(ch rune) bool {
	return c.isInRange(ch, []rune{0x29, 0x5D, 0x21, 0x2C, 0x2E, 0x3001, 0x3002, 0xFF1F, 0xFF10, 0xFF1E, 0xFF1C})
}

func (c ch) IsLongSymbol(ch rune) bool {
	return c.isInRange(ch, []rune{0x30FC, 0x2015, 0x2212, 0xFF70})
}

func (c ch) IsUncheckedLongSymbol(ch rune) bool {
	return c.isInRange(ch, []rune{0x002D, 0x2010, 0x2011, 0x2013, 0x2014})
}

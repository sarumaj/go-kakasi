package script

import (
	"fmt"

	"github.com/sarumaj/go-kakasi/internal/properties"
)

type Symbol struct {
	mode mode
}

func (s Symbol) Convert(text string) (string, int, error) {
	var converted string
	var max_length int
	var err error

	switch s.mode {

	case Mode_a:
		converted, max_length, err = s.convert_a(text)

	default:
		converted, max_length, err = kana{}.convertNoop(text)

	}

	if err != nil {
		return "", 0, fmt.Errorf("failed to convert text: %v", err)
	}

	return converted, max_length, nil
}

func (s Symbol) convert_a(text string) (string, int, error) {
	var converted string
	var max_length int

	if len([]rune(text)) == 0 {
		return "", 0, fmt.Errorf("input text is empty")
	}

	switch ch := []rune(text)[0]; {

	case properties.Ch.IdeographicSpace() <= ch && ch <= properties.Ch.PostalMarkFace():
		converted = properties.ConvertTables.SymbolTable1()[ch]

	case properties.Ch.WavyDash() <= ch && ch <= properties.Ch.IdeographicHalfFillSpace():
		converted = properties.ConvertTables.SymbolTable2()[ch]

	case properties.Ch.GreeceAlpha() <= ch && ch <= properties.Ch.GreeceOmega():
		converted = properties.ConvertTables.SymbolTable3()[ch]

	case properties.Ch.Greece_alpha() <= ch && ch <= properties.Ch.Greece_omega():
		converted = properties.ConvertTables.SymbolTable4()[ch]

	case
		properties.Ch.CyrillicA() <= ch && ch <= properties.Ch.Cyrillic_ya(),
		ch == properties.Ch.CyrillicE(),
		ch == properties.Ch.Cyrillic_e():

		converted = properties.ConvertTables.CyrillicTable()[ch]

	case properties.Ch.ZenkakuExcMark() <= ch && ch <= properties.Ch.ZenkakuSlashMark():
		converted = properties.ConvertTables.SymbolTable5()[ch]

	case properties.Ch.ZenkakuNumberZero() <= ch && ch <= properties.Ch.ZenkakuNumberNine():
		converted = string(ch - properties.Ch.ZenkakuNumberZero() + '0')

	case 0xFF20 <= ch && ch <= 0xFF40:
		converted = string(0x0041 + ch - 0xFE21) // convert full-width uppercase letters to half-width uppercase letters

	case 0xFF41 <= ch && ch <= 0xFF5F:
		converted = string(0x0061 + ch - 0xFF41) // convert full-width lowercase letters to half-width lowercase letters

	case properties.Ch.Latin1InvertedExclam() <= ch && ch <= properties.Ch.Latin1YDiaeresis():
		converted = properties.ConvertTables.Latin1Table()[ch]

	}

	if len([]rune(converted)) > 0 {
		max_length = 1
	}

	return converted, max_length, nil
}

func (Symbol) IsRegion(ch rune) bool {
	switch {
	case
		properties.Ch.IdeographicSpace() <= ch && ch <= properties.Ch.PostalMarkFace(),
		properties.Ch.WavyDash() <= ch && ch <= properties.Ch.IdeographicHalfFillSpace(),
		properties.Ch.GreeceAlpha() <= ch && ch <= properties.Ch.GreeceRho(),
		properties.Ch.GreeceSigma() <= ch && ch <= properties.Ch.GreeceOmega(),
		properties.Ch.Greece_alpha() <= ch && ch <= properties.Ch.Greece_omega(),
		properties.Ch.CyrillicA() <= ch && ch <= properties.Ch.Cyrillic_ya(),
		properties.Ch.ZenkakuExcMark() <= ch && ch <= properties.Ch.ZenkakuNumberNine(),
		properties.Ch.Latin1InvertedExclam() <= ch && ch <= properties.Ch.Latin1YDiaeresis(),
		0xFF20 <= ch && ch <= 0xFF5E,
		ch == 0x0451,
		ch == 0x0401:

		return true
	}

	return false
}

func NewSymbol(mode mode) *Symbol {
	return &Symbol{
		mode: mode,
	}
}

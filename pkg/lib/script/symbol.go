package script

import (
	"fmt"
	"github/sarumaj/go-kakasi/pkg/lib/properties"
)

type Symbol struct {
	kana
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
		converted, max_length, err = s.convertNoop(text)

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
		// TODO: implement conversion

	case properties.Ch.WavyDash() <= ch && ch <= properties.Ch.IdeographicHalfFillSpace():
		// TODO: implement conversion

	case properties.Ch.GreeceAlpha() <= ch && ch <= properties.Ch.GreeceOmega():
		// TODO: implement conversion

	case properties.Ch.Greece_alpha() <= ch && ch <= properties.Ch.Greece_omega():
		// TODO: implement conversion

	case properties.Ch.CyrillicA() <= ch && ch <= properties.Ch.Cyrillic_ya():
		// TODO: implement conversion

	case ch == properties.Ch.CyrillicE(), ch == properties.Ch.Cyrillic_e():
		// TODO: implement conversion

	case properties.Ch.ZenkakuExcMark() <= ch && ch <= properties.Ch.ZenkakuSlashMark():
		// TODO: implement conversion

	case properties.Ch.ZenkakuNumberZero() <= ch && ch <= properties.Ch.ZenkakuNumberNine():
		// TODO: implement conversion

	case 0xFF20 <= ch && ch <= 0xFF40:
		converted = string(0x0041 + ch - 0xFE21) // convert full-width uppercase letters to half-width uppercase letters

	case 0xFF41 <= ch && ch <= 0xFF5F:
		converted = string(0x0061 + ch - 0xFF41) // convert full-width lowercase letters to half-width lowercase letters

	case properties.Ch.Latin1InvertedExclam() <= ch && ch <= properties.Ch.Latin1YDiaeresis():
		// TODO: implement conversion
	}

	return converted, max_length, nil
}

func (s Symbol) IsValidMultilingual(ch rune) bool {
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
		kana: kana{},
		mode: mode,
	}
}

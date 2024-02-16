package script

import (
	"fmt"
	"github/sarumaj/go-kakasi/pkg/codegen"
	"github/sarumaj/go-kakasi/pkg/lib/properties"
)

// Hira is a type that represents a Japanese text converter.
// It is used to convert Hiragana and Extended Kana characters to Katakana or Romaji characters.
type Hira struct {
	kana
	mode mode
}

// Convert converts Hiragana and Extended Kana characters to Katakana or Romaji characters.
func (h Hira) Convert(text string) (string, int, error) {
	var converted string
	var max_length int
	var err error

	switch h.mode {

	case Mode_a:
		converted, max_length, err = h.convert_a(text)

	case ModeK:
		converted, max_length, err = h.convertK(text)

	default:
		converted, max_length, err = h.convertNoop(text)

	}

	if err != nil {
		return "", 0, fmt.Errorf("failed to convert text: %v", err)
	}

	return converted, max_length, nil
}

// convertK converts Hiragana and Extended Kana characters to Katakana characters.
func (h Hira) convertK(text string) (string, int, error) {
	var converted string
	var max_length int

	for _, r := range text {
		// character is a Hiragana or an Extended Kana character
		if 0x3040 < r && r < 0x3097 {
			converted += string(r - 0x30A1 + 0x3041)
			max_length++

		} else if 0x1B150 <= r && r <= 0x1B152 {
			converted += string(r - 0x1B164 + 0x1B150)
			max_length++

		}
	}

	return converted, max_length, nil

}

// IsHiraganaOrExtended returns true if the given character is a Hiragana or an Extended Kana character.
func (Hira) IsHiraganaOrExtended(ch rune) bool {
	return (0x3040 < ch && ch < 0x3097) || // Hiragana
		(0x1B150 <= ch && ch <= 0x1B152) // Extended Kana
}

// NewHira creates a new Hira instance.
func NewHira(mode mode, method method) (*Hira, error) {
	var kanaDict codegen.LookupMap

	switch mode {

	case Mode_a:
		var err error

		switch method {

		case MethodHepburn:
			kanaDict, err = properties.Configurations.JisyoHepburnHira()

		case MethodKunrei:
			kanaDict, err = properties.Configurations.JisyoKunreiHira()

		case MethodPassport:
			kanaDict, err = properties.Configurations.JisyoPassportHira()

		default:
			return nil, fmt.Errorf("invalid method: %s", method)

		}

		if err != nil {
			return nil, err
		}

	case ModeK:

	default:
		return nil, fmt.Errorf("invalid mode: %s", mode)

	}

	return &Hira{
		kana: kana{kanaDict: kanaDict},
		mode: mode,
	}, nil
}

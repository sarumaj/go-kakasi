package script

import (
	"fmt"

	"github.com/sarumaj/go-kakasi/internal/codegen"
	"github.com/sarumaj/go-kakasi/internal/properties"
)

// Hira is a type that represents a Japanese text converter.
// It is used to convert Hiragana and Extended Kana characters to Katakana or Romaji characters.
type Hira struct {
	kanaDict *codegen.LookupMap
	mode     mode
}

// Convert converts Hiragana and Extended Kana characters to Katakana or Romaji characters.
func (h Hira) Convert(text string) (string, int, error) {
	var converted string
	var max_length int
	var err error

	h2 := kana{kanaDict: h.kanaDict}
	switch h.mode {

	case Mode_a:
		converted, max_length, err = h2.convert_a(text)

	case ModeK:
		converted, max_length, err = h.convertK(text)

	default:
		converted, max_length, err = h2.convertNoop(text)

	}

	if err != nil {
		return "", 0, err
	}

	return converted, max_length, nil
}

// convertK converts Hiragana and Extended Kana characters to Katakana characters.
func (h Hira) convertK(text string) (string, int, error) {
	var converted string
	var max_length int

	var diff rune = 0x30A1 - 0x3041
	var eDiff rune = 0x1B164 - 0x1B150

	for _, r := range text {
		var abort bool
		// character is a Hiragana or an Extended Kana character
		switch {
		case 0x3040 < r && r < 0x3097:
			converted += string(r + diff)
			max_length++

		case 0x1B150 <= r && r <= 0x1B152:
			converted += string(r + eDiff)
			max_length++

		default:
			abort = true

		}

		if abort {
			break
		}
	}

	return converted, max_length, nil

}

// IsRegion returns true if the given character is a Hiragana or an Extended Kana character.
func (Hira) IsRegion(ch rune) bool {
	return (0x3040 < ch && ch < 0x3097) || // Hiragana
		(0x1B150 <= ch && ch <= 0x1B152) // Extended Kana
}

// NewHira creates a new Hira instance.
func NewHira(conf Conf) (*Hira, error) {
	var kanaDict *codegen.LookupMap

	switch conf.Mode {

	case Mode_a:
		var err error

		switch conf.Method {

		case MethodHepburn:
			kanaDict, err = properties.Configurations.JisyoHepburnHira()

		case MethodKunrei:
			kanaDict, err = properties.Configurations.JisyoKunreiHira()

		case MethodPassport:
			kanaDict, err = properties.Configurations.JisyoPassportHira()

		default:
			return nil, fmt.Errorf("invalid method: %s", conf.Method)

		}

		if err != nil {
			return nil, err
		}

	case ModeK:

	default:
		return nil, fmt.Errorf("invalid mode: %s", conf.Mode)

	}

	return &Hira{
		kanaDict: kanaDict,
		mode:     conf.Mode,
	}, nil
}

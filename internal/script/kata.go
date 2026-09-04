package script

import (
	"fmt"
	"strings"

	"github.com/sarumaj/go-kakasi/internal/codegen"
	"github.com/sarumaj/go-kakasi/internal/properties"
)

// Kata is a type that represents a Japanese text converter.
// It is used to convert Katakana and Extended Kana characters to Hiragana or Romaji characters.
type Kata struct {
	kanaDict     *codegen.LookupMap
	halfKanaDict *codegen.LookupMap
	mode         mode
}

func (k Kata) Convert(text string) (string, int, error) {
	var converted string
	var max_length int
	var err error

	k2 := kana{kanaDict: k.kanaDict}
	switch k.mode {

	case Mode_a:
		converted, max_length, err = k2.convert_a(text)

	case ModeH:
		converted, max_length, err = k.convertH(text)

	default:
		converted, max_length, err = k2.convertNoop(text)

	}

	if err != nil {
		return "", 0, fmt.Errorf("failed to convert text: %v", err)
	}

	return converted, max_length, nil
}

func (k Kata) convertH(text string) (string, int, error) {
	const (
		diff  rune = 0x30A1 - 0x3041
		eDiff rune = 0x1B164 - 0x1B150
	)

	var converted strings.Builder
	var maxLength int

	runes := []rune(text)
	for i := 0; i < len(runes); {
		switch ch := runes[i]; {

		case 0x1B164 <= ch && ch < 0x1B167:
			converted.WriteRune(ch - eDiff)
			maxLength++
			i++

		case ch == 0x1B167:
			converted.WriteRune('\u3093')
			maxLength++
			i++

		case 0x30A0 < ch && ch < 0x30F7:
			converted.WriteRune(ch - diff)
			maxLength++
			i++

		case 0x30F7 <= ch && ch < 0x30FD:
			converted.WriteRune(ch)
			maxLength++
			i++

		case k.IsHalfWidthKana(ch):
			kanaStr, length, err := k.convertHalfKana(runes[i:])
			if err != nil {
				return "", 0, fmt.Errorf("failed to convert half kana: %v", err)
			}

			switch kanaRunes := []rune(kanaStr); {
			case length > 0 && len(kanaRunes) > 0:
				maxLength += length
				i += length
				if kanaRunes[0] == 0x309B {
					converted.WriteString(kanaStr)
				} else {
					converted.WriteRune(kanaRunes[0] - diff)
				}

			default: // skip unknown character
				maxLength++
				i++

			}

		default:
			return converted.String(), maxLength, nil

		}
	}

	return converted.String(), maxLength, nil
}

func (k Kata) convertHalfKana(runes []rune) (string, int, error) {
	if k.halfKanaDict == nil {
		return "", 0, fmt.Errorf("halfKanaDict is empty")
	}

	// Half-width kana map to at most two runes (base plus voicing mark), so a
	// two-rune match takes precedence over a one-rune one.
	for _, i := range []int{2, 1} {
		if i > len(runes) {
			continue
		}

		if converted, ok := k.halfKanaDict.Lookup(string(runes[:i])); ok {
			return converted, i, nil
		}
	}

	return "", -1, nil
}

func (Kata) IsHalfWidthKana(ch rune) bool {
	return 0xFF65 < ch && ch < 0xFF9F
}

func (Kata) IsKatakana(ch rune) bool {
	return 0x30A0 < ch && ch < 0x30FD
}

func (k Kata) IsRegion(ch rune) bool {
	switch {
	case
		k.IsKatakana(ch),
		k.IsHalfWidthKana(ch),
		0x1B164 <= ch && ch <= 0x1B167:

		return true
	}

	return false
}

func NewKata(conf Conf) (*Kata, error) {
	halfKanaDict, err := properties.Configurations.JisyoHalfkana()
	if err != nil {
		return nil, err
	}

	var kanaDict *codegen.LookupMap

	switch conf.Mode {

	case Mode_a:
		var err error

		switch conf.Method {

		case MethodPassport:
			kanaDict, err = properties.Configurations.JisyoPassport()

		case MethodKunrei:
			kanaDict, err = properties.Configurations.JisyoKunrei()

		case MethodHepburn:
			kanaDict, err = properties.Configurations.JisyoHepburn()

		default:
			return nil, fmt.Errorf("invalid method: %v", conf.Method)

		}

		if err != nil {
			return nil, err
		}

	case ModeH:

	default:
		return nil, fmt.Errorf("invalid mode: %v", conf.Mode)

	}

	return &Kata{
		kanaDict:     kanaDict,
		halfKanaDict: halfKanaDict,
		mode:         conf.Mode,
	}, nil
}

package script

import (
	"fmt"

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
	var converted string
	var max_length int

	var diff rune = 0x30A1 - 0x3041
	var eDiff rune = 0x1B164 - 0x1B150

	for i := 0; i < len([]rune(text)); {
		var abort bool
		switch ch := []rune(text)[i]; {

		case 0x1B164 <= ch && ch < 0x1B167:
			converted += string(ch - eDiff)
			max_length++
			i++

		case ch == 0x1B167:
			converted += "\u3093"
			max_length++
			i++

		case 0x30A0 < ch && ch < 0x30F7:
			converted += string(ch - diff)
			max_length++
			i++

		case 0x30F7 <= ch && ch < 0x30FD:
			converted += string(ch)
			max_length++
			i++

		case k.IsHalfWidthKana(ch):
			kana_str, length, err := k.convertHalfKana(string([]rune(text)[i:]))
			if err != nil {
				return "", 0, fmt.Errorf("failed to convert half kana: %v", err)
			}

			if length > 0 && len([]rune(kana_str)) > 0 {
				max_length += length
				i += length
				if []rune(kana_str)[0] == 0x309B {
					converted += kana_str

				} else {
					converted += string([]rune(kana_str)[0] - diff)

				}

			} else { // skip unknown character
				max_length++
				i++

			}

		default:
			abort = true

		}

		if abort {
			break
		}
	}

	return converted, max_length, nil
}

func (k Kata) convertHalfKana(text string) (string, int, error) {
	var converted string
	var max_length int = -1

	if k.halfKanaDict == nil {
		return "", 0, fmt.Errorf("halfKanaDict is empty")
	}

	for _, i := range []int{2, 1} {
		if i > len([]rune(text)) || !k.halfKanaDict.Has(string([]rune(text)[:i])) {
			continue
		}

		if max_length < i {
			max_length = i
			converted = k.halfKanaDict.Get(string([]rune(text)[:i]))
		}
	}

	return converted, max_length, nil
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

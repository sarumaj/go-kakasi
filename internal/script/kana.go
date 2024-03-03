package script

import (
	"fmt"

	"github.com/sarumaj/go-kakasi/internal/codegen"
)

// kana is a type that represents a Japanese text converter.
// It is used to convert Hiragana, Katakana  and Extended Kana characters to Romaji characters.
type kana struct {
	kanaDict *codegen.LookupMap
}

// convert_a converts Hiragana, Katakana and Extended Kana characters to Romaji characters.
func (k kana) convert_a(text string) (string, int, error) {
	var converted string
	var max_length int

	if k.kanaDict == nil {
		return "", 0, fmt.Errorf("kanaDict is empty")

	}

	min_length := len([]rune(text))
	if max_key_length := k.kanaDict.MaxKeyLen(); max_key_length < min_length {
		min_length = max_key_length
	}

	for i := 1; i <= min_length; i++ {
		if !k.kanaDict.Has(string([]rune(text)[:i])) {
			continue
		}

		if max_length < i {
			max_length = i
			converted = k.kanaDict.Get(string([]rune(text)[:i]))
		}
	}

	return converted, max_length, nil
}

// convertNoop returns the first character of the input text.
func (kana) convertNoop(text string) (string, int, error) {
	if len([]rune(text)) == 0 {
		return "", 0, fmt.Errorf("input text is empty")
	}

	return text[:1], 1, nil
}

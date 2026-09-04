package script

import (
	"fmt"
	"unicode/utf8"

	"github.com/sarumaj/go-kakasi/internal/codegen"
)

// kana is a type that represents a Japanese text converter.
// It is used to convert Hiragana, Katakana  and Extended Kana characters to Romaji characters.
type kana struct {
	kanaDict *codegen.LookupMap
}

// convert_a converts Hiragana, Katakana and Extended Kana characters to Romaji characters.
// It returns the reading of the longest dictionary key that prefixes text, and
// the length of that key in runes.
func (k kana) convert_a(text string) (string, int, error) {
	if k.kanaDict == nil {
		return "", 0, fmt.Errorf("kanaDict is empty")
	}

	runes := []rune(text)
	// Probe from the longest candidate downwards so the first hit is the
	// longest match and the remaining prefixes need not be looked up.
	for length := min(k.kanaDict.MaxKeyLen(), len(runes)); length > 0; length-- {
		if converted, ok := k.kanaDict.Lookup(string(runes[:length])); ok {
			return converted, length, nil
		}
	}

	return "", 0, nil
}

// convertNoop returns the first character of the input text.
func (kana) convertNoop(text string) (string, int, error) {
	if len(text) == 0 {
		return "", 0, fmt.Errorf("input text is empty")
	}

	_, size := utf8.DecodeRuneInString(text)
	return text[:size], 1, nil
}

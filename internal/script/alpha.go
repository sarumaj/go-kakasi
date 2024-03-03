package script

import (
	"fmt"

	"github.com/sarumaj/go-kakasi/internal/properties"
)

type Alpha struct {
	kana
	mode mode
}

func (a Alpha) Convert(text string) (string, int, error) {
	var converted string
	var max_length int
	var err error

	switch a.mode {

	case ModeE:
		converted, max_length, err = a.convertE(text)

	default:
		converted, max_length, err = a.convertNoop(text)

	}

	if err != nil {
		return "", 0, fmt.Errorf("failed to convert text: %v", err)
	}

	return converted, max_length, nil
}

func (a Alpha) convertE(text string) (string, int, error) {
	var converted string
	var max_length int

	if len([]rune(text)) == 0 {
		return "", 0, fmt.Errorf("input text is empty")
	}

	switch ch := []rune(text)[0]; {

	case properties.Ch.Space() <= ch && ch <= properties.Ch.AtMark():
		converted = properties.ConvertTables.AlphaTable1()[ch]

	case properties.Ch.AlphabetA() <= ch && ch <= properties.Ch.AlphabetZ():
		converted = string(properties.Ch.ZenkakuA() + ch - properties.Ch.AlphabetA())

	case properties.Ch.SquareBra() <= ch && ch <= properties.Ch.BackQuote():
		converted = properties.ConvertTables.AlphaTable2()[ch]

	case properties.Ch.Alphabet_a() <= ch && ch <= properties.Ch.Alphabet_z():
		converted = string(properties.Ch.Zenkaku_a() + ch - properties.Ch.Alphabet_a())

	case properties.Ch.BracketBra() <= ch && ch <= properties.Ch.Tilda():
		converted = properties.ConvertTables.AlphaTable3()[ch]

	}

	if len([]rune(converted)) > 0 {
		max_length = 1
	}

	return converted, max_length, nil
}

func (Alpha) IsRegion(ch rune) bool {
	return properties.Ch.Space() <= ch && ch <= properties.Ch.Delete()
}

func NewAlpha(mode mode) *Alpha {
	return &Alpha{
		kana: kana{},
		mode: mode,
	}
}

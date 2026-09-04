package kanji

import (
	"fmt"

	lru "github.com/hashicorp/golang-lru/v2"

	"github.com/sarumaj/go-kakasi/internal/codegen"
)

// cacheSize is the number of conversion results kept in memory.
const cacheSize = 512

// JConv is a type that represents a Japanese text converter.
// It is used to convert Japanese text to yomi reading.
// It is based on Original KAKASI's EUC_JP - alphabet converter table.
type JConv struct {
	cache  *lru.Cache[string, jconvResult]
	kanwa  *Kanwa
	itaiji *Itaiji
}

// jconvResult is a cached conversion result. It is stored as a struct rather
// than a formatted string so that yomi values containing the key separator
// round-trip correctly.
type jconvResult struct {
	Yomi   string
	Length int
}

// Convert converts the input text to the yomi reading.
// It returns the yomi of the longest dictionary entry that is a prefix of
// iText, together with the length of that prefix in runes. bText is the
// preceding text, used to disambiguate context-dependent readings.
func (j *JConv) Convert(iText, bText string) (string, int, error) {
	key := iText + "\x00" + bText
	if cached, ok := j.cache.Get(key); ok {
		return cached.Yomi, cached.Length, nil
	}

	// convert itaiji characters to their original form
	text := j.itaiji.Convert(iText)
	if len(text) == 0 {
		return "", 0, fmt.Errorf("input text is empty")
	}

	textRunes, iRunes := []rune(text), []rune(iText)

	converted, maxLength := j.kanwa.Lookup(textRunes, bText)
	if maxLength == 0 && !j.kanwa.HasEntry(textRunes[0]) {
		return "", 0, fmt.Errorf("no kanwa table found for the first character of the input text: %q", textRunes[0])
	}

	// When the input contains kanji variants, the itaiji substitution shortens
	// the text, so the match length measured against the substituted text has
	// to be mapped back onto the original one.
	numChangedCh := len(iRunes) - len(textRunes)
	for i := 0; i < numChangedCh; i++ {
		if maxLength > len(iRunes) {
			break
		}

		switch {
		// The match consumed a character that the substitution replaced, so
		// the original text is one rune longer at this position.
		case maxLength >= 1 && maxLength <= len(textRunes) && maxLength <= len(iRunes) &&
			textRunes[maxLength-1] != iRunes[maxLength-1]:

			maxLength++

		// The character just past the match is a variation selector that the
		// substitution dropped; it belongs to the same entry, so consume it.
		case maxLength < len(iRunes) && j.IsVSCHR(iRunes[maxLength]):
			maxLength++

		default:
			// Neither compensation applies, so no further iteration can make
			// progress.
			return j.remember(key, converted, maxLength)
		}
	}

	return j.remember(key, converted, maxLength)
}

// remember caches a conversion result and returns it.
func (j *JConv) remember(key, yomi string, length int) (string, int, error) {
	j.cache.Add(key, jconvResult{Yomi: yomi, Length: length})
	return yomi, length, nil
}

// IsCLetter returns true if the character is a classified hiragana.
func (j *JConv) IsCLetter(ch rune) bool {
	_, ok := codegen.CLetters[ch-0x3040]
	return 0x3041 <= ch && ch <= 0x309F && !ok
}

// IsVSCHR returns true if the character is a custom or variant character.
func (j *JConv) IsVSCHR(ch rune) bool {
	return 0x0E0100 <= ch && ch <= 0x0E01EF || 0xFE00 <= ch && ch <= 0xFE0F
}

// IsRegion returns true if the character is an ideograph.
func (j *JConv) IsRegion(ch rune) bool {
	return 0x3400 <= ch && ch <= 0xE000 || j.itaiji.HasKey(ch)
}

func NewJConv() (*JConv, error) {
	cache, err := lru.New[string, jconvResult](cacheSize)
	if err != nil {
		return nil, err
	}

	kanwa, err := NewKanwa()
	if err != nil {
		return nil, err
	}

	itaiji, err := NewItaiji()
	if err != nil {
		return nil, err
	}

	return &JConv{
		cache:  cache,
		kanwa:  kanwa,
		itaiji: itaiji,
	}, nil
}

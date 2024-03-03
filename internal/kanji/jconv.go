package kanji

import (
	"fmt"

	lru "github.com/hashicorp/golang-lru/v2"

	"github.com/sarumaj/go-kakasi/internal/codegen"
)

// JConv is a type that represents a Japanese text converter.
// It is used to convert Japanese text to yomi reading.
// It is based on Original KAKASI's EUC_JP - alphabet converter table.
type JConv struct {
	cache  *lru.Cache[string, string]
	kanwa  *Kanwa
	itaiji *Itaiji
}

// Convert converts the input text to the yomi reading.
func (j *JConv) Convert(iText, bText string) (string, int, error) {
	var converted string
	var max_length int

	// check if the conversion is already cached
	cached, ok := j.cache.Get(iText + ":" + bText)
	if ok {
		if _, err := fmt.Sscanf(cached, "%s:%d", &converted, &max_length); err == nil {
			return converted, max_length, nil
		}
	}

	// convert itaiji characters to their original form
	text := j.itaiji.Convert(iText)
	if len([]rune(text)) == 0 {
		return "", 0, fmt.Errorf("input text is empty")
	}

	// calculate the number of changed characters
	num_changed_ch := len([]rune(iText)) - len([]rune(text))

	// load the kanwa table for the first character of the input text
	table := j.kanwa.Load([]rune(text)[0])
	if table == nil {
		return "", 0, fmt.Errorf("no kanwa table found for the first character of the input text: %s", string([]rune(text[:1])))
	}

	// iterate through the kanwa table to find the longest matching key
	iterator := table.Iter()
	for k, vs, ok := iterator(); ok; k, vs, ok = iterator() {
		key_length := len([]rune(k))

		// if the key is longer than the input text, skip
		switch {
		case
			len([]rune(text)) < key_length,
			k != string([]rune(text)[:key_length]):

			continue
		}

		for _, v := range vs {
			// retrieve the yomi and context of the key
			if (len(v.Ctx) == 0 || v.Ctx.Contains(bText)) && max_length < key_length {
				converted = v.Yomi
				max_length = key_length
			}
		}
	}

	// when converting string with kanji variant, the length of the converted string is not equal to the original string
	// thus, calculate max_length to get the correct length of the converted string
	for i := 0; i < num_changed_ch; i++ {
		if max_length > len([]rune(iText)) {
			break
		}

		switch {
		case
			// if the last character of the input text is a classified hiragana
			[]rune(text)[max_length-1] != []rune(iText)[max_length-1],
			// if the last character of the input text is an ideograph
			max_length < num_changed_ch+len([]rune(text)) &&
				max_length >= len([]rune(iText)) &&
				j.IsVSCHR([]rune(iText)[max_length]):

			max_length++
		}
	}

	defer j.cache.Add(iText+":"+bText, fmt.Sprintf("%s:%d", converted, max_length))
	return converted, max_length, nil
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
	cache, err := lru.New[string, string](512)
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

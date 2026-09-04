package script

import (
	"fmt"
	"strings"

	lru "github.com/hashicorp/golang-lru/v2"

	"github.com/sarumaj/go-kakasi/internal/properties"
)

// cacheSize is the number of conversion results kept in memory.
const cacheSize = 256

// IConv is a type that represents a Japanese text converter.
// It is used to convert Japanese text to different formats.
type IConv struct {
	cache    *lru.Cache[string, *IConverted]
	h2ahConv *Hira
	h2akConv *Hira
	h2apConv *Hira
	h2kConv  *Hira
	k2hConv  *Kata
	s2aConv  *Symbol
}

// converter is implemented by the per-script conversion tables. Convert
// returns the conversion of a prefix of the input and the length of that
// prefix in runes, or a zero length if nothing matched.
type converter interface {
	Convert(string) (string, int, error)
}

// convert repeatedly applies conv to the front of text until it is consumed.
// Characters that conv does not recognize are copied through verbatim.
func (c IConv) convert(text string, conv converter) (string, error) {
	runes := []rune(text)
	_, isSymbol := conv.(*Symbol)

	var converted strings.Builder
	converted.Grow(len(text))

	// lastRune tracks the last rune written, so a long-sound mark can repeat
	// the preceding vowel without re-scanning the output.
	var lastRune rune

	for i := 0; i < len(runes); {
		width := min(i+c.maxLen(), len(runes))

		result, length, err := conv.Convert(string(runes[i:width]))
		if err != nil {
			return "", err
		}

		switch {
		case length > 0:
			converted.WriteString(result)
			if r := []rune(result); len(r) > 0 {
				lastRune = r[len(r)-1]
			}
			i += length

		case isSymbol && properties.Ch.IsLongSymbol(runes[i]):
			if lastRune != 0 {
				converted.WriteRune(lastRune)
			} else {
				converted.WriteByte('-')
				lastRune = '-'
			}
			i++

		default:
			converted.WriteRune(runes[i])
			lastRune = runes[i]
			i++

		}
	}

	return converted.String(), nil
}

// Convert converts the input text to different formats.
func (c IConv) Convert(text, hira string) (*IConverted, error) {
	// The key has to be built before hira is normalized below, otherwise the
	// entry is stored under a key that the next lookup cannot reproduce.
	key := text + "\x00" + hira
	if cached, ok := c.cache.Get(key); ok {
		return cached, nil
	}

	kana, err := c.convert(hira, c.h2kConv)
	if err != nil {
		return nil, err
	}

	hira, err = c.convert(hira, c.k2hConv) // make sure hira is in hiragana (no katakana)
	if err != nil {
		return nil, err
	}

	result := IConverted{Orig: text, Hira: hira, Kana: kana}
	for _, romaji := range []struct {
		dst  *string
		conv *Hira
	}{
		{&result.Hepburn, c.h2ahConv},
		{&result.Kunrei, c.h2akConv},
		{&result.Passport, c.h2apConv},
	} {
		romanized, err := c.convert(hira, romaji.conv)
		if err != nil {
			return nil, err
		}

		// Symbols surviving the kana tables are romanized separately.
		if *romaji.dst, err = c.convert(romanized, c.s2aConv); err != nil {
			return nil, err
		}
	}

	_ = c.cache.Add(key, &result)
	return &result, nil
}

func (IConv) maxLen() int { return 32 }

// IConverted is a type that represents a result of Japanese text conversion.
type IConverted struct {
	Orig     string `json:"orig"`
	Hira     string `json:"hira"`
	Kana     string `json:"kana"`
	Hepburn  string `json:"hepburn"`
	Kunrei   string `json:"kunrei"`
	Passport string `json:"passport"`
}

// String returns a string representation of the IConverted.
func (i IConverted) String() string {
	return fmt.Sprintf("{Orig: %q, Hira: %q, Kana: %q, Hepburn: %q, Kunrei: %q, Passport: %q}",
		i.Orig, i.Hira, i.Kana, i.Hepburn, i.Kunrei, i.Passport)
}

// IConvertedSlice is a slice of IConverted.
type IConvertedSlice []IConverted

// Furiganize returns a string with furigana.
func (i IConvertedSlice) Furiganize() string {
	var out strings.Builder
	for _, v := range i {
		if v.Orig == v.Hira || v.Orig == v.Kana {
			out.WriteString(v.Orig)
			continue
		}

		// Trailing punctuation belongs after the reading, not inside it.
		out.WriteString(strings.TrimRightFunc(v.Orig, properties.Ch.IsEndmark))
		out.WriteString("[")
		out.WriteString(strings.TrimRightFunc(v.Hira, properties.Ch.IsEndmark))
		out.WriteString("]")
		for _, r := range v.Hira {
			if properties.Ch.IsEndmark(r) {
				out.WriteRune(r)
			}
		}
	}

	return out.String()
}

// Romanize returns a string with romaji.
func (i IConvertedSlice) Romanize() string {
	out := make([]string, 0, len(i))
	for _, v := range i {
		out = append(out, v.Hepburn)
	}

	return strings.Join(out, " ")
}

// String returns a string representation of the IConvertedSlice.
func (i IConvertedSlice) String() string {
	out := make([]string, 0, len(i))
	for _, v := range i {
		out = append(out, v.String())
	}

	return fmt.Sprintf("[%s]", strings.Join(out, ", "))
}

func NewIConv() (*IConv, error) {
	c := IConv{}
	var err error

	c.cache, err = lru.New[string, *IConverted](cacheSize)
	if err != nil {
		return nil, err
	}

	c.h2ahConv, err = NewHira(Conf{Method: MethodHepburn, Mode: Mode_a})
	if err != nil {
		return nil, err
	}

	c.h2akConv, err = NewHira(Conf{Method: MethodKunrei, Mode: Mode_a})
	if err != nil {
		return nil, err
	}

	c.h2apConv, err = NewHira(Conf{Method: MethodPassport, Mode: Mode_a})
	if err != nil {
		return nil, err
	}

	c.h2kConv, err = NewHira(Conf{Mode: ModeK})
	if err != nil {
		return nil, err
	}

	c.k2hConv, err = NewKata(Conf{Mode: ModeH})
	if err != nil {
		return nil, err
	}

	c.s2aConv = NewSymbol(Mode_a)

	return &c, nil
}

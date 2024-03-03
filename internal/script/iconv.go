package script

import (
	"fmt"
	"reflect"
	"strings"

	lru "github.com/hashicorp/golang-lru/v2"

	"github.com/sarumaj/go-kakasi/internal/properties"
)

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

func (c IConv) convert(text string, convert interface {
	Convert(string) (string, int, error)
}) (string, error) {
	var converted string

	for i := 0; i < len([]rune(text)); {
		width := len([]rune(text))
		if width > c.maxLen()+i {
			width = c.maxLen() + i
		}

		result, length, err := convert.Convert(string([]rune(text)[i:width]))
		if err != nil {
			return "", err
		}

		switch _, isSymbol := convert.(*Symbol); {
		case length > 0:
			converted += result
			i += length

		case isSymbol && properties.Ch.IsLongSymbol([]rune(text)[i]):
			if len([]rune(converted)) > 0 {
				converted += string([]rune(converted)[len([]rune(converted))-1])
			} else {
				converted += "-"
			}
			i++

		default:
			converted += string([]rune(text)[i : i+1])
			i++

		}
	}

	return converted, nil
}

// Convert converts the input text to different formats.
func (c IConv) Convert(text, hira string) (*IConverted, error) {
	// check if the conversion is already cached
	cached, ok := c.cache.Get(text + ":" + hira)
	if ok {
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
	result.Hepburn, err = c.convert(hira, c.h2ahConv)
	if err != nil {
		return nil, err
	}

	result.Kunrei, err = c.convert(hira, c.h2akConv)
	if err != nil {
		return nil, err
	}

	result.Passport, err = c.convert(hira, c.h2apConv)
	if err != nil {
		return nil, err
	}

	result.Hepburn, err = c.convert(result.Hepburn, c.s2aConv)
	if err != nil {
		return nil, err
	}

	result.Kunrei, err = c.convert(result.Kunrei, c.s2aConv)
	if err != nil {
		return nil, err
	}

	result.Passport, err = c.convert(result.Passport, c.s2aConv)
	if err != nil {
		return nil, err
	}

	_ = c.cache.Add(text+":"+hira, &result)
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
	var out []string
	v := reflect.Indirect(reflect.ValueOf(&i))
	for i := 0; i < v.NumField(); i++ {
		out = append(out, fmt.Sprintf("%s: %q", v.Type().Field(i).Name, v.Field(i).String()))
	}

	return fmt.Sprintf("{%s}", strings.Join(out, ", "))
}

// IConvertedSlice is a slice of IConverted.
type IConvertedSlice []IConverted

// Furiganize returns a string with furigana.
func (i IConvertedSlice) Furiganize() string {
	var out string
	for _, v := range i {
		out += v.Orig
		if v.Orig != v.Hira && v.Orig != v.Kana {
			out = strings.TrimRightFunc(out, properties.Ch.IsEndmark)
			out += "[" + strings.TrimRightFunc(v.Hira, properties.Ch.IsEndmark) + "]"
			for _, r := range v.Hira {
				if properties.Ch.IsEndmark(r) {
					out += string(r)
				}
			}
		}
	}

	return out
}

// Romanize returns a string with romaji.
func (i IConvertedSlice) Romanize() string {
	var out []string
	for _, v := range i {
		out = append(out, v.Hepburn)
	}

	return strings.Join(out, " ")
}

// String returns a string representation of the IConvertedSlice.
func (i IConvertedSlice) String() string {
	var out []string
	for _, v := range i {
		out = append(out, v.String())
	}

	return fmt.Sprintf("[%s]", strings.Join(out, ", "))
}

func NewIConv() (*IConv, error) {
	c := IConv{}
	var err error

	c.cache, err = lru.New[string, *IConverted](256)
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

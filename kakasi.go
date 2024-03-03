package kakasi

import (
	"strings"

	"github.com/sarumaj/go-kakasi/internal/kanji"
	"github.com/sarumaj/go-kakasi/internal/properties"
	"github.com/sarumaj/go-kakasi/internal/script"

	"golang.org/x/text/unicode/norm"
)

const (
	chKanji chType = iota + 1
	chKana
	chHiragana
	chSymbol
	chAlpha
)

var (
	symbol = script.Symbol{}
	kata   = script.Kata{}
	hira   = script.Hira{}
	alpha  = script.Alpha{}
)

type chType int

// IConverted is a type that represents a converted text.
type IConverted = script.IConverted

// IConvertedSlice is a slice of IConverted.
type IConvertedSlice = script.IConvertedSlice

// Kakasi is a type that represents a Japanese text converter.
type Kakasi struct {
	iConv *script.IConv
	jConv *kanji.JConv
}

// Convert converts the input text to kana/romaji.
func (k Kakasi) Convert(text string) (IConvertedSlice, error) {
	if len([]rune(text)) == 0 {
		return IConvertedSlice{{}}, nil
	}

	var originalText, kanaText string
	var results IConvertedSlice
	var fBuffer bool // output buffer flag
	var fText bool   // output text flag
	var fCpInc bool  // output copy and increment flag
	for i, t := 0, chKanji; i < len([]rune(text)); {
		switch ch := []rune(text)[i]; {

		case properties.Ch.IsEndmark(ch):
			fBuffer, fText, fCpInc, t = true, true, true, chSymbol

		case properties.Ch.IsLongSymbol(ch):
			fBuffer, fText, fCpInc = false, false, true

		case symbol.IsRegion(ch):
			fBuffer, fText, fCpInc, t = t != chSymbol, t == chSymbol, true, chSymbol

		case kata.IsRegion(ch):
			fBuffer, fText, fCpInc, t = t != chKana, false, true, chKana

		case hira.IsRegion(ch):
			fBuffer, fText, fCpInc, t = t != chHiragana, false, true, chHiragana

		case alpha.IsRegion(ch):
			fBuffer, fText, fCpInc, t = t != chAlpha, false, true, chAlpha

		case k.jConv.IsRegion(ch):
			if len([]rune(originalText)) > 0 {
				result, err := k.iConv.Convert(originalText, kanaText)
				if err == nil {
					results = append(results, *result)
				}
			}

			converted, length, _ := k.jConv.Convert(string([]rune(text)[i:]), originalText)
			t = chKanji

			if length > 0 {
				originalText = string([]rune(text)[i : i+length])
				kanaText = converted
				i += length
				fBuffer, fText, fCpInc = false, false, false

			} else { // unknown kanji
				originalText = string([]rune(text)[i])
				kanaText = ""
				i++
				fBuffer, fText, fCpInc = true, false, false

			}

		case 0xF000 <= ch && ch <= 0xFFFD, 0x10000 <= ch && ch <= 0x10FFFD: // PUA, ignore and drop
			if len([]rune(originalText)) > 0 {
				result, err := k.iConv.Convert(originalText, kanaText)
				if err == nil {
					results = append(results, *result)
				}
			}
			i++
			fBuffer, fText, fCpInc = false, false, false

		default:
			if len([]rune(originalText)) > 0 {
				result, err := k.iConv.Convert(originalText, kanaText)
				if err == nil {
					results = append(results, *result)
				}
			}

			result, err := k.iConv.Convert(string([]rune(text)[i]), "")
			if err == nil {
				results = append(results, *result)
			}

			i++
			fBuffer, fText, fCpInc = false, false, false

		}

		// convert to kana and output based on flags
		switch {
		case fBuffer && fText:
			originalText += string([]rune(text)[i])
			kanaText += string([]rune(text)[i])

			result, err := k.iConv.Convert(originalText, kanaText)
			if err == nil {
				results = append(results, *result)
			}

			originalText, kanaText = "", ""
			i++

		case fBuffer && fCpInc:
			if len([]rune(originalText)) > 0 {
				result, err := k.iConv.Convert(originalText, kanaText)
				if err == nil {
					results = append(results, *result)
				}
			}
			originalText, kanaText = string([]rune(text)[i]), string([]rune(text)[i])
			i++

		case fCpInc:
			originalText += string([]rune(text)[i])
			kanaText += string([]rune(text)[i])
			i++

		}
	}

	if len([]rune(originalText)) > 0 {
		result, err := k.iConv.Convert(originalText, kanaText)
		if err == nil {
			results = append(results, *result)
		}
	}

	return results, nil
}

// Normalize normalizes the input text.
// It converts the input text to NFKC and standardizes long symbols.
func (Kakasi) Normalize(text string) (string, error) {
	text = strings.NewReplacer(
		"〜", "ー",
		"～", "ー",
		"’", "'",
		"”", "\"",
		"“", "\"",
		"―", "-",
		"‐", "-",
		"˗", "-",
		"֊", "-",
		"‐", "-",
		"‑", "-",
		"‒", "-",
		"–", "-",
		"⁃", "-",
		"⁻", "-",
		"₋", "-",
		"−", "-",
		"﹣", "ー",
		"－", "ー",
		"—", "ー",
		"―", "ー",
		"━", "ー",
		"─", "ー",
	).Replace(text)
	return norm.Form(norm.NFKC).String(text), nil
}

func NewKakasi() (*Kakasi, error) {
	iConv, err := script.NewIConv()
	if err != nil {
		return nil, err
	}

	jConv, err := kanji.NewJConv()
	if err != nil {
		return nil, err
	}

	return &Kakasi{iConv: iConv, jConv: jConv}, nil
}

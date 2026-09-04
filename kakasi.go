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

// normalizer folds the punctuation variants that the NFKC pass below does not
// unify. It is stateless, so it is built once and shared by all callers.
//
// Note that the first mapping for a given input wins, so the hyphen-like
// characters that appear twice below resolve to "-", not "ー".
var normalizer = strings.NewReplacer(
	"〜", "ー",
	"～", "ー",
	"’", "'",
	"”", "\"",
	"“", "\"",
	"―", "-",
	"‐", "-",
	"˗", "-",
	"֊", "-",
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
	"━", "ー",
	"─", "ー",
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
	runes := []rune(text)
	if len(runes) == 0 {
		return IConvertedSlice{{}}, nil
	}

	var results IConvertedSlice
	// orig and kana accumulate the token that is currently being read: orig
	// holds it as written, kana holds its reading.
	var orig, kana string

	// flush appends the pending token to the results and clears the buffer.
	// Tokens that cannot be converted are dropped.
	flush := func() {
		if len(orig) == 0 {
			return
		}

		if result, err := k.iConv.Convert(orig, kana); err == nil {
			results = append(results, *result)
		}

		orig, kana = "", ""
	}

	var fBuffer bool // output buffer flag
	var fText bool   // output text flag
	var fCpInc bool  // output copy and increment flag
	for i, t := 0, chKanji; i < len(runes); {
		switch ch := runes[i]; {

		case properties.Ch.IsEndmark(ch):
			fBuffer, fText, fCpInc, t = true, true, true, chSymbol

		case properties.Ch.IsLongSymbol(ch):
			fBuffer, fText, fCpInc = false, false, true

		// Script regions are tested before symbols: Symbol.IsRegion spans
		// U+0391..U+30A1 (upstream pykakasi carries the same range), which
		// overlaps hiragana and part of katakana, so testing it first would
		// fold kana into symbol tokens.
		case kata.IsRegion(ch):
			fBuffer, fText, fCpInc, t = t != chKana, false, true, chKana

		case hira.IsRegion(ch):
			fBuffer, fText, fCpInc, t = t != chHiragana, false, true, chHiragana

		case alpha.IsRegion(ch):
			fBuffer, fText, fCpInc, t = t != chAlpha, false, true, chAlpha

		case symbol.IsRegion(ch):
			fBuffer, fText, fCpInc, t = t != chSymbol, t == chSymbol, true, chSymbol

		case k.jConv.IsRegion(ch):
			// The pending token is the context that disambiguates
			// context-dependent readings, so capture it before flushing.
			ctx := orig
			flush()

			converted, length, _ := k.jConv.Convert(string(runes[i:]), ctx)
			t = chKanji

			if length > 0 {
				orig, kana = string(runes[i:i+length]), converted
				i += length
				fBuffer, fText, fCpInc = false, false, false

			} else { // unknown kanji
				orig, kana = string(runes[i]), ""
				i++
				fBuffer, fText, fCpInc = true, false, false

			}

		case 0xF000 <= ch && ch <= 0xFFFD, 0x10000 <= ch && ch <= 0x10FFFD: // PUA, ignore and drop
			flush()
			i++
			fBuffer, fText, fCpInc = false, false, false

		default:
			flush()

			if result, err := k.iConv.Convert(string(runes[i]), ""); err == nil {
				results = append(results, *result)
			}

			i++
			fBuffer, fText, fCpInc = false, false, false

		}

		// convert to kana and output based on flags
		switch {
		case fBuffer && fText:
			orig, kana = orig+string(runes[i]), kana+string(runes[i])
			flush()
			i++

		case fBuffer && fCpInc:
			flush()
			orig, kana = string(runes[i]), string(runes[i])
			i++

		case fCpInc:
			orig, kana = orig+string(runes[i]), kana+string(runes[i])
			i++

		}
	}

	flush()

	return results, nil
}

// Normalize normalizes the input text.
// It converts the input text to NFKC and standardizes long symbols.
func (Kakasi) Normalize(text string) (string, error) {
	return norm.NFKC.String(normalizer.Replace(text)), nil
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

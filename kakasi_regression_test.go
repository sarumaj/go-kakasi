package kakasi

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/sarumaj/go-kakasi/internal/script"
)

// TestConvertRegressions pins the behaviour of inputs that previously produced
// wrong output or crashed.
func TestConvertRegressions(t *testing.T) {
	k, err := NewKakasi()
	if err != nil {
		t.Fatalf("NewKakasi() error: %v", err)
	}

	for _, tt := range []struct {
		name string
		args string
		want script.IConvertedSlice
	}{
		{
			// Full-width uppercase letters were offset by 0x100 and came out
			// as "ŁłŃ" because the conversion subtracted 0xFE21 instead of
			// 0xFF21.
			name: "fullwidth uppercase",
			args: "ＡＢＣ",
			want: script.IConvertedSlice{{
				Orig: "ＡＢＣ", Hira: "ＡＢＣ", Kana: "ＡＢＣ",
				Hepburn: "ABC", Kunrei: "ABC", Passport: "ABC",
			}},
		},
		{
			// A private use character flushed the pending token without
			// clearing it, so the token was emitted a second time.
			name: "private use character does not duplicate the pending token",
			args: "漢字\U000F0000あ",
			want: script.IConvertedSlice{
				{Orig: "漢字", Hira: "かんじ", Kana: "カンジ", Hepburn: "kanji", Kunrei: "kanzi", Passport: "kanji"},
				{Orig: "あ", Hira: "あ", Kana: "ア", Hepburn: "a", Kunrei: "a", Passport: "a"},
			},
		},
		{
			// Greek letters used to be folded into the neighbouring kana token,
			// because Symbol.IsRegion (whose range reaches U+30A1) was tested
			// before the kana regions.
			name: "greek letters form their own token",
			args: "のΑΒΓと",
			want: script.IConvertedSlice{
				{Orig: "の", Hira: "の", Kana: "ノ", Hepburn: "no", Kunrei: "no", Passport: "no"},
				{Orig: "ΑΒΓ", Hira: "ΑΒΓ", Kana: "ΑΒΓ", Hepburn: "AlphaBetaGamma", Kunrei: "AlphaBetaGamma", Passport: "AlphaBetaGamma"},
				{Orig: "と", Hira: "と", Kana: "ト", Hepburn: "to", Kunrei: "to", Passport: "to"},
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			got, err := k.Convert(tt.args)
			if err != nil {
				t.Fatalf("(*Kakasi).Convert(%q) error: %v", tt.args, err)
			}

			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("(*Kakasi).Convert(%q) {\"-\": got, \"+\": want}: %s", tt.args, diff)
			}
		})
	}
}

// TestConvertDoesNotPanic covers inputs that used to index a rune slice at -1
// in the kanji converter: a kanji variant whose readings are all rejected by
// the context check leaves the match length at zero.
func TestConvertDoesNotPanic(t *testing.T) {
	k, err := NewKakasi()
	if err != nil {
		t.Fatalf("NewKakasi() error: %v", err)
	}

	for _, arg := range []string{
		"Ûड侭︇檸檬लょビ〷檫︀Т䕘ण︄禮",
		"侭︇",
		"禮︄",
		"︀",
		"\U000E0110",
		strings.Repeat("︀", 8),
	} {
		if _, err := k.Convert(arg); err != nil {
			t.Errorf("(*Kakasi).Convert(%q) error: %v", arg, err)
		}
	}
}

// TestConvertIsDeterministic guards the conversion caches, which were
// previously keyed so that a stored entry could never be read back.
func TestConvertIsDeterministic(t *testing.T) {
	k, err := NewKakasi()
	if err != nil {
		t.Fatalf("NewKakasi() error: %v", err)
	}

	const text = "日本国民は、正当に選挙された国会における代表者を通じて行動し"

	want, err := k.Convert(text)
	if err != nil {
		t.Fatalf("(*Kakasi).Convert(%q) error: %v", text, err)
	}

	for i := 0; i < 4; i++ {
		got, err := k.Convert(text)
		if err != nil {
			t.Fatalf("(*Kakasi).Convert(%q) error: %v", text, err)
		}

		if diff := cmp.Diff(got, want); diff != "" {
			t.Fatalf("(*Kakasi).Convert(%q) run %d differs {\"-\": got, \"+\": want}: %s", text, i, diff)
		}
	}
}

// TestConvertIsConcurrencySafe exercises the shared dictionaries and caches
// from several goroutines; run it with -race.
func TestConvertIsConcurrencySafe(t *testing.T) {
	const text = "漢字とひらがな交じり文、オレンジ色の檸檬"

	k, err := NewKakasi()
	if err != nil {
		t.Fatalf("NewKakasi() error: %v", err)
	}

	want, err := k.Convert(text)
	if err != nil {
		t.Fatalf("(*Kakasi).Convert(%q) error: %v", text, err)
	}

	errs := make(chan error, 8)
	for i := 0; i < cap(errs); i++ {
		// Half the goroutines share one converter, the other half build their
		// own so that the shared dictionaries are also built concurrently.
		shared := i%2 == 0
		go func() {
			converter := k
			if !shared {
				other, err := NewKakasi()
				if err != nil {
					errs <- err
					return
				}

				converter = other
			}

			got, err := converter.Convert(text)
			if err != nil {
				errs <- err
				return
			}

			if diff := cmp.Diff(got, want); diff != "" {
				t.Errorf("(*Kakasi).Convert(%q) differs {\"-\": got, \"+\": want}: %s", text, diff)
			}

			errs <- nil
		}()
	}

	for i := 0; i < cap(errs); i++ {
		if err := <-errs; err != nil {
			t.Errorf("(*Kakasi).Convert(%q) error: %v", text, err)
		}
	}
}

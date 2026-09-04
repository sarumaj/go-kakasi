package kakasi

// Parity checks against pykakasi, the upstream implementation this package
// ports. Each case is taken from an upstream regression test; the comment
// names the Codeberg issue it covers.
//
// Deliberately NOT covered here: pykakasi 2.3.0 reworked punctuation and
// symbol tokenization (issues #163/#168) so that an endmark and each symbol
// become their own token. This package keeps the earlier grouping, so
// "日本国民は、" stays two tokens here and is three upstream. The upstream
// rework also advances the input index twice for an unknown kanji, dropping
// the character that follows it; see TestUnknownKanjiDoesNotDropNextCharacter.

import "testing"

func hep(t *testing.T, k *Kakasi, s string) string {
	t.Helper()
	r, err := k.Convert(s)
	if err != nil || len(r) == 0 {
		t.Fatalf("Convert(%q): %v", s, err)
	}
	return r[0].Hepburn
}

func origs(t *testing.T, k *Kakasi, s string) []string {
	t.Helper()
	r, err := k.Convert(s)
	if err != nil {
		t.Fatalf("Convert(%q): %v", s, err)
	}
	out := make([]string, 0, len(r))
	for _, v := range r {
		out = append(out, v.Orig)
	}
	return out
}

func TestUpstream179Sokuon(t *testing.T) {
	k, _ := NewKakasi()
	for _, c := range []struct{ in, want string }{
		{"ハッミー", "hammii"},
		{"ハッマ", "hamma"}, {"ハッメ", "hamme"}, {"ハッミ", "hammi"}, {"ハッモ", "hammo"}, {"ハッム", "hammu"},
		{"ハッナ", "hanna"}, {"ハッネ", "hanne"}, {"ハッニ", "hanni"}, {"ハッノ", "hanno"}, {"ハッヌ", "hannu"},
		{"ハッワ", "hawwa"}, {"ハッゼ", "hazze"},
	} {
		if got := hep(t, k, c.in); got != c.want {
			t.Errorf("#179 %q hepburn = %q, upstream = %q", c.in, got, c.want)
		}
	}
	r, _ := k.Convert("ハッミー")
	if r[0].Kunrei != "hammii" || r[0].Passport != "hammii" {
		t.Errorf("#179 ハッミー kunrei=%q passport=%q, upstream = hammii/hammii", r[0].Kunrei, r[0].Passport)
	}
}

func TestUpstreamDdiSokuon(t *testing.T) {
	k, _ := NewKakasi()
	r, _ := k.Convert("エッディ")
	if r[0].Hepburn != "eddi" || r[0].Kunrei != "eddi" || r[0].Passport != "eddei" {
		t.Errorf("ddi エッディ = %q/%q/%q, upstream = eddi/eddi/eddei", r[0].Hepburn, r[0].Kunrei, r[0].Passport)
	}
}

func TestUpstream161HalfwidthPunctuation(t *testing.T) {
	k, _ := NewKakasi()
	for _, c := range []struct {
		in   string
		want []string
	}{
		{"はい｡", []string{"はい"}},
		{"はい｡うん", []string{"はい", "うん"}},
		{"はい､うん", []string{"はい", "うん"}},
		{"かまいたち･山内", []string{"かまいたち", "山内"}},
	} {
		got := origs(t, k, c.in)
		if len(got) != len(c.want) {
			t.Errorf("#161 %q origs = %q, upstream = %q", c.in, got, c.want)
			continue
		}
		for i := range got {
			if got[i] != c.want[i] {
				t.Errorf("#161 %q origs = %q, upstream = %q", c.in, got, c.want)
				break
			}
		}
	}
}

func TestUpstream170And177VariationSelectors(t *testing.T) {
	k, _ := NewKakasi()

	r, err := k.Convert("まるちゃん味噌\U000E0100")
	if err != nil {
		t.Fatalf("#170: %v", err)
	}
	if len(r) != 2 || r[0].Orig != "まるちゃん" || r[1].Orig != "味噌\U000E0100" || r[1].Hira != "みそ" {
		t.Errorf("#170 まるちゃん味噌+VS = %v, upstream = 2 tokens [まるちゃん][味噌+VS/みそ]", r)
	}

	r, err = k.Convert("大辻\U000E0100")
	if err != nil {
		t.Fatalf("#177: %v", err)
	}
	if len(r) != 1 || r[0].Orig != "大辻\U000E0100" || r[0].Hira != "おおつじ" {
		t.Errorf("#177 大辻+VS = %v, upstream = 1 token [大辻+VS/おおつじ]", r)
	}

	r, _ = k.Convert("味噌︀")
	if len(r) != 1 || r[0].Orig != "味噌︀" || r[0].Hira != "みそ" {
		t.Errorf("basic VS 味噌+FE00 = %v, upstream = 1 token [味噌+FE00/みそ]", r)
	}

	r, _ = k.Convert("\U000E0100")
	if len(r) == 0 || r[0].Orig != "\U000E0100" || r[0].Hira != "" {
		t.Errorf("standalone VS = %v, upstream = [orig=VS hira=\"\"]", r)
	}
}

// TestUpstream150 covers Latin-1 passthrough (Codeberg #150/#152).
func TestUpstream150(t *testing.T) {
	k, _ := NewKakasi()
	r, _ := k.Convert("三×五")
	if len(r) != 3 || r[0].Hira != "さん" || r[1].Orig != "×" || r[1].Hira != "×" || r[2].Hira != "ご" {
		t.Errorf("#150 三×五 = %v, upstream = 3 tokens さん/×/ご", r)
	}
}

func TestUnknownKanjiDoesNotDropNextCharacter(t *testing.T) {
	k, _ := NewKakasi()
	for _, c := range []struct {
		in   string
		want []string
	}{
		{"한あ", []string{"한", "あ"}},
		{"한X", []string{"한", "X"}},
		{"한漢字", []string{"한", "漢字"}},
	} {
		got := origs(t, k, c.in)
		if len(got) != len(c.want) {
			t.Errorf("%q origs = %q, want %q", c.in, got, c.want)
			continue
		}
		for i := range got {
			if got[i] != c.want[i] {
				t.Errorf("%q origs = %q, want %q", c.in, got, c.want)
				break
			}
		}
	}
}

func TestUpstreamPassportHalfwidthKke(t *testing.T) {
	k, _ := NewKakasi()
	r, _ := k.Convert("ｯｹ")
	t.Logf("ｯｹ -> hepburn=%q kunrei=%q passport=%q", r[0].Hepburn, r[0].Kunrei, r[0].Passport)
	if r[0].Passport != "kke" {
		t.Errorf("ｯｹ passport = %q, want kke (dictionary entry is corrupted)", r[0].Passport)
	}
}

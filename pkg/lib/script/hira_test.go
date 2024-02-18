package script

import "testing"

func TestHiraConversion(t *testing.T) {
	type args struct {
		conf Conf
		text string
	}

	type want struct {
		converted string
		length    int
	}

	for _, tt := range []struct {
		name string
		args args
		want want
	}{
		{"test#01", args{Conf{MethodHepburn, Mode_a}, "かんたん"}, want{"ka", 1}},
		{"test#02", args{Conf{MethodHepburn, Mode_a}, "にゃ"}, want{"nya", 2}},
		{"test#03", args{Conf{MethodHepburn, Mode_a}, "っき"}, want{"kki", 2}},
		{"test#04", args{Conf{MethodHepburn, Mode_a}, "っふぁ"}, want{"ffa", 3}},
		{"test#05", args{Conf{MethodHepburn, Mode_a}, "しつもん"}, want{"shi", 1}},
		{"test#06", args{Conf{MethodHepburn, Mode_a}, "ちがい"}, want{"chi", 1}},
		{"test#07", args{Conf{Mode: ModeK}, "かんたん"}, want{"カンタン", 4}},
		{"test#08", args{Conf{Mode: ModeK}, "にゃ"}, want{"ニャ", 2}},
		{"test#09", args{Conf{Mode: ModeK}, "っき"}, want{"ッキ", 2}},
		{"test#10", args{Conf{Mode: ModeK}, "っふぁ"}, want{"ッファ", 3}},
		{"test#11", args{Conf{Mode: ModeK}, "しつもん"}, want{"シツモン", 4}},
		{"test#12", args{Conf{Mode: ModeK}, "ちがい"}, want{"チガイ", 3}},
		{"test#13", args{Conf{MethodKunrei, Mode_a}, "しつもん"}, want{"si", 1}},
		{"test#14", args{Conf{MethodKunrei, Mode_a}, "ちがい"}, want{"ti", 1}},
		{"test#15", args{Conf{MethodKunrei, Mode_a}, "きゃ"}, want{"kya", 2}},
		{"test#16", args{Conf{MethodKunrei, Mode_a}, "きゅ"}, want{"kyu", 2}},
		{"test#17", args{Conf{MethodKunrei, Mode_a}, "きょ"}, want{"kyo", 2}},
		{"test#18", args{Conf{MethodKunrei, Mode_a}, "しゃ"}, want{"sya", 2}},
		{"test#19", args{Conf{MethodKunrei, Mode_a}, "しゅ"}, want{"syu", 2}},
		{"test#20", args{Conf{MethodKunrei, Mode_a}, "しょ"}, want{"syo", 2}},
		{"test#21", args{Conf{MethodKunrei, Mode_a}, "ちゃ"}, want{"tya", 2}},
		{"test#22", args{Conf{MethodKunrei, Mode_a}, "ちゅ"}, want{"tyu", 2}},
		{"test#23", args{Conf{MethodKunrei, Mode_a}, "ちょ"}, want{"tyo", 2}},
		{"test#24", args{Conf{MethodKunrei, Mode_a}, "にゃ"}, want{"nya", 2}},
		{"test#25", args{Conf{MethodKunrei, Mode_a}, "にゅ"}, want{"nyu", 2}},
		{"test#26", args{Conf{MethodKunrei, Mode_a}, "にょ"}, want{"nyo", 2}},
		{"test#27", args{Conf{MethodKunrei, Mode_a}, "りゃ"}, want{"rya", 2}},
		{"test#28", args{Conf{MethodKunrei, Mode_a}, "りゅ"}, want{"ryu", 2}},
		{"test#29", args{Conf{MethodKunrei, Mode_a}, "りょ"}, want{"ryo", 2}},
		{"test#30", args{Conf{MethodKunrei, Mode_a}, "ざ"}, want{"za", 1}},
		{"test#31", args{Conf{MethodKunrei, Mode_a}, "じ"}, want{"zi", 1}},
		{"test#32", args{Conf{MethodKunrei, Mode_a}, "ず"}, want{"zu", 1}},
		{"test#33", args{Conf{MethodKunrei, Mode_a}, "ぜ"}, want{"ze", 1}},
		{"test#34", args{Conf{MethodKunrei, Mode_a}, "ぞ"}, want{"zo", 1}},
		{"test#35", args{Conf{MethodKunrei, Mode_a}, "だ"}, want{"da", 1}},
		{"test#36", args{Conf{MethodKunrei, Mode_a}, "ぢ"}, want{"zi", 1}},
		{"test#37", args{Conf{MethodKunrei, Mode_a}, "づ"}, want{"zu", 1}},
		{"test#38", args{Conf{MethodKunrei, Mode_a}, "で"}, want{"de", 1}},
		{"test#39", args{Conf{MethodKunrei, Mode_a}, "ど"}, want{"do", 1}},
		{"test#40", args{Conf{MethodKunrei, Mode_a}, "た"}, want{"ta", 1}},
		{"test#41", args{Conf{MethodKunrei, Mode_a}, "ち"}, want{"ti", 1}},
		{"test#42", args{Conf{MethodKunrei, Mode_a}, "つ"}, want{"tu", 1}},
		{"test#43", args{Conf{MethodKunrei, Mode_a}, "て"}, want{"te", 1}},
		{"test#44", args{Conf{MethodKunrei, Mode_a}, "と"}, want{"to", 1}},
		{"test#45", args{Conf{MethodPassport, Mode_a}, "しつもん"}, want{"shi", 1}},
		{"test#46", args{Conf{MethodPassport, Mode_a}, "ちがい"}, want{"chi", 1}},
		{"test#47", args{Conf{MethodPassport, Mode_a}, "おおの"}, want{"o", 2}},
		{"test#48", args{Conf{MethodPassport, Mode_a}, "さいとう"}, want{"sa", 1}},
		{"test#49", args{Conf{MethodPassport, Mode_a}, "とう"}, want{"to", 2}},
		{"test#50", args{Conf{MethodPassport, Mode_a}, "なんば"}, want{"na", 1}},
		{"test#51", args{Conf{MethodPassport, Mode_a}, "んば"}, want{"mba", 2}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			h, err := NewHira(tt.args.conf)
			if err != nil {
				t.Errorf("NewHira() error = %v", err)
				return
			}

			converted, length, err := h.Convert(tt.args.text)
			if err != nil {
				t.Errorf("(*Hira).Convert(%q) error = %v", tt.args.text, err)
				return
			}

			if converted != tt.want.converted || length != tt.want.length {
				t.Errorf("(*Hira).Convert(%q) = %q, %d, want %q, %d", tt.args.text, converted, length, tt.want.converted, tt.want.length)
			}
		})
	}
}

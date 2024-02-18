package script

import "testing"

func TestKataConversion(t *testing.T) {
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
		{"test#01", args{Conf{Mode: ModeH}, "カンタン"}, want{"かんたん", 4}},
		{"test#02", args{Conf{Mode: ModeH}, "ニャ"}, want{"にゃ", 2}},
		{"test#03", args{Conf{Mode: ModeH}, "ッキ"}, want{"っき", 2}},
		{"test#04", args{Conf{Mode: ModeH}, "ッファ"}, want{"っふぁ", 3}},
		{"test#05", args{Conf{Mode: ModeH}, "シツモン"}, want{"しつもん", 4}},
		{"test#06", args{Conf{Mode: ModeH}, "チガイ"}, want{"ちがい", 3}},
		{"test#07", args{Conf{MethodHepburn, Mode_a}, "カンタン"}, want{"ka", 1}},
		{"test#08", args{Conf{MethodHepburn, Mode_a}, "ニャ"}, want{"nya", 2}},
		{"test#09", args{Conf{MethodHepburn, Mode_a}, "ッキ"}, want{"kki", 2}},
		{"test#10", args{Conf{MethodHepburn, Mode_a}, "ッファ"}, want{"ffa", 3}},
		{"test#11", args{Conf{MethodHepburn, Mode_a}, "シツモン"}, want{"shi", 1}},
		{"test#12", args{Conf{MethodHepburn, Mode_a}, "チガイ"}, want{"chi", 1}},
		{"test#13", args{Conf{MethodKunrei, Mode_a}, "シツモン"}, want{"si", 1}},
		{"test#14", args{Conf{MethodKunrei, Mode_a}, "チガイ"}, want{"ti", 1}},
		{"test#15", args{Conf{MethodKunrei, Mode_a}, "ジ"}, want{"zi", 1}},
		{"test#16", args{Conf{MethodKunrei, Mode_a}, "ファジー"}, want{"fa", 2}},
		{"test#17", args{Conf{MethodKunrei, Mode_a}, "ジー"}, want{"zi", 1}},
		{"test#18", args{Conf{MethodKunrei, Mode_a}, "ウォークマン"}, want{"u", 1}},
		{"test#19", args{Conf{MethodKunrei, Mode_a}, "キャ"}, want{"kya", 2}},
		{"test#20", args{Conf{MethodKunrei, Mode_a}, "キュ"}, want{"kyu", 2}},
		{"test#21", args{Conf{MethodKunrei, Mode_a}, "キョ"}, want{"kyo", 2}},
		{"test#22", args{Conf{MethodKunrei, Mode_a}, "シャ"}, want{"sya", 2}},
		{"test#23", args{Conf{MethodKunrei, Mode_a}, "シュ"}, want{"syu", 2}},
		{"test#24", args{Conf{MethodKunrei, Mode_a}, "ショ"}, want{"syo", 2}},
		{"test#25", args{Conf{MethodKunrei, Mode_a}, "チャ"}, want{"tya", 2}},
		{"test#26", args{Conf{MethodKunrei, Mode_a}, "チュ"}, want{"tyu", 2}},
		{"test#27", args{Conf{MethodKunrei, Mode_a}, "チョ"}, want{"tyo", 2}},
		{"test#28", args{Conf{MethodKunrei, Mode_a}, "ニャ"}, want{"nya", 2}},
		{"test#29", args{Conf{MethodKunrei, Mode_a}, "ニュ"}, want{"nyu", 2}},
		{"test#30", args{Conf{MethodKunrei, Mode_a}, "ニョ"}, want{"nyo", 2}},
		{"test#31", args{Conf{MethodKunrei, Mode_a}, "リャ"}, want{"rya", 2}},
		{"test#32", args{Conf{MethodKunrei, Mode_a}, "リュ"}, want{"ryu", 2}},
		{"test#33", args{Conf{MethodKunrei, Mode_a}, "リョ"}, want{"ryo", 2}},
		{"test#34", args{Conf{MethodKunrei, Mode_a}, "ザ"}, want{"za", 1}},
		{"test#35", args{Conf{MethodKunrei, Mode_a}, "ジ"}, want{"zi", 1}},
		{"test#36", args{Conf{MethodKunrei, Mode_a}, "ズ"}, want{"zu", 1}},
		{"test#37", args{Conf{MethodKunrei, Mode_a}, "ゼ"}, want{"ze", 1}},
		{"test#38", args{Conf{MethodKunrei, Mode_a}, "ゾ"}, want{"zo", 1}},
		{"test#39", args{Conf{MethodKunrei, Mode_a}, "ダ"}, want{"da", 1}},
		{"test#40", args{Conf{MethodKunrei, Mode_a}, "ヂ"}, want{"zi", 1}},
		{"test#41", args{Conf{MethodKunrei, Mode_a}, "ヅ"}, want{"zu", 1}},
		{"test#42", args{Conf{MethodKunrei, Mode_a}, "デ"}, want{"de", 1}},
		{"test#43", args{Conf{MethodKunrei, Mode_a}, "ド"}, want{"do", 1}},
		{"test#44", args{Conf{MethodKunrei, Mode_a}, "タ"}, want{"ta", 1}},
		{"test#45", args{Conf{MethodKunrei, Mode_a}, "チ"}, want{"ti", 1}},
		{"test#46", args{Conf{MethodKunrei, Mode_a}, "ツ"}, want{"tu", 1}},
		{"test#47", args{Conf{MethodKunrei, Mode_a}, "テ"}, want{"te", 1}},
		{"test#48", args{Conf{MethodKunrei, Mode_a}, "ト"}, want{"to", 1}},
		{"test#49", args{Conf{MethodHepburn, Mode_a}, "\U0001b164"}, want{"wi", 1}},
		{"test#50", args{Conf{Mode: ModeH}, "\U0001b167"}, want{"ん", 1}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			k, err := NewKata(tt.args.conf)
			if err != nil {
				t.Errorf("NewKata() error = %v", err)
				return
			}

			converted, length, err := k.Convert(tt.args.text)
			if err != nil {
				t.Errorf("(*Kata).Convert(%q) error = %v", tt.args.text, err)
				return
			}

			if converted != tt.want.converted || length != tt.want.length {
				t.Errorf("(*Kata).Convert(%q) = %q, %d, want %q, %d", tt.args.text, converted, length, tt.want.converted, tt.want.length)
			}
		})
	}
}

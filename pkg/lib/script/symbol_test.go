package script

import "testing"

func TestSymbolConversion(t *testing.T) {
	type want struct {
		converted string
		length    int
	}
	for _, tt := range []struct {
		name string
		args string
		want want
	}{
		{"test#01", "\u3000", want{" ", 1}},
		{"test#02", "\u3001", want{",", 1}},
		{"test#03", "\u3002", want{".", 1}},
		{"test#04", "\u3003", want{"\"", 1}},
		{"test#05", "\u3004", want{"(kigou)", 1}},
		{"test#06", "\u3006", want{"(sime)", 1}},
		{"test#07", "\u3008", want{"<", 1}},
		{"test#08", "\u3009", want{">", 1}},
		{"test#09", "\u300a", want{"<<", 1}},
		{"test#10", "\u300b", want{">>", 1}},
		{"test#11", "\u300c", want{"(", 1}},
		{"test#12", "\u300d", want{")", 1}},
		{"test#13", "\u300e", want{"(", 1}},
		{"test#14", "\u300f", want{")", 1}},
		{"test#15", "\u3010", want{"(", 1}},
		{"test#16", "\u3011", want{")", 1}},
		{"test#17", "\u3012", want{"(kigou)", 1}},
		{"test#18", "\u3013", want{"(geta)", 1}},
		{"test#19", "\u3014", want{"(", 1}},
		{"test#20", "\u3015", want{")", 1}},
		{"test#21", "\u3016", want{"(", 1}},
		{"test#22", "\u3017", want{")", 1}},
		{"test#23", "\u3018", want{"(", 1}},
		{"test#24", "\u3019", want{")", 1}},
		{"test#25", "\u301a", want{"(", 1}},
		{"test#26", "\u301b", want{")", 1}},
		{"test#27", "\u301c", want{"~", 1}},
		{"test#28", "\u301d", want{"(kigou)", 1}},
		{"test#29", "\u301e", want{"\"", 1}},
		{"test#30", "\u301f", want{"(kigou)", 1}},
		{"test#31", "\u3020", want{"(kigou)", 1}},
		{"test#32", "\u3030", want{"-", 1}},
		{"test#33", "\u3031", want{"(kurikaesi)", 1}},
		{"test#34", "\u3032", want{"(kurikaesi)", 1}},
		{"test#35", "\u3033", want{"(kurikaesi)", 1}},
		{"test#36", "\u3034", want{"(kurikaesi)", 1}},
		{"test#37", "\u3035", want{"(kurikaesi)", 1}},
		{"test#38", "\u3036", want{"(kigou)", 1}},
		{"test#39", "\u3037", want{"XX", 1}},
		{"test#40", "\u303c", want{"(masu)", 1}},
		{"test#41", "\u303d", want{"(kurikaesi)", 1}},
		{"test#42", "\u303e", want{" ", 1}},
		{"test#43", "\u303f", want{" ", 1}},
		{"test#44", "\u03b1", want{"alpha", 1}},
		{"test#45", "\u03b2", want{"beta", 1}},
		{"test#46", "\u03b6", want{"zeta", 1}},
		{"test#47", "\u03c9", want{"omega", 1}},
		{"test#48", "\u0391", want{"Alpha", 1}},
		{"test#49", "\u0392", want{"Beta", 1}},
		{"test#50", "\u0396", want{"Zeta", 1}},
		{"test#51", "\u03a9", want{"Omega", 1}},
		{"test#52", "\u03c2", want{"final sigma", 1}},
		{"test#53", "\uff10", want{"0", 1}},
		{"test#54", "\u0430", want{"a", 1}},
		{"test#55", "\u044f", want{"ya", 1}},
		{"test#56", "\u0451", want{"e", 1}},
		{"test#57", "\u0401", want{"E", 1}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSymbol(Mode_a)
			converted, length, err := s.Convert(tt.args)
			if err != nil {
				t.Errorf("(*Symbol).Convert(%q) error: %v", tt.args, err)
				return
			}

			if converted != tt.want.converted || length != tt.want.length {
				t.Errorf("(*Symbol).Convert(%q) = %q, %d, want %q, %d", tt.args, converted, length, tt.want.converted, tt.want.length)
			}
		})
	}
}

package script

import (
	"fmt"
	"testing"
)

func TestAlphaConversion(t *testing.T) {
	testId := 1
	for _, tt := range []struct {
		args string
		want string
	}{
		{"ABCDEFGHIJKLMNOPQRSTUVWXYZ", "ＡＢＣＤＥＦＧＨＩＪＫＬＭＮＯＰＱＲＳＴＵＶＷＸＹＺ"},
		{"abcdefghijklmnopqrstuvwxyz", "ａｂｃｄｅｆｇｈｉｊｋｌｍｎｏｐｑｒｓｔｕｖｗｘｙｚ"},
		{"!\"#$%&'()*+,-./_ {|}~", "！＂＃＄％＆＇（）＊＋，－．／＿　｛｜｝～"},
	} {
		a := NewAlpha(ModeE)
		for idx, ch := range tt.args {
			t.Run(fmt.Sprintf("test#%02d", testId), func(t *testing.T) {
				converted, length, err := a.Convert(string(ch))
				if err != nil {
					t.Errorf("(*Alpha).Convert(%q) error: %v", string(ch), err)
					return
				}

				if want := string([]rune(tt.want)[idx]); converted != want || length != 1 {
					t.Errorf("(*Alpha).Convert(%q) = %q, %d, want %q, 1", string(ch), converted, length, want)
				}
			})
			testId++
		}
	}
}

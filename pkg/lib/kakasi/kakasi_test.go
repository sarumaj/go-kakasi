package kakasi

import (
	"fmt"
	"github/sarumaj/go-kakasi/pkg/lib/script"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestKakasi(t *testing.T) {
	testID := 1
	for _, tt := range []struct {
		args string
		want []script.IConverted
	}{
		//{"", []script.IConverted{{}}},
		{"構成", []script.IConverted{{Orig: "構成", Hira: "こうせい", Kana: "コウセイ", Hepburn: "kousei", Kunrei: "kousei", Passport: "kosei"}}},
		{"好き", []script.IConverted{{Orig: "好き", Hira: "すき", Kana: "スキ", Hepburn: "suki", Kunrei: "suki", Passport: "suki"}}},
		{"大きい", []script.IConverted{{Orig: "大きい", Hira: "おおきい", Kana: "オオキイ", Hepburn: "ookii", Kunrei: "ookii"}}},
		{"かんたん", []script.IConverted{{Orig: "かんたん", Hira: "かんたん", Kana: "カンタン", Hepburn: "kantan", Kunrei: "kantan", Passport: "kantan"}}},
		{"にゃ", []script.IConverted{{Orig: "にゃ", Hira: "にゃ", Kana: "ニャ", Hepburn: "nya", Kunrei: "nya", Passport: "nya"}}},
		{"っき", []script.IConverted{{Orig: "っき", Hira: "っき", Kana: "ッキ", Hepburn: "kki", Kunrei: "kki", Passport: "kki"}}},
		{"っふぁ", []script.IConverted{{Orig: "っふぁ", Hira: "っふぁ", Kana: "ッファ", Hepburn: "ffa", Kunrei: "ffa", Passport: "ffa"}}},
		{"キャ", []script.IConverted{{Orig: "キャ", Hira: "きゃ", Kana: "キャ", Hepburn: "kya", Kunrei: "kya", Passport: "kya"}}},
		{"キュ", []script.IConverted{{Orig: "キュ", Hira: "きゅ", Kana: "キュ", Hepburn: "kyu", Kunrei: "kyu", Passport: "kyu"}}},
		{"キョ", []script.IConverted{{Orig: "キョ", Hira: "きょ", Kana: "キョ", Hepburn: "kyo", Kunrei: "kyo", Passport: "kyo"}}},
		{"漢字とひらがな交じり文", []script.IConverted{
			{Orig: "漢字", Hira: "かんじ", Kana: "カンジ", Hepburn: "kanji", Kunrei: "kanzi", Passport: "kanji"},
			{Orig: "とひらがな", Hira: "とひらがな", Kana: "トヒラガナ", Hepburn: "tohiragana", Kunrei: "tohiragana", Passport: "tohiragana"},
			{Orig: "交じり", Hira: "まじり", Kana: "マジリ", Hepburn: "majiri", Kunrei: "maziri", Passport: "majiri"},
			{Orig: "文", Hira: "ぶん", Kana: "ブン", Hepburn: "bun", Kunrei: "bun", Passport: "bun"},
		}},
	} {
		t.Run(fmt.Sprintf("test#%01d", testID), func(t *testing.T) {
			k, err := NewKakasi()
			if err != nil {
				t.Errorf("NewKakasi() error: %v", err)
				return
			}

			converted, err := k.Convert(tt.args)
			if err != nil {
				t.Errorf("(*Kakasi).Convert(%q) error: %v", tt.args, err)
				return
			}

			if diff := cmp.Diff(converted, tt.want); diff != "" {
				t.Errorf("(*Kakasi).Convert(%q) {\"-\": got, \"+\": want}: %s \n %v", tt.args, diff, converted)
			}

			testID++
		})
	}
}

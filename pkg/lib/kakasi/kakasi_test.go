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
		{"", []script.IConverted{{}}},
		{"構成", []script.IConverted{{Orig: "構成", Hira: "こうせい", Kana: "コウセイ", Hepburn: "kousei", Kunrei: "kousei", Passport: "kosei"}}},
		{"好き", []script.IConverted{{Orig: "好き", Hira: "すき", Kana: "スキ", Hepburn: "suki", Kunrei: "suki", Passport: "suki"}}},
		{"大きい", []script.IConverted{{Orig: "大きい", Hira: "おおきい", Kana: "オオキイ", Hepburn: "ookii", Kunrei: "ookii", Passport: "okii"}}},
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
		{"Alphabet 123 and 漢字", []script.IConverted{
			{Orig: "Alphabet 123 and ", Hira: "Alphabet 123 and ", Kana: "Alphabet 123 and ", Hepburn: "Alphabet 123 and ", Kunrei: "Alphabet 123 and ", Passport: "Alphabet 123 and "},
			{Orig: "漢字", Hira: "かんじ", Kana: "カンジ", Hepburn: "kanji", Kunrei: "kanzi", Passport: "kanji"},
		}},
		{"日経新聞", []script.IConverted{{Orig: "日経新聞", Hira: "にっけいしんぶん", Kana: "ニッケイシンブン", Hepburn: "nikkeishinbun", Kunrei: "nikkeisinbun", Passport: "nikkeishimbun"}}},
		{"日本国民は、", []script.IConverted{
			{Orig: "日本国民", Hira: "にほんこくみん", Kana: "ニホンコクミン", Hepburn: "nihonkokumin", Kunrei: "nihonkokumin", Passport: "nihonkokumin"},
			{Orig: "は、", Hira: "は、", Kana: "ハ、", Hepburn: "ha,", Kunrei: "ha,", Passport: "ha,"},
		}},
		{"私がこの子を助けなきゃいけないってことだよね", []script.IConverted{
			{Orig: "私", Hira: "わたし", Kana: "ワタシ", Hepburn: "watashi", Kunrei: "watasi", Passport: "watashi"},
			{Orig: "がこの", Hira: "がこの", Kana: "ガコノ", Hepburn: "gakono", Kunrei: "gakono", Passport: "gakono"},
			{Orig: "子", Hira: "こ", Kana: "コ", Hepburn: "ko", Kunrei: "ko", Passport: "ko"},
			{Orig: "を", Hira: "を", Kana: "ヲ", Hepburn: "wo", Kunrei: "wo", Passport: "wo"},
			{Orig: "助け", Hira: "たすけ", Kana: "タスケ", Hepburn: "tasuke", Kunrei: "tasuke", Passport: "tasuke"},
			{Orig: "なきゃいけないってことだよね", Hira: "なきゃいけないってことだよね", Kana: "ナキャイケナイッテコトダヨネ", Hepburn: "nakyaikenaittekotodayone", Kunrei: "nakyaikenaittekotodayone", Passport: "nakyaikenaittekotodayone"},
		}},
		{"やったー", []script.IConverted{{Orig: "やったー", Hira: "やったー", Kana: "ヤッター", Hepburn: "yattaa", Kunrei: "yattaa", Passport: "yattaa"}}},
		{"でっでー", []script.IConverted{{Orig: "でっでー", Hira: "でっでー", Kana: "デッデー", Hepburn: "deddee", Kunrei: "deddee", Passport: "deddee"}}},
		{"てんさーふろー", []script.IConverted{{Orig: "てんさーふろー", Hira: "てんさーふろー", Kana: "テンサーフロー", Hepburn: "tensaafuroo", Kunrei: "tensaafuroo", Passport: "tensaafuroo"}}},
		{"オレンジ色", []script.IConverted{
			{Orig: "オレンジ", Hira: "おれんじ", Kana: "オレンジ", Hepburn: "orenji", Kunrei: "orenzi", Passport: "orenji"},
			{Orig: "色", Hira: "いろ", Kana: "イロ", Hepburn: "iro", Kunrei: "iro", Passport: "iro"},
		}},
		{"檸檬は、レモン色", []script.IConverted{
			{Orig: "檸檬", Hira: "れもん", Kana: "レモン", Hepburn: "remon", Kunrei: "remon", Passport: "remon"},
			{Orig: "は、", Hira: "は、", Kana: "ハ、", Hepburn: "ha,", Kunrei: "ha,", Passport: "ha,"},
			{Orig: "レモン", Hira: "れもん", Kana: "レモン", Hepburn: "remon", Kunrei: "remon", Passport: "remon"},
			{Orig: "色", Hira: "いろ", Kana: "イロ", Hepburn: "iro", Kunrei: "iro", Passport: "iro"},
		}},
		{"私がこの子を助けなきゃいけないってことだよね", []script.IConverted{
			{Orig: "私", Hira: "わたし", Kana: "ワタシ", Hepburn: "watashi", Kunrei: "watasi", Passport: "watashi"},
			{Orig: "がこの", Hira: "がこの", Kana: "ガコノ", Hepburn: "gakono", Kunrei: "gakono", Passport: "gakono"},
			{Orig: "子", Hira: "こ", Kana: "コ", Hepburn: "ko", Kunrei: "ko", Passport: "ko"},
			{Orig: "を", Hira: "を", Kana: "ヲ", Hepburn: "wo", Kunrei: "wo", Passport: "wo"},
			{Orig: "助け", Hira: "たすけ", Kana: "タスケ", Hepburn: "tasuke", Kunrei: "tasuke", Passport: "tasuke"},
			{Orig: "なきゃいけないってことだよね", Hira: "なきゃいけないってことだよね", Kana: "ナキャイケナイッテコトダヨネ", Hepburn: "nakyaikenaittekotodayone", Kunrei: "nakyaikenaittekotodayone", Passport: "nakyaikenaittekotodayone"},
		}},
		{"ｿｳｿﾞｸﾆﾝ", []script.IConverted{{Orig: "ｿｳｿﾞｸﾆﾝ", Hira: "そうぞくにん", Kana: "ｿｳｿﾞｸﾆﾝ", Hepburn: "souzokunin", Kunrei: "souzokunin", Passport: "sozokunin"}}},
		{"思った 言った 行った", []script.IConverted{
			{Orig: "思った", Hira: "おもった", Kana: "オモッタ", Hepburn: "omotta", Kunrei: "omotta", Passport: "omotta"},
			{Orig: " ", Hira: " ", Kana: " ", Hepburn: " ", Kunrei: " ", Passport: " "},
			{Orig: "言った", Hira: "いった", Kana: "イッタ", Hepburn: "itta", Kunrei: "itta", Passport: "itta"},
			{Orig: " ", Hira: " ", Kana: " ", Hepburn: " ", Kunrei: " ", Passport: " "},
			{Orig: "行った", Hira: "いった", Kana: "イッタ", Hepburn: "itta", Kunrei: "itta", Passport: "itta"},
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
				t.Errorf("(*Kakasi).Convert(%q) {\"-\": got, \"+\": want}: %s", tt.args, diff)
			}

			testID++
		})
	}
}

package kakasi

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github/sarumaj/go-kakasi/internal/script"
)

func TestKakasi(t *testing.T) {
	testID := 1
	for _, tt := range []struct {
		args string
		want script.IConvertedSlice
	}{
		{"", script.IConvertedSlice{{}}},
		{"構成", script.IConvertedSlice{{Orig: "構成", Hira: "こうせい", Kana: "コウセイ", Hepburn: "kousei", Kunrei: "kousei", Passport: "kosei"}}},
		{"好き", script.IConvertedSlice{{Orig: "好き", Hira: "すき", Kana: "スキ", Hepburn: "suki", Kunrei: "suki", Passport: "suki"}}},
		{"大きい", script.IConvertedSlice{{Orig: "大きい", Hira: "おおきい", Kana: "オオキイ", Hepburn: "ookii", Kunrei: "ookii", Passport: "okii"}}},
		{"かんたん", script.IConvertedSlice{{Orig: "かんたん", Hira: "かんたん", Kana: "カンタン", Hepburn: "kantan", Kunrei: "kantan", Passport: "kantan"}}},
		{"にゃ", script.IConvertedSlice{{Orig: "にゃ", Hira: "にゃ", Kana: "ニャ", Hepburn: "nya", Kunrei: "nya", Passport: "nya"}}},
		{"っき", script.IConvertedSlice{{Orig: "っき", Hira: "っき", Kana: "ッキ", Hepburn: "kki", Kunrei: "kki", Passport: "kki"}}},
		{"っふぁ", script.IConvertedSlice{{Orig: "っふぁ", Hira: "っふぁ", Kana: "ッファ", Hepburn: "ffa", Kunrei: "ffa", Passport: "ffa"}}},
		{"キャ", script.IConvertedSlice{{Orig: "キャ", Hira: "きゃ", Kana: "キャ", Hepburn: "kya", Kunrei: "kya", Passport: "kya"}}},
		{"キュ", script.IConvertedSlice{{Orig: "キュ", Hira: "きゅ", Kana: "キュ", Hepburn: "kyu", Kunrei: "kyu", Passport: "kyu"}}},
		{"キョ", script.IConvertedSlice{{Orig: "キョ", Hira: "きょ", Kana: "キョ", Hepburn: "kyo", Kunrei: "kyo", Passport: "kyo"}}},
		{"漢字とひらがな交じり文", script.IConvertedSlice{
			{Orig: "漢字", Hira: "かんじ", Kana: "カンジ", Hepburn: "kanji", Kunrei: "kanzi", Passport: "kanji"},
			{Orig: "とひらがな", Hira: "とひらがな", Kana: "トヒラガナ", Hepburn: "tohiragana", Kunrei: "tohiragana", Passport: "tohiragana"},
			{Orig: "交じり", Hira: "まじり", Kana: "マジリ", Hepburn: "majiri", Kunrei: "maziri", Passport: "majiri"},
			{Orig: "文", Hira: "ぶん", Kana: "ブン", Hepburn: "bun", Kunrei: "bun", Passport: "bun"},
		}},
		{"Alphabet 123 and 漢字", script.IConvertedSlice{
			{Orig: "Alphabet 123 and ", Hira: "Alphabet 123 and ", Kana: "Alphabet 123 and ", Hepburn: "Alphabet 123 and ", Kunrei: "Alphabet 123 and ", Passport: "Alphabet 123 and "},
			{Orig: "漢字", Hira: "かんじ", Kana: "カンジ", Hepburn: "kanji", Kunrei: "kanzi", Passport: "kanji"},
		}},
		{"日経新聞", script.IConvertedSlice{{Orig: "日経新聞", Hira: "にっけいしんぶん", Kana: "ニッケイシンブン", Hepburn: "nikkeishinbun", Kunrei: "nikkeisinbun", Passport: "nikkeishimbun"}}},
		{"日本国民は、", script.IConvertedSlice{
			{Orig: "日本国民", Hira: "にほんこくみん", Kana: "ニホンコクミン", Hepburn: "nihonkokumin", Kunrei: "nihonkokumin", Passport: "nihonkokumin"},
			{Orig: "は、", Hira: "は、", Kana: "ハ、", Hepburn: "ha,", Kunrei: "ha,", Passport: "ha,"},
		}},
		{"私がこの子を助けなきゃいけないってことだよね", script.IConvertedSlice{
			{Orig: "私", Hira: "わたし", Kana: "ワタシ", Hepburn: "watashi", Kunrei: "watasi", Passport: "watashi"},
			{Orig: "がこの", Hira: "がこの", Kana: "ガコノ", Hepburn: "gakono", Kunrei: "gakono", Passport: "gakono"},
			{Orig: "子", Hira: "こ", Kana: "コ", Hepburn: "ko", Kunrei: "ko", Passport: "ko"},
			{Orig: "を", Hira: "を", Kana: "ヲ", Hepburn: "wo", Kunrei: "wo", Passport: "wo"},
			{Orig: "助け", Hira: "たすけ", Kana: "タスケ", Hepburn: "tasuke", Kunrei: "tasuke", Passport: "tasuke"},
			{Orig: "なきゃいけないってことだよね", Hira: "なきゃいけないってことだよね", Kana: "ナキャイケナイッテコトダヨネ", Hepburn: "nakyaikenaittekotodayone", Kunrei: "nakyaikenaittekotodayone", Passport: "nakyaikenaittekotodayone"},
		}},
		{"やったー", script.IConvertedSlice{{Orig: "やったー", Hira: "やったー", Kana: "ヤッター", Hepburn: "yattaa", Kunrei: "yattaa", Passport: "yattaa"}}},
		{"でっでー", script.IConvertedSlice{{Orig: "でっでー", Hira: "でっでー", Kana: "デッデー", Hepburn: "deddee", Kunrei: "deddee", Passport: "deddee"}}},
		{"てんさーふろー", script.IConvertedSlice{{Orig: "てんさーふろー", Hira: "てんさーふろー", Kana: "テンサーフロー", Hepburn: "tensaafuroo", Kunrei: "tensaafuroo", Passport: "tensaafuroo"}}},
		{"オレンジ色", script.IConvertedSlice{
			{Orig: "オレンジ", Hira: "おれんじ", Kana: "オレンジ", Hepburn: "orenji", Kunrei: "orenzi", Passport: "orenji"},
			{Orig: "色", Hira: "いろ", Kana: "イロ", Hepburn: "iro", Kunrei: "iro", Passport: "iro"},
		}},
		{"檸檬は、レモン色", script.IConvertedSlice{
			{Orig: "檸檬", Hira: "れもん", Kana: "レモン", Hepburn: "remon", Kunrei: "remon", Passport: "remon"},
			{Orig: "は、", Hira: "は、", Kana: "ハ、", Hepburn: "ha,", Kunrei: "ha,", Passport: "ha,"},
			{Orig: "レモン", Hira: "れもん", Kana: "レモン", Hepburn: "remon", Kunrei: "remon", Passport: "remon"},
			{Orig: "色", Hira: "いろ", Kana: "イロ", Hepburn: "iro", Kunrei: "iro", Passport: "iro"},
		}},
		{"私がこの子を助けなきゃいけないってことだよね", script.IConvertedSlice{
			{Orig: "私", Hira: "わたし", Kana: "ワタシ", Hepburn: "watashi", Kunrei: "watasi", Passport: "watashi"},
			{Orig: "がこの", Hira: "がこの", Kana: "ガコノ", Hepburn: "gakono", Kunrei: "gakono", Passport: "gakono"},
			{Orig: "子", Hira: "こ", Kana: "コ", Hepburn: "ko", Kunrei: "ko", Passport: "ko"},
			{Orig: "を", Hira: "を", Kana: "ヲ", Hepburn: "wo", Kunrei: "wo", Passport: "wo"},
			{Orig: "助け", Hira: "たすけ", Kana: "タスケ", Hepburn: "tasuke", Kunrei: "tasuke", Passport: "tasuke"},
			{Orig: "なきゃいけないってことだよね", Hira: "なきゃいけないってことだよね", Kana: "ナキャイケナイッテコトダヨネ", Hepburn: "nakyaikenaittekotodayone", Kunrei: "nakyaikenaittekotodayone", Passport: "nakyaikenaittekotodayone"},
		}},
		{"ｿｳｿﾞｸﾆﾝ", script.IConvertedSlice{{Orig: "ｿｳｿﾞｸﾆﾝ", Hira: "そうぞくにん", Kana: "ｿｳｿﾞｸﾆﾝ", Hepburn: "souzokunin", Kunrei: "souzokunin", Passport: "sozokunin"}}},
		{"思った 言った 行った", script.IConvertedSlice{
			{Orig: "思った", Hira: "おもった", Kana: "オモッタ", Hepburn: "omotta", Kunrei: "omotta", Passport: "omotta"},
			{Orig: " ", Hira: " ", Kana: " ", Hepburn: " ", Kunrei: " ", Passport: " "},
			{Orig: "言った", Hira: "いった", Kana: "イッタ", Hepburn: "itta", Kunrei: "itta", Passport: "itta"},
			{Orig: " ", Hira: " ", Kana: " ", Hepburn: " ", Kunrei: " ", Passport: " "},
			{Orig: "行った", Hira: "いった", Kana: "イッタ", Hepburn: "itta", Kunrei: "itta", Passport: "itta"},
		}},
		{"ﾞっ、", script.IConvertedSlice{
			{Orig: "ﾞ", Hira: "゛", Kana: "ﾞ", Hepburn: "\"", Kunrei: "゛", Passport: "゛"},
			{Orig: "っ、", Hira: "っ、", Kana: "ッ、", Hepburn: "tsu,", Kunrei: "tu,", Passport: "tsu,"},
		}},
		{"藍之介", script.IConvertedSlice{{Orig: "藍之介", Hira: "あいのすけ", Kana: "アイノスケ", Hepburn: "ainosuke", Kunrei: "ainosuke", Passport: "ainosuke"}}},
		{"藍水", script.IConvertedSlice{{Orig: "藍水", Hira: "らんすい", Kana: "ランスイ", Hepburn: "ransui", Kunrei: "ransui", Passport: "ransui"}}},
		{"見えますか？", script.IConvertedSlice{
			{Orig: "見え", Hira: "みえ", Kana: "ミエ", Hepburn: "mie", Kunrei: "mie", Passport: "mie"},
			{Orig: "ますか？", Hira: "ますか？", Kana: "マスカ？", Hepburn: "masuka？", Kunrei: "masuka？", Passport: "masuka？"},
		}},
		{"バニーちゃんちのシャワーノズルの先端", script.IConvertedSlice{
			{Orig: "バニー", Hira: "ばにー", Kana: "バニー", Hepburn: "banii", Kunrei: "banii", Passport: "banii"},
			{Orig: "ちゃんちの", Hira: "ちゃんちの", Kana: "チャンチノ", Hepburn: "chanchino", Kunrei: "tyantino", Passport: "chanchino"},
			{Orig: "シャワーノズル", Hira: "しゃわーのずる", Kana: "シャワーノズル", Hepburn: "shawaanozuru", Kunrei: "syawaanozuru", Passport: "shawaanozuru"},
			{Orig: "の", Hira: "の", Kana: "ノ", Hepburn: "no", Kunrei: "no", Passport: "no"},
			{Orig: "先端", Hira: "せんたん", Kana: "センタン", Hepburn: "sentan", Kunrei: "sentan", Passport: "sentan"},
		}},
		{"明日は明日の風が吹く", script.IConvertedSlice{
			{Orig: "明日", Hira: "あした", Kana: "アシタ", Hepburn: "ashita", Kunrei: "asita", Passport: "ashita"},
			{Orig: "は", Hira: "は", Kana: "ハ", Hepburn: "ha", Kunrei: "ha", Passport: "ha"},
			{Orig: "明日", Hira: "あした", Kana: "アシタ", Hepburn: "ashita", Kunrei: "asita", Passport: "ashita"},
			{Orig: "の", Hira: "の", Kana: "ノ", Hepburn: "no", Kunrei: "no", Passport: "no"},
			{Orig: "風", Hira: "かぜ", Kana: "カゼ", Hepburn: "kaze", Kunrei: "kaze", Passport: "kaze"},
			{Orig: "が", Hira: "が", Kana: "ガ", Hepburn: "ga", Kunrei: "ga", Passport: "ga"},
			{Orig: "吹く", Hira: "ふく", Kana: "フク", Hepburn: "fuku", Kunrei: "fuku", Passport: "fuku"},
		}},
		{"\uF862\u6709\u9650\u4F1A\u793E", script.IConvertedSlice{{Orig: "有限会社", Hira: "ゆうげんがいしゃ", Kana: "ユウゲンガイシャ", Hepburn: "yuugengaisha", Kunrei: "yuugengaisya", Passport: "yuugengaisha"}}},
		{"三\u00D7五", script.IConvertedSlice{
			{Orig: "三", Hira: "さん", Kana: "サン", Hepburn: "san", Kunrei: "san", Passport: "san"},
			{Orig: "×", Hira: "×", Kana: "×", Hepburn: "x", Kunrei: "x", Passport: "x"},
			{Orig: "五", Hira: "ご", Kana: "ゴ", Hepburn: "go", Kunrei: "go", Passport: "go"},
		}},
		{"日本国民は、正当に選挙された国会における代表者を通じて行動し、われらとわれらの子孫のために、" +
			"諸国民との協和による成果と、わが国全土にわたつて自由のもたらす恵沢を確保し、" +
			"政府の行為によつて再び戦争の惨禍が起ることのないやうにすることを決意し、ここに主権が国民に存することを宣言し、" +
			"この憲法を確定する。そもそも国政は、国民の厳粛な信託によるものであつて、その権威は国民に由来し、" +
			"その権力は国民の代表者がこれを行使し、その福利は国民がこれを享受する。これは人類普遍の原理であり、" +
			"この憲法は、かかる原理に基くものである。われらは、これに反する一切の憲法、法令及び詔勅を排除する。", script.IConvertedSlice{
			{Orig: "日本国民", Hira: "にほんこくみん", Kana: "ニホンコクミン", Hepburn: "nihonkokumin", Kunrei: "nihonkokumin", Passport: "nihonkokumin"},
			{Orig: "は、", Hira: "は、", Kana: "ハ、", Hepburn: "ha,", Kunrei: "ha,", Passport: "ha,"},
			{Orig: "正当", Hira: "せいとう", Kana: "セイトウ", Hepburn: "seitou", Kunrei: "seitou", Passport: "seito"},
			{Orig: "に", Hira: "に", Kana: "ニ", Hepburn: "ni", Kunrei: "ni", Passport: "ni"},
			{Orig: "選挙", Hira: "せんきょ", Kana: "センキョ", Hepburn: "senkyo", Kunrei: "senkyo", Passport: "senkyo"},
			{Orig: "された", Hira: "された", Kana: "サレタ", Hepburn: "sareta", Kunrei: "sareta", Passport: "sareta"},
			{Orig: "国会", Hira: "こっかい", Kana: "コッカイ", Hepburn: "kokkai", Kunrei: "kokkai", Passport: "kokkai"},
			{Orig: "における", Hira: "における", Kana: "ニオケル", Hepburn: "niokeru", Kunrei: "niokeru", Passport: "niokeru"},
			{Orig: "代表者", Hira: "だいひょうしゃ", Kana: "ダイヒョウシャ", Hepburn: "daihyousha", Kunrei: "daihyousya", Passport: "daihyousha"},
			{Orig: "を", Hira: "を", Kana: "ヲ", Hepburn: "wo", Kunrei: "wo", Passport: "wo"},
			{Orig: "通じ", Hira: "つうじ", Kana: "ツウジ", Hepburn: "tsuuji", Kunrei: "tuuzi", Passport: "tsuuji"},
			{Orig: "て", Hira: "て", Kana: "テ", Hepburn: "te", Kunrei: "te", Passport: "te"},
			{Orig: "行動", Hira: "こうどう", Kana: "コウドウ", Hepburn: "koudou", Kunrei: "koudou", Passport: "kodou"},
			{Orig: "し、", Hira: "し、", Kana: "シ、", Hepburn: "shi,", Kunrei: "si,", Passport: "shi,"},
			{Orig: "われらとわれらの", Hira: "われらとわれらの", Kana: "ワレラトワレラノ", Hepburn: "wareratowarerano", Kunrei: "wareratowarerano", Passport: "wareratowarerano"},
			{Orig: "子孫", Hira: "しそん", Kana: "シソン", Hepburn: "shison", Kunrei: "sison", Passport: "shison"},
			{Orig: "のために、", Hira: "のために、", Kana: "ノタメニ、", Hepburn: "notameni,", Kunrei: "notameni,", Passport: "notameni,"},
			{Orig: "諸国民", Hira: "しょこくみん", Kana: "ショコクミン", Hepburn: "shokokumin", Kunrei: "syokokumin", Passport: "shokokumin"},
			{Orig: "との", Hira: "との", Kana: "トノ", Hepburn: "tono", Kunrei: "tono", Passport: "tono"},
			{Orig: "協和", Hira: "きょうわ", Kana: "キョウワ", Hepburn: "kyouwa", Kunrei: "kyouwa", Passport: "kyouwa"},
			{Orig: "による", Hira: "による", Kana: "ニヨル", Hepburn: "niyoru", Kunrei: "niyoru", Passport: "niyoru"},
			{Orig: "成果", Hira: "せいか", Kana: "セイカ", Hepburn: "seika", Kunrei: "seika", Passport: "seika"},
			{Orig: "と、", Hira: "と、", Kana: "ト、", Hepburn: "to,", Kunrei: "to,", Passport: "to,"},
			{Orig: "わが", Hira: "わが", Kana: "ワガ", Hepburn: "waga", Kunrei: "waga", Passport: "waga"},
			{Orig: "国", Hira: "くに", Kana: "クニ", Hepburn: "kuni", Kunrei: "kuni", Passport: "kuni"},
			{Orig: "全土", Hira: "ぜんど", Kana: "ゼンド", Hepburn: "zendo", Kunrei: "zendo", Passport: "zendo"},
			{Orig: "にわたつて", Hira: "にわたつて", Kana: "ニワタツテ", Hepburn: "niwatatsute", Kunrei: "niwatatute", Passport: "niwatatsute"},
			{Orig: "自由", Hira: "じゆう", Kana: "ジユウ", Hepburn: "jiyuu", Kunrei: "ziyuu", Passport: "jiyuu"},
			{Orig: "のもたらす", Hira: "のもたらす", Kana: "ノモタラス", Hepburn: "nomotarasu", Kunrei: "nomotarasu", Passport: "nomotarasu"},
			{Orig: "恵沢", Hira: "けいたく", Kana: "ケイタク", Hepburn: "keitaku", Kunrei: "keitaku", Passport: "keitaku"},
			{Orig: "を", Hira: "を", Kana: "ヲ", Hepburn: "wo", Kunrei: "wo", Passport: "wo"},
			{Orig: "確保", Hira: "かくほ", Kana: "カクホ", Hepburn: "kakuho", Kunrei: "kakuho", Passport: "kakuho"},
			{Orig: "し、", Hira: "し、", Kana: "シ、", Hepburn: "shi,", Kunrei: "si,", Passport: "shi,"},
			{Orig: "政府", Hira: "せいふ", Kana: "セイフ", Hepburn: "seifu", Kunrei: "seifu", Passport: "seifu"},
			{Orig: "の", Hira: "の", Kana: "ノ", Hepburn: "no", Kunrei: "no", Passport: "no"},
			{Orig: "行為", Hira: "こうい", Kana: "コウイ", Hepburn: "koui", Kunrei: "koui", Passport: "koi"},
			{Orig: "によつて", Hira: "によつて", Kana: "ニヨツテ", Hepburn: "niyotsute", Kunrei: "niyotute", Passport: "niyotsute"},
			{Orig: "再び", Hira: "ふたたび", Kana: "フタタビ", Hepburn: "futatabi", Kunrei: "futatabi", Passport: "futatabi"},
			{Orig: "戦争", Hira: "せんそう", Kana: "センソウ", Hepburn: "sensou", Kunrei: "sensou", Passport: "senso"},
			{Orig: "の", Hira: "の", Kana: "ノ", Hepburn: "no", Kunrei: "no", Passport: "no"},
			{Orig: "惨禍", Hira: "さんか", Kana: "サンカ", Hepburn: "sanka", Kunrei: "sanka", Passport: "sanka"},
			{Orig: "が", Hira: "が", Kana: "ガ", Hepburn: "ga", Kunrei: "ga", Passport: "ga"},
			{Orig: "起る", Hira: "おこる", Kana: "オコル", Hepburn: "okoru", Kunrei: "okoru", Passport: "okoru"},
			{Orig: "ことのないやうにすることを", Hira: "ことのないやうにすることを", Kana: "コトノナイヤウニスルコトヲ", Hepburn: "kotononaiyaunisurukotowo", Kunrei: "kotononaiyaunisurukotowo", Passport: "kotononaiyaunisurukotowo"},
			{Orig: "決意", Hira: "けつい", Kana: "ケツイ", Hepburn: "ketsui", Kunrei: "ketui", Passport: "ketsui"},
			{Orig: "し、", Hira: "し、", Kana: "シ、", Hepburn: "shi,", Kunrei: "si,", Passport: "shi,"},
			{Orig: "ここに", Hira: "ここに", Kana: "ココニ", Hepburn: "kokoni", Kunrei: "kokoni", Passport: "kokoni"},
			{Orig: "主権", Hira: "しゅけん", Kana: "シュケン", Hepburn: "shuken", Kunrei: "syuken", Passport: "shuken"},
			{Orig: "が", Hira: "が", Kana: "ガ", Hepburn: "ga", Kunrei: "ga", Passport: "ga"},
			{Orig: "国民", Hira: "こくみん", Kana: "コクミン", Hepburn: "kokumin", Kunrei: "kokumin", Passport: "kokumin"},
			{Orig: "に", Hira: "に", Kana: "ニ", Hepburn: "ni", Kunrei: "ni", Passport: "ni"},
			{Orig: "存す", Hira: "そんす", Kana: "ソンス", Hepburn: "sonsu", Kunrei: "sonsu", Passport: "sonsu"},
			{Orig: "ることを", Hira: "ることを", Kana: "ルコトヲ", Hepburn: "rukotowo", Kunrei: "rukotowo", Passport: "rukotowo"},
			{Orig: "宣言", Hira: "せんげん", Kana: "センゲン", Hepburn: "sengen", Kunrei: "sengen", Passport: "sengen"},
			{Orig: "し、", Hira: "し、", Kana: "シ、", Hepburn: "shi,", Kunrei: "si,", Passport: "shi,"},
			{Orig: "この", Hira: "この", Kana: "コノ", Hepburn: "kono", Kunrei: "kono", Passport: "kono"},
			{Orig: "憲法", Hira: "けんぽう", Kana: "ケンポウ", Hepburn: "kenpou", Kunrei: "kenpou", Passport: "kempou"},
			{Orig: "を", Hira: "を", Kana: "ヲ", Hepburn: "wo", Kunrei: "wo", Passport: "wo"},
			{Orig: "確定す", Hira: "かくていす", Kana: "カクテイス", Hepburn: "kakuteisu", Kunrei: "kakuteisu", Passport: "kakuteisu"},
			{Orig: "る。", Hira: "る。", Kana: "ル。", Hepburn: "ru.", Kunrei: "ru.", Passport: "ru."},
			{Orig: "そもそも", Hira: "そもそも", Kana: "ソモソモ", Hepburn: "somosomo", Kunrei: "somosomo", Passport: "somosomo"},
			{Orig: "国政", Hira: "こくせい", Kana: "コクセイ", Hepburn: "kokusei", Kunrei: "kokusei", Passport: "kokusei"},
			{Orig: "は、", Hira: "は、", Kana: "ハ、", Hepburn: "ha,", Kunrei: "ha,", Passport: "ha,"},
			{Orig: "国民", Hira: "こくみん", Kana: "コクミン", Hepburn: "kokumin", Kunrei: "kokumin", Passport: "kokumin"},
			{Orig: "の", Hira: "の", Kana: "ノ", Hepburn: "no", Kunrei: "no", Passport: "no"},
			{Orig: "厳粛", Hira: "げんしゅく", Kana: "ゲンシュク", Hepburn: "genshuku", Kunrei: "gensyuku", Passport: "genshuku"},
			{Orig: "な", Hira: "な", Kana: "ナ", Hepburn: "na", Kunrei: "na", Passport: "na"},
			{Orig: "信託", Hira: "しんたく", Kana: "シンタク", Hepburn: "shintaku", Kunrei: "sintaku", Passport: "shintaku"},
			{Orig: "によるものであつて、", Hira: "によるものであつて、", Kana: "ニヨルモノデアツテ、", Hepburn: "niyorumonodeatsute,", Kunrei: "niyorumonodeatute,", Passport: "niyorumonodeatsute,"},
			{Orig: "その", Hira: "その", Kana: "ソノ", Hepburn: "sono", Kunrei: "sono", Passport: "sono"},
			{Orig: "権威", Hira: "けんい", Kana: "ケンイ", Hepburn: "ken'i", Kunrei: "ken'i", Passport: "keni"},
			{Orig: "は", Hira: "は", Kana: "ハ", Hepburn: "ha", Kunrei: "ha", Passport: "ha"},
			{Orig: "国民", Hira: "こくみん", Kana: "コクミン", Hepburn: "kokumin", Kunrei: "kokumin", Passport: "kokumin"},
			{Orig: "に", Hira: "に", Kana: "ニ", Hepburn: "ni", Kunrei: "ni", Passport: "ni"},
			{Orig: "由来", Hira: "ゆらい", Kana: "ユライ", Hepburn: "yurai", Kunrei: "yurai", Passport: "yurai"},
			{Orig: "し、", Hira: "し、", Kana: "シ、", Hepburn: "shi,", Kunrei: "si,", Passport: "shi,"},
			{Orig: "その", Hira: "その", Kana: "ソノ", Hepburn: "sono", Kunrei: "sono", Passport: "sono"},
			{Orig: "権力", Hira: "けんりょく", Kana: "ケンリョク", Hepburn: "kenryoku", Kunrei: "kenryoku", Passport: "kenryoku"},
			{Orig: "は", Hira: "は", Kana: "ハ", Hepburn: "ha", Kunrei: "ha", Passport: "ha"},
			{Orig: "国民", Hira: "こくみん", Kana: "コクミン", Hepburn: "kokumin", Kunrei: "kokumin", Passport: "kokumin"},
			{Orig: "の", Hira: "の", Kana: "ノ", Hepburn: "no", Kunrei: "no", Passport: "no"},
			{Orig: "代表者", Hira: "だいひょうしゃ", Kana: "ダイヒョウシャ", Hepburn: "daihyousha", Kunrei: "daihyousya", Passport: "daihyousha"},
			{Orig: "がこれを", Hira: "がこれを", Kana: "ガコレヲ", Hepburn: "gakorewo", Kunrei: "gakorewo", Passport: "gakorewo"},
			{Orig: "行使", Hira: "こうし", Kana: "コウシ", Hepburn: "koushi", Kunrei: "kousi", Passport: "koshi"},
			{Orig: "し、", Hira: "し、", Kana: "シ、", Hepburn: "shi,", Kunrei: "si,", Passport: "shi,"},
			{Orig: "その", Hira: "その", Kana: "ソノ", Hepburn: "sono", Kunrei: "sono", Passport: "sono"},
			{Orig: "福利", Hira: "ふくり", Kana: "フクリ", Hepburn: "fukuri", Kunrei: "fukuri", Passport: "fukuri"},
			{Orig: "は", Hira: "は", Kana: "ハ", Hepburn: "ha", Kunrei: "ha", Passport: "ha"},
			{Orig: "国民", Hira: "こくみん", Kana: "コクミン", Hepburn: "kokumin", Kunrei: "kokumin", Passport: "kokumin"},
			{Orig: "がこれを", Hira: "がこれを", Kana: "ガコレヲ", Hepburn: "gakorewo", Kunrei: "gakorewo", Passport: "gakorewo"},
			{Orig: "享受", Hira: "きょうじゅ", Kana: "キョウジュ", Hepburn: "kyouju", Kunrei: "kyouju", Passport: "kyouju"},
			{Orig: "する。", Hira: "する。", Kana: "スル。", Hepburn: "suru.", Kunrei: "suru.", Passport: "suru."},
			{Orig: "これは", Hira: "これは", Kana: "コレハ", Hepburn: "koreha", Kunrei: "koreha", Passport: "koreha"},
			{Orig: "人類普遍", Hira: "じんるいふへん", Kana: "ジンルイフヘン", Hepburn: "jinruifuhen", Kunrei: "zinruifuhen", Passport: "jinruifuhen"},
			{Orig: "の", Hira: "の", Kana: "ノ", Hepburn: "no", Kunrei: "no", Passport: "no"},
			{Orig: "原理", Hira: "げんり", Kana: "ゲンリ", Hepburn: "genri", Kunrei: "genri", Passport: "genri"},
			{Orig: "であり、", Hira: "であり、", Kana: "デアリ、", Hepburn: "deari,", Kunrei: "deari,", Passport: "deari,"},
			{Orig: "この", Hira: "この", Kana: "コノ", Hepburn: "kono", Kunrei: "kono", Passport: "kono"},
			{Orig: "憲法", Hira: "けんぽう", Kana: "ケンポウ", Hepburn: "kenpou", Kunrei: "kenpou", Passport: "kempou"},
			{Orig: "は、", Hira: "は、", Kana: "ハ、", Hepburn: "ha,", Kunrei: "ha,", Passport: "ha,"},
			{Orig: "かかる", Hira: "かかる", Kana: "カカル", Hepburn: "kakaru", Kunrei: "kakaru", Passport: "kakaru"},
			{Orig: "原理", Hira: "げんり", Kana: "ゲンリ", Hepburn: "genri", Kunrei: "genri", Passport: "genri"},
			{Orig: "に", Hira: "に", Kana: "ニ", Hepburn: "ni", Kunrei: "ni", Passport: "ni"},
			{Orig: "基く", Hira: "もとづく", Kana: "モトヅク", Hepburn: "motozuku", Kunrei: "motozuku", Passport: "motozuku"},
			{Orig: "ものである。", Hira: "ものである。", Kana: "モノデアル。", Hepburn: "monodearu.", Kunrei: "monodearu.", Passport: "monodearu."},
			{Orig: "われらは、", Hira: "われらは、", Kana: "ワレラハ、", Hepburn: "wareraha,", Kunrei: "wareraha,", Passport: "wareraha,"},
			{Orig: "これに", Hira: "これに", Kana: "コレニ", Hepburn: "koreni", Kunrei: "koreni", Passport: "koreni"},
			{Orig: "反す", Hira: "はんす", Kana: "ハンス", Hepburn: "hansu", Kunrei: "hansu", Passport: "hansu"},
			{Orig: "る", Hira: "る", Kana: "ル", Hepburn: "ru", Kunrei: "ru", Passport: "ru"},
			{Orig: "一切", Hira: "いっさい", Kana: "イッサイ", Hepburn: "issai", Kunrei: "issai", Passport: "issai"},
			{Orig: "の", Hira: "の", Kana: "ノ", Hepburn: "no", Kunrei: "no", Passport: "no"},
			{Orig: "憲法、", Hira: "けんぽう、", Kana: "ケンポウ、", Hepburn: "kenpou,", Kunrei: "kenpou,", Passport: "kempou,"},
			{Orig: "法令", Hira: "ほうれい", Kana: "ホウレイ", Hepburn: "hourei", Kunrei: "hourei", Passport: "hourei"},
			{Orig: "及び", Hira: "および", Kana: "オヨビ", Hepburn: "oyobi", Kunrei: "oyobi", Passport: "oyobi"},
			{Orig: "詔勅", Hira: "しょうちょく", Kana: "ショウチョク", Hepburn: "shouchoku", Kunrei: "syoutyoku", Passport: "shouchoku"},
			{Orig: "を", Hira: "を", Kana: "ヲ", Hepburn: "wo", Kunrei: "wo", Passport: "wo"},
			{Orig: "排除", Hira: "はいじょ", Kana: "ハイジョ", Hepburn: "haijo", Kunrei: "haijo", Passport: "haijo"},
			{Orig: "する。", Hira: "する。", Kana: "スル。", Hepburn: "suru.", Kunrei: "suru.", Passport: "suru."},
		}},
	} {
		t.Run(fmt.Sprintf("test#%02d", testID), func(t *testing.T) {
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
				return
			}

			t.Log(converted.Furiganize())
			t.Log(converted.Romanize())

			testID++
		})
	}
}

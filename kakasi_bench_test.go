package kakasi

import "testing"

var benchText = "" +
	"日本国民は、正当に選挙された国会における代表者を通じて行動し、われらとわれらの子孫のために、" +
	"諸国民との協和による成果と、わが国全土にわたつて自由のもたらす恵沢を確保し、政府の行為によつて" +
	"再び戦争の惨禍が起ることのないやうにすることを決意し、ここに主権が国民に存することを宣言し、" +
	"この憲法を確定する。Alphabet 123 and 漢字とひらがな交じり文、オレンジ色の檸檬は、レモン色。" +
	"私がこの子を助けなきゃいけないってことだよね ｿｳｿﾞｸﾆﾝ てんさーふろー でっでー やったー"

func BenchmarkNewKakasi(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if _, err := NewKakasi(); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkConvert(b *testing.B) {
	k, err := NewKakasi()
	if err != nil {
		b.Fatal(err)
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := k.Convert(benchText); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkNormalize(b *testing.B) {
	var k Kakasi

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if _, err := k.Normalize(benchText); err != nil {
			b.Fatal(err)
		}
	}
}

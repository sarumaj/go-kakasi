[![test_and_report](https://github.com/sarumaj/go-kakasi/actions/workflows/test_and_report.yml/badge.svg)](https://github.com/sarumaj/go-kakasi/actions/workflows/test_and_report.yml)
[![build_and_release](https://github.com/sarumaj/go-kakasi/actions/workflows/build_and_release.yml/badge.svg)](https://github.com/sarumaj/go-kakasi/actions/workflows/build_and_release.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/sarumaj/go-kakasi)](https://goreportcard.com/report/github.com/sarumaj/go-kakasi)
[![Maintainability](https://qlty.sh/gh/sarumaj/projects/go-kakasi/maintainability.svg)](https://qlty.sh/gh/sarumaj/projects/go-kakasi)
[![Code Coverage](https://qlty.sh/gh/sarumaj/projects/go-kakasi/coverage.svg)](https://qlty.sh/gh/sarumaj/projects/go-kakasi)
[![Go Reference](https://pkg.go.dev/badge/github.com/sarumaj/go-kakasi.svg)](https://pkg.go.dev/github.com/sarumaj/go-kakasi)
[![Go version](https://img.shields.io/github/go-mod/go-version/sarumaj/go-kakasi?logo=go&label=&labelColor=gray)](https://go.dev)

---

# go-kakasi

Natural Language Processing Library for Japanese.
Based on the work of Hiroshi Miura: [pykakasi](https://codeberg.org/miurahr/pykakasi), who transcoded [kakasi](http://kakasi.namazu.org/index.html.en) library to Python.

There already exists a C-binding-library of [kakasi](http://kakasi.namazu.org/index.html.en): [go-kakasi](https://github.com/ysugimoto/go-kakasi) developed by Yoshiaki Sugimoto.
This library is a pure Go implementation of [kakasi](http://kakasi.namazu.org/index.html.en), which is not platform-dependent and does not require CGO to compile.

## Usage

```Go
package main

import (
    "fmt"

    "github.com/sarumaj/go-kakasi"
)

func main() {
    k, err := kakasi.NewKakasi()
    if err != nil {
        panic(err)
    }

    converted, err := k.Convert("日本国民は、正当に選挙された国会における代表者を通じて行動し、われらとわれらの子孫のために、" +
        "諸国民との協和による成果と、わが国全土にわたつて自由のもたらす恵沢を確保し、" +
        "政府の行為によつて再び戦争の惨禍が起ることのないやうにすることを決意し、ここに主権が国民に存することを宣言し、" +
        "この憲法を確定する。そもそも国政は、国民の厳粛な信託によるものであつて、その権威は国民に由来し、" +
        "その権力は国民の代表者がこれを行使し、その福利は国民がこれを享受する。これは人類普遍の原理であり、" +
        "この憲法は、かかる原理に基くものである。われらは、これに反する一切の憲法、法令及び詔勅を排除する。")
    if err != nil {
        panic(err)
    }

    // Prints:
    // 日本国民[にほんこくみん]は、正当[せいとう]に選挙[せんきょ]された国会[こっかい]における代表者[だいひょうしゃ]を
    // 通じ[つうじ]て行動[こうどう]し、われらとわれらの子孫[しそん]のために、諸国民[しょこくみん]との協和[きょうわ]による
    // 成果[せいか]と、わが国[くに]全土[ぜんど]にわたつて自由[じゆう]のもたらす恵沢[けいたく]を確保[かくほ]し、政府[せいふ]の
    // 行為[こうい]によつて再び[ふたたび]戦争[せんそう]の惨禍[さんか]が起る[おこる]ことのないやうにすることを決意[けつい]し、
    // ここに主権[しゅけん]が国民[こくみん]に存す[そんす]ることを宣言[せんげん]し、この憲法[けんぽう]を確定す[かくていす]る。
    // そもそも国政[こくせい]は、国民[こくみん]の厳粛[げんしゅく]な信託[しんたく]によるものであつて、その権威[けんい]は
    // 国民[こくみん]に由来[ゆらい]し、その権力[けんりょく]は国民[こくみん]の代表者[だいひょうしゃ]がこれを行使[こうし]し、
    // その福利[ふくり]は国民[こくみん]がこれを享受[きょうじゅ]する。これは人類普遍[じんるいふへん]の原理[げんり]であり、
    // この憲法[けんぽう]は、かかる原理[げんり]に基く[もとづく]ものである。われらは、これに反す[はんす]る一切[いっさい]の
    // 憲法[けんぽう]、法令[ほうれい]及び[および]詔勅[しょうちょく]を排除[はいじょ]する。
    fmt.Println(converted.Furiganize())

    // Prints:
    // nihonkokumin ha, seitou ni senkyo sareta kokkai niokeru daihyousha wo tsuuji te koudou shi,
    // wareratowarerano shison notameni, shokokumin tono kyouwa niyoru seika to, waga kuni zendo
    // niwatatsute jiyuu nomotarasu keitaku wo kakuho shi, seifu no koui niyotsute futatabi sensou
    // no sanka ga okoru kotononaiyaunisurukotowo ketsui shi, kokoni shuken ga kokumin ni sonsu rukotowo
    // sengen shi, kono kenpou wo kakuteisu ru. somosomo kokusei ha, kokumin no genshuku na shintaku
    // niyorumonodeatsute, sono ken'i ha kokumin ni yurai shi, sono kenryoku ha kokumin no daihyousha
    // gakorewo koushi shi, sono fukuri ha kokumin gakorewo kyouju suru. koreha jinruifuhen no genri deari,
    // kono kenpou ha, kakaru genri ni motozuku monodearu. wareraha, koreni hansu ru issai no kenpou,
    // hourei oyobi shouchoku wo haijo suru.
    fmt.Println(converted.Romanize())

    // Prints:
    // {
    //     Orig: "日本国民",
    //     Hira: "にほんこくみん",
    //     Kana: "ニホンコクミン",
    //     Hepburn: "nihonkokumin",
    //     Kunrei: "nihonkokumin",
    //     Passport: "nihonkokumin",
    // }
    fmt.Println(converted[0])
}
```

## Projects

- [bing-wallpaper-changer](https://github.com/sarumaj/bing-wallpaper-changer) uses **go-kakasi** to add Furigana annotations to image descriptions for the Japanese Bing wallpapers.

## Copyright and License

**PyKakasi:**

Copyright (C) 2010-2021 Hiroshi Miura and his contributors (see [AUTHORS](internal/codegen/data/AUTHORS.md))

**KAKASI Dictionary**:

Copyright (C) 2010-2021 Hiroshi Miura and his contributors (see [AUTHORS](internal/codegen/data/AUTHORS.md))

Copyright (C) 1992 1993 1994 Hironobu Takahashi, Masahiko Sato, Yukiyoshi Kameyama, Miki Inooka, Akihiko Sasaki, Dai Ando, Junichi Okukawa, Katsushi Sato and Nobuhiro Yamagishi

**UniDic**:

Copyright (c) 2011-2021, The UniDic Consortium

All rights reserved.

Unidic is released under any of the GPL2, the LGPL2.1, or the 3-clause BSD License (see [internal/codegen/data/LICENSE](internal/codegen/data/LICENSE)). [pykakasi](https://codeberg.org/miurahr/pykakasi) and this project relicense a part of the unidic with GPL3+.

## Contribution

Since I am not a native speaker, feel free to open issues/pull requests anytime. I am grateful for any help. よろしくお願いします。

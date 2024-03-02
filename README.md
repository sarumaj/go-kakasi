# go-kakasi

Natural Language Processing Library for Japanese.
Based on the work of Hiroshi Miuara [pykakasi](https://codeberg.org/miurahr/pykakasi).

There exists a C binding library of [kakasi](http://kakasi.namazu.org/index.html.en): [go-kakasi](https://github.com/ysugimoto/go-kakasi).
This library is a pure Go implementation of [kakasi](http://kakasi.namazu.org/index.html.en), which is not platform dependent and does not require CGO to compile.

## Usage

```Go
package main

import (
    "fmt"

    "github/sarumaj/go-kakasi"
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

    // Prints: 日本国民(にほんこくみん)は、正当(せいとう)に選挙(せんきょ)された国会(こっかい)における代表者(だいひょうしゃ)を通じ(つうじ)て行動(こうどう)し、われらとわれらの子孫(しそん)のために、諸国民(しょこくみん)との協和(きょうわ)による成果(せいか)と、わが国(くに)全土(ぜんど)にわたつて自由(じゆう)のもたらす恵沢(けいたく)を確保(かくほ)し、政府(せいふ)の行為(こうい)によつて再び(ふたたび)戦争(せんそう)の惨禍(さんか)が起る(おこる)ことのないやうにすることを決意(けつい)し、ここに主権(しゅけん)が国民(こくみん)に存す(そんす)ることを宣言(せんげん)し、この憲法(けんぽう)を確定す(かくていす)る。そもそも国政(こくせい)は、国民(こくみん)の厳粛(げんしゅく)な信託(しんたく)によるものであつて、その権威(けんい)は国民(こくみん)に由来(ゆらい)し、その権力(けんりょく)は国民(こくみん)の代表者(だいひょうしゃ)がこれを行使(こうし)し、その福利(ふくり)は国民(こくみん)がこれを享受(きょうじゅ)する。これは人類普遍(じんるいふへん)の原理(げんり)であり、この憲法(けんぽう)は、かかる原理(げんり)に基く(もとづく)ものである。われらは、これに反す(はんす)る一切(いっさい)の憲法、(けんぽう、)法令(ほうれい)及び(および)詔勅(しょうちょく)を排除(はいじょ)する。
    fmt.Println(converted.Furiganize())
}
```

## Copyright and License

**PyKakasi:**
Copyright (C) 2010-2021 Hiroshi Miura and his contributors (see [AUTHORS](https://codeberg.org/miurahr/pykakasi/src/branch/master/AUTHORS))

**KAKASI Dictionary**:
Copyright (C) 2010-2021 Hiroshi Miura and his contributors (see [AUTHORS](https://codeberg.org/miurahr/pykakasi/src/branch/master/AUTHORS))

Copyright (C) 1992 1993 1994 Hironobu Takahashi, Masahiko Sato, Yukiyoshi Kameyama, Miki Inooka, Akihiko Sasaki, Dai Ando, Junichi Okukawa, Katsushi Sato and Nobuhiro Yamagishi

**UniDic**:
Copyright (c) 2011-2021, The UniDic Consortium

All rights reserved.

Unidic is released under any of the GPL2, the LGPL2.1, or the 3-clause BSD License. (See src/data/unidic/BSD.txt) PyKakasi relicenses a part of the unidic with GPL3+.

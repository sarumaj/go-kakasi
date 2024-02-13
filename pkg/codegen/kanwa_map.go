package codegen

import (
	"context"
	"fmt"
	"os"
	"strings"
)

// cLetters is a map of runes to a list of strings.
// It is used to generate the kanwa map.
// Each rune represents a list of sounds in the Japanese language.
var cLetters = map[rune][]string{
	'a': {"あ", "ぁ", "っ", "わ", "ゎ"},
	'i': {"い", "ぃ", "っ", "ゐ"},
	'u': {"う", "ぅ", "っ"},
	'e': {"え", "ぇ", "っ", "ゑ"},
	'o': {"お", "ぉ", "っ"},
	'k': {"か", "ゕ", "き", "く", "け", "ゖ", "こ", "っ"},
	'g': {"が", "ぎ", "ぐ", "げ", "ご", "っ"},
	's': {"さ", "し", "す", "せ", "そ", "っ"},
	'z': {"ざ", "じ", "ず", "ぜ", "ぞ", "っ"},
	'j': {"ざ", "じ", "ず", "ぜ", "ぞ", "っ"},
	't': {"た", "ち", "つ", "て", "と", "っ"},
	'd': {"だ", "ぢ", "づ", "で", "ど", "っ"},
	'c': {"ち", "っ"},
	'n': {"な", "に", "ぬ", "ね", "の", "ん"},
	'h': {"は", "ひ", "ふ", "へ", "ほ", "っ"},
	'b': {"ば", "び", "ぶ", "べ", "ぼ", "っ"},
	'f': {"ふ", "っ"},
	'p': {"ぱ", "ぴ", "ぷ", "ぺ", "ぽ", "っ"},
	'm': {"ま", "み", "む", "め", "も"},
	'y': {"や", "ゃ", "ゆ", "ゅ", "よ", "ょ"},
	'r': {"ら", "り", "る", "れ", "ろ"},
	'w': {"わ", "ゐ", "ゑ", "ゎ", "を", "っ"},
	'v': {"ゔ"},
}

// kanwaMapResources is a map of target and source files.
// The target file is the destination file.
var kanwaMapResources = map[string][]string{
	"kanwadict4.json": {
		"data/kakasidict.utf8",
		"data/unidict_noun.utf8",
		"data/unidict_adj.utf8",
	},
}

// kanjiCtx is a kanji character or a phrase in the Japanese language.
// It has a yomi and a list of contexts.
// yomi is the reading of the kanji character or phrase.
// ctx is a list of contexts in which the kanji character or phrase is used.
type kanjiCtx [2]any

func (m kanjiCtx) AppendCtx(v ...string) { r, _ := m[1].([]string); m[1] = append(r, v...) }
func (m kanjiCtx) GetCtx() []string      { r, _ := m[1].([]string); return r }
func (m kanjiCtx) GetYomi() string       { r, _ := m[0].(string); return r }
func (m *kanjiCtx) SetCtx(v []string)    { m[1] = v }
func (m *kanjiCtx) SetYomi(v string)     { m[0] = v }

// kanjiCtxMap is a map of strings to kanjiCtx.
// It is used to generate the kanwa map.
// Each string represents a kanji character.
type kanjiCtxMap map[string]kanjiCtx

func (m kanjiCtxMap) Get(k string) kanjiCtx    { return mapGet(m, k) }
func (m kanjiCtxMap) Has(k string) bool        { return mapHas(m, k) }
func (m kanjiCtxMap) Keys() []string           { return mapKeys(m) }
func (m kanjiCtxMap) Set(k string, v kanjiCtx) { m = mapSet(m, k, v) }

// kanwaMap is a map of runes to a map of strings to kanjiCtx.
// It is used to generate the kanwa map.
// Each rune represents a beginning of a kanji character or a phrase in the Japanese language.
type kanwaMap map[rune]kanjiCtxMap

func (m kanwaMap) Get(k rune) kanjiCtxMap    { return mapGet(m, k) }
func (m kanwaMap) Has(k rune) bool           { return mapHas(m, k) }
func (m kanwaMap) Keys() []rune              { return mapKeys(m) }
func (m kanwaMap) Set(k rune, v kanjiCtxMap) { m = mapSet(m, k, v) }

// parseLine parses a line from a file and updates the kanwa map.
// The line is expected to have the format "yomi kanji [ctx ...]".
// The yomi is the reading of the kanji character or phrase.
// The kanji is the kanji character or phrase.
func (m kanwaMap) parseLine(line string) {
	token := strings.Split(line, " ")
	yomi, kanji := token[0], token[1]

	yomi_runes := []rune(yomi)

	var tail []rune
	if yomi_runes[len(yomi_runes)-1] <= 'z' {
		tail = yomi_runes[len(yomi_runes)-1:]
		yomi_runes = yomi_runes[: len(yomi_runes)-1 : len(yomi_runes)-1]
	}

	var token_ctx []string
	if len(token) > 2 {
		token_ctx = token[2:]
	}

	m.update(kanji, string(yomi_runes), string(tail), token_ctx...)
}

// update updates the kanwa map with a kanji character or phrase.
// The kanji is the kanji character or phrase.
// The yomi is the reading of the kanji character or phrase.
// The tail is the last character of the reading.
// The token_ctx is a list of contexts in which the kanji character or phrase is used.
func (m kanwaMap) update(kanji, yomi, tail string, token_ctx ...string) {
	if len(tail) == 0 {
		kanji_runes := []rune(kanji)
		c := kanji_runes[0]
		if !m.Has(c) {
			m.Set(c, kanjiCtxMap{kanji: kanjiCtx{yomi, token_ctx}})
			return
		}

		gotKanjiCtxMap := m.Get(c)
		if gotKanjiCtxMap.Has(kanji) {
			gotKanjiCtx := gotKanjiCtxMap.Get(kanji)
			gotKanjiCtx.SetYomi(yomi)
			gotKanjiCtx.AppendCtx(token_ctx...)
			gotKanjiCtxMap.Set(kanji, gotKanjiCtx)

		} else {
			gotKanjiCtxMap.Set(kanji, kanjiCtx{yomi, token_ctx})

		}

		m.Set(c, gotKanjiCtxMap)
		return
	}

	got, ok := cLetters[rune(tail[0])]
	if !ok {
		return
	}

	for _, v := range got {
		m.update(kanji+v, yomi+v, "", token_ctx...)
	}
}

// makeKanwa creates a kanwa map from a list of source files and writes it to a destination file.
// It returns the kanwa map and an error if any.
// The source files are expected to have lines in the format "yomi kanji [ctx ...]".
// The yomi is the reading of the kanji character or phrase.
// The kanji is the kanji character or phrase.
// The ctx is a list of contexts in which the kanji character or phrase is used.
func makeKanwaMap(src_list []string) (kanwaMap, error) {
	if err := verifyKanwaMapSourceList(src_list); err != nil {
		return nil, err
	}

	m := make(kanwaMap)
	for _, src := range src_list {
		f, err := os.OpenFile(src, os.O_RDONLY, os.ModePerm)
		if err != nil {
			return nil, err
		}

		defer f.Close()

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		for line := range traverseFile(ctx, f) {
			m.parseLine(line)
		}
	}

	return m, nil
}

// verifyKanwaMapSourceList verifies the source files.
// It returns an error if the source files are invalid.
// The source files are invalid if they are not in the list of resources or if they do not exist.
func verifyKanwaMapSourceList(src_list []string) error {
	if len(src_list) == 0 {
		return nil
	}

	for _, v_list := range kanwaMapResources {
		for _, v := range v_list {
			if v == src_list[0] {
				_, err := os.Stat(src_list[0])
				return err
			}
		}
	}
	if len(src_list) > 1 {
		return verifyKanwaMapSourceList(src_list[1:])
	}

	return fmt.Errorf("invalid source: %v", src_list)
}

package codegen

import "testing"

func Test_makeLookupMap(t *testing.T) {
	for _, tt := range []struct {
		name string
		src  string
		dst  string
	}{
		{"test#1", "data/hepburndict.utf8", "build/hepburndict3.json"},
		{"test#2", "data/kunreidict.utf8", "build/kunreidict3.json"},
		{"test#3", "data/passportdict.utf8", "build/passportdict3.json"},
		{"test#4", "data/hepburnhira.utf8", "build/hepburnhira3.json"},
		{"test#5", "data/kunreihira.utf8", "build/kunreihira3.json"},
		{"test#6", "data/passporthira.utf8", "build/passporthira3.json"},
		{"test#7", "data/halfkana.utf8", "build/halfkana3.json"},
	} {
		t.Run(tt.src, func(t *testing.T) {
			m, err := makeLookupMap(tt.src, tt.dst)
			if err != nil {
				t.Fatal(err)
				return
			}

			t.Log(m)
		})
	}
}

func Test_makeTransTable(t *testing.T) {
	for _, tt := range []struct {
		name string
		src  string
		dst  string
	}{
		{"test#1", "data/itaijidict.utf8", "build/itaijidict4.json"},
	} {
		t.Run(tt.src, func(t *testing.T) {
			m, err := makeTransTable(tt.src, tt.dst)
			if err != nil {
				t.Fatal(err)
				return
			}

			t.Log(m)
		})
	}
}

func Test_makeKanwa(t *testing.T) {
	for _, tt := range []struct {
		name string
		src  []string
		dst  string
	}{
		{"test#1", []string{
			"data/kakasidict.utf8",
			"data/unidict_noun.utf8",
			"data/unidict_adj.utf8",
		}, "build/kanwadict4.json"},
	} {
		t.Run(tt.name, func(t *testing.T) {
			m, err := makeKanwa(tt.src, tt.dst)
			if err != nil {
				t.Fatal(err)
				return
			}

			t.Log(m)
		})
	}
}

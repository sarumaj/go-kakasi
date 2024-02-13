package codegen

import (
	"path/filepath"
	"testing"
)

func Test_makeKanwa(t *testing.T) {
	tmpDir := t.TempDir()

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
			m, err := makeKanwaMap(tt.src)
			if err != nil {
				t.Errorf("makeKanwa() error = %v", err)
				return
			}

			if err := dumpJSON(filepath.Join(tmpDir, tt.dst), m, ""); err != nil {
				t.Errorf("dumpJSON() error = %v", err)
			}
		})
	}
}

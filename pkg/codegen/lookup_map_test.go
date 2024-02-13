package codegen

import (
	"path/filepath"
	"testing"
)

func Test_makeLookupMap(t *testing.T) {
	tmpDir := t.TempDir()

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
			m, err := makeLookupMap(tt.src)
			if err != nil {
				t.Errorf("makeLookupMap() error = %v", err)
				return
			}

			if err := dumpJSON(filepath.Join(tmpDir, tt.dst), m, ""); err != nil {
				t.Errorf("dumpJSON() error = %v", err)
			}
		})
	}
}

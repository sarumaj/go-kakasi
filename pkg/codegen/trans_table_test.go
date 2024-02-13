package codegen

import (
	"path/filepath"
	"testing"
)

func Test_makeTransTable(t *testing.T) {
	tmpDir := t.TempDir()

	for _, tt := range []struct {
		name string
		src  string
		dst  string
	}{
		{"test#1", "data/itaijidict.utf8", "build/itaijidict4.json"},
	} {
		t.Run(tt.src, func(t *testing.T) {
			m, err := makeTransTable(tt.src)
			if err != nil {
				t.Errorf("makeTransTable() error = %v", err)
				return
			}

			if err := dumpJSON(filepath.Join(tmpDir, tt.dst), m, ""); err != nil {
				t.Errorf("dumpJSON() error = %v", err)
			}
		})
	}
}

package codegen

import (
	"path/filepath"
	"testing"
)

func Test_makeTransTable(t *testing.T) {
	tmpDir := t.TempDir()

	for dst, src := range transTableResources {
		t.Run(src, func(t *testing.T) {
			m, err := makeTransTable(src)
			if err != nil {
				t.Errorf("makeTransTable() error = %v", err)
				return
			}

			if err := dumpJSON(filepath.Join(tmpDir, dst), m, ""); err != nil {
				t.Errorf("dumpJSON() error = %v", err)
			}
		})
	}
}

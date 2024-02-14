package codegen

import (
	"path/filepath"
	"testing"
)

func Test_makeKanwa(t *testing.T) {
	tmpDir := t.TempDir()

	for dst, src_list := range kanwaMapResources {
		t.Run(dst, func(t *testing.T) {
			m, err := makeKanwaMap(src_list)
			if err != nil {
				t.Errorf("makeKanwa() error = %v", err)
				return
			}

			if err := dumpJSON(filepath.Join(tmpDir, dst), m, ""); err != nil {
				t.Errorf("dumpJSON() error = %v", err)
			}
		})
	}
}

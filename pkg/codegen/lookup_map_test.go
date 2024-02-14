package codegen

import (
	"path/filepath"
	"testing"
)

func Test_makeLookupMap(t *testing.T) {
	tmpDir := t.TempDir()

	for dst, src := range lookupMapResources {
		t.Run(src, func(t *testing.T) {
			m, err := makeLookupMap(src)
			if err != nil {
				t.Errorf("makeLookupMap() error = %v", err)
				return
			}

			if err := dumpJSON(filepath.Join(tmpDir, dst), m, ""); err != nil {
				t.Errorf("dumpJSON() error = %v", err)
			}
		})
	}
}

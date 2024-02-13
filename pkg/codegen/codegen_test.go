package codegen

import (
	"testing"
)

func TestGenerate(t *testing.T) {
	tmpDir := t.TempDir()

	if err := Generate(tmpDir); err != nil {
		t.Errorf("Generate() error = %v", err)
	}
}

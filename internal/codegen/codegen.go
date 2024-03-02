//go:generate go run ../../generate.go -buildDir=../properties/data

package codegen

import (
	"path/filepath"
)

// Generate is a code generation function that generates lookup maps, translation tables, and kanwa maps.
// The generated files are written to the specified directory.
func Generate(dst, indent string) error {
	for tgt, src := range lookupMapResources {
		m, err := makeLookupMap(src)
		if err != nil {
			return err
		}

		if err := dumpJSON(filepath.Join(dst, tgt), m, indent); err != nil {
			return err
		}
	}

	for tgt, src := range transTableResources {
		m, err := makeTransTable(src)
		if err != nil {
			return err
		}

		if err := dumpJSON(filepath.Join(dst, tgt), m, indent); err != nil {
			return err
		}
	}

	for tgt, src_list := range kanwaMapResources {
		m, err := makeKanwaMap(src_list)
		if err != nil {
			return err
		}

		if err := dumpJSON(filepath.Join(dst, tgt), m, indent); err != nil {
			return err
		}
	}

	return nil
}

package codegen

import (
	"context"
	"fmt"
	"os"
	"strings"
)

// lookupMap is a lookup table.
// It maps a string to a string.
type lookupMap map[string]string

func (m lookupMap) Get(k string) string { return mapGet(m, k) }
func (m lookupMap) Has(k string) bool   { return mapHas(m, k) }
func (m lookupMap) Keys() []string      { return mapKeys(m) }
func (m lookupMap) Set(k, v string)     { m = mapSet(m, k, v) }

// MaxKeyLen returns the length of the longest key in the map.
func (m lookupMap) MaxKeyLen() int {
	l := 0
	for k := range m {
		if len(k) > l {
			l = len(k)
		}
	}

	return l
}

// makeLookupMap creates a lookup table from a source file and writes it to a destination file.
// It returns the lookup table and an error if any.
// The source file is expected to have lines in the format "value key".
func makeLookupMap(src, dst string) (lookupMap, error) {
	f, err := os.OpenFile(src, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	m := make(lookupMap)
	for line := range traverseFile(ctx, f) {
		v, k, _ := strings.Cut(line, " ")
		m.Set(k, v)
	}

	m.Set("_max_key_len_", fmt.Sprintf("%d", m.MaxKeyLen()))

	if err := dumpJSON(dst, m, ""); err != nil {
		return nil, err
	}

	return m, nil
}

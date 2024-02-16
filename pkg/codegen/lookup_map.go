package codegen

import (
	"context"
	"fmt"
	"os"
	"strings"
)

// LookupMapResources is a map of target and source files.
// The target file is the destination file.
var lookupMapResources = map[string]string{
	"halfkana3.json":     "data/halfkana.utf8",
	"hepburndict3.json":  "data/hepburndict.utf8",
	"hepburnhira3.json":  "data/hepburnhira.utf8",
	"kunreidict3.json":   "data/kunreidict.utf8",
	"kunreihira3.json":   "data/kunreihira.utf8",
	"passportdict3.json": "data/passportdict.utf8",
	"passporthira3.json": "data/passporthira.utf8",
}

// LookupMap is a lookup table.
// It maps a string to a string.
type LookupMap map[string]string

func (m LookupMap) Get(k string) string { return mapGet(m, k) }
func (m LookupMap) Has(k string) bool   { return mapHas(m, k) }
func (m LookupMap) Keys() []string      { return mapKeys(m) }
func (m LookupMap) Set(k, v string)     { m = mapSet(m, k, v) }

// MaxKeyLen returns the length of the longest key in the map.
func (m LookupMap) MaxKeyLen() int {
	if m.Has("_max_key_len_") {
		var l int
		if _, err := fmt.Sscanf(m.Get("_max_key_len_"), "%d", &l); err == nil {
			return l
		}
	}

	l := 0
	for k := range m {
		if len([]rune(k)) > l {
			l = len([]rune(k))
		}
	}

	return l
}

// makeLookupMap creates a lookup table from a source file and writes it to a destination file.
// It returns the lookup table and an error if any.
// The source file is expected to have lines in the format "value key".
func makeLookupMap(src string) (LookupMap, error) {
	if err := verifyLookupMapSource(src); err != nil {
		return nil, err
	}

	f, err := os.OpenFile(src, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	m := make(LookupMap)
	for line := range traverseFile(ctx, f) {
		v, k, _ := strings.Cut(line, " ")
		m.Set(k, v)
	}

	m.Set("_max_key_len_", fmt.Sprintf("%d", m.MaxKeyLen()))

	return m, nil
}

// verifyLookupMapSource verifies the source file.
// It returns an error if the source file is invalid.
// The source file is invalid if it is not in the LookupMapResources.
func verifyLookupMapSource(src string) error {
	for _, v := range lookupMapResources {
		if v == src {
			_, err := os.Stat(src)
			return err
		}
	}

	return fmt.Errorf("invalid source: %s", src)
}

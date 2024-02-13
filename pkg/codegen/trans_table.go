package codegen

import (
	"context"
	"fmt"
	"os"
	"strings"
)

// transTableResources is a map of target and source files.
// The target file is the destination file.
var transTableResources = map[string]string{
	"itaijidict4.json": "data/itaijidict.utf8",
}

// transTable is a translation table.
// It maps a rune to a list of strings.
type transTable map[rune]*string

func (m transTable) Has(c rune) bool      { return mapHas(m, c) }
func (m transTable) Get(c rune) string    { return deref[string](mapGet(m, c)) }
func (m transTable) Keys() []rune         { return mapKeys(m) }
func (m transTable) Set(c rune, v string) { m = mapSet(m, c, &v) }

// spoof adds a range of runes to the table.
// If a rune is already in the table, it is skipped.
// If a rune is not in the table, it is added with an empty slice of strings.
func (m transTable) spoof(lo, hi int64) {
	for i := lo; i <= hi; i++ {
		c := rune(i)
		if _, ok := m[c]; ok {
			continue
		}

		m[c] = nil
	}
}

// makeTransTable creates a translation table from a source file and writes it to a destination file.
// It returns the translation table and an error if any.
// The source file is expected to have lines in the format "value key".
func makeTransTable(src string) (transTable, error) {
	if err := verifyTransTableSource(src); err != nil {
		return nil, err
	}

	f, err := os.OpenFile(src, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	m := make(transTable)
	for line := range traverseFile(ctx, f) {
		v, k, ok := strings.Cut(line, " ")
		if !ok {
			return nil, fmt.Errorf("invalid line: %s", line)
		}

		if l := len([]rune(k)); l > 1 || l == 0 {
			return nil, fmt.Errorf("invalid key: %q", k)
		}

		c := []rune(k)[0]
		m.Set(c, v)
	}

	m.spoof(0xFE00, 0xFE02)
	m.spoof(0xE0110, 0xE01EF)

	return m, nil
}

// verifyTransTableSource verifies the source file.
// It returns an error if the source file is invalid.
// The source file is invalid if it is not in the list of resources or if it does not exist.
func verifyTransTableSource(src string) error {
	for _, v := range transTableResources {
		if v == src {
			_, err := os.Stat(src)
			return err
		}
	}

	return fmt.Errorf("invalid source: %s", src)
}

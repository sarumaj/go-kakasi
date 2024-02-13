package codegen

import (
	"context"
	"fmt"
	"os"
	"strings"
)

// transTable is a translation table.
// It maps a rune to a list of strings.
type transTable map[rune][]string

func (m transTable) Add(c rune, v string)   { m = mapAdd(m, c, v) }
func (m transTable) Has(c rune) bool        { return mapHas(m, c) }
func (m transTable) Get(c rune) []string    { return mapGet(m, c) }
func (m transTable) Keys() []rune           { return mapKeys(m) }
func (m transTable) Set(c rune, v []string) { m = mapSet(m, c, v) }

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
func makeTransTable(src, dst string) (transTable, error) {
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
		m.Add(c, v)
	}

	m.spoof(0xFE00, 0xFE02)
	m.spoof(0xE0110, 0xE01EF)

	if err := dumpJSON(dst, m, ""); err != nil {
		return nil, err
	}

	return m, nil
}

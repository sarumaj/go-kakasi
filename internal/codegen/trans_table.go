package codegen

import (
	"context"
	"fmt"
	"os"
	"strings"

	ordered "github.com/wk8/go-ordered-map/v2"
)

// TransTableResources is a map of target and source files.
// The target file is the destination file.
var transTableResources = map[string]string{
	"itaijidict4.json": "data/itaijidict.utf8",
}

// TransTable is a translation table.
// It maps a rune to a list of strings.
type TransTable ordered.OrderedMap[rune, *string]

func (m TransTable) Has(c rune) bool { return mapHas(ordered.OrderedMap[rune, *string](m), c) }

func (m TransTable) Iter() func() (rune, *string, bool) {
	return mapIter(ordered.OrderedMap[rune, *string](m))
}

func (m TransTable) Get(c rune) string {
	return deref[string](mapGet(ordered.OrderedMap[rune, *string](m), c))
}

func (m TransTable) Keys() []rune { return mapKeys(ordered.OrderedMap[rune, *string](m)) }
func (m TransTable) Len() int     { return mapLen(ordered.OrderedMap[rune, *string](m)) }

func (m TransTable) MarshalJSON() ([]byte, error) {
	return (*ordered.OrderedMap[rune, *string])(&m).MarshalJSON()
}

func (m *TransTable) Set(c rune, v string) *TransTable {
	return (*TransTable)(mapSet((*ordered.OrderedMap[rune, *string])(m), c, &v))
}

// spoof adds a range of runes to the table.
// If a rune is already in the table, it is skipped.
// If a rune is not in the table, it is added with an empty slice of strings.
func (m *TransTable) spoof(lo, hi int64) {
	for i := lo; i <= hi; i++ {
		c := rune(i)
		if m.Has(c) {
			continue
		}

		o := (*ordered.OrderedMap[rune, *string])(m)
		o.Set(c, nil)
		m = (*TransTable)(o)
	}
}

func (m *TransTable) UnmarshalJSON(data []byte) error {
	return (*ordered.OrderedMap[rune, *string])(m).UnmarshalJSON(data)
}

// makeTransTable creates a translation table from a source file and writes it to a destination file.
// It returns the translation table and an error if any.
// The source file is expected to have lines in the format "value key".
func makeTransTable(src string) (*TransTable, error) {
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

	m := (*TransTable)(ordered.New[rune, *string]())
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

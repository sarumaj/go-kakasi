package kanji

import (
	"strings"
	"sync"

	"github.com/sarumaj/go-kakasi/internal/codegen"
	"github.com/sarumaj/go-kakasi/internal/properties"
)

// Itaiji is a type that represents a map of Itaiji characters.
// It is used to convert Itaiji characters to their original form.
type Itaiji struct {
	sync.Mutex
	table *codegen.TransTable
}

// Convert converts Itaiji characters to their original form.
func (t *Itaiji) Convert(s string) string {
	t.Lock()
	defer t.Unlock()

	var replacements []string
	iterator := t.table.Iter()
	for k, v, ok := iterator(); ok; k, v, ok = iterator() {
		if v == nil {
			replacements = append(replacements, string(k), "")
			continue
		}

		replacements = append(replacements, string(k), *v)
	}

	return strings.NewReplacer(replacements...).Replace(s)
}

// HasKey returns true if the given key exists in the Itaiji map.
func (t *Itaiji) HasKey(key rune) bool {
	t.Lock()
	defer t.Unlock()

	return t.table.Has(key)
}

// NewItaiji returns a new Itaiji instance.
func NewItaiji() (*Itaiji, error) {
	t, err := properties.Configurations.JisyoItaiji()
	if err != nil {
		return nil, err
	}

	return &Itaiji{table: t}, nil
}

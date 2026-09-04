package kanji

import (
	"strings"
	"sync"

	"github.com/sarumaj/go-kakasi/internal/properties"
)

// Itaiji is a type that represents a map of Itaiji characters.
// It is used to convert Itaiji characters to their original form.
//
// The lookup table is static, so the replacer and the key set are built once
// at construction time. The instance is immutable afterwards and therefore
// safe for concurrent use without further synchronization.
type Itaiji struct {
	replacer *strings.Replacer
	keys     map[rune]struct{}
}

// Convert converts Itaiji characters to their original form.
func (t *Itaiji) Convert(s string) string { return t.replacer.Replace(s) }

// HasKey returns true if the given key exists in the Itaiji map.
func (t *Itaiji) HasKey(key rune) bool {
	_, ok := t.keys[key]
	return ok
}

// NewItaiji returns the shared Itaiji instance, building it on first use.
// The instance is read-only, so all callers can share it.
var NewItaiji = sync.OnceValues(newItaiji)

func newItaiji() (*Itaiji, error) {
	table, err := properties.Configurations.JisyoItaiji()
	if err != nil {
		return nil, err
	}

	keys := make(map[rune]struct{}, table.Len())
	replacements := make([]string, 0, 2*table.Len())

	iterator := table.Iter()
	for k, v, ok := iterator(); ok; k, v, ok = iterator() {
		keys[k] = struct{}{}
		// A nil value marks a rune that is dropped rather than replaced.
		replacements = append(replacements, string(k), deref(v))
	}

	return &Itaiji{replacer: strings.NewReplacer(replacements...), keys: keys}, nil
}

// deref returns the value a pointer points to, or the zero value if it is nil.
func deref[T any](v *T) T {
	if v == nil {
		var zero T
		return zero
	}

	return *v
}

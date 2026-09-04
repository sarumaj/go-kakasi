package kanji

import (
	"sync"

	"github.com/sarumaj/go-kakasi/internal/codegen"
	"github.com/sarumaj/go-kakasi/internal/properties"
)

// Kanwa is a type that represents a map of Kanwa characters.
//
// The dictionary is read-only after construction, so the instance is safe for
// concurrent use without further synchronization.
type Kanwa struct {
	kanwa *codegen.KanwaMap
	// maxKeyLen caches, per leading character, the rune length of the longest
	// key in that character's table. It bounds the prefix probing in Lookup.
	maxKeyLen map[rune]int
}

// Lookup returns the yomi of the longest entry in the Kanwa dictionary that is
// a prefix of text and that is valid in the context ctx, together with the
// length of the matched prefix in runes. It returns ("", 0) if nothing matches.
func (k *Kanwa) Lookup(text []rune, ctx string) (string, int) {
	if len(text) == 0 {
		return "", 0
	}

	if !k.kanwa.Has(text[0]) {
		return "", 0
	}

	table := k.kanwa.Get(text[0])

	// Probe from the longest possible key downwards: at most one key can equal
	// a given prefix, so the first hit is the longest match.
	length := min(k.maxKeyLen[text[0]], len(text))
	for ; length > 0; length-- {
		pairs, ok := table.Lookup(string(text[:length]))
		if !ok {
			continue
		}

		for _, pair := range pairs {
			if len(pair.Ctx) == 0 || pair.Ctx.Contains(ctx) {
				return pair.Yomi, length
			}
		}
	}

	return "", 0
}

// HasEntry reports whether the dictionary has a table for the leading
// character c.
func (k *Kanwa) HasEntry(c rune) bool { return k.kanwa.Has(c) }

// NewKanwa returns the shared Kanwa instance, building it on first use.
// The instance is read-only, so all callers can share it.
var NewKanwa = sync.OnceValues(newKanwa)

func newKanwa() (*Kanwa, error) {
	m, err := properties.Configurations.JisyoKanwa()
	if err != nil {
		return nil, err
	}

	maxKeyLen := make(map[rune]int, m.Len())
	outer := m.Iter()
	for c, table, ok := outer(); ok; c, table, ok = outer() {
		inner := table.Iter()
		for k, _, ok := inner(); ok; k, _, ok = inner() {
			if l := len([]rune(k)); l > maxKeyLen[c] {
				maxKeyLen[c] = l
			}
		}
	}

	return &Kanwa{kanwa: m, maxKeyLen: maxKeyLen}, nil
}

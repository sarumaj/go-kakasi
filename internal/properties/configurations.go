package properties

import (
	"sync"

	"github.com/goccy/go-json"

	"github.com/sarumaj/go-kakasi/internal/codegen"
)

// Configurations is a set of configuration values.
// It is used to get the path to the dictionary files.
// It is also used to read files from the file system.
var Configurations = configurations{}

// The dictionaries are embedded, immutable and expensive to decode (the kanwa
// dictionary alone is ~10 MiB of JSON), so each one is decoded at most once and
// shared by every caller. Callers must treat the returned maps as read-only.
var (
	jisyoHalfkana     = loadOnce[codegen.LookupMap]("data/halfkana3.json")
	jisyoHepburn      = loadOnce[codegen.LookupMap]("data/hepburndict3.json")
	jisyoHepburnHira  = loadOnce[codegen.LookupMap]("data/hepburnhira3.json")
	jisyoItaiji       = loadOnce[codegen.TransTable]("data/itaijidict4.json")
	jisyoKanwa        = loadOnce[codegen.KanwaMap]("data/kanwadict4.json")
	jisyoKunrei       = loadOnce[codegen.LookupMap]("data/kunreidict3.json")
	jisyoKunreiHira   = loadOnce[codegen.LookupMap]("data/kunreihira3.json")
	jisyoPassport     = loadOnce[codegen.LookupMap]("data/passportdict3.json")
	jisyoPassportHira = loadOnce[codegen.LookupMap]("data/passporthira3.json")
)

type configurations struct{}

func (configurations) JisyoHalfkana() (*codegen.LookupMap, error)     { return jisyoHalfkana() }
func (configurations) JisyoHepburn() (*codegen.LookupMap, error)      { return jisyoHepburn() }
func (configurations) JisyoHepburnHira() (*codegen.LookupMap, error)  { return jisyoHepburnHira() }
func (configurations) JisyoItaiji() (*codegen.TransTable, error)      { return jisyoItaiji() }
func (configurations) JisyoKanwa() (*codegen.KanwaMap, error)         { return jisyoKanwa() }
func (configurations) JisyoKunrei() (*codegen.LookupMap, error)       { return jisyoKunrei() }
func (configurations) JisyoKunreiHira() (*codegen.LookupMap, error)   { return jisyoKunreiHira() }
func (configurations) JisyoPassport() (*codegen.LookupMap, error)     { return jisyoPassport() }
func (configurations) JisyoPassportHira() (*codegen.LookupMap, error) { return jisyoPassportHira() }

// loadOnce returns an accessor that decodes the embedded dictionary at path on
// its first call and returns the same value on every subsequent call.
func loadOnce[T any](path string) func() (*T, error) {
	return sync.OnceValues(func() (*T, error) {
		f, err := dataFS.Open(path)
		if err != nil {
			return nil, err
		}

		defer f.Close()

		var v T
		if err := json.NewDecoder(f).Decode(&v); err != nil {
			return nil, err
		}

		return &v, nil
	})
}

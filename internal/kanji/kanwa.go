package kanji

import (
	"sync"

	"github.com/sarumaj/go-kakasi/internal/codegen"
	"github.com/sarumaj/go-kakasi/internal/properties"
)

// Kanwa is a type that represents a map of Kanwa characters.
type Kanwa struct {
	sync.Mutex
	kanwa *codegen.KanwaMap
}

// Load returns the KanjiCtxMap for the given key.
// The key is the first character of the kanji character or phrase.
// The KanjiCtxMap contains the reading and the contexts in which the kanji character or phrase is used.
func (k *Kanwa) Load(key rune) *codegen.KanjiCtxMap {
	k.Lock()
	defer k.Unlock()

	if k.kanwa.Has(key) {
		v := k.kanwa.Get(key)
		return &v
	}

	return nil
}

// NewKanwa returns a new Kanwa instance.
func NewKanwa() (*Kanwa, error) {
	k, err := properties.Configurations.JisyoKanwa()
	if err != nil {
		return nil, err
	}

	return &Kanwa{kanwa: k}, nil
}

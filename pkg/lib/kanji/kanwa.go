package kanji

import (
	"github/sarumaj/go-kakasi/pkg/codegen"
	"github/sarumaj/go-kakasi/pkg/lib/properties"
	"sync"
)

// Kanwa is a type that represents a map of Kanwa characters.
type Kanwa struct {
	sync.Mutex
	kanwa codegen.KanwaMap
}

// Load returns the KanjiCtxMap for the given key.
// The key is the first character of the kanji character or phrase.
// The KanjiCtxMap contains the reading and the contexts in which the kanji character or phrase is used.
func (k *Kanwa) Load(key rune) codegen.KanjiCtxMap {
	k.Lock()
	defer k.Unlock()

	return k.kanwa.Get(key)
}

// NewKanwa returns a new Kanwa instance.
func NewKanwa() (*Kanwa, error) {
	k, err := properties.Configurations.JisyoKanwa()
	if err != nil {
		return nil, err
	}

	return &Kanwa{kanwa: k}, nil
}

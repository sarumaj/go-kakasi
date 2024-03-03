package properties

import (
	"github.com/goccy/go-json"

	"github.com/sarumaj/go-kakasi/internal/codegen"
)

// Configurations is a set of configuration values.
// It is used to get the path to the dictionary files.
// It is also used to read files from the file system.
var Configurations = configurations{}

type configurations struct{}

func (configurations) jisyoHalfkana() string        { return "data/halfkana3.json" }
func (configurations) jisyoHepburn() string         { return "data/hepburndict3.json" }
func (configurations) jisyoHepburnHira() (v string) { return "data/hepburnhira3.json" }
func (configurations) jisyoItaiji() string          { return "data/itaijidict4.json" }
func (configurations) jisyoKanwa() string           { return "data/kanwadict4.json" }
func (configurations) jisyoKunrei() string          { return "data/kunreidict3.json" }
func (configurations) jisyoKunreiHira() string      { return "data/kunreihira3.json" }
func (configurations) jisyoPassport() string        { return "data/passportdict3.json" }
func (configurations) jisyoPassportHira() string    { return "data/passporthira3.json" }

func (configurations) decode(path string, v any) error {
	f, err := dataFS.Open(path)
	if err != nil {
		return err
	}

	defer f.Close()

	return json.NewDecoder(f).Decode(v)
}

func (c configurations) JisyoHalfkana() (*codegen.LookupMap, error) {
	var v codegen.LookupMap
	if err := c.decode(c.jisyoHalfkana(), &v); err != nil {
		return nil, err
	}

	return &v, nil
}

func (c configurations) JisyoHepburn() (*codegen.LookupMap, error) {
	var v codegen.LookupMap
	if err := c.decode(c.jisyoHepburn(), &v); err != nil {
		return nil, err
	}

	return &v, nil
}

func (c configurations) JisyoHepburnHira() (*codegen.LookupMap, error) {
	var v codegen.LookupMap
	if err := c.decode(c.jisyoHepburnHira(), &v); err != nil {
		return nil, err
	}

	return &v, nil
}

func (c configurations) JisyoItaiji() (*codegen.TransTable, error) {
	var v codegen.TransTable
	if err := c.decode(c.jisyoItaiji(), &v); err != nil {
		return nil, err
	}

	return &v, nil
}

func (c configurations) JisyoKanwa() (*codegen.KanwaMap, error) {
	var v codegen.KanwaMap
	if err := c.decode(c.jisyoKanwa(), &v); err != nil {
		return nil, err
	}

	return &v, nil
}

func (c configurations) JisyoKunrei() (*codegen.LookupMap, error) {
	var v codegen.LookupMap
	if err := c.decode(c.jisyoKunrei(), &v); err != nil {
		return nil, err
	}

	return &v, nil
}

func (c configurations) JisyoKunreiHira() (*codegen.LookupMap, error) {
	var v codegen.LookupMap
	if err := c.decode(c.jisyoKunreiHira(), &v); err != nil {
		return nil, err
	}

	return &v, nil
}

func (c configurations) JisyoPassport() (*codegen.LookupMap, error) {
	var v codegen.LookupMap
	if err := c.decode(c.jisyoPassport(), &v); err != nil {
		return nil, err
	}

	return &v, nil
}

func (c configurations) JisyoPassportHira() (*codegen.LookupMap, error) {
	var v codegen.LookupMap
	if err := c.decode(c.jisyoPassportHira(), &v); err != nil {
		return nil, err
	}

	return &v, nil
}

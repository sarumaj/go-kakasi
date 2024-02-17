package script

import "github/sarumaj/go-kakasi/pkg/lib/properties"

type Conv struct {
	hahConv *Hira
	hakConv *Hira
	hapConv *Hira
	hkConv  *Hira
	khConv  *Kata
	saConv  *Symbol
}

func (Conv) maxLen() int { return 32 }

func (c Conv) s2a(text string) (string, error) {
	var converted string

	for i := 0; i < len([]rune(text)); {
		width := len([]rune(text))
		if width > c.maxLen()+i {
			width = c.maxLen() + i
		}

		result, length, err := c.saConv.Convert(text[i:width])
		if err != nil {
			return "", err
		}

		switch {
		case length > 0:
			converted += result
			i += length

		case properties.Ch.IsLongSymbol([]rune(text)[i]):
			if len([]rune(result)) > 0 {
				converted += string([]rune(result)[len([]rune(result))-1])
			} else {
				converted += "-"
			}
			i++

		default:
			converted += string([]rune(text)[i : i+1])
			i++

		}
	}

	return converted, nil
}

func NewConv() (*Conv, error) {
	c := Conv{}

	var err error
	c.hahConv, err = NewHira(Conf{Method: MethodHepburn, Mode: Mode_a})
	if err != nil {
		return nil, err
	}

	c.hakConv, err = NewHira(Conf{Method: MethodKunrei, Mode: Mode_a})
	if err != nil {
		return nil, err
	}

	c.hapConv, err = NewHira(Conf{Method: MethodPassport, Mode: Mode_a})
	if err != nil {
		return nil, err
	}

	c.hkConv, err = NewHira(Conf{Mode: ModeK})
	if err != nil {
		return nil, err
	}

	c.khConv, err = NewKata(Conf{Mode: ModeH})
	if err != nil {
		return nil, err
	}

	c.saConv = NewSymbol(Mode_a)

	return &c, nil
}

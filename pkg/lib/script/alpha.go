package script

type Alpha struct {
	kana
	mode mode
}

func NewAlpha(mode mode) *Alpha {
	return &Alpha{
		kana: kana{},
		mode: mode,
	}
}

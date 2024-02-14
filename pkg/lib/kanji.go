package lib

import (
	"github/sarumaj/go-kakasi/pkg/codegen"
	"github/sarumaj/go-kakasi/pkg/lib/properties"
)

type Itaiji struct {
	codegen.TransTable
}

func NewItaiji() (*Itaiji, error) {
	t, err := properties.Configurations.JisyoItaiji()
	if err != nil {
		return nil, err
	}

	return &Itaiji{TransTable: t}, nil
}

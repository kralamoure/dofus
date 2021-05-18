package dofussvc

import (
	"errors"

	"github.com/kralamoure/dofus"
)

type Service struct {
	storer dofus.Storer
}

func NewService(storer dofus.Storer) (*Service, error) {
	if storer == nil {
		return nil, errors.New("nil storer")
	}

	svc := &Service{
		storer: storer,
	}

	return svc, nil
}

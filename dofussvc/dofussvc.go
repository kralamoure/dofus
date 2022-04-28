package dofussvc

import (
	"errors"

	"github.com/kralamoure/dofus"
)

type Service struct {
	store dofus.Store
}

func NewService(store dofus.Store) (*Service, error) {
	if store == nil {
		return nil, errors.New("nil store")
	}

	svc := &Service{
		store: store,
	}

	return svc, nil
}

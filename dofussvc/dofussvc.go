package dofussvc

import (
	"errors"

	"github.com/happybydefault/logger"

	"github.com/kralamoure/dofus"
)

type Config struct {
	Repo   dofus.Repo
	Logger logger.Logger
}

type Service struct {
	repo   dofus.Repo
	logger logger.Logger
}

func NewService(c Config) (*Service, error) {
	if c.Repo == nil {
		return nil, errors.New("nil repository")
	}
	if c.Logger == nil {
		c.Logger = logger.Noop{}
	}
	svc := &Service{
		repo:   c.Repo,
		logger: c.Logger,
	}
	return svc, nil
}

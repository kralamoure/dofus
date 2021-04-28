package dofussvc

import (
	"context"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"

	"github.com/kralamoure/dofus"
)

func (svc Service) CreateAccount(ctx context.Context, account dofus.Account) (id string, err error) {
	err = validation.Validate(account)
	if err != nil {
		return
	}

	return svc.repo.CreateAccount(ctx, account)
}

func (svc Service) Accounts(ctx context.Context) (accounts map[string]dofus.Account, err error) {
	return svc.repo.Accounts(ctx)
}

func (svc Service) AccountsByUserId(ctx context.Context, userId string) (accounts map[string]dofus.Account, err error) {
	return svc.repo.AccountsByUserId(ctx, userId)
}

func (svc Service) Account(ctx context.Context, id string) (account dofus.Account, err error) {
	return svc.repo.Account(ctx, id)
}

func (svc Service) AccountByName(ctx context.Context, name string) (account dofus.Account, err error) {
	return svc.repo.AccountByName(ctx, name)
}

func (svc Service) SetAccountLastAccessAndLastIP(ctx context.Context, id string, lastAccess time.Time, lastIP string) error {
	return svc.repo.SetAccountLastAccessAndLastIP(ctx, id, lastAccess, lastIP)
}

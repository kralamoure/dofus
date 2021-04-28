package dofussvc

import (
	"context"

	validation "github.com/go-ozzo/ozzo-validation/v4"

	"github.com/kralamoure/dofus"
	"github.com/kralamoure/dofus/dofustyp"
)

func (svc Service) CreateUser(ctx context.Context, user dofus.User) (id string, err error) {
	err = validation.Validate(user)
	if err != nil {
		return
	}

	return svc.repo.CreateUser(ctx, user)
}

func (svc Service) Users(ctx context.Context) (users map[string]dofus.User, err error) {
	return svc.repo.Users(ctx)
}

func (svc Service) User(ctx context.Context, id string) (user dofus.User, err error) {
	return svc.repo.User(ctx, id)
}

func (svc Service) UserByNickname(ctx context.Context, nickname string) (user dofus.User, err error) {
	return svc.repo.UserByNickname(ctx, nickname)
}

func (svc Service) UserAddChatChannel(ctx context.Context, id string, chatChannels ...dofustyp.ChatChannel) error {
	return svc.repo.UserAddChatChannels(ctx, id, chatChannels...)
}

func (svc Service) UserRemoveChatChannel(ctx context.Context, id string, chatChannels ...dofustyp.ChatChannel) error {
	return svc.repo.UserRemoveChatChannels(ctx, id, chatChannels...)
}

package dofus

import (
	"context"
	"errors"
	"time"

	"github.com/kralamoure/dofus/dofustyp"
)

var (
	ErrNotFound      = errors.New("not found")
	ErrAlreadyExists = errors.New("already exists")

	ErrUserEmailAlreadyExists    = errors.New("user email already exists")
	ErrUserNicknameAlreadyExists = errors.New("user nickname already exists")
	ErrAccountNameAlreadyExists  = errors.New("account name already exists")
)

type Repo interface {
	CreateUser(ctx context.Context, user User) (id string, err error)
	Users(ctx context.Context) (users map[string]User, err error)
	User(ctx context.Context, id string) (user User, err error)
	UserByNickname(ctx context.Context, nickname string) (user User, err error)
	UserAddChatChannels(ctx context.Context, id string, chatChannels ...dofustyp.ChatChannel) error
	UserRemoveChatChannels(ctx context.Context, id string, chatChannels ...dofustyp.ChatChannel) error

	CreateAccount(ctx context.Context, account Account) (id string, err error)
	Accounts(ctx context.Context) (accounts map[string]Account, err error)
	AccountsByUserId(ctx context.Context, userId string) (accounts map[string]Account, err error)
	Account(ctx context.Context, id string) (account Account, err error)
	AccountByName(ctx context.Context, name string) (account Account, err error)
	SetAccountLastAccessAndLastIP(ctx context.Context, id string, lastAccess time.Time, lastIP string) error
}

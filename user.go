package dofus

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"

	"github.com/kralamoure/dofus/dofustyp"
)

type User struct {
	Id             string
	Email          dofustyp.Email
	Nickname       dofustyp.Nickname
	Hash           dofustyp.Hash
	SecretQuestion string
	SecretAnswer   string
	Gender         dofustyp.Gender
	Community      dofustyp.Community
	ChatChannels   UserChatChannels
}

type UserChatChannels struct {
	Admin       bool
	Info        bool
	Public      bool
	Private     bool
	Group       bool
	Team        bool
	Guild       bool
	Alignment   bool
	Recruitment bool
	Trading     bool
	Newbies     bool
}

func (u User) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Email,
			validation.Required,
		),
		validation.Field(&u.Nickname,
			validation.Required,
		),
		validation.Field(&u.Hash,
			validation.Required,
		),
		validation.Field(&u.SecretQuestion,
			validation.Required,
		),
		validation.Field(&u.SecretAnswer,
			validation.Required,
		),
		validation.Field(&u.Gender),
		validation.Field(&u.Community),
		validation.Field(&u.ChatChannels),
	)
}

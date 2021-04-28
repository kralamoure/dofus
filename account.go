package dofus

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"

	"github.com/kralamoure/dofus/dofustyp"
)

type Account struct {
	Id           string
	UserId       string
	Name         dofustyp.AccountName
	Subscription time.Time
	Admin        bool
	LastAccess   time.Time
	LastIP       string
}

func (a Account) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Name,
			validation.Required,
		),
		validation.Field(&a.LastIP, is.IP),
	)
}

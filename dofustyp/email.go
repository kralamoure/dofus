package dofustyp

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type Email string

func (e Email) Validate() error {
	v := string(e)

	return validation.Validate(v,
		is.Email,
	)
}

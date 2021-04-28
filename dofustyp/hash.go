package dofustyp

import (
	"github.com/alexedwards/argon2id"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Hash string

func (h Hash) Validate() error {
	v := string(h)

	return validation.Validate(v,
		validation.By(func(value interface{}) error {
			_, err := argon2id.ComparePasswordAndHash("", v)
			return err
		}),
	)
}

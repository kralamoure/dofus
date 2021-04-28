package dofustyp

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type AccountName string

func (n AccountName) Validate() error {
	v := string(n)

	return validation.Validate(v,
		validation.Length(6, 19),
		validation.Match(regexp.MustCompile(`^[0-9a-zA-Z-]*$`)),
	)
}

package dofustyp

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Nickname string

func (n Nickname) Validate() error {
	v := string(n)

	return validation.Validate(v,
		validation.Length(6, 19),
		validation.Match(regexp.MustCompile(`^[a-zA-Z-]*$`)),
	)
}

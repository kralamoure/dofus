package dofustyp

import (
	"errors"
	"strings"
	"unicode"
)

type Password string

func (p Password) Validate() error {
	v := string(p)

	if len(v) < 8 {
		return errors.New("not enough characters")
	} else if len(v) > 49 {
		return errors.New("too many characters")
	}

	if strings.ContainsAny(v, " <>") {
		return errors.New("contains invalid characters")
	}

	var containsLetters bool
	var containsDigits bool
	for _, r := range v {
		if unicode.IsLetter(r) {
			containsLetters = true
		} else if unicode.IsDigit(r) {
			containsDigits = true
		}
		if containsLetters && containsDigits {
			break
		}
	}

	if !containsLetters {
		return errors.New("does not contain letters")
	}

	if !containsDigits {
		return errors.New("does not contain digits")
	}

	return nil
}

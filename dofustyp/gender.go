package dofustyp

const (
	GenderMale Gender = iota
	GenderFemale
)

type Gender int

var Genders = map[Gender]string{
	GenderMale:   "Male",
	GenderFemale: "Female",
}

func (g Gender) Validate() error {
	_, ok := Genders[g]
	if !ok {
		return ErrInvalidValue
	}

	return nil
}

func (g Gender) String() string {
	name, ok := Genders[g]
	if ok {
		return name
	}

	return nameUnknown
}

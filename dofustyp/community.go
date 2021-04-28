package dofustyp

const (
	CommunityFrenchSpeaking Community = iota
	CommunityUnitedKingdomAndIreland
	CommunityInternational
	CommunityGerman
	CommunitySpanishSpeaking
	CommunityRussian
	CommunityBrazilian
	CommunityDutch
	CommunityItalian Community = iota + 1
	CommunityJapanese
	CommunityDebug Community = 99
)

type Community int

var Communities = map[Community]string{
	CommunityFrenchSpeaking:          "French-speaking",
	CommunityUnitedKingdomAndIreland: "United Kingdom & Ireland",
	CommunityInternational:           "International",
	CommunityGerman:                  "German",
	CommunitySpanishSpeaking:         "Spanish-speaking",
	CommunityRussian:                 "Russian",
	CommunityBrazilian:               "Brazilian",
	CommunityDutch:                   "Dutch",
	CommunityItalian:                 "Italian",
	CommunityJapanese:                "Japanese",
	CommunityDebug:                   "Debug",
}

func (c Community) Validate() error {
	_, ok := Communities[c]
	if !ok {
		return ErrInvalidValue
	}

	return nil
}

func (c Community) String() string {
	name, ok := Communities[c]
	if ok {
		return name
	}

	return nameUnknown
}

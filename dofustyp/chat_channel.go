package dofustyp

const (
	ChatChannelInfo        ChatChannel = 'i'
	ChatChannelPublic      ChatChannel = '*'
	ChatChannelPrivate     ChatChannel = 'p'
	ChatChannelGroup       ChatChannel = '$'
	ChatChannelTeam        ChatChannel = '#'
	ChatChannelGuild       ChatChannel = '%'
	ChatChannelAlignment   ChatChannel = '!'
	ChatChannelRecruitment ChatChannel = '?'
	ChatChannelTrading     ChatChannel = ':'
	ChatChannelNewbies     ChatChannel = '^'
	ChatChannelAdmin       ChatChannel = '@'
)

type ChatChannel rune

var ChatChannels = map[ChatChannel]string{
	ChatChannelInfo:        "Info",
	ChatChannelPublic:      "Public",
	ChatChannelPrivate:     "Private",
	ChatChannelGroup:       "Group",
	ChatChannelTeam:        "Team",
	ChatChannelGuild:       "Guild",
	ChatChannelAlignment:   "Alignment",
	ChatChannelRecruitment: "Recruitment",
	ChatChannelTrading:     "Trading",
	ChatChannelNewbies:     "Newbies",
	ChatChannelAdmin:       "Admin",
}

func (a ChatChannel) Validate() error {
	_, ok := ChatChannels[a]
	if !ok {
		return ErrInvalidValue
	}

	return nil
}

func (a ChatChannel) String() string {
	name, ok := ChatChannels[a]
	if ok {
		return name
	}

	return nameUnknown
}

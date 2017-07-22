package user

import (
	"github.com/satori/go.uuid"
	"github.com/bwmarrin/discordgo"
)

type User struct{
	ID uuid.UUID  // v4
	JBID string
	DiscordID string
	IrcID string
	TwitchID string
	TelegramID string

	// Temp until the relationships are more clear
	DiscordUser *discordgo.User
}

func New(jbID string, discordID string, ircID string,
		 twitchID string, telegramID string) *User {
		 	id := uuid.NewV4()
		 	return &User{
		 		ID: id,
		 		JBID: jbID,
		 		DiscordID: discordID,
		 		IrcID: ircID,
		 		TwitchID: twitchID,
		 		TelegramID: telegramID,
			}
}
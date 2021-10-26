package good

import "github.com/bwmarrin/discordgo"

type DiscordSource interface {
	ChannelMessageSend(string, string) (*discordgo.Message, error)
}

type DiscordSession struct {
	*discordgo.Session
}

func NewDiscord(authToken string) (DiscordSession, error) {
	ds, err := discordgo.New("Bot " + authToken)
	return DiscordSession{ds}, err
}

func SendMsg(ds DiscordSource, channelID string, content string) error {
	_, err := ds.ChannelMessageSend(channelID, content)
	return err
}

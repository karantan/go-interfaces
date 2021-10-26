package good

import (
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockedDiscordSource struct {
	mock.Mock
}

func (m *mockedDiscordSource) ChannelMessageSend(channelID, content string) (*discordgo.Message, error) {
	args := m.Called(channelID, content)
	msg := &discordgo.Message{}
	return msg, args.Error(0)
}

func TestNewDiscord(t *testing.T) {
	assert := assert.New(t)
	got, err := NewDiscord("secret")
	assert.NoError(err)
	assert.IsType(got.Session, &discordgo.Session{})
}

func TestSendMsg(t *testing.T) {
	assert := assert.New(t)
	ds := new(mockedDiscordSource)
	defer ds.AssertExpectations(t)
	ds.On("ChannelMessageSend", "123", "foo").Return(nil)

	err := SendMsg(ds, "123", "foo")
	assert.NoError(err)
}

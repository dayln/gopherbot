package main

import (
	"github.com/bwmarrin/discordgo"
	"github.com/gopherbot/types"
)

//Bot type that handles all the discodgo methods
type Bot struct {
	uname         string
	password      string
	dg            *discordgo.Session
	users         []types.User
	msgResPulgins map[string]types.Plugin
}

//New creates a new instance of a bot
func (b *Bot) New(username, password string) {
	b.uname = username
	b.password = password
}

//Login and store the session token. Returns error if failed to authenticate
func (b *Bot) Login() error {
	var err error
	b.dg, err = discordgo.New(b.uname, b.password)
	return err
}

//SendMsg Sends a message on a specified channel.
func (b *Bot) SendMsg(msg types.Response) (err error) {
	rt, err := b.dg.ChannelMessageSend(msg.OriginalPost.ChannelID, msg.Content)

	return

}

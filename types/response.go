package types

import "github.com/bwmarrin/discordgo"

//Response is a bot response to a message
type Response struct {
	Content      string
	OriginalPost discordgo.Message
}

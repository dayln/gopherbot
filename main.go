package main

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
)

func main() {
	dg, err := discordgo.New(os.Getenv("DIS_TOKEN"))
	if err != nil {
		fmt.Println(err)
		return
	}
	//set a call back func OnMessageCreate
	dg.AddHandler(messageCreate)
	err = dg.Open()
	if err != nil {
		fmt.Println(err)
		return
	}
	st, err := dg.UserGuilds()
	fmt.Printf("%s", st[0].Name)
	// Simple way to keep program running until any key press.
	var input string
	fmt.Scanln(&input)
	return
}

// This function will be called (due to above assignment) every time a new
// message is created on any channel that the autenticated user has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if !m.Author.Bot {
		s.ChannelMessageSend(m.ChannelID, m.Content)
	}
}

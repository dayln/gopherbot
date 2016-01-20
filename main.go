package main

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

func main() {
	dg, err := discordgo.New("uname", "password")
	if err != nil {
		fmt.Println(err)
		return
	}
	//set a call back func OnMessageCreate
	dg.OnMessageCreate = messageCreate
	// Simple way to keep program running until any key press.
	var input string
	fmt.Scanln(&input)
	return
}

// This function will be called (due to above assignment) every time a new
// message is created on any channel that the autenticated user has access to.
func messageCreate(s *discordgo.Session, m *discordgo.Message) {

	// Print message to stdout.
	fmt.Printf("%20s %20s %20s > %s\n", m.ChannelID, time.Now().Format(time.Stamp), m.Author.Username, m.Content)
}

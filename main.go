package main

import (
	"flag"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var (
	Email    string
	Password string
	Token    string
	BotID    string
)

func init() {

	flag.StringVar(&Email, "e", "", "Account Email")
	flag.StringVar(&Password, "p", "", "Account Password")
	flag.StringVar(&Token, "t", "", "Account Token")
	flag.Parse()
}

func main() {
	dg, err := discordgo.New("Insert token here")
	if err != nil {
		fmt.Println(err)
		return
	}
	//set a call back func OnMessageCreate
	dg.AddHandler(messageCreate)
	err = dg.Open()
	_, err = dg.InviteAccept("0mBNH4vETz8nN1Pv")
	if err != nil {
		fmt.Println(err)
		return
	}
	st, err := dg.Guilds()
	fmt.Printf("%d", len(st))
	// Simple way to keep program running until any key press.
	var input string
	fmt.Scanln(&input)
	return
}

// This function will be called (due to above assignment) every time a new
// message is created on any channel that the autenticated user has access to.
func messageCreate(s *discordgo.Session, m *discordgo.Message) {
	fmt.Printf("s", m.ChannelID)
	if m.Author.Email != "dayln.limesand@gmail.com" {
		s.ChannelMessageSend(m.ChannelID, m.Content)
	}

}

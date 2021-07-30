package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var token string     //token used for logging in as client
var botPrefix string //prefix which will be used for users to call commands

func init() { //go run . -t [token] -p [PREFIX]
	flag.StringVar(&token, "t", "", "Bot token")
	flag.StringVar(&botPrefix, "p", "", "Bot prefix")

	flag.Parse()
}

func main() {
	d, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println(err)
	}

	d.AddHandler(onReady)
	d.AddHandler(onMessage)

	d.Open()

	fmt.Scanln() //Many advise against using fmt.Scanln() to keep the client open, but I do not know why so I will use it.

	d.Close()
}

func onReady(s *discordgo.Session, r *discordgo.Ready) {
	s.UpdateGameStatus(0, botPrefix+"help")
	fmt.Println("Bot is now running")
	fmt.Println("Logged in as: " + r.User.String())
	fmt.Println("Session ID: " + r.SessionID)

}

func onMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	//simple
	if strings.HasPrefix(m.Content, botPrefix+"ping") || strings.HasPrefix(m.Content, botPrefix+"pong") {
		ping(s, m)
	}
	if strings.HasPrefix(m.Content, botPrefix+"help") {
		help(s, m)
	}

	//moderation
	if strings.HasPrefix(m.Content, botPrefix+"ban") {
		ban(s, m)
	}
	if strings.HasPrefix(m.Content, botPrefix+"unban") {
		unban(s, m)
	}
	if strings.HasPrefix(m.Content, botPrefix+"kick") {
		kick(s, m)
	}
	if strings.HasPrefix(m.Content, botPrefix+"setup") {
		setup(s, m)
	}
}

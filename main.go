//This files contains the mandatory functions and handler functions.
//The handler functions call more functions in botFunctions.go for the sake of making the main file look nicer.
//The functions in botFunctions.go sometimes call functions in utils.go which exist to be more aesthetically pleasing.
package main

import(
  "fmt"
  "flag"
  "strings"
  "github.com/bwmarrin/discordgo"
)

var Token string
var BotPrefix string

func init() { //go run . -t [TOKEN] -p [PREFIX]
  flag.StringVar(&Token, "t", "", "Bot token")
  flag.StringVar(&BotPrefix, "p", "", "Bot prefix")

  flag.Parse()
}

func main() {
  d, err := discordgo.New("Bot " + Token)
  if err != nil {
    fmt.Println(err)
  }

  d.AddHandler(onReady)
  d.AddHandler(onMessage)

  d.Open()
  fmt.Println("Bot is now running")

  fmt.Scanln() //Many advise against using fmt.Scanln() to keep the client open, but I do not know why so I will use it.

  d.Close()
}

func onReady(s *discordgo.Session, event *discordgo.Ready) {
  s.UpdateGameStatus(Token + "help")
}

func onMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
  if m.Author.ID == s.State.User.ID {
		return
	}

  if strings.HasPrefix(m.Content, BotPrefix + "ping") {
    ping(s, m)
  } else if strings.HasPrefix(m.Content, BotPrefix + "ban") {
    ban(s, m)
  }
}

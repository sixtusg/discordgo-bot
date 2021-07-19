package main

import(
  "fmt"
  "flag"

  "github.com/bwmarrin/discordgo"
)

var Token string

func init() {
  flag.StringVar(&Token, "t", "", "Bot Token")

  flag.Parse()
}

func main() {
  d, err := discordgo.New("Bot " + Token)

  if err != nil {
    fmt.Println(err)
  }

  d.Open()

  fmt.Println("Bot is now running")

  fmt.Scanln()

  d.Close()
}

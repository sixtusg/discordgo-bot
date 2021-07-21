package main

import(
  "strings"
  "strconv"
  "fmt"
  "github.com/bwmarrin/discordgo"
)

func ping(s *discordgo.Session, m *discordgo.MessageCreate) {
  s.ChannelMessageSend(m.ChannelID, "Pong!")
}

func ban(s *discordgo.Session, m *discordgo.MessageCreate) {
  arg := strings.Fields(m.Content)

  if discordgo.PermissionBanMembers != 3 {
    s.ChannelMessageSend(m.ChannelID, "You do not have permission to issue this command.")
    return
  } else if discordgo.PermissionBanMembers != 4 {
    s.ChannelMessageSend(m.ChannelID, "You do not have permission to issue this command.")
    return
  }

  if len(arg) < 2 {
    s.ChannelMessageSend(m.ChannelID, "Syntax: `" + BotPrefix + "ban [mention user] [duration in days]`")
    return
  }

  if arg[1] == "@everyone" {
    s.ChannelMessageSend(m.ChannelID, "You can not issue this command with argument `@everyone`")
    return
  }

  if strings.HasPrefix(arg[1], "<@!") {

    id := getIDFromMention(arg[1])

    days, err := strconv.Atoi(arg[2])

    if err != nil {
      fmt.Println(err)
      return
    }

    s.GuildBanCreate(m.GuildID, id, days)
  } else {
    s.ChannelMessageSend(m.ChannelID, "Syntax: `" + BotPrefix + "ban [mention user] [duration in days]`")
  }

  fmt.Println(discordgo.PermissionBanMembers)
}

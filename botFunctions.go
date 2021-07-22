package main

import(
  "strings"
  "strconv"
  "fmt"
  "github.com/bwmarrin/discordgo"
)
func help(s *discordgo.Session, m *discordgo.MessageCreate) {
    s.ChannelMessageSend(m.ChannelID, "Commands: `" + BotPrefix + "help, " + BotPrefix + "ping, " + BotPrefix + "ban" + "`") //note to self: this is horrendously ugly and must be fixed, especially when considering that it is likely the first command users will issue. in the future, add embeds and order in a hierarchal manner
}

func ping(s *discordgo.Session, m *discordgo.MessageCreate) {
  s.ChannelMessageSend(m.ChannelID, "Pong!")
}

func ban(s *discordgo.Session, m *discordgo.MessageCreate) {
  p, err := s.UserChannelPermissions(m.Author.ID, m.ChannelID)

  if err != nil {
    fmt.Println(err)
  }

  if p&discordgo.PermissionBanMembers == discordgo.PermissionBanMembers {
    arg := strings.Fields(m.Content)

    if len(arg) < 2 {
      s.ChannelMessageSend(m.ChannelID, "Syntax: `" + BotPrefix + "ban [mention user] [duration in days]`")
    }

    if arg[1] == "@everyone" {
      s.ChannelMessageSend(m.ChannelID, "You can not issue this command with argument `@everyone`")
    }

    if strings.HasPrefix(arg[1], "<@!") {

      id := getIDFromMention(arg[1])

      days, err := strconv.Atoi(arg[2])

      if err != nil {
        fmt.Println(err)
      }

      s.GuildBanCreate(m.GuildID, id, days)
      } else {
        s.ChannelMessageSend(m.ChannelID, "Syntax: `" + BotPrefix + "ban [mention user] [duration in days]`")
      }
  } else {
    s.ChannelMessageSend(m.ChannelID, "You do not have permission to issue this command.")
  }
}

func kick(s *discordgo.Session, m *discordgo.MessageCreate) {
  p, err := s.UserChannelPermissions(m.Author.ID, m.ChannelID)

  if err != nil {
    fmt.Println(err)
  }

  if p&discordgo.PermissionKickMembers == discordgo.PermissionKickMembers {
    arg := strings.Fields(m.Content)

    if len(arg) < 2 {
      s.ChannelMessageSend(m.ChannelID, "Syntax: `" + BotPrefix + "kick [mention user]`")
      return
    }

    if arg[1] == "@everyone" {
      s.ChannelMessageSend(m.ChannelID, "You can not issue this command with argument `@everyone`")
    }

    if strings.HasPrefix(arg[1], "<@!") {
      id := getIDFromMention(arg[1])

      s.GuildMemberDelete(m.GuildID, id)
    }
  } else {
    s.ChannelMessageSend(m.ChannelID, "You do not have permission to issue this command.")
  }
}

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
  genericEmbed("Pong!", "", s, m)
}

func ban(s *discordgo.Session, m *discordgo.MessageCreate) {
  p, err := s.UserChannelPermissions(m.Author.ID, m.ChannelID)

  if err != nil {
    fmt.Println(err)
  }

  if p&discordgo.PermissionBanMembers == discordgo.PermissionBanMembers {
    arg := strings.Fields(m.Content)

    if len(arg) < 2 {
      errEmbed("Syntax error", BotPrefix + "ban [mention user] [number of days]", s, m)
      return
    }

    if arg[1] == "@everyone" {
      errEmbed("Syntax error", "You cannot issue this command with argument @everyone", s, m)
    }

    if strings.HasPrefix(arg[1], "<@!") {
      id := getIDFromMention(arg[1])

      days, err := strconv.Atoi(arg[2])

      if err != nil {
        fmt.Println(err)
      }

      s.GuildBanCreate(m.GuildID, id, days)
      } else {
        errEmbed("Syntax error", BotPrefix + "ban [mention user] [number of days]", s, m)
      }
  } else {
    errEmbed("Error", "You do not have permission to issue this command.", s, m)
  }
}

func unban(s *discordgo.Session, m *discordgo.MessageCreate) {
  p, err := s.UserChannelPermissions(m.Author.ID, m.ChannelID)

  if err != nil {
    fmt.Println(err)
  }

  if p&discordgo.PermissionBanMembers == discordgo.PermissionBanMembers {
    arg := strings.Fields(m.Content)

    if len(arg) < 2 {
      errEmbed("Syntax error", BotPrefix + "unban [mention user]", s, m)
      return
    }

    if arg[1] == "@everyone" {
      errEmbed("Syntax error", "You cannot issue this command with argument @everyone", s, m)
    }

    if strings.HasPrefix(arg[1], "<@!") {
      id := getIDFromMention(arg[1])

      s.GuildBanDelete(m.GuildID, id)
    } else {
      errEmbed("Syntax error", BotPrefix + "unban [mention user]", s, m)
    }
  } else {
    errEmbed("Error", "You do not have permission to issue t his command.", s, m)
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

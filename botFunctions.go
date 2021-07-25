package main

import(
  "strings"
  "strconv"
  "fmt"
  "github.com/bwmarrin/discordgo"
)

//simple
func help(s *discordgo.Session, m *discordgo.MessageCreate) {
  //if this isn't the most tedious thing i've ever done...
  //is there any way to not make this code ugly?
  //yours truly, past Sixtus
  banHelp := "**Ban**\n`" + BotPrefix + "ban help`\n\n"
  unbanHelp := "**Unban**\n`" + BotPrefix + "unban help`\n\n"
  kickHelp := "**Kick**\n`" + BotPrefix + "kick help`\n\n"

  genericEmbed("Commands", banHelp + unbanHelp + kickHelp, s, m)
}

func ping(s *discordgo.Session, m *discordgo.MessageCreate) {
  genericEmbed("Pong!", "", s, m)
  fmt.Println(s.HeartbeatLatency())
}

//moderation
func ban(s *discordgo.Session, m *discordgo.MessageCreate) { //spaghetti function
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
      return
    }

    if arg[1] == "help" {
      genericEmbed("Help", "Use this command to ban a user\n\n**Usage**\n`" + BotPrefix + "ban [mention user] [number of days for user to be banned]`", s, m)
      return
    }

    if len(arg) < 3 {
      errEmbed("Syntax error", BotPrefix + "ban [mention user] [number of days]", s, m)
      return
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
      return
    }

    if arg[1] == "help" {
      genericEmbed("Help", "Use this command to unban a user\n\n**Usage**\n`" + BotPrefix + "unban <@![user id]`", s, m)
      return
    }

    if strings.HasPrefix(arg[1], "<@!") {
      id := getIDFromMention(arg[1])

      s.GuildBanDelete(m.GuildID, id)
    } else {
      errEmbed("Syntax error", BotPrefix + "unban [mention user]", s, m)
    }
  } else {
    errEmbed("Error", "You do not have permission to issue this command.", s, m)
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
      errEmbed("Syntax error", BotPrefix + "kick [mention user]", s, m)
      return
    }

    if arg[1] == "@everyone" {
      errEmbed("Syntax error", "You cannot issue this command with argument @everyone", s, m)
    }

    if arg[1] == "help" {
      genericEmbed("Help", "Use this command to kick a user\n\n**Usage**\n`" + BotPrefix + "kick [mention a user]`", s, m)
    }

    if strings.HasPrefix(arg[1], "<@!") {
      id := getIDFromMention(arg[1])

      s.GuildMemberDelete(m.GuildID, id)
    } else {
      errEmbed("Syntax error", BotPrefix + "unban [mention user]", s, m)
    }
  } else {
    errEmbed("Error", "You do not have permission to issue this command.", s, m)
  }
}

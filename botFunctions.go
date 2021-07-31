package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

//simple
func help(s *discordgo.Session, m *discordgo.MessageCreate) {
	//if this isn't the most tedious thing i've ever done...
	//is there any way to not make this code ugly?
	//yours truly, past Sixtus
	banHelp := "**Ban**\n`" + botPrefix + "ban help`\n\n"
	unbanHelp := "**Unban**\n`" + botPrefix + "unban help`\n\n"
	kickHelp := "**Kick**\n`" + botPrefix + "kick help`\n\n"

	genericEmbed("Commands", banHelp+unbanHelp+kickHelp, s, m)
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
			errEmbed("Syntax error", botPrefix+"ban [mention user] [number of days]", s, m)
			return
		}

		if arg[1] == "@everyone" {
			errEmbed("Syntax error", "You cannot issue this command with argument @everyone", s, m)
			return
		}

		if arg[1] == "help" {
			genericEmbed("Help", "Use this command to ban a user\n\n**Usage**\n`"+botPrefix+"ban [mention user] [number of days for user to be banned]`", s, m)
			return
		}

		if len(arg) < 3 {
			errEmbed("Syntax error", botPrefix+"ban [mention user] [number of days]", s, m)
			return
		}

		if strings.HasPrefix(arg[1], "<@!") {
			id := getIDFromMention(arg[1])

			days, err := strconv.Atoi(arg[2])

			if err != nil {
				fmt.Println(err)
			}

			s.GuildBanCreate(m.GuildID, id, days)
			successEmbed("Success", "User successfully banned.", s, m)
		} else {
			errEmbed("Syntax error", botPrefix+"ban [mention user] [number of days]", s, m)
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
			errEmbed("Syntax error", botPrefix+"unban [mention user]", s, m)
			return
		}

		if arg[1] == "@everyone" {
			errEmbed("Syntax error", "You cannot issue this command with argument @everyone", s, m)
			return
		}

		if arg[1] == "help" {
			genericEmbed("Help", "Use this command to unban a user\n\n**Usage**\n`"+botPrefix+"unban <@![user id]`", s, m)
			return
		}

		if strings.HasPrefix(arg[1], "<@!") {
			id := getIDFromMention(arg[1])

			s.GuildBanDelete(m.GuildID, id)
			successEmbed("Success", "User successfully unbanned.", s, m)
		} else {
			errEmbed("Syntax error", botPrefix+"unban [mention user]", s, m)
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
			errEmbed("Syntax error", botPrefix+"kick [mention user]", s, m)
			return
		}

		if arg[1] == "@everyone" {
			errEmbed("Syntax error", "You cannot issue this command with argument @everyone", s, m)
		}

		if arg[1] == "help" {
			genericEmbed("Help", "Use this command to kick a user\n\n**Usage**\n`"+botPrefix+"kick [mention a user]`", s, m)
		}

		if strings.HasPrefix(arg[1], "<@!") {
			id := getIDFromMention(arg[1])

			s.GuildMemberDelete(m.GuildID, id)
			successEmbed("Success", "User successfully kicked.", s, m)
		} else {
			errEmbed("Syntax error", botPrefix+"unban [mention user]", s, m)
		}
	} else {
		errEmbed("Error", "You do not have permission to issue this command.", s, m)
	}
}

func setup(s *discordgo.Session, m *discordgo.MessageCreate) {
	p, err := s.UserChannelPermissions(m.Author.ID, m.ChannelID)

	if err != nil {
		fmt.Println(err)
	}

	if p&discordgo.PermissionManageRoles == discordgo.PermissionManageRoles {
		arg := strings.Fields(m.Content)

		if len(arg) > 1 {
			if arg[1] == "help" {
				genericEmbed("Help", "Use this command to create a muted role for muting users\n\n**Usage**\n`"+botPrefix+"setup`", s, m)
				return
			}
		}

		role, err := s.GuildRoleCreate(m.GuildID)

		if err != nil {
			fmt.Println(err)
		}

		if getRoleIDFromMutedRole(s, m) != "" {
			s.GuildRoleDelete(m.GuildID, getRoleIDFromMutedRole(s, m))
		}

		s.GuildRoleEdit(m.GuildID, role.ID, "Muted", 0x646464, false, 1024, false)
		successEmbed("Success", "Muted role successfuly created.", s, m)
	} else {
		errEmbed("Error", "You do not have permission to issue this command.", s, m)
	}
}

func mute(s *discordgo.Session, m *discordgo.MessageCreate) {
	p, err := s.UserChannelPermissions(m.Author.ID, m.ChannelID)

	if err != nil {
		fmt.Println(err)
	}

	if p&discordgo.PermissionManageRoles == discordgo.PermissionManageRoles {
		arg := strings.Fields(m.Content)

		if len(arg) < 2 {
			errEmbed("Syntax error", "`"+botPrefix+"mute [mention user]", s, m)
			return
		}

		if arg[1] == "@everyone" {
			errEmbed("Syntax error", "You cannot issue this command with argument @everyone", s, m)
			return
		}

		if arg[1] == "help" {
			genericEmbed("Help", "Use this command to mute a user\n\n**Usage**\n`"+botPrefix+"mute [mention user]`", s, m)
			return
		}

		if strings.HasPrefix(arg[1], "<@!") {
			id := getIDFromMention(arg[1])

			if getRoleIDFromMutedRole(s, m) == "" {
				errEmbed("Error", "You do have not run the `"+botPrefix+"setup` command yet. Muted role cannot be added", s, m)
				return
			}

			s.GuildMemberRoleAdd(m.GuildID, id, getRoleIDFromMutedRole(s, m))
		} else {
			errEmbed("Syntax error", botPrefix+"mute [mention user]", s, m)
		}

	} else {
		errEmbed("Error", "You do not have permission to issue this command.", s, m)
	}
}

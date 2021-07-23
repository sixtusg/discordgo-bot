package main
import(
  "github.com/bwmarrin/discordgo"
)

func getIDFromMention(content string) string {
  return content[3 : len(content)-1]
}

func genericEmbed(title string, description string, s *discordgo.Session, m *discordgo.MessageCreate) {
  msgEmb := &discordgo.MessageEmbed{}
  msgEmb.Title = title
  msgEmb.Description = description
  msgEmb.Color = 0x0096ff
  s.ChannelMessageSendEmbed(m.ChannelID, msgEmb)
}

func errEmbed(title string, description string, s *discordgo.Session, m *discordgo.MessageCreate) {
  msgEmb := &discordgo.MessageEmbed{}
  msgEmb.Title = title
  msgEmb.Description = description
  msgEmb.Color = 0xff0000
  s.ChannelMessageSendEmbed(m.ChannelID, msgEmb)
}

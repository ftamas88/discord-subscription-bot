package bot

import "github.com/bwmarrin/discordgo"

func (b *Bot) messageReceive(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	b.logger.Infof("Message: %+v | Channel: %s", m.Content, m.ChannelID)
	/*
		if m.ChannelID == strconv.Itoa(b.cfg.NotificationReceiver.TargetID) {
			_, _ = s.ChannelMessageSend(m.ChannelID, "message received")
			return
		}
	*/
}

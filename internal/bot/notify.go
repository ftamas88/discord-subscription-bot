package bot

import (
	"fmt"
	"strconv"

	"github.com/bwmarrin/discordgo"
	"github.com/ftamas88/discord-subscription-bot/internal/config"
)

func (b *Bot) notify(s *discordgo.Session, m *discordgo.Member, new bool, role string) {
	switch b.cfg.NotificationReceiver.TargetType {
	case config.NotificationTargetChannel:
		msg := fmt.Sprintf("User [%s] has gained a role: %s", m.User.Username, role)
		if !new {
			msg = fmt.Sprintf("User [%s] has lost a role: %s", m.User.Username, role)
		}

		if _, err := s.ChannelMessageSend(
			strconv.Itoa(b.cfg.NotificationReceiver.TargetID),
			msg,
		); err != nil {
			b.logger.Warnf("unable to notify the channel: %v", err)
		}
		b.logger.Info(msg)

		break
	case config.NotificationTargetUser:
		b.logger.Warn("TODO: DM type notifications are not supported YET!")
		break
	default:
		b.logger.Warn(
			"Invalid target: supported options: [%s|%s]",
			config.NotificationTargetChannel,
			config.NotificationTargetUser,
		)
	}
}

package bot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/ftamas88/discord-subscription-bot/internal/model"
)

func (b *Bot) startup(s *discordgo.Session, m *discordgo.GuildCreate) {
	for _, r := range m.Guild.Roles {
		b.roles[r.ID] = r.Name
	}

	if b.fetchUsers() {
		return
	}

	for _, discordUser := range m.Guild.Members {
		userRoles := make([]model.Role, 0)
		for _, r := range discordUser.Roles {
			userRoles = append(userRoles, model.Role{
				RoleID: r,
				Name:   b.roles[r],
			})
		}

		usr, found := b.users[discordUser.User.ID]
		if !found {
			b.logger.Debugf(
				"New user found, storing in the db: %s [%s]",
				discordUser.User.Username,
				discordUser.User.ID,
			)

			newUser := &model.Member{
				DiscordID: discordUser.User.ID,
				Name:      discordUser.User.Mention(),
				Username:  discordUser.User.Username,
				Email:     discordUser.User.Email,
				Roles:     userRoles,
			}
			if err := b.repository.CreateUser(newUser); err != nil {
				b.logger.Errorf("unable to store user: %d, error: %s", newUser.ID, err)
			}

			continue
		}

		b.roleUpdates(s, usr, discordUser, userRoles)
	}

	b.logger.Info("== Server Online members ==")
	for _, v := range m.Guild.Presences {
		b.logger.Infof("%s\t(#%s)\t[%s]", b.users[v.User.ID].Username, v.User.ID, v.Status)
	}
}

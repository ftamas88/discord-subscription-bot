package bot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/ftamas88/discord-subscription-bot/internal/model"
)

func (b *Bot) roleChange(s *discordgo.Session, m *discordgo.GuildMemberUpdate) {
	userRoles := make([]model.Role, 0)
	for _, r := range m.Roles {
		userRoles = append(userRoles, model.Role{
			RoleID: r,
			Name:   b.roles[r],
		})
	}

	if b.fetchUsers() {
		return
	}

	usr, found := b.users[m.User.ID]
	if !found {
		return
	}

	b.roleUpdates(s, usr, m.Member, userRoles)
}

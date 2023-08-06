package bot

import (
	"strconv"

	"github.com/bwmarrin/discordgo"
	"github.com/ftamas88/discord-subscription-bot/internal/model"
)

type userRoleStatus struct {
	RoleID    string
	FoundInDB bool
	FoundNow  bool
}

func (b *Bot) fetchUsers() bool {
	users, err := b.repository.Users()
	if err != nil {
		b.logger.Errorf("unable to fetch members from the database: %v", err)

		return true
	}

	for _, v := range users {
		b.users[v.DiscordID] = v
	}

	return false
}

func (b *Bot) roleUpdates(
	session *discordgo.Session,
	usr model.Member,
	discordUser *discordgo.Member,
	userRoles []model.Role,
) {
	userRoleStatuses := map[string]userRoleStatus{}

	// update user
	if len(usr.Roles) != len(userRoles) {
		b.logger.Infof("Roles has been updated: %s [%s]", usr.Username, usr.DiscordID)

		// default statuses
		for _, subRole := range b.cfg.SubscribedRoles {
			userRoleStatuses[strconv.Itoa(subRole.ID)] = userRoleStatus{
				RoleID:    strconv.Itoa(subRole.ID),
				FoundInDB: false,
				FoundNow:  false,
			}
		}

		for _, r := range usr.Roles {
			for _, subRole := range b.cfg.SubscribedRoles {
				if r.RoleID == strconv.Itoa(subRole.ID) {
					userRoleStatuses[r.RoleID] = userRoleStatus{
						RoleID:    r.RoleID,
						FoundInDB: true,
						FoundNow:  userRoleStatuses[r.RoleID].FoundNow,
					}
				}
			}
		}

		usr.Roles = userRoles
		if err := b.repository.UpdateUser(&usr); err != nil {
			b.logger.Errorf("unable to update user: %s, error: %s", usr.DiscordID, err)
		}

		for _, r := range userRoles {
			for _, subRole := range b.cfg.SubscribedRoles {
				if r.RoleID == strconv.Itoa(subRole.ID) {
					userRoleStatuses[r.RoleID] = userRoleStatus{
						RoleID:    r.RoleID,
						FoundInDB: userRoleStatuses[r.RoleID].FoundInDB,
						FoundNow:  true,
					}
				}
			}
		}

		for _, st := range userRoleStatuses {
			if st.FoundInDB && !st.FoundNow {
				b.notify(session, discordUser, false, b.roles[st.RoleID])

				continue
			}

			if !st.FoundInDB && st.FoundNow {
				b.notify(session, discordUser, true, b.roles[st.RoleID])

				continue
			}
		}
	}
}

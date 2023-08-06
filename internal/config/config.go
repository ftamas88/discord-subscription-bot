// Package config It's a package containing the configuration for the application
package config

import (
	"log"
	"os"
	"strconv"
	"strings"
)

// AppConfig contains the necessary configuration to run the bot
type AppConfig struct {
	Token                string
	SubscribedRoles      []Role
	NotificationReceiver NotificationReceiver
}

func NewConfig() *AppConfig {
	nID, err := strconv.Atoi(os.Getenv("NOTIFICATION_ID"))
	if err != nil {
		log.Fatal("unable to parse notification id")
	}

	return &AppConfig{
		Token:           os.Getenv("BOT_TOKEN"),
		SubscribedRoles: loadRoles(),
		NotificationReceiver: NotificationReceiver{
			TargetType: NotificationTarget(os.Getenv("NOTIFICATION_TYPE")),
			TargetID:   nID,
		},
	}
}

func loadRoles() []Role {
	var subscribedRoles []Role
	roles := strings.Split(os.Getenv("SUBSCRIBED_ROLE_IDS"), "|")
	for _, r := range roles {
		rID, err := strconv.Atoi(r)
		if err != nil {
			log.Fatal("subscribed role id is not a number")
		}

		subscribedRoles = append(subscribedRoles, Role{
			ID:   rID,
			Name: "",
		})
	}

	return subscribedRoles
}

type Role struct {
	ID   int
	Name string
}

type NotificationReceiver struct {
	TargetType NotificationTarget
	TargetID   int
}

const (
	NotificationTargetUser    NotificationTarget = "user"
	NotificationTargetChannel NotificationTarget = "channel"
)

type NotificationTarget string

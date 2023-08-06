package model

import (
	"gorm.io/gorm"
)

type Member struct {
	gorm.Model
	ID        uint   `gorm:"primary_key" json:"Id"`
	DiscordID string `json:"DiscordId"`
	Name      string `json:"Name"`
	Username  string `json:"Username"`
	Email     string `json:"Email"`
	Roles     []Role `gorm:"ForeignKey:MemberID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"Roles"`
}

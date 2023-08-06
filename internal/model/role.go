package model

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	RoleID   string `json:"RoleID"`
	Name     string `json:"Name"`
	MemberID uint   `json:"MemberID"`
}

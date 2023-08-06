package model

import (
	"time"

	"gorm.io/gorm"
)

type State struct {
	gorm.Model
	LastUpdated time.Time
}

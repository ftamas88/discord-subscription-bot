package storage

import (
	"github.com/ftamas88/discord-subscription-bot/internal/model"
)

//go:generate mockery -name=Repository
type Repository interface {
	Users() ([]model.Member, error)
	CreateUser(user *model.Member) error
	UpdateUser(user *model.Member) error
}

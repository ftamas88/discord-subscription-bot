package storage

import (
	"github.com/ftamas88/discord-subscription-bot/internal/model"
	"go.uber.org/zap"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Sqlite struct {
	log *zap.SugaredLogger
	db  *gorm.DB
}

func NewSqlite(l *zap.SugaredLogger) *Sqlite {
	db, err := gorm.Open(sqlite.Open("data/database.db"), &gorm.Config{})
	if err != nil {
		l.Fatal("failed to connect database")
	}

	// Migrate the schemas
	err = db.AutoMigrate(&model.Member{}, &model.Role{}, &model.State{})
	if err != nil {
		l.Fatalf("unable to create database scheme: %v", err)
	}

	return &Sqlite{
		db:  db,
		log: l,
	}
}

func (s *Sqlite) Users() ([]model.Member, error) {
	var members []model.Member

	result := s.db.Model(&model.Member{}).Preload("Roles").Find(&members)

	return members, result.Error
}

func (s *Sqlite) CreateUser(user *model.Member) error {
	result := s.db.Create(user)

	return result.Error
}

func (s *Sqlite) UpdateUser(user *model.Member) error {
	result := s.db.
		Unscoped().
		Model(&user).
		Association("Roles").
		Unscoped().
		Replace(user.Roles)

	return result
}

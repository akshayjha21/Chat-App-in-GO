package postgres

import (
	// "fmt"
	"log"
	// "log/slog"

	"github.com/akshayjha21/Chat-App-in-GO/Backend/internal/config"
	"github.com/akshayjha21/Chat-App-in-GO/Backend/internal/types"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	Db *gorm.DB
}

func New(cfg *config.Config) (*Postgres, error) {
	db, err := gorm.Open(postgres.Open(cfg.DatabaseURL), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	log.Println("Database connection has been established");
	err = db.AutoMigrate(
		&types.User{},
		&types.Room{},
		&types.RoomMember{},
		&types.Message{},
	)
	if err != nil {
		return nil, err
	}

	return &Postgres{Db: db}, nil
}
func (p *Postgres) RegisterUser(user *types.User) (*types.User, error) {
	if err := p.Db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

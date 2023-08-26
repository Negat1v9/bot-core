package storage

import (
	"context"

	"github.com/Negat1v9/bot-core/models"
)

type UserRepository interface {
	SaveUser(ctx context.Context, u *models.User) error
	Find(ctx context.Context, ID int) (*models.User, error)
	ChangeUser(ctx context.Context, ID int) error
	CreateTable() error
}

type Storage interface {
	User() UserRepository
}

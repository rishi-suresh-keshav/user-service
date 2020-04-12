package repository

import (
	"context"
	"github.com/rishi-suresh-keshav/user-service/models"
)

type UserRepo interface {
	GetUserById(ctx context.Context, id int64) (*models.User, error)
	CreateUser(ctx context.Context, user *models.User) (int64, error)
}

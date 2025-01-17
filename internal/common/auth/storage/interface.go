package storage

import (
	"context"

	"github.com/erupshis/golang-integration-developer-test/internal/common/auth/models"
)

type BaseAuthStorage interface {
	AddUser(ctx context.Context, user *models.User) error
	GetUserByLogin(ctx context.Context, login string) (*models.User, error)
}

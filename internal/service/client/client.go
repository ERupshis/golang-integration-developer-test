package client

import (
	"context"

	"github.com/erupshis/golang-integration-developer-test/internal/service/models"
)

type BaseClient interface {
	GetGames(ctx context.Context, platform string) (models.Games, error)
	GetBalance(ctx context.Context, playerID string) (int64, error)
}

package storage

import (
	"github.com/erupshis/golang-integration-developer-test/internal/players/models"
)

type BaseUserStorage interface {
	GetUserByID(id int64) (*models.UserData, error)
	WithdrawBalance(ID, amount int64) error
}

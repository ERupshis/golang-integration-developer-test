package storage

import (
	"github.com/erupshis/golang-integration-developer-test/internal/players/models"
)

type BaseUserStorage interface {
	GetUserByID(id int64) (*models.UserDataP, error)
	WithdrawBalance(ID, amount int64) (int64, error)
}

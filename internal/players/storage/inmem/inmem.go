package inmem

import (
	"sync"

	"github.com/erupshis/golang-integration-developer-test/internal/players/models"
	"github.com/erupshis/golang-integration-developer-test/internal/players/storage"
)

var (
	_ storage.BaseUserStorage = (*UserStorage)(nil)
)

type UserStorage struct {
	users map[int64]models.UserData
	mu    sync.RWMutex
}

func NewUserStorage(users map[int64]models.UserData) *UserStorage {
	return &UserStorage{users: users}
}

func (us *UserStorage) GetUserByID(ID int64) (*models.UserData, error) {
	us.mu.RLock()
	defer us.mu.RUnlock()

	userData, ok := us.users[ID]
	if !ok {
		return nil, storage.ErrUserNotFound
	}
	return &userData, nil
}

func (us *UserStorage) WithdrawBalance(ID, amount int64) error {
	us.mu.RLock()
	userData, ok := us.users[ID]
	if !ok {
		return storage.ErrUserNotFound
	}

	if userData.Balance < amount {
		return storage.ErrUserInSufficientFunds
	}
	us.mu.RUnlock()

	us.mu.Lock()
	defer us.mu.Unlock()

	userData.Balance -= amount
	us.users[ID] = userData
	return nil
}

package storage

import (
	"fmt"
)

var (
	ErrUserNotFound          = fmt.Errorf("user not found")
	ErrUserInSufficientFunds = fmt.Errorf("balance is too low")
)

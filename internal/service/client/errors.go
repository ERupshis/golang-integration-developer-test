package client

import (
	"fmt"
)

var (
	ErrInvalidPlatform   = fmt.Errorf("invalid game platform")
	ErrUserNotFound      = fmt.Errorf("user not found")
	ErrInsufficientFunds = fmt.Errorf("insufficient funds")
)

package models

//go:generate easyjson -all models.go
type UserDataP struct {
	ID      int64 `json:"id"`
	Balance int64 `json:"balance"`
}

type UserWithdrawP struct {
	ID     int64 `json:"id"`
	Amount int64 `json:"amount"`
}

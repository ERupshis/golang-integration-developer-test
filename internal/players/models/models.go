package models

//go:generate easyjson -all models.go
type UserData struct {
	ID      int64
	Balance int64
}

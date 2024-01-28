package models

import (
	pb_integr "github.com/erupshis/golang-integration-developer-test/pb/integration"
)

//go:generate easyjson -all models.go
type Currency struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

func ConvertCurrencyFromGRPC(currency *pb_integr.Currency) *Currency {
	return &Currency{
		Code: currency.GetCode(),
		Name: currency.GetName(),
	}
}

type Game struct {
	ID               string
	Title            string
	ShortDescription string
	GameURL          string
}

func ConvertGameToGRPC(game *Game) *pb_integr.Game {
	return &pb_integr.Game{
		Id:               game.ID,
		Title:            game.Title,
		ShortDescription: game.ShortDescription,
		GameUrl:          game.GameURL,
	}
}

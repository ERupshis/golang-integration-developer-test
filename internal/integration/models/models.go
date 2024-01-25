package models

import (
	"github.com/erupshis/golang-integration-developer-test/pb"
)

//go:generate easyjson -all models.go
type Currency struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

func ConvertCurrencyFromGRPC(currency *pb.Currency) *Currency {
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

func ConvertGameToGRPC(game *Game) *pb.Game {
	return &pb.Game{
		Id:               game.ID,
		Title:            game.Title,
		ShortDescription: game.ShortDescription,
		GameUrl:          game.GameURL,
	}
}

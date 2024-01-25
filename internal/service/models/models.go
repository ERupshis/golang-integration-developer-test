package models

type Games []Game

//go:generate easyjson -all models.go
type Game struct {
	ID                   int64  `json:"id"`
	Title                string `json:"title"`
	Thumbnail            string `json:"thumbnail"`
	ShortDescription     string `json:"short_description"`
	GameURL              string `json:"game_url"`
	Genre                string `json:"genre"`
	Platform             string `json:"platform"`
	Publisher            string `json:"publisher"`
	Developer            string `json:"developer"`
	ReleaseDate          string `json:"release_date"`
	FreeToGameProfileURL string `json:"freetogame_profile_url"`
}

func (g *Games) FindGameByID(ID int64) *Game {
	for _, game := range *g {
		if game.ID == ID {
			return &game
		}
	}

	return nil
}

type UserData struct {
	ID      int64 `json:"id"`
	Balance int64 `json:"balance"`
}

type UserWithdraw struct {
	ID     int64 `json:"id"`
	Amount int64 `json:"amount"`
}

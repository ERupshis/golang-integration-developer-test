package models

type Games []Game

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

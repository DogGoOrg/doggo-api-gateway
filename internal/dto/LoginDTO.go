package dto

type LoginDto struct {
	Id           string `json:"id"`
	Email        string `json:"email"`
	RefreshToken string `json:"refreshToken"`
	AccessToken  string `json:"accessToken"`
}

package domain

import "golang.org/x/net/context"

type RefreshTokenRequest struct {
	RefreshToken string `json:"refreshToken"`
}

type RefreshTokenResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type RefreshTokenUsecase interface {
	GetUserByID(c context.Context, id string) (*User, error)
	CreateAccessToken(user *User, expiry int) (accessToken string, err error)
	CreateRefreshToken(user *User, expiry int) (refreshToken string, err error)
	ExtractIDFromToken(requestToken string) (string, error)
}

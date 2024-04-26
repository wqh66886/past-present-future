package domain

import "context"

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type LoginUsecase interface {
	GetUserByEmail(c context.Context, email string) (*User, error)
	CreateAccessToken(user *User) (string, error)
	CreateRefreshToken(user *User) (string, error)
}

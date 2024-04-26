package domain

import "context"

type SignUpRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignupResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type SignupUsecase interface {
	Create(c context.Context, user *User) error
	GetUserByEmail(c context.Context, email string) (*User, error)
	CreateAccessToken(user *User, expiry int) (accessToken string, err error)
	CreateRefreshToken(user *User, expiry int) (refreshToken string, err error)
}

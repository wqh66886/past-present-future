package mapper

import (
	"github.com/wqh66886/past-present-future/domain"
	"golang.org/x/net/context"
)

type loginMapper struct {
	userRepository domain.UserRepository
}

func NewLoginMapper(userRepository domain.UserRepository) domain.LoginUsecase {
	return &loginMapper{
		userRepository: userRepository,
	}
}

func (l loginMapper) GetUserByEmail(c context.Context, email string) (domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (l loginMapper) CreateAccessToken(user *domain.User) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (l loginMapper) CreateRefreshToken(user *domain.User) (string, error) {
	//TODO implement me
	panic("implement me")
}

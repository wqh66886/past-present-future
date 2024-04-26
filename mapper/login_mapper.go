package mapper

import (
	"github.com/wqh66886/past-present-future/define"
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

func (l loginMapper) GetUserByEmail(c context.Context, email string) (*domain.User, error) {
	user, err := l.userRepository.GetByEmail(c, email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (l loginMapper) CreateAccessToken(user *domain.User) (string, error) {
	token, err := define.CreateToken(user.Name, user.ID)
	return token, err
}

func (l loginMapper) CreateRefreshToken(user *domain.User) (string, error) {
	token, err := define.CreateToken(user.Name, user.ID)
	return token, err
}

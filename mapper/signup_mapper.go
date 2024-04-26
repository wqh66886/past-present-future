package mapper

import (
	"context"
	"github.com/wqh66886/past-present-future/define"
	"github.com/wqh66886/past-present-future/domain"
)

type signupMapper struct {
	userRepository domain.UserRepository
}

func NewSignupMapper(userRepository domain.UserRepository) domain.SignupUsecase {
	return &signupMapper{userRepository: userRepository}
}

func (s signupMapper) Create(c context.Context, user *domain.User) error {
	err := s.userRepository.Create(c, user)
	if err != nil {
		return err
	}
	return nil
}

func (s signupMapper) GetUserByEmail(c context.Context, email string) (*domain.User, error) {
	user, err := s.userRepository.GetByEmail(c, email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s signupMapper) CreateAccessToken(user *domain.User, expiry int) (accessToken string, err error) {
	token, err := define.CreateToken(user.Name, user.ID, expiry)
	return token, err
}

func (s signupMapper) CreateRefreshToken(user *domain.User, expiry int) (refreshToken string, err error) {
	token, err := define.CreateToken(user.Name, user.ID, expiry)
	return token, err
}

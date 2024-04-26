package mapper

import (
	"github.com/wqh66886/past-present-future/define"
	"github.com/wqh66886/past-present-future/domain"
	"golang.org/x/net/context"
)

type refreshTokenMapper struct {
	userRepository domain.UserRepository
}

func NewRefreshTokenMapper(userRepository domain.UserRepository) domain.RefreshTokenUsecase {
	return &refreshTokenMapper{userRepository: userRepository}
}

func (r refreshTokenMapper) GetUserByID(c context.Context, id string) (*domain.User, error) {
	user, err := r.userRepository.GetByID(c, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r refreshTokenMapper) CreateAccessToken(user *domain.User, expiry int) (string, error) {
	token, err := define.CreateToken(user.Name, user.ID, expiry)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r refreshTokenMapper) CreateRefreshToken(user *domain.User, expiry int) (string, error) {
	token, err := define.CreateToken(user.Name, user.ID, expiry)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r refreshTokenMapper) ExtractIDFromToken(requestToken string) (string, error) {
	userId, err := define.ExtractIDFromToken(requestToken)
	if err != nil {
		return "", err
	}
	return userId, nil
}

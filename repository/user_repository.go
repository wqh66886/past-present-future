package repository

import (
	"context"
	"github.com/wqh66886/past-present-future/define"
	"github.com/wqh66886/past-present-future/domain"
)

type userRepository struct {
	ctx context.Context
}

func NewUserRepository(ctx context.Context) domain.UserRepository {
	return &userRepository{
		ctx: ctx,
	}
}

func (u userRepository) Create(c context.Context, user *domain.User) error {
	err := define.DB.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

func (u userRepository) Fetch(c context.Context) (users []domain.User, err error) {
	err = define.DB.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return
}

func (u userRepository) GetByEmail(c context.Context, email string) (user *domain.User, err error) {
	err = define.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u userRepository) GetByID(c context.Context, id string) (user *domain.User, err error) {
	err = define.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

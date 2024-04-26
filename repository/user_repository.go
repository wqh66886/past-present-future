package repository

import (
	"context"
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
	//TODO implement me
	panic("implement me")
}

func (u userRepository) Fetch(c context.Context) ([]domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u userRepository) GetByEmail(c context.Context, email string) (domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u userRepository) GetByID(c context.Context, id string) (domain.User, error) {
	//TODO implement me
	panic("implement me")
}

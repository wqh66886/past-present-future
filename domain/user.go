package domain

import (
	"context"
	"time"
)

type User struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	CreateTime time.Time `json:"createTime"`
	UpdateTime time.Time `json:"updateTime"`
}

type UserRepository interface {
	Create(c context.Context, user *User) error
	Fetch(c context.Context) ([]User, error)
	GetByEmail(c context.Context, email string) (User, error)
	GetByID(c context.Context, id string) (User, error)
}

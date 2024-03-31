package domain

import (
	"TaskManger/models"
	"context"
)

type User struct {
	Id       string `json:"id" bson:"_id,omitempty"`
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}

type UserRepository interface {
	Create(c context.Context, user *User) (*User, error)
	Delete(c context.Context, user *User) (*User, error)
	UpdatePassword(c context.Context, user *User) (*User, error)
	UpdateUsername(c context.Context, user *User) (*User, error)
	GetByUsername(c context.Context,username string) (*User, error)
	GetById(c context.Context, id string) (*User, error)
}

type UserUseCase interface {
	Register(c context.Context, payload *models.UserCreate) (*User, error)
	Login(c context.Context, payload *models.UserCreate) (string, error)
	Delete(c context.Context, userId string) (*User, error)
	UpdatePassword(c context.Context, userId string, payload *models.UserUpdatePassword) (*User, error)
	UpdateUsername(c context.Context, userId string, payload *models.UserUpdateUsername) (*User, error)
	GetById(c context.Context, id string) (*User, error)
	GetByUsername(c context.Context, username string) (*User, error)
}
package repository

import (
	"TaskManger/domain"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	database   mongo.Database
	collection string
}

func NewUserRepository(db mongo.Database, collection string) domain.UserRepository {
	return &UserRepository{
		database:   db,
		collection: collection,
	}
}

// GetById implements domain.UserRepository.
func (u *UserRepository) GetById(id string) (domain.User, error) {
	panic("unimplemented")
}

// Update implements domain.UserRepository.
func (u *UserRepository) Update(user *domain.User) (domain.User, error) {
	panic("unimplemented")
}



// Create implements domain.UserRepository.
func (u *UserRepository) Create(user *domain.User) (domain.User, error) {
	panic("unimplemented")
}

// GetByUsername implements domain.UserRepository.
func (u *UserRepository) GetByUsername(username string) (domain.User, error) {
	panic("unimplemented")
}

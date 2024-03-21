package repository

import (
	"TaskManger/domain"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	database   *mongo.Database
	collection string
}

func NewUserRepository(db *mongo.Database, collection string) domain.UserRepository {
	return &UserRepository{
		database:   db,
		collection: collection,
	}
}


// Create implements domain.UserRepository.
func (u *UserRepository) Create(c context.Context, user *domain.User) (*domain.User, error) {
	user.Id = primitive.NewObjectID().Hex()
	result, err := u.database.Collection(u.collection).InsertOne(c, user)
	return &domain.User{
		Id:       result.InsertedID.(string),
		Username: user.Username,
		Password: user.Password,
	}, err
}

// Delete implements domain.UserRepository.
func (u *UserRepository) Delete(c context.Context, user *domain.User) (*domain.User ,error) {
	filter := bson.M{"_id": user.Id}
	_, err := u.database.Collection(u.collection).DeleteOne(c, filter)
	return user, err

}

// UpdatePassword implements domain.UserRepository.
func (u *UserRepository) UpdatePassword(c context.Context, user *domain.User) (*domain.User, error) {
	filter := bson.M{"_id": user.Id}
    update := bson.M{"$set": bson.M{"password": user.Password}}
    _, err := u.database.Collection(u.collection).UpdateOne(c, filter, update)
	return user, err

}

// UpdateUsername implements domain.UserRepository.
func (u *UserRepository) UpdateUsername(c context.Context, user *domain.User) (*domain.User, error) {
	filter := bson.M{"_id": user.Id}
    update := bson.M{"$set": bson.M{"username": user.Password}}
    _, err := u.database.Collection(u.collection).UpdateOne(c, filter, update)
	return user, err
}

// GetById implements domain.UserRepository.
func (u *UserRepository) GetById(c context.Context, id string) (*domain.User, error) {
	filter := bson.M{"_id": id}
	result := u.database.Collection(u.collection).FindOne(c, filter)
	var user domain.User
	err := result.Decode(&user);
	return &user, err
}

// GetByUsername implements domain.UserRepository.
func (u *UserRepository) GetByUsername(c context.Context, username string) (*domain.User, error) {
	filter := bson.M{"username": username}
	result := u.database.Collection(u.collection).FindOne(c, filter)
	var user domain.User
	err := result.Decode(&user);
	return &user, err
}

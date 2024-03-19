package domain

type User struct {
	Id       string `json:"id" bson:"_id"`
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}


type UserRepository interface {
	Create(user *User) (User, error)
	GetByUsername(username string) (User, error)
	GetById(id string) (User, error)
	Update(user *User) (User, error)
}


type UserUseCase interface {
	Create(user *User) (User, error)
	GetByUsername(username string) (User, error)
	GetById(id string) (User, error)
	Update(user *User) (User, error)
}
package models

type UserCreate struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserUpdateUsername struct {
	Username string `json:"username"`
}

type UserUpdatePassword struct {
	Password string `json:"password"`
}
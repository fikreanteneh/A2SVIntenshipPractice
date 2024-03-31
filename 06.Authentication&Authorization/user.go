package main

import (

	"github.com/gin-gonic/gin"
)

var client = GetMongoClient()
var collection = client.Database("taskmanager").Collection("users")

type User struct {
	id       string `bson:"_id" json:"_id,omitempty"`
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}

type UserCreate struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}

func CreateUser(context *gin.Context){
	var user UserCreate
	context.Bind(&user)
	user.Password, _ = EncryptPassword(user.Password)
	createUser,err := collection.FindOne(context, User{Username: user.Username})
	if  createUser != nil{
		context.JSON(400, gin.H{"status": "error", "message": "Username already exists"})
	}
	createUser, err = collection.InsertOne(context, user)
	if err != nil {
		context.JSON(500, gin.H{"status": "error", "message": "Error while creating user"})
		return
	}
	context.JSON(200, gin.H{"status": "success", "message": "User created successfully"})
}

func LoginUser(context *gin.Context){
	var user UserCreate
	var databaseUser User
	context.Bind(&user)
	err := collection.FindOne(context, User{Username: user.Username}).Decode(&databaseUser)
	if err := ComparePasswords(databaseUser.Password, user.Password); err != nil {
		context.JSON(400, gin.H{"status": "error", "message": "Invalid username/password"})
		return
	}
	
	context.JSON(200, gin.H{"status": "success", "message": "User logged in successfully", "token": })

}
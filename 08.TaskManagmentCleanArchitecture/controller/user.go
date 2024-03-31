package controller

import (
	"TaskManger/domain"
	"TaskManger/middleware"
	"TaskManger/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)


type UserController struct {
	userUseCase domain.UserUseCase
}

func NewUserController(userUseCase domain.UserUseCase) *UserController {
	return &UserController{
		userUseCase: userUseCase,
	}
}


func (u *UserController) Register(c *gin.Context) {
	//TODO: Error Handling
	var user models.UserCreate
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	_, err := u.userUseCase.Register(c, &user)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	middleware.SuccessResponseHandler(c, http.StatusCreated, "Account Created Successfully", struct{}{})
}

func (u *UserController) Login(c *gin.Context) {
	var user models.UserCreate
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	token, _ := u.userUseCase.Login(c, &user)
	fmt.Println(token, "", user)
	middleware.SuccessResponseHandler(c, http.StatusOK, "Login Successful", gin.H{"token": token})
}

func (u *UserController) Delete(c *gin.Context) {
	username, _ := c.Get("username")
	user, err := u.userUseCase.GetByUsername(c, username.(string))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	deletedUser, err := u.userUseCase.Delete(c, user)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	middleware.SuccessResponseHandler(c, 200, "Task Deleted Successfully", deletedUser)
}
func (u *UserController) UpdateUsername(c *gin.Context) {


}
func (u *UserController) UpdatePassword(c *gin.Context) {

}
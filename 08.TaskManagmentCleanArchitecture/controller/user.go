package controller

import (
	"TaskManger/domain"

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


func (t *UserController) Create(c *gin.Context) {
}

func (t *UserController) Delete(c *gin.Context) {

}

func (t *UserController) UpdatePassword(c *gin.Context) {

}

func (t *UserController) UpdateUsername(c *gin.Context) {

}

func (t *UserController) Login(c *gin.Context) {

}

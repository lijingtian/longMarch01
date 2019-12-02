package controllers

import(
	"gopkg.in/gin-gonic/gin.v1"
)

type UserLogin struct{}

func NewUserLogin() *UserLogin{
	return new(UserLogin)
}

func (u *UserLogin) Login(c *gin.Context){
	c.JSON(200, gin.H{"message": "this is web site login"})
}
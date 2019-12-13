package controllers

import(
	"github.com/gin-gonic/gin"
	"longMarch01/webSite/models"
)

type User struct{}

func NewUser() *User{
	return new(User)
}

func (u *User) Login(c *gin.Context){
	c.JSON(200, gin.H{"message": "this is web site login"})
}

func (u *User) Register(c *gin.Context) {
	c.HTML(200, "User/Register.html", gin.H{})
}

func (u *User) RegisterJson(c *gin.Context) {
	var msg string
	m := models.NewUser()
	err := c.Bind(m)
	if err != nil{
		msg =  err.Error()
	} else {
		msg = m.Add()
	}
	c.JSON(200, msg)
}
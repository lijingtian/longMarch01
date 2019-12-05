package controllers

import(
	"gopkg.in/gin-gonic/gin.v1"
	"longMarch01/manageSystem/models"
)

type User struct{}

func NewUser() *User{
	return new(User)
}

func (u *User) Login(c *gin.Context){
	models.NewUser().GetUserInfo()
	c.JSON(200, gin.H{"message": "this is manage system login"})
}

func (u *User) Register(c *gin.Context) {
	c.HTML(200, "User/Register.html", gin.H{})
}

func (u *User) RegisterJson(c *gin.Context) {
	m := models.NewUser()
	var msg string
	if err := c.Bind(m); err != nil{
		msg = err.Error()
	} else {
		msg = m.InsertInto()
	}
	c.JSON(200, msg)
}
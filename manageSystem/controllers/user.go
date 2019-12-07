package controllers

import(
	"github.com/gin-gonic/gin"
	"longMarch01/manageSystem/models"

)

type User struct{}

func NewUser() *User{
	return new(User)
}

func (u *User) Login(c *gin.Context){
	c.HTML(200, "User/Login.html", gin.H{})
}

func (u *User) LoginJson(c *gin.Context) {
	var msg string
	name := c.DefaultQuery("name", "")
	password := c.DefaultQuery("password", "")
	m := models.NewUser()
	m.Name = name
	m.Password = password
	m.GetUserInfo()
	if m.Id > 0{
		m.SaveSession(c)
		msg = "success"
	} else {
		msg = "账号或者密码错误"
	}
	c.JSON(200, msg)
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
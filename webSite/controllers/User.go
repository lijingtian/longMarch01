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
	c.HTML(200, "User/Login.html", gin.H{})
}

func (u *User) LoginJson(c *gin.Context) {
	m := models.NewUser()
	m.Name = c.DefaultQuery("name", "")
	m.Password = c.DefaultQuery("password", "")
	m.CheckLogin()
	var msg string
	if m.Id > 0{
		msg = "success"
		m.SaveSession(c)
	} else {
		msg = "账号或密码错误"
	}
	c.JSON(200, msg)
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
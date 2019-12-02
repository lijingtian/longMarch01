package services

import(
	"gopkg.in/gin-gonic/gin.v1"
	"longMarch01/webSite/controllers"
)

func StartWebSite(){
	r := gin.Default()
	{
		c := controllers.NewUserLogin()
		r.GET("./login", c.Login)
	}
	r.Run(":8081")
}
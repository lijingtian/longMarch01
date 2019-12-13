package services

import(
	"github.com/gin-gonic/gin"
	"longMarch01/webSite/controllers"
)

func StartWebSite(){
	r := gin.Default()
	r.Static("/assets", "assets")
	r.LoadHTMLGlob("website/views/**/*")
	
	{
		g := r.Group("/user")
		c := controllers.NewUser()
		g.GET("./login", c.Login)
		g.GET("./register", c.Register)
		g.GET("./registerJson", c.RegisterJson)
	}
	r.Run(":8081")
}
package services

import(
	"github.com/gin-gonic/gin"
	"longMarch01/webSite/controllers"
	"longMarch01/webSite/middleware"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
)

func StartWebSite(){
	r := gin.Default()
	r.Static("/assets", "assets")
	r.LoadHTMLGlob("website/views/**/*")
	
	store := cookie.NewStore([]byte("webSite"))
	r.Use(sessions.Sessions("mySession", store))
	
	r.Use(middleware.CheckLogin)
	{
		g := r.Group("/user")
		c := controllers.NewUser()
		g.GET("/login", c.Login)
		g.GET("/loginJson", c.LoginJson)
		g.GET("/register", c.Register)
		g.GET("/registerJson", c.RegisterJson)
	}
	
	{
		g := r.Group("/goods")
		c := controllers.NewGoods()
		g.GET("/list", c.List)
		g.GET("/buy", c.Buy)
	}
	
	r.Run(":8081")
}
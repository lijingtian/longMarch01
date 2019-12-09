package services
import(
	"github.com/gin-gonic/gin"
	"longMarch01/manageSystem/controllers"
	"longMarch01/manageSystem/middleware"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
)


func StartManageSystem(){
	r := gin.Default()
	r.Static("/assets", "assets")
	r.LoadHTMLGlob("manageSystem/views/**/*")
	//设置session
	store := cookie.NewStore([]byte("manageSystem"))
	r.Use(sessions.Sessions("mySession", store))
	
	//检测登录
	r.Use(middleware.NewMiddleware().CheckLogin)
	
	{
		g := r.Group("/user")
		{
			c := controllers.NewUser()
			g.GET("/login", c.Login)
			g.GET("/loginJson", c.LoginJson)
			g.GET("/register", c.Register)
			g.GET("/registerJson", c.RegisterJson)
		}
	}
	{
		g := r.Group("/goods")
		{
			c := controllers.NewGoods()
			g.GET("/list", c.List)
			g.GET("/add", c.Add)
			g.GET("/addJson", c.AddJson)
		}
	}
	
	r.Run(":8082")
}
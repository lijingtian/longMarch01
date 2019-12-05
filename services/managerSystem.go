package services
import(
	"gopkg.in/gin-gonic/gin.v1"
	"longMarch01/manageSystem/controllers"
)

func StartManageSystem(){
	r := gin.Default()
	r.Static("/assets", "assets")
	r.LoadHTMLGlob("manageSystem/views/**/*")
	{
		g := r.Group("/user")
		{
			c := controllers.NewUser()
			g.GET("/login", c.Login)
			g.GET("/register", c.Register)
			g.GET("/registerJson", c.RegisterJson)
		}
	}
	
	r.Run(":8082")
}
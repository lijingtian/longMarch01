package services
import(
	"gopkg.in/gin-gonic/gin.v1"
	"longMarch01/manageSystem/controllers"
)

func StartManageSystem(){
	r := gin.Default()
	{
		c := controllers.NewUserLogin()
		r.GET("./login", c.Login)
	}
	r.Run(":8082")
}
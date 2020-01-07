package middleware

import(
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"strings"
)

type Middleware struct{}

func NewMiddleware() *Middleware{
	return new(Middleware)
}

func CheckLogin(c *gin.Context){
	url := c.Request.URL.Path
	if !strings.HasPrefix(url, "/user/login") {
		s := sessions.Default(c)
		if s.Get("userInfo") == nil{
			c.Redirect(302, "/user/login")
		}
	}
	c.Next()
}
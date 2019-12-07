package middleware

import(
	"strings"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"fmt"
)

type middleware struct{}

func NewMiddleware() *middleware{
	return new(middleware)
}

func (m *middleware) CheckLogin(c *gin.Context) {
	url := c.Request.URL.Path
	if !strings.HasPrefix(url, "/user/login") {
		session := sessions.Default(c)
		if session.Get("userInfo") == nil{
			c.Redirect(302, "/user/login")
		}
		fmt.Println("middleware.CheckLogin: ", session.Get("userInfo"))
	}
}
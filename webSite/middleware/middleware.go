package middleware

import (
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type Middleware struct{}

func NewMiddleware() *Middleware {
	return new(Middleware)
}

func CheckLogin(c *gin.Context) {
	url := c.Request.URL.Path
	if !strings.HasPrefix(url, "/user/login") && !strings.HasPrefix(url, "/user/register") {
		s := sessions.Default(c)
		if s.Get("userInfo") == nil {
			c.Redirect(302, "/user/login")
		}
	}
	c.Next()
}

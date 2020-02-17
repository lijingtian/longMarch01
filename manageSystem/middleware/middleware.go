package middleware

import (
	"fmt"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type middleware struct{}

func NewMiddleware() *middleware {
	return new(middleware)
}

func (m *middleware) CheckLogin(c *gin.Context) {
	url := c.Request.URL.Path
	if !strings.HasPrefix(url, "/user/login") && !strings.HasPrefix(url, "/user/register") {
		session := sessions.Default(c)
		if session.Get("userInfo") == nil {
			c.Redirect(302, "/user/login")
		}
		fmt.Println("middleware.CheckLogin: ", session.Get("userInfo"))
	}
}

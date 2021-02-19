package middleware

import (
	"regexp"
	"net/http"
	"auth-go/provider"
	"github.com/gin-gonic/gin"
)

func HandleTokenMiddleware(c *gin.Context) {
	// unAuth 接口不做 token 校验
	if regexp.MustCompile(`^/unAuth/`).FindString(c.Request.URL.Path) != "" {
		c.Next()
		return
	}
	if c.GetHeader("token") == "" {
		c.AbortWithStatusJSON(http.StatusOK, gin.H { "state": 500, "msg": "auth-server error: token isn't exist", "data": nil })
		return
	}
	if _, error := provider.ParseTokenString(c.GetHeader("token")); error != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H { "state": 500, "msg": "token is error", "data": nil })
		return
	}
	c.Next()
}

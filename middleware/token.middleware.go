package middleware

import (
	"auth-go/provider"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleTokenMiddleware(c *gin.Context) {
	if c.GetHeader("token") == "" {
		c.AbortWithStatusJSON(http.StatusOK, gin.H { "state": 500, "msg": "auth-server error: token isn't exist" })
		return
	}

	if _, error := provider.ParseTokenString(c.GetHeader("token")); error != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H { "state": 500, "msg": "token is error" })
		return
	}
	c.Next()
}

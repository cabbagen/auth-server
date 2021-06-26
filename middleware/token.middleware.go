package middleware

import (
	"net/http"
	"go-gateway/provider"
	"github.com/gin-gonic/gin"
)

var whitelist []string = []string {
	"/gateway/login",
	"/gateway/captcha",
	"/gateway/signToken",
	"/gateway/parseToken",
}

func isInWhitelistUrl(url string) bool {
	for _, value := range whitelist {
		if value == url {
			return true
		}
	}
	return false
}
func HandleTokenMiddleware(c *gin.Context) {
	if isInWhitelistUrl(c.Request.URL.String()) {
		c.Next()
		return
	}

	if c.GetHeader("token") == "" {
		c.AbortWithStatusJSON(http.StatusOK, gin.H { "state": 500, "msg": "auth-server error: token isn't exist", "data": nil })
		return
	}
	if tokenString, error := provider.ParseTokenString(c.GetHeader("token")); error != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H { "state": 500, "msg": "token is error", "data": nil })
		return
	} else {
		c.Set("parsed-token", tokenString)
	}
	c.Next()
}

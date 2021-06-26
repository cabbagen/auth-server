package middleware

import (
	"strings"
	"net/http"
	"github.com/gin-gonic/gin"
)

var defaultCorsOptions map[string]string = map[string]string {
	"Access-Control-Allow-Origin": "*",
	"Access-Control-Allow-Methods": "get, post, delete, put, options",
	"Access-Control-Allow-Headers": "content-type, token, app-key, x-requested-with",
}

func HandleCorsMiddleware(c *gin.Context) {
	for key, value := range defaultCorsOptions {
		c.Header(key, value)
	}

	if strings.ToUpper(c.Request.Method) == "OPTIONS" {
		c.String(http.StatusOK, "true")
		c.AbortWithStatus(http.StatusOK)
		return
	}
	c.Next()
}

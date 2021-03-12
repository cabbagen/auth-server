package middleware

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

var defaultCorsOptions map[string]string = map[string]string {
	"Access-Control-Allow-Origin": "*",
	"Access-Control-Allow-Headers": "content-type, token, app, x-requested-with",
	"Access-Control-Allow-Methods": "GET, POST, DELETE, PUT, OPTIONS",
}

func HandleCorsMiddleware(c *gin.Context) {
	for key, value := range defaultCorsOptions {
		c.Header(key, value)
	}

	if c.Request.Method == "OPTIONS" {
		c.String(http.StatusOK, "true")
		c.AbortWithStatus(http.StatusOK)
		return
	}
	c.Next()
}

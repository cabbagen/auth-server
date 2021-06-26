package middleware

import (
	"go-gateway/provider"
	"github.com/gin-gonic/gin"
)

func HandleTranceMiddleware(c *gin.Context) {
	tranceId := provider.NewTranceId()

	c.Set("tranceId", tranceId)
	c.Next()
}

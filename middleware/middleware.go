package middleware

import "github.com/gin-gonic/gin"

var globalMiddleware []gin.HandlerFunc = []gin.HandlerFunc {
	HandleCorsMiddleware,
	HandlePanicRecover,
	HandleAppKeyMiddleware,
	HandleTokenMiddleware,
	HandleTranceMiddleware,
	//HandleLoggerMiddleware(),
}

func RegisterMiddleware(app *gin.Engine) {
	app.Use(globalMiddleware...)
}

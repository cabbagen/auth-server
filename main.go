package main

import (
	"go-gateway/cache"
	"go-gateway/router"
	"go-gateway/database"
	"go-gateway/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)

	engine := gin.Default()

	// database
	database.Connect()

	defer database.Destroy()

	// cache
	redis := cache.NewRedisCache()

	redis.Connect()

	defer redis.Destroy()

	// application middleware
	middleware.RegisterMiddleware(engine)

	// application router
	router.RegisterRouter(engine)

	// run
	engine.Run(":7000") // listen and serve on 0.0.0.0:7000 (for windows "localhost:7000")
}

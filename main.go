package main

import (
	"auth-go/cache"
	"auth-go/database"
	"auth-go/middleware"
	"auth-go/router"
	"github.com/gin-gonic/gin"
)

func main() {
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

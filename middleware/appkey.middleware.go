package middleware

import (
	"net/http"
	"auth-go/cache"
	"auth-go/model"
	"auth-go/utils"
	"github.com/gin-gonic/gin"
)

func HandleAppKeyMiddleware(c *gin.Context) {
	if c.GetHeader("app_key") == "" {
		c.AbortWithStatusJSON(http.StatusOK, gin.H { "status": 500, "msg": "auth-server error: app_key isn't exist", "data": nil })
		return
	}

	appKeys, error := cache.GetRedisCacheInstance().GetSet("app_keys")

	if error != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H { "status": 500, "msg": "get cache app_keys error", "data": nil })
		return
	}

	if len(appKeys) == 0 {
		appKeysFromDatabase, error :=  model.NewApplicationModel().GetAllApplicationKey()

		if error != nil {
			c.JSON(http.StatusOK, gin.H { "status": 500, "msg": "get database app_keys error", "data": nil })
			return
		}

		cache.GetRedisCacheInstance().PushSet("app_keys", appKeysFromDatabase...)

		appKeys = appKeysFromDatabase
	}

	targetAppKey := utils.SliceFind(utils.StringSliceToInterfaces(appKeys), func(item interface{}, index int) bool {
		return item != nil && item.(string) == c.GetHeader("app_key")
	})

	if targetAppKey == nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H { "status": 500, "msg": "the header app_key error", "data": nil })
		return
	}
	c.Next()
}

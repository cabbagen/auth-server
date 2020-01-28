package router

import (
	"auth-go/controller"
	"github.com/gin-gonic/gin"
)

var proxyController controller.ProxyController

var proxyRoutes []description = []description {
	description {
		path: "/proxy/*url",
		method: "ANY",
		handlers: []gin.HandlerFunc { proxyController.HandleProxyRequest },
	},
}

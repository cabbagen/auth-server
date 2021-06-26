package router

import (
	"go-gateway/controller"
	"github.com/gin-gonic/gin"
)

var applicationController controller.ApplicationController

var applicationRoutes []description = []description {
	description {
		path: "/gateway/applications",
		method: "GET",
		handlers: []gin.HandlerFunc{ applicationController.HandleGetApplications },
	},
	description {
		path: "/gateway/application/:appId",
		method: "GET",
		handlers: []gin.HandlerFunc{ applicationController.HandleGetApplicationDetail },
	},
	description {
		path: "/gateway/application",
		method: "POST",
		handlers: []gin.HandlerFunc{ applicationController.HandleUpdateApplication },
	},
}

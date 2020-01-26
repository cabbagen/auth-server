package router

import (
	"auth-go/controller"
	"github.com/gin-gonic/gin"
)

var applicationController controller.ApplicationController

var applicationRoutes []description = []description {
	description {
		path: "/auth/applications",
		method: "GET",
		handlers: []gin.HandlerFunc{ applicationController.HandleGetApplications },
	},
	description {
		path: "/auth/application/:appId",
		method: "GET",
		handlers: []gin.HandlerFunc{ applicationController.HandleGetApplicationDetail },
	},
	description {
		path: "/auth/application",
		method: "POST",
		handlers: []gin.HandlerFunc{ applicationController.HandleUpdateApplication },
	},
}

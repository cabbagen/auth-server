package router

import (
	"go-gateway/controller"
	"github.com/gin-gonic/gin"
)

var loginController controller.LoginController

var loginRoutes []description = []description {
	description {
		path: "/gateway/captcha",
		method: "GET",
		handlers: []gin.HandlerFunc { loginController.HandleGenerateCaptcha },
	},
	description {
		path: "/gateway/login",
		method: "POST",
		handlers: []gin.HandlerFunc { loginController.HandleLogin },
	},
	description {
		path: "/gateway/parseToken",
		method: "POST",
		handlers: []gin.HandlerFunc { loginController.HandleParseTokenString },
	},
	description {
		path: "/gateway/signToken",
		method: "GET",
		handlers: []gin.HandlerFunc { loginController.HandleSignTokenString },
	},
}

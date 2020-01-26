package router

import (
	"auth-go/controller"
	"github.com/gin-gonic/gin"
)

var loginController controller.LoginController

var loginRoutes []description = []description {
	description {
		path: "/auth/captcha",
		method: "GET",
		handlers: []gin.HandlerFunc { loginController.HandleGenerateCaptcha },
	},
	description {
		path: "/auth/login",
		method: "POST",
		handlers: []gin.HandlerFunc { loginController.HandleLogin },
	},
	description {
		path: "/auth/parseToken",
		method: "POST",
		handlers: []gin.HandlerFunc { loginController.HandleParseTokenString },
	},
	description {
		path: "/auth/signToken",
		method: "GET",
		handlers: []gin.HandlerFunc { loginController.HandleSignTokenString },
	},
}

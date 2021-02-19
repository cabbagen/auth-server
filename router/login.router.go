package router

import (
	"auth-go/controller"
	"github.com/gin-gonic/gin"
)

var loginController controller.LoginController

var loginRoutes []description = []description {
	description {
		path: "/unAuth/captcha",
		method: "GET",
		handlers: []gin.HandlerFunc { loginController.HandleGenerateCaptcha },
	},
	description {
		path: "/unAuth/login",
		method: "POST",
		handlers: []gin.HandlerFunc { loginController.HandleLogin },
	},
	description {
		path: "/unAuth/parseToken",
		method: "POST",
		handlers: []gin.HandlerFunc { loginController.HandleParseTokenString },
	},
	description {
		path: "/unAuth/signToken",
		method: "GET",
		handlers: []gin.HandlerFunc { loginController.HandleSignTokenString },
	},
}

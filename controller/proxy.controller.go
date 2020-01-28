package controller

import (
	"auth-go/provider"
	"github.com/gin-gonic/gin"
	"strings"
)

type ProxyController struct {
	BaseController
}

func (pc ProxyController) HandleProxyRequest(c *gin.Context) {
	var requestUrl string = strings.Replace(c.Request.URL.String(), "/proxy", "", 1)

	content, error := provider.NewHttpProxy(c.Request.Method, requestUrl, c.Request.Body).Request()

	if error != nil {
		pc.HandleFailResponse(c, error)
	}
	c.String(200, string(content))
}

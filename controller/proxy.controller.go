package controller

import (
	"strings"
	"auth-go/provider"
	"github.com/gin-gonic/gin"
)

type ProxyController struct {
	BaseController
}

func (pc ProxyController) HandleProxyRequest(c *gin.Context) {
	var headers provider.ProxyHeader
	var requestUrl string = strings.Replace(c.Request.URL.String(), "/proxy", "", 1)

	if error := c.BindHeader(&headers); error != nil {
		pc.HandleFailResponse(c, error)
		return
	}

	content, error := provider.NewHttpProxy(c.Request.Method, requestUrl, c.Request.Body).Request(headers)

	if error != nil {
		pc.HandleFailResponse(c, error)
	}
	c.String(200, string(content))
}

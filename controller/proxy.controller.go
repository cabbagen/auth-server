package controller

import (
	"strings"
	"go-gateway/provider"
	"github.com/gin-gonic/gin"
)

type ProxyController struct {
	BaseController
}

func (pc ProxyController) HandleProxyRequest(c *gin.Context) {
	var headers provider.ProxyHeader
	var proxyHeaders map[string]string = make(map[string]string)
	var requestUrl string = strings.Replace(c.Request.URL.String(), "/proxy", "", 1)

	if error := c.BindHeader(&headers); error != nil {
		pc.HandleFailResponse(c, error)
		return
	}

	if value, isExist := c.Get("parsed-token"); isExist {
		proxyHeaders["parsed-token"] = value.(string)
	}

	content, error := provider.NewHttpProxy(c.Request.Method, requestUrl, c.Request.Body, proxyHeaders).Request(headers)

	if error != nil {
		pc.HandleFailResponse(c, error)
	}
	c.String(200, string(content))
}

package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type BaseController struct {
}

func (bc BaseController) HandleSuccessResponse(c *gin.Context, data interface{}) {
	tranceId, _ := c.Get("tranceId")
	c.JSON(http.StatusOK, gin.H { "status": 200, "msg": nil, "data": data, "tranceId": tranceId })
}

func (bc BaseController) HandleFailResponse(c *gin.Context, error error) {
	tranceId, _ := c.Get("tranceId")
	c.JSON(http.StatusOK, gin.H { "status": 500, "msg": error.Error(), "data": nil, "tranceId": tranceId })
}

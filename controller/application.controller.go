package controller

import (
	"auth-go/model"
	"auth-go/schema"
	"github.com/gin-gonic/gin"
	"strconv"
)

type ApplicationController struct {
	BaseController
}

func (ac ApplicationController) HandleGetApplicationDetail(c *gin.Context) {
	appId, error := strconv.Atoi(c.Param("appId"))

	if error != nil {
		ac.HandleFailResponse(c, error)
		return
	}

	info, error := model.NewApplicationModel().GetApplicationDetail(appId)

	if error != nil {
		ac.HandleFailResponse(c, error)
		return
	}
	ac.HandleSuccessResponse(c, info)
}

type HandleGetApplicationsParams struct {
	Name        string       `form:"name"`
	PageNo      int          `form:"pageNo"`
	PageSize    int          `form:"pageSize"`
}
func (ac ApplicationController) HandleGetApplications(c *gin.Context) {
	var params HandleGetApplicationsParams = HandleGetApplicationsParams {"", 0, 10}

	if error := c.BindQuery(&params); error != nil {
		ac.HandleFailResponse(c, error)
		return
	}
	applications, total, error := model.NewApplicationModel().GetApplications(params.Name, params.PageNo, params.PageSize)

	if error != nil {
		ac.HandleFailResponse(c, error)
		return
	}
	ac.HandleSuccessResponse(c, map[string]interface{} { "applications": applications, "total": total })
}

func (ac ApplicationController) HandleUpdateApplication(c *gin.Context) {
	var params schema.ApplicationSchema

	if error := c.BindJSON(&params); error != nil {
		ac.HandleFailResponse(c, error)
		return
	}
	if error := model.NewApplicationModel().UpdateApplication(params); error != nil {
		ac.HandleFailResponse(c, error)
		return
	}
	ac.HandleSuccessResponse(c, "更新成功")
}

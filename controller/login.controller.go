package controller

import (
	"auth-go/provider"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
)

type LoginController struct {
	BaseController
}

func (lc LoginController) HandleGenerateCaptcha(c *gin.Context) {
	captcha, error := provider.GenerateCaptcha()

	if error != nil {
		lc.HandleFailResponse(c, error)
		return
	}
	lc.HandleSuccessResponse(c, captcha)
}

type LoginParams struct {
	Username        string    `json:"username"`
	Password        string    `json:"password"`
	CaptchaId       string    `json:"captchaId"`
	Answer          string    `json:"answer"`
}
func (lc LoginController) HandleLogin(c *gin.Context) {
	var params LoginParams

	if error := c.BindJSON(&params); error != nil {
		lc.HandleFailResponse(c, error)
		return
	}

	if isOk := provider.ValidateCaptcha(params.CaptchaId, params.Answer); !isOk {
		lc.HandleFailResponse(c, errors.New("验证码不正确"))
		return
	}

	tokenString, error := provider.SignToken(fmt.Sprintf("%s-%s", params.Username, params.Password))

	if error != nil {
		lc.HandleFailResponse(c, error)
		return
	}

	lc.HandleSuccessResponse(c, gin.H { "token": tokenString })
}

type ParseTokenStringParams struct {
	Token           string      `json:"token"`
}
func (lc LoginController) HandleParseTokenString(c *gin.Context) {
	var params ParseTokenStringParams

	if error := c.BindJSON(&params); error != nil {
		lc.HandleFailResponse(c, error)
		return
	}

	responseString, error := provider.ParseTokenString(params.Token)

	if error != nil {
		lc.HandleFailResponse(c, error)
		return
	}
	lc.HandleSuccessResponse(c, responseString)
}

func (lc LoginController) HandleSignTokenString(c *gin.Context) {
	token, error := provider.SignToken("hello world xia")

	if error != nil {
		lc.HandleFailResponse(c, error)
		return
	}
	lc.HandleSuccessResponse(c, token)
}
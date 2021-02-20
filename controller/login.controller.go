package controller

import (
	"bytes"
	"errors"
	"encoding/json"
	"auth-go/provider"
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
	Username        string          `json:"username"`
	Password        string          `json:"password"`
	CaptchaId       string          `json:"captchaId"`
	Answer          string          `json:"answer"`
}

type ProxyLoginRequest struct {
	Username        string          `json:"username"`
	Password        string          `json:"password"`
}

type ProxyLoginResponse struct {
	Status          int             `json:"status"`
	Data            string          `json:"data"`
	Msg             string          `json:"msg"`
}
func (lc LoginController) HandleLogin(c *gin.Context) {
	var params LoginParams
	var headers provider.ProxyHeader

	if error := c.BindJSON(&params); error != nil {
		lc.HandleFailResponse(c, error)
		return
	}
	if error := c.BindHeader(&headers); error != nil {
		lc.HandleFailResponse(c, error)
		return
	}

	// 验证码校验
	if isOk := provider.ValidateCaptcha(params.CaptchaId, params.Answer); !isOk {
		lc.HandleFailResponse(c, errors.New("验证码不正确"))
		return
	}

	// 项目登录接口，校验用户身份
	var proxyLoginRequestParams ProxyLoginRequest = ProxyLoginRequest{params.Username, params.Password }

	proxyLoginParamsBytes, _ := json.Marshal(proxyLoginRequestParams)

	content, error := provider.NewHttpProxy("POST", "/handle/login", bytes.NewBuffer(proxyLoginParamsBytes)).Request(headers)

	if error != nil {
		lc.HandleFailResponse(c, error)
		return
	}

	var proxyLoginResponse ProxyLoginResponse

	if error := json.Unmarshal(content, &proxyLoginResponse); error != nil {
		lc.HandleFailResponse(c, error)
		return
	}

	if proxyLoginResponse.Status != 200 {
		lc.HandleFailResponse(c, errors.New("remote server handle/login error"))
		return
	}

	// 签发用户 token
	tokenString, error := provider.SignToken(proxyLoginResponse.Data)

	if error != nil {
		lc.HandleFailResponse(c, error)
		return
	}
	lc.HandleSuccessResponse(c, gin.H { "token": tokenString, "rawResponse":  proxyLoginResponse.Data })
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

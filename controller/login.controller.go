package controller

import (
	"bytes"
	"errors"
	"encoding/json"
	"go-gateway/provider"
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
	Code            int             `json:"code"`
	Data            string          `json:"data"`
	Message         string          `json:"message"`
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
	var proxyHeaders map[string]string = make(map[string]string)
	var proxyLoginRequestParams ProxyLoginRequest = ProxyLoginRequest{params.Username, params.Password }

	proxyLoginParamsBytes, _ := json.Marshal(proxyLoginRequestParams)

	if value, isExist := c.Get("parsed-token"); isExist {
		proxyHeaders["parsed-token"] = value.(string)
	}

	content, error := provider.NewHttpProxy("POST", "/" + c.Param("server") + "/handle/login", bytes.NewBuffer(proxyLoginParamsBytes), proxyHeaders).Request(headers)

	if error != nil {
		lc.HandleFailResponse(c, error)
		return
	}

	var proxyLoginResponse ProxyLoginResponse

	if error := json.Unmarshal(content, &proxyLoginResponse); error != nil {
		lc.HandleFailResponse(c, error)
		return
	}

	if proxyLoginResponse.Code != 100200 {
		lc.HandleFailResponse(c, errors.New(proxyLoginResponse.Message))
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

package provider

import "github.com/mojocn/base64Captcha"

var defaultCaptchaInstance *base64Captcha.Captcha = base64Captcha.NewCaptcha(
	base64Captcha.DefaultDriverDigit,
	base64Captcha.DefaultMemStore,
)

func GenerateCaptcha() (map[string]string, error) {
	var captchaInfo map[string]string = make(map[string]string)

	id, b64s, error := defaultCaptchaInstance.Generate()

	if error != nil {
		return nil, error
	}

	captchaInfo["b64s"] = b64s
	captchaInfo["captchaId"] = id

	return captchaInfo, nil
}

func ValidateCaptcha(id, answer string) bool {
	return defaultCaptchaInstance.Verify(id, answer, true)
}

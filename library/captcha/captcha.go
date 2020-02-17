package captcha

import (
	"github.com/mojocn/base64Captcha"
	"github.com/yamakiller/magicWeb/library/log"
)

//GenerateCaptcha doc
//@Method GenerateCaptcha @Summary Generate Image Captcha
//@Param (*gin.Context)
func GenerateCaptcha(w, h int) (string, string) {
	var id string = ""

	config := base64Captcha.ConfigCharacter{
		Height:             h,
		Width:              w,
		Mode:               base64Captcha.CaptchaModeNumberAlphabet,
		IsUseSimpleFont:    false,
		ComplexOfNoiseText: base64Captcha.CaptchaComplexLower,
		ComplexOfNoiseDot:  base64Captcha.CaptchaComplexLower,
		IsShowHollowLine:   false,
		IsShowNoiseDot:     true,
		IsShowNoiseText:    true,
		IsShowSlimeLine:    false,
		IsShowSineLine:     false,
		CaptchaLen:         6,
	}

	captchaID, digitCap := base64Captcha.GenerateCaptcha(id, config)
	base64Png := base64Captcha.CaptchaWriteToBase64Encoding(digitCap)

	log.Debug("Captcha Image Width:%d Height:%d", w, h)
	return captchaID, base64Png
}

//VerfiyCaptcha doc
//@Method VerfiyCaptcha @Summary Verfiy Captchea
//@Param (string) captcha ID
//@Param (string) submit value
func VerfiyCaptcha(captchaID, verifyValue string) bool {
	verifyResult := base64Captcha.VerifyCaptcha(captchaID, verifyValue)
	if verifyResult {
		return true
	}

	return false
}

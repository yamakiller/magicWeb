package captcha

import (
	"fmt"

	"github.com/mojocn/base64Captcha"
)

//GenerateCaptcha desc
//@method GenerateCaptcha desc: Generate Image Captcha
//@param (*gin.Context)
func GenerateCaptcha(w, h, m int) (string, string) {
	var id string

	config := base64Captcha.ConfigCharacter{
		Height:             w,
		Width:              h,
		Mode:               m,
		IsUseSimpleFont:    false,
		ComplexOfNoiseText: 0,
		ComplexOfNoiseDot:  0,
		IsShowHollowLine:   false,
		IsShowNoiseDot:     false,
		IsShowNoiseText:    false,
		IsShowSlimeLine:    false,
		IsShowSineLine:     false,
		CaptchaLen:         0,
	}

	captchaID, digitCap := base64Captcha.GenerateCaptcha(id, config)
	base64Png := base64Captcha.CaptchaWriteToBase64Encoding(digitCap)

	return captchaID， base64Png
}

//VerfiyCaptcha desc
//@method VerfiyCaptcha desc: Verfiy Captchea
//@param (string) captcha ID
//@param (string) submit value
func VerfiyCaptcha(captchaID, verifyValue string) bool {
	verifyResult := base64Captcha.VerifyCaptcha(captchaID, verifyValue)
	if verifyResult {
		return true
	}

	return false
}

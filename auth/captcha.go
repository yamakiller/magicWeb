package auth

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"github.com/yamakiller/magicWeb/param"
)

//GenerateCaptcha desc
//@method GenerateCaptcha desc: Generate Image Captcha
//@param (*gin.Context)
func GenerateCaptcha(c *gin.Context) {
	session := sessions.Default(c)
	var id string

	w := param.GetQueryInt(c, "width", 240)
	h := param.GetQueryInt(c, "height", 60)
	m := param.GetQueryInt(c, "mode", 2)

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
	session.Set("captchaID", captchaID)
	c.String(http.StatusOK, base64Png)
}

//VerfiyCaptcha desc
//@method VerfiyCaptcha desc: Verfiy Captchea
//@param (*gin.Context) captcha ID
//@param (string) submit value
func VerfiyCaptcha(c *gin.Context, verifyValue string) error {
	session := sessions.Default(c)
	verifyResult := base64Captcha.VerifyCaptcha(session.Get("captchaID"), verifyValue)
	if verifyResult {
		return nil
	}

	return fmt.Errorf("captcha is error")
}

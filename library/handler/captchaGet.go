package handler

import (
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/yamakiller/magicWeb/library/captcha"
	"github.com/yamakiller/magicWeb/library/code"
)

//GetCaptcha Return CaptchaImage
func GetCaptcha(context *gin.Context, width, height, mode int, captchaKey string) {
	captchID, base64Png := captcha.GenerateCaptcha(width, height, mode)
	session := sessions.Default(context)
	session.Set(captchaKey, captchID)

	context.JSON(http.StatusOK, code.SpawnResponse(base64Png))
}

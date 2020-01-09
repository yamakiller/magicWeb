package handler

import (
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/yamakiller/carPlays/app/config"
	"github.com/yamakiller/magicWeb/library/captcha"
)

//GetCaptcha Return CaptchaImage
func GetCaptcha(context *gin.Context, width, height, mode int, captchaKey string) {

	captchID, base64Png := captcha.GenerateCaptcha(width, height, mode)
	session := sessions.Default(context)
	session.Set(config.CaptchaSessionKey, captchID)

	context.String(http.StatusOK, code.SpawnResponse(base64Png))
}

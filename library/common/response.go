package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yamakiller/magicWeb/library/protocol"
)

//ResponseError doc
//Summary common method output error informat
//Param (*gin.Context) http context
//Param (protocol.Response) error message json
func ResponseError(context *gin.Context, resp protocol.Response) {
	context.JSON(http.StatusOK, resp)
	context.Abort()
}

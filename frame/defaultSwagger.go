package frame

import (
	"fmt"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//RegisterSwagger register swagger service
func RegisterSwagger(addr string, engine *gin.Engine) {
	url := ginSwagger.URL(fmt.Sprintf("http://%s/swagger/doc.json", addr))
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}

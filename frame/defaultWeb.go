package frame

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yamakiller/magicLibs/args"
	"github.com/yamakiller/magicLibs/envs"
	"github.com/yamakiller/magicWeb/library/log"
)

//DefaultWeb doc
//@Struct DefaultWeb @Summary Default web framework
type DefaultWeb struct {
	_release bool
	_router  *gin.Engine
	_start   func(IMagicWeb) error
}

//WithStart doc
//@Summary job start function to frame
//@Param (func()error)
func (slf *DefaultWeb) WithStart(f func(IMagicWeb) error) {
	slf._start = f
}

//Engine doc
//@Summary Returns gin engine
func (slf *DefaultWeb) Engine() *gin.Engine {
	return slf._router
}

//Start doc
//@Summary start system
//@Return (error) start fail returns error
func (slf *DefaultWeb) Start() error {
	slf._release = args.Instance().GetBoolean("-release", false)
	//------------------------------
	//create http router
	slf._router = gin.Default()
	if !slf._release {
		slf._router.Use(slf.cors())
	}
	slf._router.Use(slf.logMaps())
	if slf._start != nil {
		if err := slf._start(slf); err != nil {
			return err
		}
	}
	//------------------------
	addr := args.Instance().GetString("-addr", "0.0.0.0:8080")
	if slf._release {
		log.Info("HTTP on %s", addr)
	}

	slf._router.Run(addr)

	return nil
}

func (slf *DefaultWeb) cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}

func (slf *DefaultWeb) logMaps() gin.HandlerFunc {
	return func(c *gin.Context) {
		if slf._release {
			startTime := time.Now()
			c.Next()
			endTime := time.Now()
			latencyTime := endTime.Sub(startTime)
			reqMethod := c.Request.Method
			reqURI := c.Request.RequestURI
			statusCode := c.Writer.Status()
			clientIP := c.ClientIP()
			log.Info("%s %s %3d =>client:%15s time:%13v", reqURI, reqMethod, statusCode, clientIP, latencyTime)
		}
	}
}

//Shutdown doc
//@Summary shutdown system
func (slf *DefaultWeb) Shutdown() {
	envs.Instance().UnLoad()
}

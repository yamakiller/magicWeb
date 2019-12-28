package frame

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
	"github.com/yamakiller/magicLibs/args"
	"github.com/yamakiller/magicLibs/envs"
	"github.com/yamakiller/magicLibs/logger"
)

//DefaultWeb doc
//@Struct DefaultWeb @Summary Default web framework
type DefaultWeb struct {
	_log    logger.Logger
	_router *gin.Engine
	_start  func(IMagicWeb) error
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
	release := args.Instance().GetBoolean("-release", false)
	logPath := args.Instance().GetString("-logPath", "")
	logSize := args.Instance().GetInt("-logSize", 128)

	slf._log = logger.New(func() logger.Logger {
		l := logger.LogContext{}
		l.SetFilPath(logPath)
		l.SetHandle(logrus.New())
		l.SetMailMax(logSize)
		if release {
			l.SetLevel(logrus.ErrorLevel)
		} else {
			l.SetLevel(logrus.InfoLevel)
		}

		formatter := new(prefixed.TextFormatter)
		formatter.FullTimestamp = true
		formatter.TimestampFormat = "2006-01-02 15:04:05"
		formatter.DisableColors = true
		formatter.SetColorScheme(&prefixed.ColorScheme{
			PrefixStyle: "blue+b"})

		l.SetFormatter(formatter)
		l.Initial()
		l.Redirect()
		return &l
	})

	logger.WithDefault(slf._log)
	slf._log.Mount()

	if !release {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	//------------------------------
	//create http router
	slf._router = gin.Default()
	slf._router.Use(slf.logmap())
	if slf._start != nil {
		if err := slf._start(slf); err != nil {
			return err
		}
	}
	//------------------------
	addr := args.Instance().GetString("-addr", "0.0.0.0:8080")
	if release {
		logger.Info(0, "HTTP on %s", addr)
	}

	slf._router.Run(addr)

	return nil
}

func (slf *DefaultWeb) logmap() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)
		reqMethod := c.Request.Method
		reqURI := c.Request.RequestURI
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		logger.Info(0, "%s %s %3d =>client:%15s time:%13v", reqURI, reqMethod, statusCode, clientIP, latencyTime)
	}
}

//Shutdown doc
//@Summary shutdown system
func (slf *DefaultWeb) Shutdown() {

	if slf._log != nil {
		slf._log.Close()
		slf._log = nil
	}
	envs.Instance().UnLoad()
}

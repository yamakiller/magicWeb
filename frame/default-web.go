package frame

import (
	"runtime"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
	"github.com/yamakiller/magicLibs/args"
	"github.com/yamakiller/magicLibs/envs"
	"github.com/yamakiller/magicLibs/logger"
)

//DefaultWeb desc
//@struct DefaultWeb desc: Default web framework
type DefaultWeb struct {
	_log    logger.Logger
	_router *gin.Engine
	_start  func(IMagicWeb) error
}

//WithStart desc
//@method WithStart desc: job start function to frame
//@param (func()error)
func (slf *DefaultWeb) WithStart(f func(IMagicWeb) error) {
	slf._start = f
}

//Engine desc
//@method Engine desc: Returns gin engine
func (slf *DefaultWeb) Engine() *gin.Engine {
	return slf._router
}

//Start desc
//@method Start desc: start system
//@return (error) start fail returns error
func (slf *DefaultWeb) Start() error {
	logEnvPath := args.Instance().GetString("-l", "./config/log.json")
	logDeploy := logger.NewDefault()

	envs.Instance().Load(logger.EnvKey, logEnvPath, logDeploy)
	slf._log = logger.New(func() logger.Logger {
		l := logger.LogContext{}
		l.SetFilPath(logDeploy.LogPath)
		l.SetHandle(logrus.New())
		l.SetMailMax(logDeploy.LogSize)
		l.SetLevel(logrus.Level(logDeploy.LogLevel))

		formatter := new(prefixed.TextFormatter)
		formatter.FullTimestamp = true
		if runtime.GOOS == "windows" {
			formatter.DisableColors = true
		} else {
			formatter.SetColorScheme(&prefixed.ColorScheme{
				PrefixStyle: "blue+b"})
		}
		l.SetFormatter(formatter)
		l.Initial()
		l.Redirect()
		return &l
	})

	logger.WithDefault(slf._log)
	slf._log.Mount()

	release := args.Instance().GetBoolean("-release", false)
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
	addr := args.Instance().GetString("-addr", "0.0.0.0")
	port := args.Instance().GetInt("-port", 8080)
	logger.Info(0, "HTTP on %s:%d", addr, port)
	slf._router.Run(addr + ":" + strconv.Itoa(port))

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
		slf._log.Info(0, "%s %s %3d/client:%15s/time:%13v", reqURI, reqMethod, statusCode, clientIP, latencyTime)
	}
}

//Shutdown desc
//@method Shutdown desc: shutdown system
func (slf *DefaultWeb) Shutdown() {

	if slf._log != nil {
		slf._log.Close()
		slf._log = nil
	}
	envs.Instance().UnLoad()
}

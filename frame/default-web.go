package frame

import (
	"runtime"
	"strconv"
	"strings"
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
	_start  func() error
}

//WithStart desc
//@method WithStart desc: job start function to frame
//@param (func()error)
func (slf *DefaultWeb) WithStart(f func() error) {
	slf._start = f
}

//RegisterGroup desc
//@method RegisterGroup desc: Create Http Router Group
//@return (gin.Group)
func (slf *DefaultWeb) RegisterGroup(URI string, handlers ...gin.HandlerFunc) *gin.RouterGroup {
	return slf._router.Group(URI, handlers...)
}

//RegisterMethod desc
//@method RegisterMethod desc: Register Http Router
//@param (*gin.RouterGroup) Router group
//@param (string) uri
//@param (string) method [get/post/put/delete/options/head]
//@param (...gin.HandlerFunc)
func (slf *DefaultWeb) RegisterMethod(g *gin.RouterGroup,
	URI string,
	method string,
	handler ...gin.HandlerFunc) {
	switch strings.ToLower(method) {
	default:
		fallthrough
	case "get":
		if g != nil {
			g.GET(URI, handler...)
			return
		}
		slf._router.GET(URI, handler...)
	case "post":
		if g != nil {
			g.POST(URI, handler...)
			return
		}
		slf._router.POST(URI, handler...)
	case "put":
		if g != nil {
			g.PUT(URI, handler...)
			return
		}
		slf._router.PUT(URI, handler...)
	case "delete":
		if g != nil {
			g.DELETE(URI, handler...)
			return
		}
		slf._router.DELETE(URI, handler...)
	case "options":
		if g != nil {
			g.OPTIONS(URI, handler...)
			return
		}
		slf._router.OPTIONS(URI, handler...)
	case "PATCH":
		if g != nil {
			g.PATCH(URI, handler...)
			return
		}
		slf._router.PATCH(URI, handler...)
	case "HEAD":
		if g != nil {
			g.HEAD(URI, handler...)
			return
		}
		slf._router.HEAD(URI, handler...)
	}
}

//LoadHTMLGlob desc
//@method LoadHTMLGlob desc: Load Html blob
//@param (string) pattern
func (slf *DefaultWeb) LoadHTMLGlob(pattern string) {
	slf._router.LoadHTMLGlob(pattern)
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
		if err := slf._start(); err != nil {
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

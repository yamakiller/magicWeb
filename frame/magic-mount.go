package frame

import (
	"runtime"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
	"github.com/yamakiller/magicLibs/args"
	"github.com/yamakiller/magicLibs/envs"
	"github.com/yamakiller/magicLibs/logger"
)

//MHttpMethod desc
//@type (func(*gin.Context) interface{}) desc: http router method
//type MHttpMethod func(*gin.Context)
//type MMiddleWare func() MHttpMethod

//MagicMount desc
//@struct MagicMount desc:
//@member (logger.Logger) log module
//@member (*gin.Engine) http gin frame
type MagicMount struct {
	_log    logger.Logger
	_router *gin.Engine
}

//RegisterGroup desc
//@method RegisterGroup desc: Create Http Router Group
//@return (gin.Group)
func (slf *MagicMount) RegisterGroup(URI string, handlers ...gin.HandlerFunc) *gin.RouterGroup {
	return slf._router.Group(URI, handlers...)
}

//RegisterMethod desc
func (slf *MagicMount) RegisterMethod(g *gin.RouterGroup,
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

//Start desc
//@method Start desc: start system
func (slf *MagicMount) Start() error {
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

	//------------------------------
	//create http router
	slf._router = gin.Default()
	slf._router.LoadHTMLGlob("template/*")
	//------------------------------
	//------------------------------
	//register router
	if err := slf.FillRouter(); err != nil {
		return err
	}
	//signal monitoring
	//------------------------
	addr := args.Instance().GetString("-addr", "0.0.0.0")
	port := args.Instance().GetInt("-port", 8080)
	logger.Info(0, "http/%s/%d Monut", addr, port)
	slf._router.Run(addr + ":" + strconv.Itoa(port))

	return nil
}

//FillRouter desc
//@method FillRouter desc: load http router informat
func (slf *MagicMount) FillRouter() error {
	return nil
}

//Shutdown desc
//@method Shutdown desc: shutdown system
func (slf *MagicMount) Shutdown() {
	if slf._log != nil {
		slf._log.Close()
		slf._log = nil
	}
	envs.Instance().UnLoad()
}

func (slf *MagicMount) signalWatch() {

}

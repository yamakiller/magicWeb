package frame

import (
	"runtime"

	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
	"github.com/yamakiller/magicLibs/args"
	"github.com/yamakiller/magicLibs/envs"
)

//MagicMount desc
//@struct MagicMount desc:
type MagicMount struct {
	sysLogger logger.Logger
}

//LoadArgs desc
//@method LoadArgs desc: load command args
//@return (error)
func (slf *MagicMount) LoadArgs() error {
	args.Instance().Parse()
	return nil
}

//LoadEnv desc
//@method LoadEnv desc: load env
func (slf *MagicMount) LoadEnv() error {
	logEnvPath := args.Instance().GetString("-l", "./config/log.json")
	logDeploy := logger.NewDefault()
	envs.Instance().Load(logger.EnvKey, logEnvPath, logDeploy)
	return ni
}

//Start desc
//@method Start desc: start system
func (slf *MagicMount) Start() error {
	logDeplay := envs.Instance().Get(logger.EnvKey).(*logger.LogDeploy)
	slf.sysLogger = logger.New(func() logger.Logger {
		l := logger.LogContext{}
		l.SetFilPath(logDeplay.LogPath)
		l.SetHandle(logrus.New())
		l.SetMailMax(logDeplay.LogSize)
		l.SetLevel(logrus.Level(logDeplay.LogLevel))

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

	logger.WithDefault(slf.sysLogger)
	slf.sysLogger.Mount()

	//注册路由
	//启动服务

	return nil
}

//Loop desc
//@method Loop desc: system loop wait
func (slf *MagicMount) Loop() {

}

//Shutdown desc
//@method Shutdown desc: shutdown system
func (slf *MagicMount) Shutdown() {
	envs.Instance().Unload()
}

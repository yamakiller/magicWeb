package boot

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
	"github.com/yamakiller/magicLibs/args"
	"github.com/yamakiller/magicLibs/logger"
	"github.com/yamakiller/magicLibs/util"
	"github.com/yamakiller/magicWeb/frame"
)

//Lanuch doc
//@Summary lanuch web system
//@Param (frame.Spawn) Spawn framework function
func Lanuch(spawn frame.Spawn) {
	frm := spawn()
	var err error
	args.Instance().Parse()

	release := args.Instance().GetBoolean("-release", false)
	logPath := args.Instance().GetString("-logPath", "")
	logSize := args.Instance().GetInt("-logSize", 128)

	log := logger.New(func() logger.Logger {
		l := logger.LogContext{}
		l.SetFilPath(logPath)
		l.SetHandle(logrus.New())
		l.SetMailMax(logSize)
		if release {
			l.SetLevel(logrus.InfoLevel)
		} else {
			l.SetLevel(logrus.DebugLevel)
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

	logger.WithDefault(log)
	log.Mount()

	mode := "DEBUG"
	if !release {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
		mode = "RELEASE"
	}

	if release {
		logger.Info(0, "                     _      __    __     _                        _       _")
		logger.Info(0, "   /\\/\\   __ _  __ _(_) ___/ / /\\ \\ \\___| |__     /\\/\\   ___   __| |_   _| | ___")
		logger.Info(0, "  /    \\ / _` |/ _` | |/ __\\ \\/  \\/ / _ \\ '_ \\   /    \\ / _ \\ / _` | | | | |/ _ \\")
		logger.Info(0, " / /\\/\\ \\ (_| | (_| | | (__ \\  /\\  /  __/ |_) | / /\\/\\ \\ (_) | (_| | |_| | |  __/")
		logger.Info(0, " \\/    \\/\\__,_|\\__, |_|\\___| \\/  \\/ \\___|_.__/  \\/    \\/\\___/ \\__,_|\\__,_|_|\\___|")
		logger.Info(0, "  ::magic net::|___/ (%s %s)", mode, util.TimeNowFormat())
		logger.Info(0, "----------------------------------------------------------------------------------")
	}

	if err = frm.Start(); err != nil {
		logger.Error(0, "%s", err.Error())
		goto exit
	}

exit:
	if log != nil {
		log.Close()
		log = nil
	}
	frm.Shutdown()
}

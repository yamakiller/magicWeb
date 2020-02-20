package boot

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/yamakiller/magicLibs/args"
	liblog "github.com/yamakiller/magicLibs/log"
	"github.com/yamakiller/magicLibs/util"
	"github.com/yamakiller/magicWeb/frame"
	"github.com/yamakiller/magicWeb/library/log"
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
	logName := args.Instance().GetString("-logName", "")

	logLevel := logrus.DebugLevel
	if release {
		logLevel = logrus.InfoLevel
	}

	tmpLog, err := liblog.SpawnFileLogrus(logLevel,
		logPath,
		logName)
	if err != nil {
		panic(err)
	}

	agentLog := &liblog.DefaultAgent{}
	agentLog.WithHandle(tmpLog)
	log.WithLog(agentLog)

	mode := "DEBUG"
	if !release {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
		mode = "RELEASE"
	}

	if release {
		log.Info("                     _      __    __     _                        _       _")
		log.Info("   /\\/\\   __ _  __ _(_) ___/ / /\\ \\ \\___| |__     /\\/\\   ___   __| |_   _| | ___")
		log.Info("  /    \\ / _` |/ _` | |/ __\\ \\/  \\/ / _ \\ '_ \\   /    \\ / _ \\ / _` | | | | |/ _ \\")
		log.Info(" / /\\/\\ \\ (_| | (_| | | (__ \\  /\\  /  __/ |_) | / /\\/\\ \\ (_) | (_| | |_| | |  __/")
		log.Info(" \\/    \\/\\__,_|\\__, |_|\\___| \\/  \\/ \\___|_.__/  \\/    \\/\\___/ \\__,_|\\__,_|_|\\___|")
		log.Info("  ::magic net::|___/ (%s %s)", mode, util.TimestampFormat())
		log.Info("----------------------------------------------------------------------------------")
	}

	if err = frm.Start(); err != nil {
		log.Error("%s", err.Error())
		goto exit
	}

exit:
	if agentLog != nil {
		agentLog.Close()
		agentLog = nil
	}
	frm.Shutdown()
}
